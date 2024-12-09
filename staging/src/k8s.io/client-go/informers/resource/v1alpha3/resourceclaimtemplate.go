/*
Copyright The Kubernetes Authors.

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

package v1alpha3

import (
	context "context"
	time "time"

	apiresourcev1alpha3 "k8s.io/api/resource/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	kubernetes "k8s.io/client-go/kubernetes"
	resourcev1alpha3 "k8s.io/client-go/listers/resource/v1alpha3"
	cache "k8s.io/client-go/tools/cache"
)

// ResourceClaimTemplateInformer provides access to a shared informer and lister for
// ResourceClaimTemplates.
type ResourceClaimTemplateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() resourcev1alpha3.ResourceClaimTemplateLister
}

type resourceClaimTemplateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewResourceClaimTemplateInformer constructs a new informer for ResourceClaimTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResourceClaimTemplateInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredResourceClaimTemplateInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredResourceClaimTemplateInformer constructs a new informer for ResourceClaimTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResourceClaimTemplateInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha3().ResourceClaimTemplates(namespace).List(context.Background(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha3().ResourceClaimTemplates(namespace).Watch(context.Background(), options)
			},
			ListWithContextFunc: func(ctx context.Context, options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha3().ResourceClaimTemplates(namespace).List(ctx, options)
			},
			WatchFuncWithContext: func(ctx context.Context, options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha3().ResourceClaimTemplates(namespace).Watch(ctx, options)
			},
		},
		&apiresourcev1alpha3.ResourceClaimTemplate{},
		resyncPeriod,
		indexers,
	)
}

func (f *resourceClaimTemplateInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredResourceClaimTemplateInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *resourceClaimTemplateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiresourcev1alpha3.ResourceClaimTemplate{}, f.defaultInformer)
}

func (f *resourceClaimTemplateInformer) Lister() resourcev1alpha3.ResourceClaimTemplateLister {
	return resourcev1alpha3.NewResourceClaimTemplateLister(f.Informer().GetIndexer())
}
