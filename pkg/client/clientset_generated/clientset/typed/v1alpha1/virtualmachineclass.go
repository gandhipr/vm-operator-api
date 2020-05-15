// Copyright (c) 2020 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/vmware-tanzu/vm-operator-api/api/v1alpha1"
	scheme "github.com/vmware-tanzu/vm-operator-api/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VirtualMachineClassesGetter has a method to return a VirtualMachineClassInterface.
// A group's client should implement this interface.
type VirtualMachineClassesGetter interface {
	VirtualMachineClasses() VirtualMachineClassInterface
}

// VirtualMachineClassInterface has methods to work with VirtualMachineClass resources.
type VirtualMachineClassInterface interface {
	Create(*v1alpha1.VirtualMachineClass) (*v1alpha1.VirtualMachineClass, error)
	Update(*v1alpha1.VirtualMachineClass) (*v1alpha1.VirtualMachineClass, error)
	UpdateStatus(*v1alpha1.VirtualMachineClass) (*v1alpha1.VirtualMachineClass, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.VirtualMachineClass, error)
	List(opts v1.ListOptions) (*v1alpha1.VirtualMachineClassList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualMachineClass, err error)
	VirtualMachineClassExpansion
}

// virtualMachineClasses implements VirtualMachineClassInterface
type virtualMachineClasses struct {
	client rest.Interface
}

// newVirtualMachineClasses returns a VirtualMachineClasses
func newVirtualMachineClasses(c *VmoperatorV1alpha1Client) *virtualMachineClasses {
	return &virtualMachineClasses{
		client: c.RESTClient(),
	}
}

// Get takes name of the virtualMachineClass, and returns the corresponding virtualMachineClass object, and an error if there is any.
func (c *virtualMachineClasses) Get(name string, options v1.GetOptions) (result *v1alpha1.VirtualMachineClass, err error) {
	result = &v1alpha1.VirtualMachineClass{}
	err = c.client.Get().
		Resource("virtualmachineclasses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualMachineClasses that match those selectors.
func (c *virtualMachineClasses) List(opts v1.ListOptions) (result *v1alpha1.VirtualMachineClassList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VirtualMachineClassList{}
	err = c.client.Get().
		Resource("virtualmachineclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualMachineClasses.
func (c *virtualMachineClasses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("virtualmachineclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a virtualMachineClass and creates it.  Returns the server's representation of the virtualMachineClass, and an error, if there is any.
func (c *virtualMachineClasses) Create(virtualMachineClass *v1alpha1.VirtualMachineClass) (result *v1alpha1.VirtualMachineClass, err error) {
	result = &v1alpha1.VirtualMachineClass{}
	err = c.client.Post().
		Resource("virtualmachineclasses").
		Body(virtualMachineClass).
		Do().
		Into(result)
	return
}

// Update takes the representation of a virtualMachineClass and updates it. Returns the server's representation of the virtualMachineClass, and an error, if there is any.
func (c *virtualMachineClasses) Update(virtualMachineClass *v1alpha1.VirtualMachineClass) (result *v1alpha1.VirtualMachineClass, err error) {
	result = &v1alpha1.VirtualMachineClass{}
	err = c.client.Put().
		Resource("virtualmachineclasses").
		Name(virtualMachineClass.Name).
		Body(virtualMachineClass).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *virtualMachineClasses) UpdateStatus(virtualMachineClass *v1alpha1.VirtualMachineClass) (result *v1alpha1.VirtualMachineClass, err error) {
	result = &v1alpha1.VirtualMachineClass{}
	err = c.client.Put().
		Resource("virtualmachineclasses").
		Name(virtualMachineClass.Name).
		SubResource("status").
		Body(virtualMachineClass).
		Do().
		Into(result)
	return
}

// Delete takes name of the virtualMachineClass and deletes it. Returns an error if one occurs.
func (c *virtualMachineClasses) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("virtualmachineclasses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualMachineClasses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("virtualmachineclasses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched virtualMachineClass.
func (c *virtualMachineClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualMachineClass, err error) {
	result = &v1alpha1.VirtualMachineClass{}
	err = c.client.Patch(pt).
		Resource("virtualmachineclasses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
