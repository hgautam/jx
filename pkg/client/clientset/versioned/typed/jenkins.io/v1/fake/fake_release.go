// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	jenkinsiov1 "github.com/jenkins-x/jx/v2/pkg/apis/jenkins.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeReleases implements ReleaseInterface
type FakeReleases struct {
	Fake *FakeJenkinsV1
	ns   string
}

var releasesResource = schema.GroupVersionResource{Group: "jenkins.io", Version: "v1", Resource: "releases"}

var releasesKind = schema.GroupVersionKind{Group: "jenkins.io", Version: "v1", Kind: "Release"}

// Get takes name of the release, and returns the corresponding release object, and an error if there is any.
func (c *FakeReleases) Get(name string, options v1.GetOptions) (result *jenkinsiov1.Release, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(releasesResource, c.ns, name), &jenkinsiov1.Release{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.Release), err
}

// List takes label and field selectors, and returns the list of Releases that match those selectors.
func (c *FakeReleases) List(opts v1.ListOptions) (result *jenkinsiov1.ReleaseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(releasesResource, releasesKind, c.ns, opts), &jenkinsiov1.ReleaseList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &jenkinsiov1.ReleaseList{ListMeta: obj.(*jenkinsiov1.ReleaseList).ListMeta}
	for _, item := range obj.(*jenkinsiov1.ReleaseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested releases.
func (c *FakeReleases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(releasesResource, c.ns, opts))

}

// Create takes the representation of a release and creates it.  Returns the server's representation of the release, and an error, if there is any.
func (c *FakeReleases) Create(release *jenkinsiov1.Release) (result *jenkinsiov1.Release, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(releasesResource, c.ns, release), &jenkinsiov1.Release{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.Release), err
}

// Update takes the representation of a release and updates it. Returns the server's representation of the release, and an error, if there is any.
func (c *FakeReleases) Update(release *jenkinsiov1.Release) (result *jenkinsiov1.Release, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(releasesResource, c.ns, release), &jenkinsiov1.Release{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.Release), err
}

// Delete takes name of the release and deletes it. Returns an error if one occurs.
func (c *FakeReleases) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(releasesResource, c.ns, name), &jenkinsiov1.Release{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReleases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(releasesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &jenkinsiov1.ReleaseList{})
	return err
}

// Patch applies the patch and returns the patched release.
func (c *FakeReleases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *jenkinsiov1.Release, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(releasesResource, c.ns, name, pt, data, subresources...), &jenkinsiov1.Release{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.Release), err
}
