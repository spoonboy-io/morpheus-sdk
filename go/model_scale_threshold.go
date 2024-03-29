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
	"time"
)

// ScaleThreshold struct for ScaleThreshold
type ScaleThreshold struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	AutoUp *bool `json:"autoUp,omitempty"`
	AutoDown *bool `json:"autoDown,omitempty"`
	MinCount *int64 `json:"minCount,omitempty"`
	MaxCount *int64 `json:"maxCount,omitempty"`
	ScaleIncrement *int64 `json:"scaleIncrement,omitempty"`
	CpuEnabled *bool `json:"cpuEnabled,omitempty"`
	MinCpu *int64 `json:"minCpu,omitempty"`
	MaxCpu *int64 `json:"maxCpu,omitempty"`
	MemoryEnabled *bool `json:"memoryEnabled,omitempty"`
	MinMemory *int64 `json:"minMemory,omitempty"`
	MaxMemory *int64 `json:"maxMemory,omitempty"`
	DiskEnabled *bool `json:"diskEnabled,omitempty"`
	MinDisk *int64 `json:"minDisk,omitempty"`
	MaxDisk *int64 `json:"maxDisk,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
}

// NewScaleThreshold instantiates a new ScaleThreshold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScaleThreshold() *ScaleThreshold {
	this := ScaleThreshold{}
	return &this
}

// NewScaleThresholdWithDefaults instantiates a new ScaleThreshold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScaleThresholdWithDefaults() *ScaleThreshold {
	this := ScaleThreshold{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ScaleThreshold) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleThreshold) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ScaleThreshold) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ScaleThreshold) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScaleThreshold) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleThreshold) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ScaleThreshold) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ScaleThreshold) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ScaleThreshold) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleThreshold) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ScaleThreshold) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ScaleThreshold) SetType(v string) {
	o.Type = &v
}

// GetAutoUp returns the AutoUp field value if set, zero value otherwise.
func (o *ScaleThreshold) GetAutoUp() bool {
	if o == nil || o.AutoUp == nil {
		var ret bool
		return ret
	}
	return *o.AutoUp
}

// GetAutoUpOk returns a tuple with the AutoUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleThreshold) GetAutoUpOk() (*bool, bool) {
	if o == nil || o.AutoUp == nil {
		return nil, false
	}
	return o.AutoUp, true
}

// HasAutoUp returns a boolean if a field has been set.
func (o *ScaleThreshold) HasAutoUp() bool {
	if o != nil && o.AutoUp != nil {
		return true
	}

	return false
}

// SetAutoUp gets a reference to the given bool and assigns it to the AutoUp field.
func (o *ScaleThreshold) SetAutoUp(v bool) {
	o.AutoUp = &v
}

// GetAutoDown returns the AutoDown field value if set, zero value otherwise.
func (o *ScaleThreshold) GetAutoDown() bool {
	if o == nil || o.AutoDown == nil {
		var ret bool
		return ret
	}
	return *o.AutoDown
}

// GetAutoDownOk returns a tuple with the AutoDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleThreshold) GetAutoDownOk() (*bool, bool) {
	if o == nil || o.AutoDown == nil {
		return nil, false
	}
	return o.AutoDown, true
}

// HasAutoDown returns a boolean if a field has been set.
func (o *ScaleThreshold) HasAutoDown() bool {
	if o != nil && o.AutoDown != nil {
		return true
	}

	return false
}

// SetAutoDown gets a reference to the given bool and assigns it to the AutoDown field.
func (o *ScaleThreshold) SetAutoDown(v bool) {
	o.AutoDown = &v
}

// GetMinCount returns the MinCount field value if set, zero value otherwise.
func (o *ScaleThreshold) GetMinCount() int64 {
	if o == nil || o.MinCount == nil {
		var ret int64
		return ret
	}
	return *o.MinCount
}

// GetMinCountOk returns a tuple with the MinCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleThreshold) GetMinCountOk() (*int64, bool) {
	if o == nil || o.MinCount == nil {
		return nil, false
	}
	return o.MinCount, true
}

// HasMinCount returns a boolean if a field has been set.
func (o *ScaleThreshold) HasMinCount() bool {
	if o != nil && o.MinCount != nil {
		return true
	}

	return false
}

// SetMinCount gets a reference to the given int64 and assigns it to the MinCount field.
func (o *ScaleThreshold) SetMinCount(v int64) {
	o.MinCount = &v
}

// GetMaxCount returns the MaxCount field value if set, zero value otherwise.
func (o *ScaleThreshold) GetMaxCount() int64 {
	if o == nil || o.MaxCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxCount
}

// GetMaxCountOk returns a tuple with the MaxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleThreshold) GetMaxCountOk() (*int64, bool) {
	if o == nil || o.MaxCount == nil {
		return nil, false
	}
	return o.MaxCount, true
}

// HasMaxCount returns a boolean if a field has been set.
func (o *ScaleThreshold) HasMaxCount() bool {
	if o != nil && o.MaxCount != nil {
		return true
	}

	return false
}

// SetMaxCount gets a reference to the given int64 and assigns it to the MaxCount field.
func (o *ScaleThreshold) SetMaxCount(v int64) {
	o.MaxCount = &v
}

// GetScaleIncrement returns the ScaleIncrement field value if set, zero value otherwise.
func (o *ScaleThreshold) GetScaleIncrement() int64 {
	if o == nil || o.ScaleIncrement == nil {
		var ret int64
		return ret
	}
	return *o.ScaleIncrement
}

// GetScaleIncrementOk returns a tuple with the ScaleIncrement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleThreshold) GetScaleIncrementOk() (*int64, bool) {
	if o == nil || o.ScaleIncrement == nil {
		return nil, false
	}
	return o.ScaleIncrement, true
}

// HasScaleIncrement returns a boolean if a field has been set.
func (o *ScaleThreshold) HasScaleIncrement() bool {
	if o != nil && o.ScaleIncrement != nil {
		return true
	}

	return false
}

// SetScaleIncrement gets a reference to the given int64 and assigns it to the ScaleIncrement field.
func (o *ScaleThreshold) SetScaleIncrement(v int64) {
	o.ScaleIncrement = &v
}

// GetCpuEnabled returns the CpuEnabled field value if set, zero value otherwise.
func (o *ScaleThreshold) GetCpuEnabled() bool {
	if o == nil || o.CpuEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CpuEnabled
}

// GetCpuEnabledOk returns a tuple with the CpuEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleThreshold) GetCpuEnabledOk() (*bool, bool) {
	if o == nil || o.CpuEnabled == nil {
		return nil, false
	}
	return o.CpuEnabled, true
}

// HasCpuEnabled returns a boolean if a field has been set.
func (o *ScaleThreshold) HasCpuEnabled() bool {
	if o != nil && o.CpuEnabled != nil {
		return true
	}

	return false
}

// SetCpuEnabled gets a reference to the given bool and assigns it to the CpuEnabled field.
func (o *ScaleThreshold) SetCpuEnabled(v bool) {
	o.CpuEnabled = &v
}

// GetMinCpu returns the MinCpu field value if set, zero value otherwise.
func (o *ScaleThreshold) GetMinCpu() int64 {
	if o == nil || o.MinCpu == nil {
		var ret int64
		return ret
	}
	return *o.MinCpu
}

// GetMinCpuOk returns a tuple with the MinCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleThreshold) GetMinCpuOk() (*int64, bool) {
	if o == nil || o.MinCpu == nil {
		return nil, false
	}
	return o.MinCpu, true
}

// HasMinCpu returns a boolean if a field has been set.
func (o *ScaleThreshold) HasMinCpu() bool {
	if o != nil && o.MinCpu != nil {
		return true
	}

	return false
}

// SetMinCpu gets a reference to the given int64 and assigns it to the MinCpu field.
func (o *ScaleThreshold) SetMinCpu(v int64) {
	o.MinCpu = &v
}

// GetMaxCpu returns the MaxCpu field value if set, zero value otherwise.
func (o *ScaleThreshold) GetMaxCpu() int64 {
	if o == nil || o.MaxCpu == nil {
		var ret int64
		return ret
	}
	return *o.MaxCpu
}

// GetMaxCpuOk returns a tuple with the MaxCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleThreshold) GetMaxCpuOk() (*int64, bool) {
	if o == nil || o.MaxCpu == nil {
		return nil, false
	}
	return o.MaxCpu, true
}

// HasMaxCpu returns a boolean if a field has been set.
func (o *ScaleThreshold) HasMaxCpu() bool {
	if o != nil && o.MaxCpu != nil {
		return true
	}

	return false
}

// SetMaxCpu gets a reference to the given int64 and assigns it to the MaxCpu field.
func (o *ScaleThreshold) SetMaxCpu(v int64) {
	o.MaxCpu = &v
}

// GetMemoryEnabled returns the MemoryEnabled field value if set, zero value otherwise.
func (o *ScaleThreshold) GetMemoryEnabled() bool {
	if o == nil || o.MemoryEnabled == nil {
		var ret bool
		return ret
	}
	return *o.MemoryEnabled
}

// GetMemoryEnabledOk returns a tuple with the MemoryEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleThreshold) GetMemoryEnabledOk() (*bool, bool) {
	if o == nil || o.MemoryEnabled == nil {
		return nil, false
	}
	return o.MemoryEnabled, true
}

// HasMemoryEnabled returns a boolean if a field has been set.
func (o *ScaleThreshold) HasMemoryEnabled() bool {
	if o != nil && o.MemoryEnabled != nil {
		return true
	}

	return false
}

// SetMemoryEnabled gets a reference to the given bool and assigns it to the MemoryEnabled field.
func (o *ScaleThreshold) SetMemoryEnabled(v bool) {
	o.MemoryEnabled = &v
}

// GetMinMemory returns the MinMemory field value if set, zero value otherwise.
func (o *ScaleThreshold) GetMinMemory() int64 {
	if o == nil || o.MinMemory == nil {
		var ret int64
		return ret
	}
	return *o.MinMemory
}

// GetMinMemoryOk returns a tuple with the MinMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleThreshold) GetMinMemoryOk() (*int64, bool) {
	if o == nil || o.MinMemory == nil {
		return nil, false
	}
	return o.MinMemory, true
}

// HasMinMemory returns a boolean if a field has been set.
func (o *ScaleThreshold) HasMinMemory() bool {
	if o != nil && o.MinMemory != nil {
		return true
	}

	return false
}

// SetMinMemory gets a reference to the given int64 and assigns it to the MinMemory field.
func (o *ScaleThreshold) SetMinMemory(v int64) {
	o.MinMemory = &v
}

// GetMaxMemory returns the MaxMemory field value if set, zero value otherwise.
func (o *ScaleThreshold) GetMaxMemory() int64 {
	if o == nil || o.MaxMemory == nil {
		var ret int64
		return ret
	}
	return *o.MaxMemory
}

// GetMaxMemoryOk returns a tuple with the MaxMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleThreshold) GetMaxMemoryOk() (*int64, bool) {
	if o == nil || o.MaxMemory == nil {
		return nil, false
	}
	return o.MaxMemory, true
}

// HasMaxMemory returns a boolean if a field has been set.
func (o *ScaleThreshold) HasMaxMemory() bool {
	if o != nil && o.MaxMemory != nil {
		return true
	}

	return false
}

// SetMaxMemory gets a reference to the given int64 and assigns it to the MaxMemory field.
func (o *ScaleThreshold) SetMaxMemory(v int64) {
	o.MaxMemory = &v
}

// GetDiskEnabled returns the DiskEnabled field value if set, zero value otherwise.
func (o *ScaleThreshold) GetDiskEnabled() bool {
	if o == nil || o.DiskEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DiskEnabled
}

// GetDiskEnabledOk returns a tuple with the DiskEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleThreshold) GetDiskEnabledOk() (*bool, bool) {
	if o == nil || o.DiskEnabled == nil {
		return nil, false
	}
	return o.DiskEnabled, true
}

// HasDiskEnabled returns a boolean if a field has been set.
func (o *ScaleThreshold) HasDiskEnabled() bool {
	if o != nil && o.DiskEnabled != nil {
		return true
	}

	return false
}

// SetDiskEnabled gets a reference to the given bool and assigns it to the DiskEnabled field.
func (o *ScaleThreshold) SetDiskEnabled(v bool) {
	o.DiskEnabled = &v
}

// GetMinDisk returns the MinDisk field value if set, zero value otherwise.
func (o *ScaleThreshold) GetMinDisk() int64 {
	if o == nil || o.MinDisk == nil {
		var ret int64
		return ret
	}
	return *o.MinDisk
}

// GetMinDiskOk returns a tuple with the MinDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleThreshold) GetMinDiskOk() (*int64, bool) {
	if o == nil || o.MinDisk == nil {
		return nil, false
	}
	return o.MinDisk, true
}

// HasMinDisk returns a boolean if a field has been set.
func (o *ScaleThreshold) HasMinDisk() bool {
	if o != nil && o.MinDisk != nil {
		return true
	}

	return false
}

// SetMinDisk gets a reference to the given int64 and assigns it to the MinDisk field.
func (o *ScaleThreshold) SetMinDisk(v int64) {
	o.MinDisk = &v
}

// GetMaxDisk returns the MaxDisk field value if set, zero value otherwise.
func (o *ScaleThreshold) GetMaxDisk() int64 {
	if o == nil || o.MaxDisk == nil {
		var ret int64
		return ret
	}
	return *o.MaxDisk
}

// GetMaxDiskOk returns a tuple with the MaxDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleThreshold) GetMaxDiskOk() (*int64, bool) {
	if o == nil || o.MaxDisk == nil {
		return nil, false
	}
	return o.MaxDisk, true
}

// HasMaxDisk returns a boolean if a field has been set.
func (o *ScaleThreshold) HasMaxDisk() bool {
	if o != nil && o.MaxDisk != nil {
		return true
	}

	return false
}

// SetMaxDisk gets a reference to the given int64 and assigns it to the MaxDisk field.
func (o *ScaleThreshold) SetMaxDisk(v int64) {
	o.MaxDisk = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *ScaleThreshold) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleThreshold) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *ScaleThreshold) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *ScaleThreshold) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ScaleThreshold) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleThreshold) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ScaleThreshold) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ScaleThreshold) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o ScaleThreshold) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
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
	if o.ScaleIncrement != nil {
		toSerialize["scaleIncrement"] = o.ScaleIncrement
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
	if o.DateCreated != nil {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	return json.Marshal(toSerialize)
}

type NullableScaleThreshold struct {
	value *ScaleThreshold
	isSet bool
}

func (v NullableScaleThreshold) Get() *ScaleThreshold {
	return v.value
}

func (v *NullableScaleThreshold) Set(val *ScaleThreshold) {
	v.value = val
	v.isSet = true
}

func (v NullableScaleThreshold) IsSet() bool {
	return v.isSet
}

func (v *NullableScaleThreshold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScaleThreshold(val *ScaleThreshold) *NullableScaleThreshold {
	return &NullableScaleThreshold{value: val, isSet: true}
}

func (v NullableScaleThreshold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScaleThreshold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


