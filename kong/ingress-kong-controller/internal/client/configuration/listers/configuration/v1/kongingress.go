/*
Copyright 2018 The Kong Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/kong/kubernetes-ingress-controller/internal/apis/configuration/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KongIngressLister helps list KongIngresses.
type KongIngressLister interface {
	// List lists all KongIngresses in the indexer.
	List(selector labels.Selector) (ret []*v1.KongIngress, err error)
	// KongIngresses returns an object that can list and get KongIngresses.
	KongIngresses(namespace string) KongIngressNamespaceLister
	KongIngressListerExpansion
}

// kongIngressLister implements the KongIngressLister interface.
type kongIngressLister struct {
	indexer cache.Indexer
}

// NewKongIngressLister returns a new KongIngressLister.
func NewKongIngressLister(indexer cache.Indexer) KongIngressLister {
	return &kongIngressLister{indexer: indexer}
}

// List lists all KongIngresses in the indexer.
func (s *kongIngressLister) List(selector labels.Selector) (ret []*v1.KongIngress, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.KongIngress))
	})
	return ret, err
}

// KongIngresses returns an object that can list and get KongIngresses.
func (s *kongIngressLister) KongIngresses(namespace string) KongIngressNamespaceLister {
	return kongIngressNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KongIngressNamespaceLister helps list and get KongIngresses.
type KongIngressNamespaceLister interface {
	// List lists all KongIngresses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.KongIngress, err error)
	// Get retrieves the KongIngress from the indexer for a given namespace and name.
	Get(name string) (*v1.KongIngress, error)
	KongIngressNamespaceListerExpansion
}

// kongIngressNamespaceLister implements the KongIngressNamespaceLister
// interface.
type kongIngressNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KongIngresses in the indexer for a given namespace.
func (s kongIngressNamespaceLister) List(selector labels.Selector) (ret []*v1.KongIngress, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.KongIngress))
	})
	return ret, err
}

// Get retrieves the KongIngress from the indexer for a given namespace and name.
func (s kongIngressNamespaceLister) Get(name string) (*v1.KongIngress, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("kongingress"), name)
	}
	return obj.(*v1.KongIngress), nil
}
