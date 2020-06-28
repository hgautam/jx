// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/jenkins-x/jx/v2/pkg/apis/jenkins.io/v1"
	scheme "github.com/jenkins-x/jx/v2/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SourceRepositoryGroupsGetter has a method to return a SourceRepositoryGroupInterface.
// A group's client should implement this interface.
type SourceRepositoryGroupsGetter interface {
	SourceRepositoryGroups(namespace string) SourceRepositoryGroupInterface
}

// SourceRepositoryGroupInterface has methods to work with SourceRepositoryGroup resources.
type SourceRepositoryGroupInterface interface {
	Create(*v1.SourceRepositoryGroup) (*v1.SourceRepositoryGroup, error)
	Update(*v1.SourceRepositoryGroup) (*v1.SourceRepositoryGroup, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.SourceRepositoryGroup, error)
	List(opts metav1.ListOptions) (*v1.SourceRepositoryGroupList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.SourceRepositoryGroup, err error)
	SourceRepositoryGroupExpansion
}

// sourceRepositoryGroups implements SourceRepositoryGroupInterface
type sourceRepositoryGroups struct {
	client rest.Interface
	ns     string
}

// newSourceRepositoryGroups returns a SourceRepositoryGroups
func newSourceRepositoryGroups(c *JenkinsV1Client, namespace string) *sourceRepositoryGroups {
	return &sourceRepositoryGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sourceRepositoryGroup, and returns the corresponding sourceRepositoryGroup object, and an error if there is any.
func (c *sourceRepositoryGroups) Get(name string, options metav1.GetOptions) (result *v1.SourceRepositoryGroup, err error) {
	result = &v1.SourceRepositoryGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sourcerepositorygroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SourceRepositoryGroups that match those selectors.
func (c *sourceRepositoryGroups) List(opts metav1.ListOptions) (result *v1.SourceRepositoryGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SourceRepositoryGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sourcerepositorygroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sourceRepositoryGroups.
func (c *sourceRepositoryGroups) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sourcerepositorygroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sourceRepositoryGroup and creates it.  Returns the server's representation of the sourceRepositoryGroup, and an error, if there is any.
func (c *sourceRepositoryGroups) Create(sourceRepositoryGroup *v1.SourceRepositoryGroup) (result *v1.SourceRepositoryGroup, err error) {
	result = &v1.SourceRepositoryGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sourcerepositorygroups").
		Body(sourceRepositoryGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sourceRepositoryGroup and updates it. Returns the server's representation of the sourceRepositoryGroup, and an error, if there is any.
func (c *sourceRepositoryGroups) Update(sourceRepositoryGroup *v1.SourceRepositoryGroup) (result *v1.SourceRepositoryGroup, err error) {
	result = &v1.SourceRepositoryGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sourcerepositorygroups").
		Name(sourceRepositoryGroup.Name).
		Body(sourceRepositoryGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the sourceRepositoryGroup and deletes it. Returns an error if one occurs.
func (c *sourceRepositoryGroups) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sourcerepositorygroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sourceRepositoryGroups) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sourcerepositorygroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sourceRepositoryGroup.
func (c *sourceRepositoryGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.SourceRepositoryGroup, err error) {
	result = &v1.SourceRepositoryGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sourcerepositorygroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
