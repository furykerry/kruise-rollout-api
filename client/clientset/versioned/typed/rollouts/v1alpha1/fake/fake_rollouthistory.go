/*
Copyright 2022 The Kruise Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/openkruise/kruise-rollout-api/rollouts/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRolloutHistories implements RolloutHistoryInterface
type FakeRolloutHistories struct {
	Fake *FakeRolloutsV1alpha1
	ns   string
}

var rollouthistoriesResource = schema.GroupVersionResource{Group: "rollouts.kruise.io", Version: "v1alpha1", Resource: "rollouthistories"}

var rollouthistoriesKind = schema.GroupVersionKind{Group: "rollouts.kruise.io", Version: "v1alpha1", Kind: "RolloutHistory"}

// Get takes name of the rolloutHistory, and returns the corresponding rolloutHistory object, and an error if there is any.
func (c *FakeRolloutHistories) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RolloutHistory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rollouthistoriesResource, c.ns, name), &v1alpha1.RolloutHistory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RolloutHistory), err
}

// List takes label and field selectors, and returns the list of RolloutHistories that match those selectors.
func (c *FakeRolloutHistories) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RolloutHistoryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rollouthistoriesResource, rollouthistoriesKind, c.ns, opts), &v1alpha1.RolloutHistoryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RolloutHistoryList{ListMeta: obj.(*v1alpha1.RolloutHistoryList).ListMeta}
	for _, item := range obj.(*v1alpha1.RolloutHistoryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rolloutHistories.
func (c *FakeRolloutHistories) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rollouthistoriesResource, c.ns, opts))

}

// Create takes the representation of a rolloutHistory and creates it.  Returns the server's representation of the rolloutHistory, and an error, if there is any.
func (c *FakeRolloutHistories) Create(ctx context.Context, rolloutHistory *v1alpha1.RolloutHistory, opts v1.CreateOptions) (result *v1alpha1.RolloutHistory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rollouthistoriesResource, c.ns, rolloutHistory), &v1alpha1.RolloutHistory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RolloutHistory), err
}

// Update takes the representation of a rolloutHistory and updates it. Returns the server's representation of the rolloutHistory, and an error, if there is any.
func (c *FakeRolloutHistories) Update(ctx context.Context, rolloutHistory *v1alpha1.RolloutHistory, opts v1.UpdateOptions) (result *v1alpha1.RolloutHistory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rollouthistoriesResource, c.ns, rolloutHistory), &v1alpha1.RolloutHistory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RolloutHistory), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRolloutHistories) UpdateStatus(ctx context.Context, rolloutHistory *v1alpha1.RolloutHistory, opts v1.UpdateOptions) (*v1alpha1.RolloutHistory, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(rollouthistoriesResource, "status", c.ns, rolloutHistory), &v1alpha1.RolloutHistory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RolloutHistory), err
}

// Delete takes name of the rolloutHistory and deletes it. Returns an error if one occurs.
func (c *FakeRolloutHistories) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(rollouthistoriesResource, c.ns, name), &v1alpha1.RolloutHistory{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRolloutHistories) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rollouthistoriesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RolloutHistoryList{})
	return err
}

// Patch applies the patch and returns the patched rolloutHistory.
func (c *FakeRolloutHistories) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RolloutHistory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rollouthistoriesResource, c.ns, name, pt, data, subresources...), &v1alpha1.RolloutHistory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RolloutHistory), err
}
