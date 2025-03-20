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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	resourcev1beta1 "k8s.io/api/resource/v1beta1"
	v1 "k8s.io/client-go/applyconfigurations/core/v1"
)

// BasicDeviceApplyConfiguration represents a declarative configuration of the BasicDevice type for use
// with apply.
type BasicDeviceApplyConfiguration struct {
	Attributes      map[resourcev1beta1.QualifiedName]DeviceAttributeApplyConfiguration `json:"attributes,omitempty"`
	Capacity        map[resourcev1beta1.QualifiedName]DeviceCapacityApplyConfiguration  `json:"capacity,omitempty"`
	ConsumesCounter []DeviceCounterConsumptionApplyConfiguration                        `json:"consumesCounter,omitempty"`
	NodeName        *string                                                             `json:"nodeName,omitempty"`
	NodeSelector    *v1.NodeSelectorApplyConfiguration                                  `json:"nodeSelector,omitempty"`
	AllNodes        *bool                                                               `json:"allNodes,omitempty"`
	Taints          []DeviceTaintApplyConfiguration                                     `json:"taints,omitempty"`
}

// BasicDeviceApplyConfiguration constructs a declarative configuration of the BasicDevice type for use with
// apply.
func BasicDevice() *BasicDeviceApplyConfiguration {
	return &BasicDeviceApplyConfiguration{}
}

// WithAttributes puts the entries into the Attributes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Attributes field,
// overwriting an existing map entries in Attributes field with the same key.
func (b *BasicDeviceApplyConfiguration) WithAttributes(entries map[resourcev1beta1.QualifiedName]DeviceAttributeApplyConfiguration) *BasicDeviceApplyConfiguration {
	if b.Attributes == nil && len(entries) > 0 {
		b.Attributes = make(map[resourcev1beta1.QualifiedName]DeviceAttributeApplyConfiguration, len(entries))
	}
	for k, v := range entries {
		b.Attributes[k] = v
	}
	return b
}

// WithCapacity puts the entries into the Capacity field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Capacity field,
// overwriting an existing map entries in Capacity field with the same key.
func (b *BasicDeviceApplyConfiguration) WithCapacity(entries map[resourcev1beta1.QualifiedName]DeviceCapacityApplyConfiguration) *BasicDeviceApplyConfiguration {
	if b.Capacity == nil && len(entries) > 0 {
		b.Capacity = make(map[resourcev1beta1.QualifiedName]DeviceCapacityApplyConfiguration, len(entries))
	}
	for k, v := range entries {
		b.Capacity[k] = v
	}
	return b
}

// WithConsumesCounter adds the given value to the ConsumesCounter field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ConsumesCounter field.
func (b *BasicDeviceApplyConfiguration) WithConsumesCounter(values ...*DeviceCounterConsumptionApplyConfiguration) *BasicDeviceApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConsumesCounter")
		}
		b.ConsumesCounter = append(b.ConsumesCounter, *values[i])
	}
	return b
}

// WithNodeName sets the NodeName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeName field is set to the value of the last call.
func (b *BasicDeviceApplyConfiguration) WithNodeName(value string) *BasicDeviceApplyConfiguration {
	b.NodeName = &value
	return b
}

// WithNodeSelector sets the NodeSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeSelector field is set to the value of the last call.
func (b *BasicDeviceApplyConfiguration) WithNodeSelector(value *v1.NodeSelectorApplyConfiguration) *BasicDeviceApplyConfiguration {
	b.NodeSelector = value
	return b
}

// WithAllNodes sets the AllNodes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AllNodes field is set to the value of the last call.
func (b *BasicDeviceApplyConfiguration) WithAllNodes(value bool) *BasicDeviceApplyConfiguration {
	b.AllNodes = &value
	return b
}

// WithTaints adds the given value to the Taints field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Taints field.
func (b *BasicDeviceApplyConfiguration) WithTaints(values ...*DeviceTaintApplyConfiguration) *BasicDeviceApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTaints")
		}
		b.Taints = append(b.Taints, *values[i])
	}
	return b
}
