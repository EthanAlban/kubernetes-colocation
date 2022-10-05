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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "node-simulator/apis/node/v1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NodeLister helps list Nodes.
// All objects returned here must be treated as read-only.
type NodeLister interface {
	// List lists all Nodes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Node, err error)
	// Nodes returns an object that can list and get Nodes.
	Nodes(namespace string) NodeNamespaceLister
	NodeListerExpansion
}

// nodeLister implements the NodeLister interface.
type nodeLister struct {
	indexer cache.Indexer
}

// NewNodeLister returns a new NodeLister.
func NewNodeLister(indexer cache.Indexer) NodeLister {
	return &nodeLister{indexer: indexer}
}

// List lists all Nodes in the indexer.
func (s *nodeLister) List(selector labels.Selector) (ret []*v1.Node, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Node))
	})
	return ret, err
}

// Nodes returns an object that can list and get Nodes.
func (s *nodeLister) Nodes(namespace string) NodeNamespaceLister {
	return nodeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NodeNamespaceLister helps list and get Nodes.
// All objects returned here must be treated as read-only.
type NodeNamespaceLister interface {
	// List lists all Nodes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Node, err error)
	// Get retrieves the Node from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Node, error)
	NodeNamespaceListerExpansion
}

// nodeNamespaceLister implements the NodeNamespaceLister
// interface.
type nodeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Nodes in the indexer for a given namespace.
func (s nodeNamespaceLister) List(selector labels.Selector) (ret []*v1.Node, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Node))
	})
	return ret, err
}

// Get retrieves the Node from the indexer for a given namespace and name.
func (s nodeNamespaceLister) Get(name string) (*v1.Node, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("node"), name)
	}
	return obj.(*v1.Node), nil
}
