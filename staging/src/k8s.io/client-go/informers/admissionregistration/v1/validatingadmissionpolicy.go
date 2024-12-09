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

package v1

import (
	context "context"
	time "time"

	apiadmissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	kubernetes "k8s.io/client-go/kubernetes"
	admissionregistrationv1 "k8s.io/client-go/listers/admissionregistration/v1"
	cache "k8s.io/client-go/tools/cache"
)

// ValidatingAdmissionPolicyInformer provides access to a shared informer and lister for
// ValidatingAdmissionPolicies.
type ValidatingAdmissionPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() admissionregistrationv1.ValidatingAdmissionPolicyLister
}

type validatingAdmissionPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewValidatingAdmissionPolicyInformer constructs a new informer for ValidatingAdmissionPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewValidatingAdmissionPolicyInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredValidatingAdmissionPolicyInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredValidatingAdmissionPolicyInformer constructs a new informer for ValidatingAdmissionPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredValidatingAdmissionPolicyInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AdmissionregistrationV1().ValidatingAdmissionPolicies().List(context.Background(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AdmissionregistrationV1().ValidatingAdmissionPolicies().Watch(context.Background(), options)
			},
			ListWithContextFunc: func(ctx context.Context, options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AdmissionregistrationV1().ValidatingAdmissionPolicies().List(ctx, options)
			},
			WatchFuncWithContext: func(ctx context.Context, options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AdmissionregistrationV1().ValidatingAdmissionPolicies().Watch(ctx, options)
			},
		},
		&apiadmissionregistrationv1.ValidatingAdmissionPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *validatingAdmissionPolicyInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredValidatingAdmissionPolicyInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *validatingAdmissionPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiadmissionregistrationv1.ValidatingAdmissionPolicy{}, f.defaultInformer)
}

func (f *validatingAdmissionPolicyInformer) Lister() admissionregistrationv1.ValidatingAdmissionPolicyLister {
	return admissionregistrationv1.NewValidatingAdmissionPolicyLister(f.Informer().GetIndexer())
}
