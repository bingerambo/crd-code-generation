/*
Copyright by bingerambo@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package fake

import (
	inspur_com_v1 "github.com/bingerambo/crd-code-generation/pkg/apis/inspur.com/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNodeCaches implements NodeCacheInterface
type FakeNodeCaches struct {
	Fake *FakeInspurV1
	ns   string
}

var nodecachesResource = schema.GroupVersionResource{Group: "inspur.com", Version: "v1", Resource: "nodecaches"}

var nodecachesKind = schema.GroupVersionKind{Group: "inspur.com", Version: "v1", Kind: "NodeCache"}

// Get takes name of the nodeCache, and returns the corresponding nodeCache object, and an error if there is any.
func (c *FakeNodeCaches) Get(name string, options v1.GetOptions) (result *inspur_com_v1.NodeCache, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(nodecachesResource, c.ns, name), &inspur_com_v1.NodeCache{})

	if obj == nil {
		return nil, err
	}
	return obj.(*inspur_com_v1.NodeCache), err
}

// List takes label and field selectors, and returns the list of NodeCaches that match those selectors.
func (c *FakeNodeCaches) List(opts v1.ListOptions) (result *inspur_com_v1.NodeCacheList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(nodecachesResource, nodecachesKind, c.ns, opts), &inspur_com_v1.NodeCacheList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &inspur_com_v1.NodeCacheList{}
	for _, item := range obj.(*inspur_com_v1.NodeCacheList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodeCaches.
func (c *FakeNodeCaches) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(nodecachesResource, c.ns, opts))

}

// Create takes the representation of a nodeCache and creates it.  Returns the server's representation of the nodeCache, and an error, if there is any.
func (c *FakeNodeCaches) Create(nodeCache *inspur_com_v1.NodeCache) (result *inspur_com_v1.NodeCache, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(nodecachesResource, c.ns, nodeCache), &inspur_com_v1.NodeCache{})

	if obj == nil {
		return nil, err
	}
	return obj.(*inspur_com_v1.NodeCache), err
}

// Update takes the representation of a nodeCache and updates it. Returns the server's representation of the nodeCache, and an error, if there is any.
func (c *FakeNodeCaches) Update(nodeCache *inspur_com_v1.NodeCache) (result *inspur_com_v1.NodeCache, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(nodecachesResource, c.ns, nodeCache), &inspur_com_v1.NodeCache{})

	if obj == nil {
		return nil, err
	}
	return obj.(*inspur_com_v1.NodeCache), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNodeCaches) UpdateStatus(nodeCache *inspur_com_v1.NodeCache) (*inspur_com_v1.NodeCache, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(nodecachesResource, "status", c.ns, nodeCache), &inspur_com_v1.NodeCache{})

	if obj == nil {
		return nil, err
	}
	return obj.(*inspur_com_v1.NodeCache), err
}

// Delete takes name of the nodeCache and deletes it. Returns an error if one occurs.
func (c *FakeNodeCaches) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(nodecachesResource, c.ns, name), &inspur_com_v1.NodeCache{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodeCaches) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(nodecachesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &inspur_com_v1.NodeCacheList{})
	return err
}

// Patch applies the patch and returns the patched nodeCache.
func (c *FakeNodeCaches) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *inspur_com_v1.NodeCache, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nodecachesResource, c.ns, name, data, subresources...), &inspur_com_v1.NodeCache{})

	if obj == nil {
		return nil, err
	}
	return obj.(*inspur_com_v1.NodeCache), err
}
