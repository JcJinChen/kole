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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	litev1alpha1 "github.com/openyurtio/kole/pkg/apis/lite/v1alpha1"
	versioned "github.com/openyurtio/kole/pkg/client/clientset/versioned"
	internalinterfaces "github.com/openyurtio/kole/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/openyurtio/kole/pkg/client/listers/lite/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// InfEdgeNodeInformer provides access to a shared informer and lister for
// InfEdgeNodes.
type InfEdgeNodeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.InfEdgeNodeLister
}

type infEdgeNodeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewInfEdgeNodeInformer constructs a new informer for InfEdgeNode type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewInfEdgeNodeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredInfEdgeNodeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredInfEdgeNodeInformer constructs a new informer for InfEdgeNode type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredInfEdgeNodeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LiteV1alpha1().InfEdgeNodes(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LiteV1alpha1().InfEdgeNodes(namespace).Watch(context.TODO(), options)
			},
		},
		&litev1alpha1.InfEdgeNode{},
		resyncPeriod,
		indexers,
	)
}

func (f *infEdgeNodeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredInfEdgeNodeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *infEdgeNodeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&litev1alpha1.InfEdgeNode{}, f.defaultInformer)
}

func (f *infEdgeNodeInformer) Lister() v1alpha1.InfEdgeNodeLister {
	return v1alpha1.NewInfEdgeNodeLister(f.Informer().GetIndexer())
}
