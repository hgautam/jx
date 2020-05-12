package amazon

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	session2 "github.com/jenkins-x/jx/pkg/cloud/amazon/session"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/google/uuid"
	"github.com/jenkins-x/jx/pkg/cloud"
	"github.com/jenkins-x/jx/pkg/config"
	"github.com/jenkins-x/jx/pkg/helm"
	"github.com/jenkins-x/jx/pkg/log"
	"github.com/jenkins-x/jx/pkg/util"
	"github.com/pkg/errors"
	"k8s.io/helm/pkg/chartutil"
)

const (
	// PoliciesTemplateName is the name of the custom policies CloudFormation stack that will be executed before
	// calling the eksctl commands
	PoliciesTemplateName = "jenkinsx-policies.yml"
	// ConfigTemplatesFolder is part of the path to the configuration templates
	ConfigTemplatesFolder = "templates"
	// IRSATemplateName is the name of the eksctl configuration file that will be processed after creating the policies
	IRSATemplateName = "irsa.tmpl.yaml"
)

// EnableIRSASupportInCluster Associates IAM as an OIDC provider so it can sign requests and assume roles
func EnableIRSASupportInCluster(requirements *config.RequirementsConfig) error {
	log.Logger().Infof("Enabling IRSA for cluster %s associating the IAM Open ID Connect provider", util.ColorInfo(requirements.Cluster.ClusterName))
	args := []string{"utils", "associate-iam-oidc-provider", "--cluster", requirements.Cluster.ClusterName, "--region", requirements.Cluster.Region, "--approve"}
	err := executeEksctlCommand(args)
	if err != nil {
		return errors.Wrap(err, "there was a porblem enabling IRSA in the cluster")
	}
	return nil
}

// CreateIRSAManagedServiceAccounts takes the KubeProviders directory and the requirements configuration and creates
// new ServiceAccounts annotated with a role ARN that is generated by eksctl. The policies attached to these roles
// are defined in the jenkinsx-policies.yml file within kubeProviders/eks/templates
// Note: this can't yet be executed in the master pipeline of the Dev Environment because in order to recreate the
// ServiceAccounts, we need to delete them and the roles first, which causes the next commands to fail
func CreateIRSAManagedServiceAccounts(requirements *config.RequirementsConfig, kubeProvidersDir string) error {
	templateValues, err := createPoliciesStack(requirements, kubeProvidersDir)
	if err != nil {
		return errors.Wrap(err, "there was a problem creating the policies stack and returning the template values")
	}

	processedTemplateFile, err := processIRSATemplateWithValues(requirements, kubeProvidersDir, templateValues)
	if err != nil {
		return errors.Wrap(err, "there was a problem processing the IRSA template with the provided values")
	}
	defer util.DeleteFile(processedTemplateFile.Name()) //nolint:errcheck

	err = deleteIAMServiceAccount(processedTemplateFile)
	if err != nil {
		return errors.Wrap(err, "failure creating the IRSA managed service accounts")
	}

	err = executeIRSAConfigFile(processedTemplateFile)
	if err != nil {
		return errors.Wrap(err, "failure creating the IRSA managed service accounts")
	}
	return nil
}

// createPoliciesStack reads the jenkinsx-policies.yml CloudFormation stack template and executes it, providing a
// random UUID as a parameter and extracting the outputs of the stack, removing the suffix from them and adding them to
// the returned map so it can be used as parameters for the Go Template irsa.tmpl.yaml
func createPoliciesStack(requirements *config.RequirementsConfig, kubeProvidersDir string) (map[string]interface{}, error) {
	eksKubeProviderDir := filepath.Join(kubeProvidersDir, cloud.EKS, ConfigTemplatesFolder)
	session, err := session2.NewAwsSession("", requirements.Cluster.Region)
	if err != nil {
		return nil, errors.Wrap(err, "error creating a new AWS Session")
	}
	cfn := cloudformation.New(session)
	policiesTemplate, err := ioutil.ReadFile(filepath.Join(eksKubeProviderDir, PoliciesTemplateName))
	if err != nil {
		return nil, err
	}
	suffix := uuid.New().String()
	describeInput := &cloudformation.DescribeStacksInput{
		StackName: aws.String(fmt.Sprintf("JenkinsXPolicies-%s", suffix)),
	}

	log.Logger().Infof("Creating CloudFormation stack %s", util.ColorInfo(*describeInput.StackName))
	_, err = cfn.CreateStack(&cloudformation.CreateStackInput{
		Capabilities: []*string{aws.String("CAPABILITY_NAMED_IAM")},
		StackName:    describeInput.StackName,
		Tags: []*cloudformation.Tag{{
			Key:   aws.String("CreatedBy"),
			Value: aws.String("Jenkins-x"),
		}},
		Parameters: []*cloudformation.Parameter{
			{
				ParameterKey:   aws.String("PoliciesSuffixParameter"),
				ParameterValue: aws.String(suffix),
			},
		},
		TemplateBody: aws.String(string(policiesTemplate)),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "there was a problem creating the %s CloudFormation stack", *describeInput.StackName)
	}

	log.Logger().Infof("Waiting until CloudFormation stack %s is created", util.ColorInfo(*describeInput.StackName))
	err = cfn.WaitUntilStackCreateComplete(describeInput)
	if err != nil {
		return nil, errors.Wrapf(err, "there was a problem waiting for the %s CloudFormation stack to be created", *describeInput.StackName)
	}

	log.Logger().Infof("Describing stack %s to extract outputs", util.ColorInfo(*describeInput.StackName))
	describeOutput, err := cfn.DescribeStacks(describeInput)
	if err != nil {
		return nil, errors.Wrapf(err, "there was a problem describing the %s CloudFormation stack to extract the outputs", *describeInput.StackName)
	}

	templateValues := make(map[string]interface{})
	if len(describeOutput.Stacks) > 0 {
		outputs := describeOutput.Stacks[0].Outputs
		log.Logger().Debugf("Exported Outputs from stack %s:", util.ColorInfo(*describeInput.StackName))
		for _, value := range outputs {
			log.Logger().Debugf("ExportName: %s, Value: %s", util.ColorInfo(*value.ExportName), util.ColorInfo(*value.OutputValue))
			exportName := strings.Replace(*value.ExportName, "-"+suffix, "", -1)
			templateValues[exportName] = *value.OutputValue
		}
	}
	return templateValues, nil
}

// processIRSATemplateWithValues processes the template irsa.tmpl.yaml using the Go templates API with the provided templateValues which will be added
// with the IAM key so it can be referenced in the template
func processIRSATemplateWithValues(requirements *config.RequirementsConfig, kubeProvidersDir string, templateValues map[string]interface{}) (*os.File, error) {
	templatePath := filepath.Join(kubeProvidersDir, cloud.EKS, ConfigTemplatesFolder, IRSATemplateName)
	tmpl, err := template.New(IRSATemplateName).Option("missingkey=error").Funcs(helm.NewFunctionMap()).ParseFiles(templatePath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse Secrets template: %s", templatePath)
	}

	requirementsMap, err := requirements.ToMap()
	if err != nil {
		return nil, errors.Wrapf(err, "failed turn requirements into a map: %+v", requirements)
	}

	templateData := map[string]interface{}{
		"Requirements": chartutil.Values(requirementsMap),
		"IAM":          chartutil.Values(templateValues),
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, templateData)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to execute Secrets template: %s", templatePath)
	}

	f, err := ioutil.TempFile("", "irsa-template-")
	if err != nil {
		return nil, errors.Wrap(err, "there was a problem creating a temp file for the IRSA template")
	}
	_, err = f.Write(buf.Bytes())
	if err != nil {
		return nil, errors.Wrap(err, "there was a problem writing the IRSA template to the temp file")
	}

	return f, nil
}

func executeEksctlCommand(args []string) error {
	eksCtlInfo := util.ColorInfo("eksctl")
	log.Logger().Debugf("executing \"%s %s\"", eksCtlInfo, util.ColorInfo(strings.Join(args, " ")))
	cmd := util.Command{
		Name: "eksctl",
		Args: args,
		Out:  os.Stdout,
		Err:  os.Stderr,
	}
	_, err := cmd.RunWithoutRetry()
	if err != nil {
		return errors.Wrapf(err, "there was a problem calling eksctl with the provided args")
	}
	return nil
}

func executeIRSAConfigFile(file *os.File) error {
	log.Logger().Info("Creating IRSA ServiceAccounts")
	args := []string{"create", "iamserviceaccount",
		"--override-existing-serviceaccounts",
		"--config-file", file.Name(),
		"--include=\"*\"",
		"--approve"}
	err := executeEksctlCommand(args)
	if err != nil {
		return errors.Wrap(err, "there was a problem executing the IRSA ConfigFile")
	}
	return nil
}

func deleteIAMServiceAccount(file *os.File) error {
	log.Logger().Info("Deleting IRSA ServiceAccounts")
	args := []string{"delete", "iamserviceaccount",
		"--config-file", file.Name(),
		"--include=\"*\"",
		"--approve",
		"--wait"}
	err := executeEksctlCommand(args)
	if err != nil {
		return errors.Wrapf(err, "there was a problem deleting IAM ServiceAccounts")
	}
	return nil
}
