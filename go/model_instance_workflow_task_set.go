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

// InstanceWorkflowTaskSet Object containing workflow configuration parameters.
type InstanceWorkflowTaskSet struct {
	// Object containing any custom option type configuration parameters
	CustomOptions *map[string]interface{} `json:"customOptions,omitempty"`
}

// NewInstanceWorkflowTaskSet instantiates a new InstanceWorkflowTaskSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceWorkflowTaskSet() *InstanceWorkflowTaskSet {
	this := InstanceWorkflowTaskSet{}
	return &this
}

// NewInstanceWorkflowTaskSetWithDefaults instantiates a new InstanceWorkflowTaskSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceWorkflowTaskSetWithDefaults() *InstanceWorkflowTaskSet {
	this := InstanceWorkflowTaskSet{}
	return &this
}

// GetCustomOptions returns the CustomOptions field value if set, zero value otherwise.
func (o *InstanceWorkflowTaskSet) GetCustomOptions() map[string]interface{} {
	if o == nil || o.CustomOptions == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.CustomOptions
}

// GetCustomOptionsOk returns a tuple with the CustomOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceWorkflowTaskSet) GetCustomOptionsOk() (*map[string]interface{}, bool) {
	if o == nil || o.CustomOptions == nil {
		return nil, false
	}
	return o.CustomOptions, true
}

// HasCustomOptions returns a boolean if a field has been set.
func (o *InstanceWorkflowTaskSet) HasCustomOptions() bool {
	if o != nil && o.CustomOptions != nil {
		return true
	}

	return false
}

// SetCustomOptions gets a reference to the given map[string]interface{} and assigns it to the CustomOptions field.
func (o *InstanceWorkflowTaskSet) SetCustomOptions(v map[string]interface{}) {
	o.CustomOptions = &v
}

func (o InstanceWorkflowTaskSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomOptions != nil {
		toSerialize["customOptions"] = o.CustomOptions
	}
	return json.Marshal(toSerialize)
}

type NullableInstanceWorkflowTaskSet struct {
	value *InstanceWorkflowTaskSet
	isSet bool
}

func (v NullableInstanceWorkflowTaskSet) Get() *InstanceWorkflowTaskSet {
	return v.value
}

func (v *NullableInstanceWorkflowTaskSet) Set(val *InstanceWorkflowTaskSet) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceWorkflowTaskSet) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceWorkflowTaskSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceWorkflowTaskSet(val *InstanceWorkflowTaskSet) *NullableInstanceWorkflowTaskSet {
	return &NullableInstanceWorkflowTaskSet{value: val, isSet: true}
}

func (v NullableInstanceWorkflowTaskSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceWorkflowTaskSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

