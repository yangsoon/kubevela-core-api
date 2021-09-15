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
// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AppRolloutLister helps list AppRollouts.
// All objects returned here must be treated as read-only.
type AppRolloutLister interface {
	// List lists all AppRollouts in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.AppRollout, err error)
	// AppRollouts returns an object that can list and get AppRollouts.
	AppRollouts(namespace string) AppRolloutNamespaceLister
	AppRolloutListerExpansion
}

// appRolloutLister implements the AppRolloutLister interface.
type appRolloutLister struct {
	indexer cache.Indexer
}

// NewAppRolloutLister returns a new AppRolloutLister.
func NewAppRolloutLister(indexer cache.Indexer) AppRolloutLister {
	return &appRolloutLister{indexer: indexer}
}

// List lists all AppRollouts in the indexer.
func (s *appRolloutLister) List(selector labels.Selector) (ret []*v1beta1.AppRollout, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.AppRollout))
	})
	return ret, err
}

// AppRollouts returns an object that can list and get AppRollouts.
func (s *appRolloutLister) AppRollouts(namespace string) AppRolloutNamespaceLister {
	return appRolloutNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AppRolloutNamespaceLister helps list and get AppRollouts.
// All objects returned here must be treated as read-only.
type AppRolloutNamespaceLister interface {
	// List lists all AppRollouts in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.AppRollout, err error)
	// Get retrieves the AppRollout from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.AppRollout, error)
	AppRolloutNamespaceListerExpansion
}

// appRolloutNamespaceLister implements the AppRolloutNamespaceLister
// interface.
type appRolloutNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AppRollouts in the indexer for a given namespace.
func (s appRolloutNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.AppRollout, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.AppRollout))
	})
	return ret, err
}

// Get retrieves the AppRollout from the indexer for a given namespace and name.
func (s appRolloutNamespaceLister) Get(name string) (*v1beta1.AppRollout, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("approllout"), name)
	}
	return obj.(*v1beta1.AppRollout), nil
}
