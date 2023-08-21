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

// ServicePlanConfig struct for ServicePlanConfig
type ServicePlanConfig struct {
	StorageSizeType NullableString `json:"storageSizeType,omitempty"`
	MemorySizeType NullableString `json:"memorySizeType,omitempty"`
	Ranges NullableServicePlanConfigRanges `json:"ranges,omitempty"`
}

// NewServicePlanConfig instantiates a new ServicePlanConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicePlanConfig() *ServicePlanConfig {
	this := ServicePlanConfig{}
	return &this
}

// NewServicePlanConfigWithDefaults instantiates a new ServicePlanConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicePlanConfigWithDefaults() *ServicePlanConfig {
	this := ServicePlanConfig{}
	return &this
}

// GetStorageSizeType returns the StorageSizeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServicePlanConfig) GetStorageSizeType() string {
	if o == nil || o.StorageSizeType.Get() == nil {
		var ret string
		return ret
	}
	return *o.StorageSizeType.Get()
}

// GetStorageSizeTypeOk returns a tuple with the StorageSizeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicePlanConfig) GetStorageSizeTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageSizeType.Get(), o.StorageSizeType.IsSet()
}

// HasStorageSizeType returns a boolean if a field has been set.
func (o *ServicePlanConfig) HasStorageSizeType() bool {
	if o != nil && o.StorageSizeType.IsSet() {
		return true
	}

	return false
}

// SetStorageSizeType gets a reference to the given NullableString and assigns it to the StorageSizeType field.
func (o *ServicePlanConfig) SetStorageSizeType(v string) {
	o.StorageSizeType.Set(&v)
}
// SetStorageSizeTypeNil sets the value for StorageSizeType to be an explicit nil
func (o *ServicePlanConfig) SetStorageSizeTypeNil() {
	o.StorageSizeType.Set(nil)
}

// UnsetStorageSizeType ensures that no value is present for StorageSizeType, not even an explicit nil
func (o *ServicePlanConfig) UnsetStorageSizeType() {
	o.StorageSizeType.Unset()
}

// GetMemorySizeType returns the MemorySizeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServicePlanConfig) GetMemorySizeType() string {
	if o == nil || o.MemorySizeType.Get() == nil {
		var ret string
		return ret
	}
	return *o.MemorySizeType.Get()
}

// GetMemorySizeTypeOk returns a tuple with the MemorySizeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicePlanConfig) GetMemorySizeTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MemorySizeType.Get(), o.MemorySizeType.IsSet()
}

// HasMemorySizeType returns a boolean if a field has been set.
func (o *ServicePlanConfig) HasMemorySizeType() bool {
	if o != nil && o.MemorySizeType.IsSet() {
		return true
	}

	return false
}

// SetMemorySizeType gets a reference to the given NullableString and assigns it to the MemorySizeType field.
func (o *ServicePlanConfig) SetMemorySizeType(v string) {
	o.MemorySizeType.Set(&v)
}
// SetMemorySizeTypeNil sets the value for MemorySizeType to be an explicit nil
func (o *ServicePlanConfig) SetMemorySizeTypeNil() {
	o.MemorySizeType.Set(nil)
}

// UnsetMemorySizeType ensures that no value is present for MemorySizeType, not even an explicit nil
func (o *ServicePlanConfig) UnsetMemorySizeType() {
	o.MemorySizeType.Unset()
}

// GetRanges returns the Ranges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServicePlanConfig) GetRanges() ServicePlanConfigRanges {
	if o == nil || o.Ranges.Get() == nil {
		var ret ServicePlanConfigRanges
		return ret
	}
	return *o.Ranges.Get()
}

// GetRangesOk returns a tuple with the Ranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicePlanConfig) GetRangesOk() (*ServicePlanConfigRanges, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Ranges.Get(), o.Ranges.IsSet()
}

// HasRanges returns a boolean if a field has been set.
func (o *ServicePlanConfig) HasRanges() bool {
	if o != nil && o.Ranges.IsSet() {
		return true
	}

	return false
}

// SetRanges gets a reference to the given NullableServicePlanConfigRanges and assigns it to the Ranges field.
func (o *ServicePlanConfig) SetRanges(v ServicePlanConfigRanges) {
	o.Ranges.Set(&v)
}
// SetRangesNil sets the value for Ranges to be an explicit nil
func (o *ServicePlanConfig) SetRangesNil() {
	o.Ranges.Set(nil)
}

// UnsetRanges ensures that no value is present for Ranges, not even an explicit nil
func (o *ServicePlanConfig) UnsetRanges() {
	o.Ranges.Unset()
}

func (o ServicePlanConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StorageSizeType.IsSet() {
		toSerialize["storageSizeType"] = o.StorageSizeType.Get()
	}
	if o.MemorySizeType.IsSet() {
		toSerialize["memorySizeType"] = o.MemorySizeType.Get()
	}
	if o.Ranges.IsSet() {
		toSerialize["ranges"] = o.Ranges.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableServicePlanConfig struct {
	value *ServicePlanConfig
	isSet bool
}

func (v NullableServicePlanConfig) Get() *ServicePlanConfig {
	return v.value
}

func (v *NullableServicePlanConfig) Set(val *ServicePlanConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableServicePlanConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableServicePlanConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicePlanConfig(val *ServicePlanConfig) *NullableServicePlanConfig {
	return &NullableServicePlanConfig{value: val, isSet: true}
}

func (v NullableServicePlanConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicePlanConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


