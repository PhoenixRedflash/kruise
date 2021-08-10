/*
Copyright 2021 The Kruise Authors.

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
	time "time"

	appsv1alpha1 "github.com/openkruise/kruise/apis/apps/v1alpha1"
	versioned "github.com/openkruise/kruise/pkg/client/clientset/versioned"
	internalinterfaces "github.com/openkruise/kruise/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/openkruise/kruise/pkg/client/listers/apps/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ResourceDistributionInformer provides access to a shared informer and lister for
// ResourceDistributions.
type ResourceDistributionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ResourceDistributionLister
}

type resourceDistributionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewResourceDistributionInformer constructs a new informer for ResourceDistribution type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResourceDistributionInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredResourceDistributionInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredResourceDistributionInformer constructs a new informer for ResourceDistribution type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResourceDistributionInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1alpha1().ResourceDistributions().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1alpha1().ResourceDistributions().Watch(options)
			},
		},
		&appsv1alpha1.ResourceDistribution{},
		resyncPeriod,
		indexers,
	)
}

func (f *resourceDistributionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredResourceDistributionInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *resourceDistributionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&appsv1alpha1.ResourceDistribution{}, f.defaultInformer)
}

func (f *resourceDistributionInformer) Lister() v1alpha1.ResourceDistributionLister {
	return v1alpha1.NewResourceDistributionLister(f.Informer().GetIndexer())
}
