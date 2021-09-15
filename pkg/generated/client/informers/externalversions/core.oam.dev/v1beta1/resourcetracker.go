/*
Copyright 2021 The KubeVela Authors.

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

package v1beta1

import (
	"context"
	time "time"

	coreoamdevv1beta1 "github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	versioned "github.com/oam-dev/kubevela-core-api/pkg/generated/client/clientset/versioned"
	internalinterfaces "github.com/oam-dev/kubevela-core-api/pkg/generated/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/oam-dev/kubevela-core-api/pkg/generated/client/listers/core.oam.dev/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ResourceTrackerInformer provides access to a shared informer and lister for
// ResourceTrackers.
type ResourceTrackerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ResourceTrackerLister
}

type resourceTrackerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewResourceTrackerInformer constructs a new informer for ResourceTracker type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResourceTrackerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredResourceTrackerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredResourceTrackerInformer constructs a new informer for ResourceTracker type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResourceTrackerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1beta1().ResourceTrackers(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1beta1().ResourceTrackers(namespace).Watch(context.TODO(), options)
			},
		},
		&coreoamdevv1beta1.ResourceTracker{},
		resyncPeriod,
		indexers,
	)
}

func (f *resourceTrackerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredResourceTrackerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *resourceTrackerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&coreoamdevv1beta1.ResourceTracker{}, f.defaultInformer)
}

func (f *resourceTrackerInformer) Lister() v1beta1.ResourceTrackerLister {
	return v1beta1.NewResourceTrackerLister(f.Informer().GetIndexer())
}
