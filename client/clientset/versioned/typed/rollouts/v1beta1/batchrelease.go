/*
Copyright 2024 The Kruise Authors.

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

package v1beta1

import (
	"context"
	"time"

	scheme "github.com/openkruise/kruise-rollout-api/client/clientset/versioned/scheme"
	v1beta1 "github.com/openkruise/kruise-rollout-api/rollouts/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BatchReleasesGetter has a method to return a BatchReleaseInterface.
// A group's client should implement this interface.
type BatchReleasesGetter interface {
	BatchReleases(namespace string) BatchReleaseInterface
}

// BatchReleaseInterface has methods to work with BatchRelease resources.
type BatchReleaseInterface interface {
	Create(ctx context.Context, batchRelease *v1beta1.BatchRelease, opts v1.CreateOptions) (*v1beta1.BatchRelease, error)
	Update(ctx context.Context, batchRelease *v1beta1.BatchRelease, opts v1.UpdateOptions) (*v1beta1.BatchRelease, error)
	UpdateStatus(ctx context.Context, batchRelease *v1beta1.BatchRelease, opts v1.UpdateOptions) (*v1beta1.BatchRelease, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.BatchRelease, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.BatchReleaseList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BatchRelease, err error)
	BatchReleaseExpansion
}

// batchReleases implements BatchReleaseInterface
type batchReleases struct {
	client rest.Interface
	ns     string
}

// newBatchReleases returns a BatchReleases
func newBatchReleases(c *RolloutsV1beta1Client, namespace string) *batchReleases {
	return &batchReleases{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the batchRelease, and returns the corresponding batchRelease object, and an error if there is any.
func (c *batchReleases) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.BatchRelease, err error) {
	result = &v1beta1.BatchRelease{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("batchreleases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BatchReleases that match those selectors.
func (c *batchReleases) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.BatchReleaseList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.BatchReleaseList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("batchreleases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested batchReleases.
func (c *batchReleases) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("batchreleases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a batchRelease and creates it.  Returns the server's representation of the batchRelease, and an error, if there is any.
func (c *batchReleases) Create(ctx context.Context, batchRelease *v1beta1.BatchRelease, opts v1.CreateOptions) (result *v1beta1.BatchRelease, err error) {
	result = &v1beta1.BatchRelease{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("batchreleases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(batchRelease).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a batchRelease and updates it. Returns the server's representation of the batchRelease, and an error, if there is any.
func (c *batchReleases) Update(ctx context.Context, batchRelease *v1beta1.BatchRelease, opts v1.UpdateOptions) (result *v1beta1.BatchRelease, err error) {
	result = &v1beta1.BatchRelease{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("batchreleases").
		Name(batchRelease.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(batchRelease).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *batchReleases) UpdateStatus(ctx context.Context, batchRelease *v1beta1.BatchRelease, opts v1.UpdateOptions) (result *v1beta1.BatchRelease, err error) {
	result = &v1beta1.BatchRelease{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("batchreleases").
		Name(batchRelease.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(batchRelease).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the batchRelease and deletes it. Returns an error if one occurs.
func (c *batchReleases) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("batchreleases").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *batchReleases) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("batchreleases").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched batchRelease.
func (c *batchReleases) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BatchRelease, err error) {
	result = &v1beta1.BatchRelease{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("batchreleases").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}