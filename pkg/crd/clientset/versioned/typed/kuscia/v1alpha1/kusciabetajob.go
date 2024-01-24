// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/secretflow/kuscia/pkg/crd/apis/kuscia/v1alpha1"
	scheme "github.com/secretflow/kuscia/pkg/crd/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KusciaBetaJobsGetter has a method to return a KusciaBetaJobInterface.
// A group's client should implement this interface.
type KusciaBetaJobsGetter interface {
	KusciaBetaJobs(namespace string) KusciaBetaJobInterface
}

// KusciaBetaJobInterface has methods to work with KusciaBetaJob resources.
type KusciaBetaJobInterface interface {
	Create(ctx context.Context, kusciaBetaJob *v1alpha1.KusciaBetaJob, opts v1.CreateOptions) (*v1alpha1.KusciaBetaJob, error)
	Update(ctx context.Context, kusciaBetaJob *v1alpha1.KusciaBetaJob, opts v1.UpdateOptions) (*v1alpha1.KusciaBetaJob, error)
	UpdateStatus(ctx context.Context, kusciaBetaJob *v1alpha1.KusciaBetaJob, opts v1.UpdateOptions) (*v1alpha1.KusciaBetaJob, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.KusciaBetaJob, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.KusciaBetaJobList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KusciaBetaJob, err error)
	KusciaBetaJobExpansion
}

// kusciaBetaJobs implements KusciaBetaJobInterface
type kusciaBetaJobs struct {
	client rest.Interface
	ns     string
}

// newKusciaBetaJobs returns a KusciaBetaJobs
func newKusciaBetaJobs(c *KusciaV1alpha1Client, namespace string) *kusciaBetaJobs {
	return &kusciaBetaJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kusciaBetaJob, and returns the corresponding kusciaBetaJob object, and an error if there is any.
func (c *kusciaBetaJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KusciaBetaJob, err error) {
	result = &v1alpha1.KusciaBetaJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kusciabetajobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KusciaBetaJobs that match those selectors.
func (c *kusciaBetaJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KusciaBetaJobList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.KusciaBetaJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kusciabetajobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kusciaBetaJobs.
func (c *kusciaBetaJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kusciabetajobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a kusciaBetaJob and creates it.  Returns the server's representation of the kusciaBetaJob, and an error, if there is any.
func (c *kusciaBetaJobs) Create(ctx context.Context, kusciaBetaJob *v1alpha1.KusciaBetaJob, opts v1.CreateOptions) (result *v1alpha1.KusciaBetaJob, err error) {
	result = &v1alpha1.KusciaBetaJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kusciabetajobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kusciaBetaJob).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a kusciaBetaJob and updates it. Returns the server's representation of the kusciaBetaJob, and an error, if there is any.
func (c *kusciaBetaJobs) Update(ctx context.Context, kusciaBetaJob *v1alpha1.KusciaBetaJob, opts v1.UpdateOptions) (result *v1alpha1.KusciaBetaJob, err error) {
	result = &v1alpha1.KusciaBetaJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kusciabetajobs").
		Name(kusciaBetaJob.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kusciaBetaJob).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *kusciaBetaJobs) UpdateStatus(ctx context.Context, kusciaBetaJob *v1alpha1.KusciaBetaJob, opts v1.UpdateOptions) (result *v1alpha1.KusciaBetaJob, err error) {
	result = &v1alpha1.KusciaBetaJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kusciabetajobs").
		Name(kusciaBetaJob.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kusciaBetaJob).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the kusciaBetaJob and deletes it. Returns an error if one occurs.
func (c *kusciaBetaJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kusciabetajobs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kusciaBetaJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kusciabetajobs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched kusciaBetaJob.
func (c *kusciaBetaJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KusciaBetaJob, err error) {
	result = &v1alpha1.KusciaBetaJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kusciabetajobs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}