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

// GuidanceSettings struct for GuidanceSettings
type GuidanceSettings struct {
	// Power Shutdown Average CPU (%). Lower limit for average CPU usage
	CpuAvgCutoffPower NullableInt32 `json:"cpuAvgCutoffPower,omitempty"`
	// Power Shutdown Maximum CPU (%). Lower limit for peak CPU usage
	CpuMaxCutoffPower NullableInt32 `json:"cpuMaxCutoffPower,omitempty"`
	// Power Shutdown Network threshold (bytes). Lower limit for average network bandwidth
	NetworkCutoffPower NullableInt32 `json:"networkCutoffPower,omitempty"`
	// CPU Up-size Average CPU (%). Upper limit for CPU usage
	CpuUpAvgStandardCutoffRightSize NullableInt32 `json:"cpuUpAvgStandardCutoffRightSize,omitempty"`
	// CPU Up-size Maximum CPU (%). Upper limit for peak CPU usage
	CpuUpMaxStandardCutoffRightSize NullableInt32 `json:"cpuUpMaxStandardCutoffRightSize,omitempty"`
	// Memory Up-size Minimum Free Memory (%). Lower limit for average free memory usage
	MemoryUpAvgStandardCutoffRightSize NullableInt32 `json:"memoryUpAvgStandardCutoffRightSize,omitempty"`
	// Memory Down-size Maximum Free Memory (%). Upper limit for average free memory
	MemoryDownAvgStandardCutoffRightSize NullableInt32 `json:"memoryDownAvgStandardCutoffRightSize,omitempty"`
	// Memory Down-size Maximum Free Memory (%). Upper limit for peak memory usage
	MemoryDownMaxStandardCutoffRightSize NullableInt32 `json:"memoryDownMaxStandardCutoffRightSize,omitempty"`
}

// NewGuidanceSettings instantiates a new GuidanceSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuidanceSettings() *GuidanceSettings {
	this := GuidanceSettings{}
	return &this
}

// NewGuidanceSettingsWithDefaults instantiates a new GuidanceSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuidanceSettingsWithDefaults() *GuidanceSettings {
	this := GuidanceSettings{}
	return &this
}

// GetCpuAvgCutoffPower returns the CpuAvgCutoffPower field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidanceSettings) GetCpuAvgCutoffPower() int32 {
	if o == nil || o.CpuAvgCutoffPower.Get() == nil {
		var ret int32
		return ret
	}
	return *o.CpuAvgCutoffPower.Get()
}

// GetCpuAvgCutoffPowerOk returns a tuple with the CpuAvgCutoffPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidanceSettings) GetCpuAvgCutoffPowerOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CpuAvgCutoffPower.Get(), o.CpuAvgCutoffPower.IsSet()
}

// HasCpuAvgCutoffPower returns a boolean if a field has been set.
func (o *GuidanceSettings) HasCpuAvgCutoffPower() bool {
	if o != nil && o.CpuAvgCutoffPower.IsSet() {
		return true
	}

	return false
}

// SetCpuAvgCutoffPower gets a reference to the given NullableInt32 and assigns it to the CpuAvgCutoffPower field.
func (o *GuidanceSettings) SetCpuAvgCutoffPower(v int32) {
	o.CpuAvgCutoffPower.Set(&v)
}
// SetCpuAvgCutoffPowerNil sets the value for CpuAvgCutoffPower to be an explicit nil
func (o *GuidanceSettings) SetCpuAvgCutoffPowerNil() {
	o.CpuAvgCutoffPower.Set(nil)
}

// UnsetCpuAvgCutoffPower ensures that no value is present for CpuAvgCutoffPower, not even an explicit nil
func (o *GuidanceSettings) UnsetCpuAvgCutoffPower() {
	o.CpuAvgCutoffPower.Unset()
}

// GetCpuMaxCutoffPower returns the CpuMaxCutoffPower field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidanceSettings) GetCpuMaxCutoffPower() int32 {
	if o == nil || o.CpuMaxCutoffPower.Get() == nil {
		var ret int32
		return ret
	}
	return *o.CpuMaxCutoffPower.Get()
}

// GetCpuMaxCutoffPowerOk returns a tuple with the CpuMaxCutoffPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidanceSettings) GetCpuMaxCutoffPowerOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CpuMaxCutoffPower.Get(), o.CpuMaxCutoffPower.IsSet()
}

// HasCpuMaxCutoffPower returns a boolean if a field has been set.
func (o *GuidanceSettings) HasCpuMaxCutoffPower() bool {
	if o != nil && o.CpuMaxCutoffPower.IsSet() {
		return true
	}

	return false
}

// SetCpuMaxCutoffPower gets a reference to the given NullableInt32 and assigns it to the CpuMaxCutoffPower field.
func (o *GuidanceSettings) SetCpuMaxCutoffPower(v int32) {
	o.CpuMaxCutoffPower.Set(&v)
}
// SetCpuMaxCutoffPowerNil sets the value for CpuMaxCutoffPower to be an explicit nil
func (o *GuidanceSettings) SetCpuMaxCutoffPowerNil() {
	o.CpuMaxCutoffPower.Set(nil)
}

// UnsetCpuMaxCutoffPower ensures that no value is present for CpuMaxCutoffPower, not even an explicit nil
func (o *GuidanceSettings) UnsetCpuMaxCutoffPower() {
	o.CpuMaxCutoffPower.Unset()
}

// GetNetworkCutoffPower returns the NetworkCutoffPower field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidanceSettings) GetNetworkCutoffPower() int32 {
	if o == nil || o.NetworkCutoffPower.Get() == nil {
		var ret int32
		return ret
	}
	return *o.NetworkCutoffPower.Get()
}

// GetNetworkCutoffPowerOk returns a tuple with the NetworkCutoffPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidanceSettings) GetNetworkCutoffPowerOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NetworkCutoffPower.Get(), o.NetworkCutoffPower.IsSet()
}

// HasNetworkCutoffPower returns a boolean if a field has been set.
func (o *GuidanceSettings) HasNetworkCutoffPower() bool {
	if o != nil && o.NetworkCutoffPower.IsSet() {
		return true
	}

	return false
}

// SetNetworkCutoffPower gets a reference to the given NullableInt32 and assigns it to the NetworkCutoffPower field.
func (o *GuidanceSettings) SetNetworkCutoffPower(v int32) {
	o.NetworkCutoffPower.Set(&v)
}
// SetNetworkCutoffPowerNil sets the value for NetworkCutoffPower to be an explicit nil
func (o *GuidanceSettings) SetNetworkCutoffPowerNil() {
	o.NetworkCutoffPower.Set(nil)
}

// UnsetNetworkCutoffPower ensures that no value is present for NetworkCutoffPower, not even an explicit nil
func (o *GuidanceSettings) UnsetNetworkCutoffPower() {
	o.NetworkCutoffPower.Unset()
}

// GetCpuUpAvgStandardCutoffRightSize returns the CpuUpAvgStandardCutoffRightSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidanceSettings) GetCpuUpAvgStandardCutoffRightSize() int32 {
	if o == nil || o.CpuUpAvgStandardCutoffRightSize.Get() == nil {
		var ret int32
		return ret
	}
	return *o.CpuUpAvgStandardCutoffRightSize.Get()
}

// GetCpuUpAvgStandardCutoffRightSizeOk returns a tuple with the CpuUpAvgStandardCutoffRightSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidanceSettings) GetCpuUpAvgStandardCutoffRightSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CpuUpAvgStandardCutoffRightSize.Get(), o.CpuUpAvgStandardCutoffRightSize.IsSet()
}

// HasCpuUpAvgStandardCutoffRightSize returns a boolean if a field has been set.
func (o *GuidanceSettings) HasCpuUpAvgStandardCutoffRightSize() bool {
	if o != nil && o.CpuUpAvgStandardCutoffRightSize.IsSet() {
		return true
	}

	return false
}

// SetCpuUpAvgStandardCutoffRightSize gets a reference to the given NullableInt32 and assigns it to the CpuUpAvgStandardCutoffRightSize field.
func (o *GuidanceSettings) SetCpuUpAvgStandardCutoffRightSize(v int32) {
	o.CpuUpAvgStandardCutoffRightSize.Set(&v)
}
// SetCpuUpAvgStandardCutoffRightSizeNil sets the value for CpuUpAvgStandardCutoffRightSize to be an explicit nil
func (o *GuidanceSettings) SetCpuUpAvgStandardCutoffRightSizeNil() {
	o.CpuUpAvgStandardCutoffRightSize.Set(nil)
}

// UnsetCpuUpAvgStandardCutoffRightSize ensures that no value is present for CpuUpAvgStandardCutoffRightSize, not even an explicit nil
func (o *GuidanceSettings) UnsetCpuUpAvgStandardCutoffRightSize() {
	o.CpuUpAvgStandardCutoffRightSize.Unset()
}

// GetCpuUpMaxStandardCutoffRightSize returns the CpuUpMaxStandardCutoffRightSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidanceSettings) GetCpuUpMaxStandardCutoffRightSize() int32 {
	if o == nil || o.CpuUpMaxStandardCutoffRightSize.Get() == nil {
		var ret int32
		return ret
	}
	return *o.CpuUpMaxStandardCutoffRightSize.Get()
}

// GetCpuUpMaxStandardCutoffRightSizeOk returns a tuple with the CpuUpMaxStandardCutoffRightSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidanceSettings) GetCpuUpMaxStandardCutoffRightSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CpuUpMaxStandardCutoffRightSize.Get(), o.CpuUpMaxStandardCutoffRightSize.IsSet()
}

// HasCpuUpMaxStandardCutoffRightSize returns a boolean if a field has been set.
func (o *GuidanceSettings) HasCpuUpMaxStandardCutoffRightSize() bool {
	if o != nil && o.CpuUpMaxStandardCutoffRightSize.IsSet() {
		return true
	}

	return false
}

// SetCpuUpMaxStandardCutoffRightSize gets a reference to the given NullableInt32 and assigns it to the CpuUpMaxStandardCutoffRightSize field.
func (o *GuidanceSettings) SetCpuUpMaxStandardCutoffRightSize(v int32) {
	o.CpuUpMaxStandardCutoffRightSize.Set(&v)
}
// SetCpuUpMaxStandardCutoffRightSizeNil sets the value for CpuUpMaxStandardCutoffRightSize to be an explicit nil
func (o *GuidanceSettings) SetCpuUpMaxStandardCutoffRightSizeNil() {
	o.CpuUpMaxStandardCutoffRightSize.Set(nil)
}

// UnsetCpuUpMaxStandardCutoffRightSize ensures that no value is present for CpuUpMaxStandardCutoffRightSize, not even an explicit nil
func (o *GuidanceSettings) UnsetCpuUpMaxStandardCutoffRightSize() {
	o.CpuUpMaxStandardCutoffRightSize.Unset()
}

// GetMemoryUpAvgStandardCutoffRightSize returns the MemoryUpAvgStandardCutoffRightSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidanceSettings) GetMemoryUpAvgStandardCutoffRightSize() int32 {
	if o == nil || o.MemoryUpAvgStandardCutoffRightSize.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MemoryUpAvgStandardCutoffRightSize.Get()
}

// GetMemoryUpAvgStandardCutoffRightSizeOk returns a tuple with the MemoryUpAvgStandardCutoffRightSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidanceSettings) GetMemoryUpAvgStandardCutoffRightSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MemoryUpAvgStandardCutoffRightSize.Get(), o.MemoryUpAvgStandardCutoffRightSize.IsSet()
}

// HasMemoryUpAvgStandardCutoffRightSize returns a boolean if a field has been set.
func (o *GuidanceSettings) HasMemoryUpAvgStandardCutoffRightSize() bool {
	if o != nil && o.MemoryUpAvgStandardCutoffRightSize.IsSet() {
		return true
	}

	return false
}

// SetMemoryUpAvgStandardCutoffRightSize gets a reference to the given NullableInt32 and assigns it to the MemoryUpAvgStandardCutoffRightSize field.
func (o *GuidanceSettings) SetMemoryUpAvgStandardCutoffRightSize(v int32) {
	o.MemoryUpAvgStandardCutoffRightSize.Set(&v)
}
// SetMemoryUpAvgStandardCutoffRightSizeNil sets the value for MemoryUpAvgStandardCutoffRightSize to be an explicit nil
func (o *GuidanceSettings) SetMemoryUpAvgStandardCutoffRightSizeNil() {
	o.MemoryUpAvgStandardCutoffRightSize.Set(nil)
}

// UnsetMemoryUpAvgStandardCutoffRightSize ensures that no value is present for MemoryUpAvgStandardCutoffRightSize, not even an explicit nil
func (o *GuidanceSettings) UnsetMemoryUpAvgStandardCutoffRightSize() {
	o.MemoryUpAvgStandardCutoffRightSize.Unset()
}

// GetMemoryDownAvgStandardCutoffRightSize returns the MemoryDownAvgStandardCutoffRightSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidanceSettings) GetMemoryDownAvgStandardCutoffRightSize() int32 {
	if o == nil || o.MemoryDownAvgStandardCutoffRightSize.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MemoryDownAvgStandardCutoffRightSize.Get()
}

// GetMemoryDownAvgStandardCutoffRightSizeOk returns a tuple with the MemoryDownAvgStandardCutoffRightSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidanceSettings) GetMemoryDownAvgStandardCutoffRightSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MemoryDownAvgStandardCutoffRightSize.Get(), o.MemoryDownAvgStandardCutoffRightSize.IsSet()
}

// HasMemoryDownAvgStandardCutoffRightSize returns a boolean if a field has been set.
func (o *GuidanceSettings) HasMemoryDownAvgStandardCutoffRightSize() bool {
	if o != nil && o.MemoryDownAvgStandardCutoffRightSize.IsSet() {
		return true
	}

	return false
}

// SetMemoryDownAvgStandardCutoffRightSize gets a reference to the given NullableInt32 and assigns it to the MemoryDownAvgStandardCutoffRightSize field.
func (o *GuidanceSettings) SetMemoryDownAvgStandardCutoffRightSize(v int32) {
	o.MemoryDownAvgStandardCutoffRightSize.Set(&v)
}
// SetMemoryDownAvgStandardCutoffRightSizeNil sets the value for MemoryDownAvgStandardCutoffRightSize to be an explicit nil
func (o *GuidanceSettings) SetMemoryDownAvgStandardCutoffRightSizeNil() {
	o.MemoryDownAvgStandardCutoffRightSize.Set(nil)
}

// UnsetMemoryDownAvgStandardCutoffRightSize ensures that no value is present for MemoryDownAvgStandardCutoffRightSize, not even an explicit nil
func (o *GuidanceSettings) UnsetMemoryDownAvgStandardCutoffRightSize() {
	o.MemoryDownAvgStandardCutoffRightSize.Unset()
}

// GetMemoryDownMaxStandardCutoffRightSize returns the MemoryDownMaxStandardCutoffRightSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidanceSettings) GetMemoryDownMaxStandardCutoffRightSize() int32 {
	if o == nil || o.MemoryDownMaxStandardCutoffRightSize.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MemoryDownMaxStandardCutoffRightSize.Get()
}

// GetMemoryDownMaxStandardCutoffRightSizeOk returns a tuple with the MemoryDownMaxStandardCutoffRightSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidanceSettings) GetMemoryDownMaxStandardCutoffRightSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MemoryDownMaxStandardCutoffRightSize.Get(), o.MemoryDownMaxStandardCutoffRightSize.IsSet()
}

// HasMemoryDownMaxStandardCutoffRightSize returns a boolean if a field has been set.
func (o *GuidanceSettings) HasMemoryDownMaxStandardCutoffRightSize() bool {
	if o != nil && o.MemoryDownMaxStandardCutoffRightSize.IsSet() {
		return true
	}

	return false
}

// SetMemoryDownMaxStandardCutoffRightSize gets a reference to the given NullableInt32 and assigns it to the MemoryDownMaxStandardCutoffRightSize field.
func (o *GuidanceSettings) SetMemoryDownMaxStandardCutoffRightSize(v int32) {
	o.MemoryDownMaxStandardCutoffRightSize.Set(&v)
}
// SetMemoryDownMaxStandardCutoffRightSizeNil sets the value for MemoryDownMaxStandardCutoffRightSize to be an explicit nil
func (o *GuidanceSettings) SetMemoryDownMaxStandardCutoffRightSizeNil() {
	o.MemoryDownMaxStandardCutoffRightSize.Set(nil)
}

// UnsetMemoryDownMaxStandardCutoffRightSize ensures that no value is present for MemoryDownMaxStandardCutoffRightSize, not even an explicit nil
func (o *GuidanceSettings) UnsetMemoryDownMaxStandardCutoffRightSize() {
	o.MemoryDownMaxStandardCutoffRightSize.Unset()
}

func (o GuidanceSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CpuAvgCutoffPower.IsSet() {
		toSerialize["cpuAvgCutoffPower"] = o.CpuAvgCutoffPower.Get()
	}
	if o.CpuMaxCutoffPower.IsSet() {
		toSerialize["cpuMaxCutoffPower"] = o.CpuMaxCutoffPower.Get()
	}
	if o.NetworkCutoffPower.IsSet() {
		toSerialize["networkCutoffPower"] = o.NetworkCutoffPower.Get()
	}
	if o.CpuUpAvgStandardCutoffRightSize.IsSet() {
		toSerialize["cpuUpAvgStandardCutoffRightSize"] = o.CpuUpAvgStandardCutoffRightSize.Get()
	}
	if o.CpuUpMaxStandardCutoffRightSize.IsSet() {
		toSerialize["cpuUpMaxStandardCutoffRightSize"] = o.CpuUpMaxStandardCutoffRightSize.Get()
	}
	if o.MemoryUpAvgStandardCutoffRightSize.IsSet() {
		toSerialize["memoryUpAvgStandardCutoffRightSize"] = o.MemoryUpAvgStandardCutoffRightSize.Get()
	}
	if o.MemoryDownAvgStandardCutoffRightSize.IsSet() {
		toSerialize["memoryDownAvgStandardCutoffRightSize"] = o.MemoryDownAvgStandardCutoffRightSize.Get()
	}
	if o.MemoryDownMaxStandardCutoffRightSize.IsSet() {
		toSerialize["memoryDownMaxStandardCutoffRightSize"] = o.MemoryDownMaxStandardCutoffRightSize.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableGuidanceSettings struct {
	value *GuidanceSettings
	isSet bool
}

func (v NullableGuidanceSettings) Get() *GuidanceSettings {
	return v.value
}

func (v *NullableGuidanceSettings) Set(val *GuidanceSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableGuidanceSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableGuidanceSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuidanceSettings(val *GuidanceSettings) *NullableGuidanceSettings {
	return &NullableGuidanceSettings{value: val, isSet: true}
}

func (v NullableGuidanceSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuidanceSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

