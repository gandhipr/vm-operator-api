// Copyright (c) 2020 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/vmware-tanzu/vm-operator-api/api/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VirtualMachineImageLister helps list VirtualMachineImages.
type VirtualMachineImageLister interface {
	// List lists all VirtualMachineImages in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.VirtualMachineImage, err error)
	// Get retrieves the VirtualMachineImage from the index for a given name.
	Get(name string) (*v1alpha1.VirtualMachineImage, error)
	VirtualMachineImageListerExpansion
}

// virtualMachineImageLister implements the VirtualMachineImageLister interface.
type virtualMachineImageLister struct {
	indexer cache.Indexer
}

// NewVirtualMachineImageLister returns a new VirtualMachineImageLister.
func NewVirtualMachineImageLister(indexer cache.Indexer) VirtualMachineImageLister {
	return &virtualMachineImageLister{indexer: indexer}
}

// List lists all VirtualMachineImages in the indexer.
func (s *virtualMachineImageLister) List(selector labels.Selector) (ret []*v1alpha1.VirtualMachineImage, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VirtualMachineImage))
	})
	return ret, err
}

// Get retrieves the VirtualMachineImage from the index for a given name.
func (s *virtualMachineImageLister) Get(name string) (*v1alpha1.VirtualMachineImage, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("virtualmachineimage"), name)
	}
	return obj.(*v1alpha1.VirtualMachineImage), nil
}
