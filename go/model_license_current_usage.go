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

// LicenseCurrentUsage struct for LicenseCurrentUsage
type LicenseCurrentUsage struct {
	Memory *int64 `json:"memory,omitempty"`
	Storage *int64 `json:"storage,omitempty"`
	Workloads *int64 `json:"workloads,omitempty"`
}

// NewLicenseCurrentUsage instantiates a new LicenseCurrentUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseCurrentUsage() *LicenseCurrentUsage {
	this := LicenseCurrentUsage{}
	return &this
}

// NewLicenseCurrentUsageWithDefaults instantiates a new LicenseCurrentUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseCurrentUsageWithDefaults() *LicenseCurrentUsage {
	this := LicenseCurrentUsage{}
	return &this
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *LicenseCurrentUsage) GetMemory() int64 {
	if o == nil || o.Memory == nil {
		var ret int64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseCurrentUsage) GetMemoryOk() (*int64, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *LicenseCurrentUsage) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int64 and assigns it to the Memory field.
func (o *LicenseCurrentUsage) SetMemory(v int64) {
	o.Memory = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *LicenseCurrentUsage) GetStorage() int64 {
	if o == nil || o.Storage == nil {
		var ret int64
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseCurrentUsage) GetStorageOk() (*int64, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *LicenseCurrentUsage) HasStorage() bool {
	if o != nil && o.Storage != nil {
		return true
	}

	return false
}

// SetStorage gets a reference to the given int64 and assigns it to the Storage field.
func (o *LicenseCurrentUsage) SetStorage(v int64) {
	o.Storage = &v
}

// GetWorkloads returns the Workloads field value if set, zero value otherwise.
func (o *LicenseCurrentUsage) GetWorkloads() int64 {
	if o == nil || o.Workloads == nil {
		var ret int64
		return ret
	}
	return *o.Workloads
}

// GetWorkloadsOk returns a tuple with the Workloads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseCurrentUsage) GetWorkloadsOk() (*int64, bool) {
	if o == nil || o.Workloads == nil {
		return nil, false
	}
	return o.Workloads, true
}

// HasWorkloads returns a boolean if a field has been set.
func (o *LicenseCurrentUsage) HasWorkloads() bool {
	if o != nil && o.Workloads != nil {
		return true
	}

	return false
}

// SetWorkloads gets a reference to the given int64 and assigns it to the Workloads field.
func (o *LicenseCurrentUsage) SetWorkloads(v int64) {
	o.Workloads = &v
}

func (o LicenseCurrentUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	if o.Workloads != nil {
		toSerialize["workloads"] = o.Workloads
	}
	return json.Marshal(toSerialize)
}

type NullableLicenseCurrentUsage struct {
	value *LicenseCurrentUsage
	isSet bool
}

func (v NullableLicenseCurrentUsage) Get() *LicenseCurrentUsage {
	return v.value
}

func (v *NullableLicenseCurrentUsage) Set(val *LicenseCurrentUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseCurrentUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseCurrentUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseCurrentUsage(val *LicenseCurrentUsage) *NullableLicenseCurrentUsage {
	return &NullableLicenseCurrentUsage{value: val, isSet: true}
}

func (v NullableLicenseCurrentUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseCurrentUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


