/*
 * Morpheus API
 *
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * API version: 6.2.1
 * Contact: dev@morpheusdata.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ServicePlanOptions Map of custom options depending on selected service plan.
type ServicePlanOptions struct {
	// Core Count
	MaxCores *int64 `json:"maxCores,omitempty"`
	// Cores Per Socket
	CoresPerSocket *int64 `json:"coresPerSocket,omitempty"`
	// Memory in bytes For backwards compatability, values less than 1048576 are treated as being in MB and will be converted to bytes
	MaxMemory *int64 `json:"maxMemory,omitempty"`
}

// NewServicePlanOptions instantiates a new ServicePlanOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicePlanOptions() *ServicePlanOptions {
	this := ServicePlanOptions{}
	return &this
}

// NewServicePlanOptionsWithDefaults instantiates a new ServicePlanOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicePlanOptionsWithDefaults() *ServicePlanOptions {
	this := ServicePlanOptions{}
	return &this
}

// GetMaxCores returns the MaxCores field value if set, zero value otherwise.
func (o *ServicePlanOptions) GetMaxCores() int64 {
	if o == nil || o.MaxCores == nil {
		var ret int64
		return ret
	}
	return *o.MaxCores
}

// GetMaxCoresOk returns a tuple with the MaxCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlanOptions) GetMaxCoresOk() (*int64, bool) {
	if o == nil || o.MaxCores == nil {
		return nil, false
	}
	return o.MaxCores, true
}

// HasMaxCores returns a boolean if a field has been set.
func (o *ServicePlanOptions) HasMaxCores() bool {
	if o != nil && o.MaxCores != nil {
		return true
	}

	return false
}

// SetMaxCores gets a reference to the given int64 and assigns it to the MaxCores field.
func (o *ServicePlanOptions) SetMaxCores(v int64) {
	o.MaxCores = &v
}

// GetCoresPerSocket returns the CoresPerSocket field value if set, zero value otherwise.
func (o *ServicePlanOptions) GetCoresPerSocket() int64 {
	if o == nil || o.CoresPerSocket == nil {
		var ret int64
		return ret
	}
	return *o.CoresPerSocket
}

// GetCoresPerSocketOk returns a tuple with the CoresPerSocket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlanOptions) GetCoresPerSocketOk() (*int64, bool) {
	if o == nil || o.CoresPerSocket == nil {
		return nil, false
	}
	return o.CoresPerSocket, true
}

// HasCoresPerSocket returns a boolean if a field has been set.
func (o *ServicePlanOptions) HasCoresPerSocket() bool {
	if o != nil && o.CoresPerSocket != nil {
		return true
	}

	return false
}

// SetCoresPerSocket gets a reference to the given int64 and assigns it to the CoresPerSocket field.
func (o *ServicePlanOptions) SetCoresPerSocket(v int64) {
	o.CoresPerSocket = &v
}

// GetMaxMemory returns the MaxMemory field value if set, zero value otherwise.
func (o *ServicePlanOptions) GetMaxMemory() int64 {
	if o == nil || o.MaxMemory == nil {
		var ret int64
		return ret
	}
	return *o.MaxMemory
}

// GetMaxMemoryOk returns a tuple with the MaxMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlanOptions) GetMaxMemoryOk() (*int64, bool) {
	if o == nil || o.MaxMemory == nil {
		return nil, false
	}
	return o.MaxMemory, true
}

// HasMaxMemory returns a boolean if a field has been set.
func (o *ServicePlanOptions) HasMaxMemory() bool {
	if o != nil && o.MaxMemory != nil {
		return true
	}

	return false
}

// SetMaxMemory gets a reference to the given int64 and assigns it to the MaxMemory field.
func (o *ServicePlanOptions) SetMaxMemory(v int64) {
	o.MaxMemory = &v
}

func (o ServicePlanOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxCores != nil {
		toSerialize["maxCores"] = o.MaxCores
	}
	if o.CoresPerSocket != nil {
		toSerialize["coresPerSocket"] = o.CoresPerSocket
	}
	if o.MaxMemory != nil {
		toSerialize["maxMemory"] = o.MaxMemory
	}
	return json.Marshal(toSerialize)
}

type NullableServicePlanOptions struct {
	value *ServicePlanOptions
	isSet bool
}

func (v NullableServicePlanOptions) Get() *ServicePlanOptions {
	return v.value
}

func (v *NullableServicePlanOptions) Set(val *ServicePlanOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableServicePlanOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableServicePlanOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicePlanOptions(val *ServicePlanOptions) *NullableServicePlanOptions {
	return &NullableServicePlanOptions{value: val, isSet: true}
}

func (v NullableServicePlanOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicePlanOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


