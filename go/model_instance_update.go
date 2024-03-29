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

// InstanceUpdate struct for InstanceUpdate
type InstanceUpdate struct {
	Instance *InstanceUpdateInstance `json:"instance,omitempty"`
	Config *InstancesConfigCustomOptions `json:"config,omitempty"`
}

// NewInstanceUpdate instantiates a new InstanceUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceUpdate() *InstanceUpdate {
	this := InstanceUpdate{}
	return &this
}

// NewInstanceUpdateWithDefaults instantiates a new InstanceUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceUpdateWithDefaults() *InstanceUpdate {
	this := InstanceUpdate{}
	return &this
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *InstanceUpdate) GetInstance() InstanceUpdateInstance {
	if o == nil || o.Instance == nil {
		var ret InstanceUpdateInstance
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceUpdate) GetInstanceOk() (*InstanceUpdateInstance, bool) {
	if o == nil || o.Instance == nil {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *InstanceUpdate) HasInstance() bool {
	if o != nil && o.Instance != nil {
		return true
	}

	return false
}

// SetInstance gets a reference to the given InstanceUpdateInstance and assigns it to the Instance field.
func (o *InstanceUpdate) SetInstance(v InstanceUpdateInstance) {
	o.Instance = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *InstanceUpdate) GetConfig() InstancesConfigCustomOptions {
	if o == nil || o.Config == nil {
		var ret InstancesConfigCustomOptions
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceUpdate) GetConfigOk() (*InstancesConfigCustomOptions, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *InstanceUpdate) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given InstancesConfigCustomOptions and assigns it to the Config field.
func (o *InstanceUpdate) SetConfig(v InstancesConfigCustomOptions) {
	o.Config = &v
}

func (o InstanceUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Instance != nil {
		toSerialize["instance"] = o.Instance
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableInstanceUpdate struct {
	value *InstanceUpdate
	isSet bool
}

func (v NullableInstanceUpdate) Get() *InstanceUpdate {
	return v.value
}

func (v *NullableInstanceUpdate) Set(val *InstanceUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceUpdate(val *InstanceUpdate) *NullableInstanceUpdate {
	return &NullableInstanceUpdate{value: val, isSet: true}
}

func (v NullableInstanceUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


