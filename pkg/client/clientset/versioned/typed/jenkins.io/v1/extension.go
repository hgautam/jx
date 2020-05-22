// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/jenkins-x/jx/v2/pkg/apis/jenkins.io/v1"
	scheme "github.com/jenkins-x/jx/v2/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ExtensionsGetter has a method to return a ExtensionInterface.
// A group's client should implement this interface.
type ExtensionsGetter interface {
	Extensions(namespace string) ExtensionInterface
}

// ExtensionInterface has methods to work with Extension resources.
type ExtensionInterface interface {
	Create(*v1.Extension) (*v1.Extension, error)
	Update(*v1.Extension) (*v1.Extension, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Extension, error)
	List(opts metav1.ListOptions) (*v1.ExtensionList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Extension, err error)
	ExtensionExpansion
}

// extensions implements ExtensionInterface
type extensions struct {
	client rest.Interface
	ns     string
}

// newExtensions returns a Extensions
func newExtensions(c *JenkinsV1Client, namespace string) *extensions {
	return &extensions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the extension, and returns the corresponding extension object, and an error if there is any.
func (c *extensions) Get(name string, options metav1.GetOptions) (result *v1.Extension, err error) {
	result = &v1.Extension{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("extensions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Extensions that match those selectors.
func (c *extensions) List(opts metav1.ListOptions) (result *v1.ExtensionList, err error) {
	result = &v1.ExtensionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("extensions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested extensions.
func (c *extensions) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("extensions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a extension and creates it.  Returns the server's representation of the extension, and an error, if there is any.
func (c *extensions) Create(extension *v1.Extension) (result *v1.Extension, err error) {
	result = &v1.Extension{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("extensions").
		Body(extension).
		Do().
		Into(result)
	return
}

// Update takes the representation of a extension and updates it. Returns the server's representation of the extension, and an error, if there is any.
func (c *extensions) Update(extension *v1.Extension) (result *v1.Extension, err error) {
	result = &v1.Extension{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("extensions").
		Name(extension.Name).
		Body(extension).
		Do().
		Into(result)
	return
}

// Delete takes name of the extension and deletes it. Returns an error if one occurs.
func (c *extensions) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("extensions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *extensions) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("extensions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched extension.
func (c *extensions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Extension, err error) {
	result = &v1.Extension{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("extensions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
