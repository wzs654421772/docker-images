// Copyright 2017-2018 The nats-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	"time"

	v1alpha2 "github.com/nats-io/nats-operator/pkg/apis/nats/v1alpha2"
	scheme "github.com/nats-io/nats-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NatsClustersGetter has a method to return a NatsClusterInterface.
// A group's client should implement this interface.
type NatsClustersGetter interface {
	NatsClusters(namespace string) NatsClusterInterface
}

// NatsClusterInterface has methods to work with NatsCluster resources.
type NatsClusterInterface interface {
	Create(*v1alpha2.NatsCluster) (*v1alpha2.NatsCluster, error)
	Update(*v1alpha2.NatsCluster) (*v1alpha2.NatsCluster, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.NatsCluster, error)
	List(opts v1.ListOptions) (*v1alpha2.NatsClusterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.NatsCluster, err error)
	NatsClusterExpansion
}

// natsClusters implements NatsClusterInterface
type natsClusters struct {
	client rest.Interface
	ns     string
}

// newNatsClusters returns a NatsClusters
func newNatsClusters(c *NatsV1alpha2Client, namespace string) *natsClusters {
	return &natsClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the natsCluster, and returns the corresponding natsCluster object, and an error if there is any.
func (c *natsClusters) Get(name string, options v1.GetOptions) (result *v1alpha2.NatsCluster, err error) {
	result = &v1alpha2.NatsCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("natsclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NatsClusters that match those selectors.
func (c *natsClusters) List(opts v1.ListOptions) (result *v1alpha2.NatsClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.NatsClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("natsclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested natsClusters.
func (c *natsClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("natsclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a natsCluster and creates it.  Returns the server's representation of the natsCluster, and an error, if there is any.
func (c *natsClusters) Create(natsCluster *v1alpha2.NatsCluster) (result *v1alpha2.NatsCluster, err error) {
	result = &v1alpha2.NatsCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("natsclusters").
		Body(natsCluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a natsCluster and updates it. Returns the server's representation of the natsCluster, and an error, if there is any.
func (c *natsClusters) Update(natsCluster *v1alpha2.NatsCluster) (result *v1alpha2.NatsCluster, err error) {
	result = &v1alpha2.NatsCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("natsclusters").
		Name(natsCluster.Name).
		Body(natsCluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the natsCluster and deletes it. Returns an error if one occurs.
func (c *natsClusters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("natsclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *natsClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("natsclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched natsCluster.
func (c *natsClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.NatsCluster, err error) {
	result = &v1alpha2.NatsCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("natsclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
