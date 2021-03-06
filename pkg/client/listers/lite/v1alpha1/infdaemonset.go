/*
Copyright 2022 The OpenYurt Authors.

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

package v1alpha1

import (
	v1alpha1 "github.com/openyurtio/kole/pkg/apis/lite/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// InfDaemonSetLister helps list InfDaemonSets.
// All objects returned here must be treated as read-only.
type InfDaemonSetLister interface {
	// List lists all InfDaemonSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InfDaemonSet, err error)
	// InfDaemonSets returns an object that can list and get InfDaemonSets.
	InfDaemonSets(namespace string) InfDaemonSetNamespaceLister
	InfDaemonSetListerExpansion
}

// infDaemonSetLister implements the InfDaemonSetLister interface.
type infDaemonSetLister struct {
	indexer cache.Indexer
}

// NewInfDaemonSetLister returns a new InfDaemonSetLister.
func NewInfDaemonSetLister(indexer cache.Indexer) InfDaemonSetLister {
	return &infDaemonSetLister{indexer: indexer}
}

// List lists all InfDaemonSets in the indexer.
func (s *infDaemonSetLister) List(selector labels.Selector) (ret []*v1alpha1.InfDaemonSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InfDaemonSet))
	})
	return ret, err
}

// InfDaemonSets returns an object that can list and get InfDaemonSets.
func (s *infDaemonSetLister) InfDaemonSets(namespace string) InfDaemonSetNamespaceLister {
	return infDaemonSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InfDaemonSetNamespaceLister helps list and get InfDaemonSets.
// All objects returned here must be treated as read-only.
type InfDaemonSetNamespaceLister interface {
	// List lists all InfDaemonSets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InfDaemonSet, err error)
	// Get retrieves the InfDaemonSet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.InfDaemonSet, error)
	InfDaemonSetNamespaceListerExpansion
}

// infDaemonSetNamespaceLister implements the InfDaemonSetNamespaceLister
// interface.
type infDaemonSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all InfDaemonSets in the indexer for a given namespace.
func (s infDaemonSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.InfDaemonSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InfDaemonSet))
	})
	return ret, err
}

// Get retrieves the InfDaemonSet from the indexer for a given namespace and name.
func (s infDaemonSetNamespaceLister) Get(name string) (*v1alpha1.InfDaemonSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("infdaemonset"), name)
	}
	return obj.(*v1alpha1.InfDaemonSet), nil
}
