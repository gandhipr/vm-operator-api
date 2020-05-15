// Copyright (c) 2020 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/vmware-tanzu/vm-operator-api/api/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVirtualMachineServices implements VirtualMachineServiceInterface
type FakeVirtualMachineServices struct {
	Fake *FakeVmoperatorV1alpha1
	ns   string
}

var virtualmachineservicesResource = schema.GroupVersionResource{Group: "vmoperator.vmware.com", Version: "v1alpha1", Resource: "virtualmachineservices"}

var virtualmachineservicesKind = schema.GroupVersionKind{Group: "vmoperator.vmware.com", Version: "v1alpha1", Kind: "VirtualMachineService"}

// Get takes name of the virtualMachineService, and returns the corresponding virtualMachineService object, and an error if there is any.
func (c *FakeVirtualMachineServices) Get(name string, options v1.GetOptions) (result *v1alpha1.VirtualMachineService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualmachineservicesResource, c.ns, name), &v1alpha1.VirtualMachineService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineService), err
}

// List takes label and field selectors, and returns the list of VirtualMachineServices that match those selectors.
func (c *FakeVirtualMachineServices) List(opts v1.ListOptions) (result *v1alpha1.VirtualMachineServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualmachineservicesResource, virtualmachineservicesKind, c.ns, opts), &v1alpha1.VirtualMachineServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VirtualMachineServiceList{ListMeta: obj.(*v1alpha1.VirtualMachineServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.VirtualMachineServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualMachineServices.
func (c *FakeVirtualMachineServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualmachineservicesResource, c.ns, opts))

}

// Create takes the representation of a virtualMachineService and creates it.  Returns the server's representation of the virtualMachineService, and an error, if there is any.
func (c *FakeVirtualMachineServices) Create(virtualMachineService *v1alpha1.VirtualMachineService) (result *v1alpha1.VirtualMachineService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualmachineservicesResource, c.ns, virtualMachineService), &v1alpha1.VirtualMachineService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineService), err
}

// Update takes the representation of a virtualMachineService and updates it. Returns the server's representation of the virtualMachineService, and an error, if there is any.
func (c *FakeVirtualMachineServices) Update(virtualMachineService *v1alpha1.VirtualMachineService) (result *v1alpha1.VirtualMachineService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualmachineservicesResource, c.ns, virtualMachineService), &v1alpha1.VirtualMachineService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualMachineServices) UpdateStatus(virtualMachineService *v1alpha1.VirtualMachineService) (*v1alpha1.VirtualMachineService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualmachineservicesResource, "status", c.ns, virtualMachineService), &v1alpha1.VirtualMachineService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineService), err
}

// Delete takes name of the virtualMachineService and deletes it. Returns an error if one occurs.
func (c *FakeVirtualMachineServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualmachineservicesResource, c.ns, name), &v1alpha1.VirtualMachineService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualMachineServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualmachineservicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.VirtualMachineServiceList{})
	return err
}

// Patch applies the patch and returns the patched virtualMachineService.
func (c *FakeVirtualMachineServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualMachineService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualmachineservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.VirtualMachineService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineService), err
}
