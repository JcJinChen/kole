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

// InfDaemonSetInformer provides access to a shared informer and lister for
// InfDaemonSets.
type InfDaemonSetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.InfDaemonSetLister
}

type infDaemonSetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewInfDaemonSetInformer constructs a new informer for InfDaemonSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewInfDaemonSetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredInfDaemonSetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredInfDaemonSetInformer constructs a new informer for InfDaemonSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredInfDaemonSetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LiteV1alpha1().InfDaemonSets(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LiteV1alpha1().InfDaemonSets(namespace).Watch(context.TODO(), options)
			},
		},
		&litev1alpha1.InfDaemonSet{},
		resyncPeriod,
		indexers,
	)
}

func (f *infDaemonSetInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredInfDaemonSetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *infDaemonSetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&litev1alpha1.InfDaemonSet{}, f.defaultInformer)
}

func (f *infDaemonSetInformer) Lister() v1alpha1.InfDaemonSetLister {
	return v1alpha1.NewInfDaemonSetLister(f.Informer().GetIndexer())
}