// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/jenkins-x/jx/v2/pkg/apis/jenkins.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EnvironmentLister helps list Environments.
type EnvironmentLister interface {
	// List lists all Environments in the indexer.
	List(selector labels.Selector) (ret []*v1.Environment, err error)
	// Environments returns an object that can list and get Environments.
	Environments(namespace string) EnvironmentNamespaceLister
	EnvironmentListerExpansion
}

// environmentLister implements the EnvironmentLister interface.
type environmentLister struct {
	indexer cache.Indexer
}

// NewEnvironmentLister returns a new EnvironmentLister.
func NewEnvironmentLister(indexer cache.Indexer) EnvironmentLister {
	return &environmentLister{indexer: indexer}
}

// List lists all Environments in the indexer.
func (s *environmentLister) List(selector labels.Selector) (ret []*v1.Environment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Environment))
	})
	return ret, err
}

// Environments returns an object that can list and get Environments.
func (s *environmentLister) Environments(namespace string) EnvironmentNamespaceLister {
	return environmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EnvironmentNamespaceLister helps list and get Environments.
type EnvironmentNamespaceLister interface {
	// List lists all Environments in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Environment, err error)
	// Get retrieves the Environment from the indexer for a given namespace and name.
	Get(name string) (*v1.Environment, error)
	EnvironmentNamespaceListerExpansion
}

// environmentNamespaceLister implements the EnvironmentNamespaceLister
// interface.
type environmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Environments in the indexer for a given namespace.
func (s environmentNamespaceLister) List(selector labels.Selector) (ret []*v1.Environment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Environment))
	})
	return ret, err
}

// Get retrieves the Environment from the indexer for a given namespace and name.
func (s environmentNamespaceLister) Get(name string) (*v1.Environment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("environment"), name)
	}
	return obj.(*v1.Environment), nil
}
