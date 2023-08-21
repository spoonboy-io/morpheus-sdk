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

// ApiInstancesIdThresholdInstanceThreshold struct for ApiInstancesIdThresholdInstanceThreshold
type ApiInstancesIdThresholdInstanceThreshold struct {
	// Auto Upscale
	AutoUp *bool `json:"autoUp,omitempty"`
	// Auto Downscale
	AutoDown *bool `json:"autoDown,omitempty"`
	// The minimum number of nodes to scale down to
	MinCount *int32 `json:"minCount,omitempty"`
	// The maximum number of nodes to scale up to
	MaxCount *int32 `json:"maxCount,omitempty"`
	// Enable CPU Threshold
	CpuEnabled *bool `json:"cpuEnabled,omitempty"`
	// Min CPU (%)
	MinCpu *float32 `json:"minCpu,omitempty"`
	// Max CPU (%)
	MaxCpu *float32 `json:"maxCpu,omitempty"`
	// Enable Memory Threshold
	MemoryEnabled *bool `json:"memoryEnabled,omitempty"`
	// Min Memory (%)
	MinMemory *float32 `json:"minMemory,omitempty"`
	// Max Memory (%)
	MaxMemory *float32 `json:"maxMemory,omitempty"`
	// Enable Disk Threshold
	DiskEnabled *bool `json:"diskEnabled,omitempty"`
	// Min Disk (%)
	MinDisk *float32 `json:"minDisk,omitempty"`
	// Max Disk (%)
	MaxDisk *float32 `json:"maxDisk,omitempty"`
}

// NewApiInstancesIdThresholdInstanceThreshold instantiates a new ApiInstancesIdThresholdInstanceThreshold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiInstancesIdThresholdInstanceThreshold() *ApiInstancesIdThresholdInstanceThreshold {
	this := ApiInstancesIdThresholdInstanceThreshold{}
	var autoUp bool = false
	this.AutoUp = &autoUp
	var autoDown bool = false
	this.AutoDown = &autoDown
	var cpuEnabled bool = false
	this.CpuEnabled = &cpuEnabled
	var minCpu float32 = 0
	this.MinCpu = &minCpu
	var maxCpu float32 = 0
	this.MaxCpu = &maxCpu
	var memoryEnabled bool = false
	this.MemoryEnabled = &memoryEnabled
	var minMemory float32 = 0
	this.MinMemory = &minMemory
	var maxMemory float32 = 0
	this.MaxMemory = &maxMemory
	var diskEnabled bool = false
	this.DiskEnabled = &diskEnabled
	var minDisk float32 = 0
	this.MinDisk = &minDisk
	var maxDisk float32 = 0
	this.MaxDisk = &maxDisk
	return &this
}

// NewApiInstancesIdThresholdInstanceThresholdWithDefaults instantiates a new ApiInstancesIdThresholdInstanceThreshold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiInstancesIdThresholdInstanceThresholdWithDefaults() *ApiInstancesIdThresholdInstanceThreshold {
	this := ApiInstancesIdThresholdInstanceThreshold{}
	var autoUp bool = false
	this.AutoUp = &autoUp
	var autoDown bool = false
	this.AutoDown = &autoDown
	var cpuEnabled bool = false
	this.CpuEnabled = &cpuEnabled
	var minCpu float32 = 0
	this.MinCpu = &minCpu
	var maxCpu float32 = 0
	this.MaxCpu = &maxCpu
	var memoryEnabled bool = false
	this.MemoryEnabled = &memoryEnabled
	var minMemory float32 = 0
	this.MinMemory = &minMemory
	var maxMemory float32 = 0
	this.MaxMemory = &maxMemory
	var diskEnabled bool = false
	this.DiskEnabled = &diskEnabled
	var minDisk float32 = 0
	this.MinDisk = &minDisk
	var maxDisk float32 = 0
	this.MaxDisk = &maxDisk
	return &this
}

// GetAutoUp returns the AutoUp field value if set, zero value otherwise.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetAutoUp() bool {
	if o == nil || o.AutoUp == nil {
		var ret bool
		return ret
	}
	return *o.AutoUp
}

// GetAutoUpOk returns a tuple with the AutoUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetAutoUpOk() (*bool, bool) {
	if o == nil || o.AutoUp == nil {
		return nil, false
	}
	return o.AutoUp, true
}

// HasAutoUp returns a boolean if a field has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) HasAutoUp() bool {
	if o != nil && o.AutoUp != nil {
		return true
	}

	return false
}

// SetAutoUp gets a reference to the given bool and assigns it to the AutoUp field.
func (o *ApiInstancesIdThresholdInstanceThreshold) SetAutoUp(v bool) {
	o.AutoUp = &v
}

// GetAutoDown returns the AutoDown field value if set, zero value otherwise.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetAutoDown() bool {
	if o == nil || o.AutoDown == nil {
		var ret bool
		return ret
	}
	return *o.AutoDown
}

// GetAutoDownOk returns a tuple with the AutoDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetAutoDownOk() (*bool, bool) {
	if o == nil || o.AutoDown == nil {
		return nil, false
	}
	return o.AutoDown, true
}

// HasAutoDown returns a boolean if a field has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) HasAutoDown() bool {
	if o != nil && o.AutoDown != nil {
		return true
	}

	return false
}

// SetAutoDown gets a reference to the given bool and assigns it to the AutoDown field.
func (o *ApiInstancesIdThresholdInstanceThreshold) SetAutoDown(v bool) {
	o.AutoDown = &v
}

// GetMinCount returns the MinCount field value if set, zero value otherwise.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetMinCount() int32 {
	if o == nil || o.MinCount == nil {
		var ret int32
		return ret
	}
	return *o.MinCount
}

// GetMinCountOk returns a tuple with the MinCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetMinCountOk() (*int32, bool) {
	if o == nil || o.MinCount == nil {
		return nil, false
	}
	return o.MinCount, true
}

// HasMinCount returns a boolean if a field has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) HasMinCount() bool {
	if o != nil && o.MinCount != nil {
		return true
	}

	return false
}

// SetMinCount gets a reference to the given int32 and assigns it to the MinCount field.
func (o *ApiInstancesIdThresholdInstanceThreshold) SetMinCount(v int32) {
	o.MinCount = &v
}

// GetMaxCount returns the MaxCount field value if set, zero value otherwise.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetMaxCount() int32 {
	if o == nil || o.MaxCount == nil {
		var ret int32
		return ret
	}
	return *o.MaxCount
}

// GetMaxCountOk returns a tuple with the MaxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetMaxCountOk() (*int32, bool) {
	if o == nil || o.MaxCount == nil {
		return nil, false
	}
	return o.MaxCount, true
}

// HasMaxCount returns a boolean if a field has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) HasMaxCount() bool {
	if o != nil && o.MaxCount != nil {
		return true
	}

	return false
}

// SetMaxCount gets a reference to the given int32 and assigns it to the MaxCount field.
func (o *ApiInstancesIdThresholdInstanceThreshold) SetMaxCount(v int32) {
	o.MaxCount = &v
}

// GetCpuEnabled returns the CpuEnabled field value if set, zero value otherwise.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetCpuEnabled() bool {
	if o == nil || o.CpuEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CpuEnabled
}

// GetCpuEnabledOk returns a tuple with the CpuEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetCpuEnabledOk() (*bool, bool) {
	if o == nil || o.CpuEnabled == nil {
		return nil, false
	}
	return o.CpuEnabled, true
}

// HasCpuEnabled returns a boolean if a field has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) HasCpuEnabled() bool {
	if o != nil && o.CpuEnabled != nil {
		return true
	}

	return false
}

// SetCpuEnabled gets a reference to the given bool and assigns it to the CpuEnabled field.
func (o *ApiInstancesIdThresholdInstanceThreshold) SetCpuEnabled(v bool) {
	o.CpuEnabled = &v
}

// GetMinCpu returns the MinCpu field value if set, zero value otherwise.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetMinCpu() float32 {
	if o == nil || o.MinCpu == nil {
		var ret float32
		return ret
	}
	return *o.MinCpu
}

// GetMinCpuOk returns a tuple with the MinCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetMinCpuOk() (*float32, bool) {
	if o == nil || o.MinCpu == nil {
		return nil, false
	}
	return o.MinCpu, true
}

// HasMinCpu returns a boolean if a field has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) HasMinCpu() bool {
	if o != nil && o.MinCpu != nil {
		return true
	}

	return false
}

// SetMinCpu gets a reference to the given float32 and assigns it to the MinCpu field.
func (o *ApiInstancesIdThresholdInstanceThreshold) SetMinCpu(v float32) {
	o.MinCpu = &v
}

// GetMaxCpu returns the MaxCpu field value if set, zero value otherwise.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetMaxCpu() float32 {
	if o == nil || o.MaxCpu == nil {
		var ret float32
		return ret
	}
	return *o.MaxCpu
}

// GetMaxCpuOk returns a tuple with the MaxCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetMaxCpuOk() (*float32, bool) {
	if o == nil || o.MaxCpu == nil {
		return nil, false
	}
	return o.MaxCpu, true
}

// HasMaxCpu returns a boolean if a field has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) HasMaxCpu() bool {
	if o != nil && o.MaxCpu != nil {
		return true
	}

	return false
}

// SetMaxCpu gets a reference to the given float32 and assigns it to the MaxCpu field.
func (o *ApiInstancesIdThresholdInstanceThreshold) SetMaxCpu(v float32) {
	o.MaxCpu = &v
}

// GetMemoryEnabled returns the MemoryEnabled field value if set, zero value otherwise.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetMemoryEnabled() bool {
	if o == nil || o.MemoryEnabled == nil {
		var ret bool
		return ret
	}
	return *o.MemoryEnabled
}

// GetMemoryEnabledOk returns a tuple with the MemoryEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetMemoryEnabledOk() (*bool, bool) {
	if o == nil || o.MemoryEnabled == nil {
		return nil, false
	}
	return o.MemoryEnabled, true
}

// HasMemoryEnabled returns a boolean if a field has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) HasMemoryEnabled() bool {
	if o != nil && o.MemoryEnabled != nil {
		return true
	}

	return false
}

// SetMemoryEnabled gets a reference to the given bool and assigns it to the MemoryEnabled field.
func (o *ApiInstancesIdThresholdInstanceThreshold) SetMemoryEnabled(v bool) {
	o.MemoryEnabled = &v
}

// GetMinMemory returns the MinMemory field value if set, zero value otherwise.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetMinMemory() float32 {
	if o == nil || o.MinMemory == nil {
		var ret float32
		return ret
	}
	return *o.MinMemory
}

// GetMinMemoryOk returns a tuple with the MinMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetMinMemoryOk() (*float32, bool) {
	if o == nil || o.MinMemory == nil {
		return nil, false
	}
	return o.MinMemory, true
}

// HasMinMemory returns a boolean if a field has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) HasMinMemory() bool {
	if o != nil && o.MinMemory != nil {
		return true
	}

	return false
}

// SetMinMemory gets a reference to the given float32 and assigns it to the MinMemory field.
func (o *ApiInstancesIdThresholdInstanceThreshold) SetMinMemory(v float32) {
	o.MinMemory = &v
}

// GetMaxMemory returns the MaxMemory field value if set, zero value otherwise.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetMaxMemory() float32 {
	if o == nil || o.MaxMemory == nil {
		var ret float32
		return ret
	}
	return *o.MaxMemory
}

// GetMaxMemoryOk returns a tuple with the MaxMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetMaxMemoryOk() (*float32, bool) {
	if o == nil || o.MaxMemory == nil {
		return nil, false
	}
	return o.MaxMemory, true
}

// HasMaxMemory returns a boolean if a field has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) HasMaxMemory() bool {
	if o != nil && o.MaxMemory != nil {
		return true
	}

	return false
}

// SetMaxMemory gets a reference to the given float32 and assigns it to the MaxMemory field.
func (o *ApiInstancesIdThresholdInstanceThreshold) SetMaxMemory(v float32) {
	o.MaxMemory = &v
}

// GetDiskEnabled returns the DiskEnabled field value if set, zero value otherwise.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetDiskEnabled() bool {
	if o == nil || o.DiskEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DiskEnabled
}

// GetDiskEnabledOk returns a tuple with the DiskEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetDiskEnabledOk() (*bool, bool) {
	if o == nil || o.DiskEnabled == nil {
		return nil, false
	}
	return o.DiskEnabled, true
}

// HasDiskEnabled returns a boolean if a field has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) HasDiskEnabled() bool {
	if o != nil && o.DiskEnabled != nil {
		return true
	}

	return false
}

// SetDiskEnabled gets a reference to the given bool and assigns it to the DiskEnabled field.
func (o *ApiInstancesIdThresholdInstanceThreshold) SetDiskEnabled(v bool) {
	o.DiskEnabled = &v
}

// GetMinDisk returns the MinDisk field value if set, zero value otherwise.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetMinDisk() float32 {
	if o == nil || o.MinDisk == nil {
		var ret float32
		return ret
	}
	return *o.MinDisk
}

// GetMinDiskOk returns a tuple with the MinDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetMinDiskOk() (*float32, bool) {
	if o == nil || o.MinDisk == nil {
		return nil, false
	}
	return o.MinDisk, true
}

// HasMinDisk returns a boolean if a field has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) HasMinDisk() bool {
	if o != nil && o.MinDisk != nil {
		return true
	}

	return false
}

// SetMinDisk gets a reference to the given float32 and assigns it to the MinDisk field.
func (o *ApiInstancesIdThresholdInstanceThreshold) SetMinDisk(v float32) {
	o.MinDisk = &v
}

// GetMaxDisk returns the MaxDisk field value if set, zero value otherwise.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetMaxDisk() float32 {
	if o == nil || o.MaxDisk == nil {
		var ret float32
		return ret
	}
	return *o.MaxDisk
}

// GetMaxDiskOk returns a tuple with the MaxDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) GetMaxDiskOk() (*float32, bool) {
	if o == nil || o.MaxDisk == nil {
		return nil, false
	}
	return o.MaxDisk, true
}

// HasMaxDisk returns a boolean if a field has been set.
func (o *ApiInstancesIdThresholdInstanceThreshold) HasMaxDisk() bool {
	if o != nil && o.MaxDisk != nil {
		return true
	}

	return false
}

// SetMaxDisk gets a reference to the given float32 and assigns it to the MaxDisk field.
func (o *ApiInstancesIdThresholdInstanceThreshold) SetMaxDisk(v float32) {
	o.MaxDisk = &v
}

func (o ApiInstancesIdThresholdInstanceThreshold) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoUp != nil {
		toSerialize["autoUp"] = o.AutoUp
	}
	if o.AutoDown != nil {
		toSerialize["autoDown"] = o.AutoDown
	}
	if o.MinCount != nil {
		toSerialize["minCount"] = o.MinCount
	}
	if o.MaxCount != nil {
		toSerialize["maxCount"] = o.MaxCount
	}
	if o.CpuEnabled != nil {
		toSerialize["cpuEnabled"] = o.CpuEnabled
	}
	if o.MinCpu != nil {
		toSerialize["minCpu"] = o.MinCpu
	}
	if o.MaxCpu != nil {
		toSerialize["maxCpu"] = o.MaxCpu
	}
	if o.MemoryEnabled != nil {
		toSerialize["memoryEnabled"] = o.MemoryEnabled
	}
	if o.MinMemory != nil {
		toSerialize["minMemory"] = o.MinMemory
	}
	if o.MaxMemory != nil {
		toSerialize["maxMemory"] = o.MaxMemory
	}
	if o.DiskEnabled != nil {
		toSerialize["diskEnabled"] = o.DiskEnabled
	}
	if o.MinDisk != nil {
		toSerialize["minDisk"] = o.MinDisk
	}
	if o.MaxDisk != nil {
		toSerialize["maxDisk"] = o.MaxDisk
	}
	return json.Marshal(toSerialize)
}

type NullableApiInstancesIdThresholdInstanceThreshold struct {
	value *ApiInstancesIdThresholdInstanceThreshold
	isSet bool
}

func (v NullableApiInstancesIdThresholdInstanceThreshold) Get() *ApiInstancesIdThresholdInstanceThreshold {
	return v.value
}

func (v *NullableApiInstancesIdThresholdInstanceThreshold) Set(val *ApiInstancesIdThresholdInstanceThreshold) {
	v.value = val
	v.isSet = true
}

func (v NullableApiInstancesIdThresholdInstanceThreshold) IsSet() bool {
	return v.isSet
}

func (v *NullableApiInstancesIdThresholdInstanceThreshold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiInstancesIdThresholdInstanceThreshold(val *ApiInstancesIdThresholdInstanceThreshold) *NullableApiInstancesIdThresholdInstanceThreshold {
	return &NullableApiInstancesIdThresholdInstanceThreshold{value: val, isSet: true}
}

func (v NullableApiInstancesIdThresholdInstanceThreshold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiInstancesIdThresholdInstanceThreshold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


