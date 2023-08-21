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

// InstanceEnvs struct for InstanceEnvs
type InstanceEnvs struct {
	Envs []InstanceEnvsEnvs `json:"envs,omitempty"`
	ReadOnlyEnvs []InstanceEnvsEnvs `json:"readOnlyEnvs,omitempty"`
	ImportedEnvs []InstanceEnvsEnvs `json:"importedEnvs,omitempty"`
}

// NewInstanceEnvs instantiates a new InstanceEnvs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceEnvs() *InstanceEnvs {
	this := InstanceEnvs{}
	return &this
}

// NewInstanceEnvsWithDefaults instantiates a new InstanceEnvs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceEnvsWithDefaults() *InstanceEnvs {
	this := InstanceEnvs{}
	return &this
}

// GetEnvs returns the Envs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceEnvs) GetEnvs() []InstanceEnvsEnvs {
	if o == nil  {
		var ret []InstanceEnvsEnvs
		return ret
	}
	return o.Envs
}

// GetEnvsOk returns a tuple with the Envs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceEnvs) GetEnvsOk() (*[]InstanceEnvsEnvs, bool) {
	if o == nil || o.Envs == nil {
		return nil, false
	}
	return &o.Envs, true
}

// HasEnvs returns a boolean if a field has been set.
func (o *InstanceEnvs) HasEnvs() bool {
	if o != nil && o.Envs != nil {
		return true
	}

	return false
}

// SetEnvs gets a reference to the given []InstanceEnvsEnvs and assigns it to the Envs field.
func (o *InstanceEnvs) SetEnvs(v []InstanceEnvsEnvs) {
	o.Envs = v
}

// GetReadOnlyEnvs returns the ReadOnlyEnvs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceEnvs) GetReadOnlyEnvs() []InstanceEnvsEnvs {
	if o == nil  {
		var ret []InstanceEnvsEnvs
		return ret
	}
	return o.ReadOnlyEnvs
}

// GetReadOnlyEnvsOk returns a tuple with the ReadOnlyEnvs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceEnvs) GetReadOnlyEnvsOk() (*[]InstanceEnvsEnvs, bool) {
	if o == nil || o.ReadOnlyEnvs == nil {
		return nil, false
	}
	return &o.ReadOnlyEnvs, true
}

// HasReadOnlyEnvs returns a boolean if a field has been set.
func (o *InstanceEnvs) HasReadOnlyEnvs() bool {
	if o != nil && o.ReadOnlyEnvs != nil {
		return true
	}

	return false
}

// SetReadOnlyEnvs gets a reference to the given []InstanceEnvsEnvs and assigns it to the ReadOnlyEnvs field.
func (o *InstanceEnvs) SetReadOnlyEnvs(v []InstanceEnvsEnvs) {
	o.ReadOnlyEnvs = v
}

// GetImportedEnvs returns the ImportedEnvs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceEnvs) GetImportedEnvs() []InstanceEnvsEnvs {
	if o == nil  {
		var ret []InstanceEnvsEnvs
		return ret
	}
	return o.ImportedEnvs
}

// GetImportedEnvsOk returns a tuple with the ImportedEnvs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceEnvs) GetImportedEnvsOk() (*[]InstanceEnvsEnvs, bool) {
	if o == nil || o.ImportedEnvs == nil {
		return nil, false
	}
	return &o.ImportedEnvs, true
}

// HasImportedEnvs returns a boolean if a field has been set.
func (o *InstanceEnvs) HasImportedEnvs() bool {
	if o != nil && o.ImportedEnvs != nil {
		return true
	}

	return false
}

// SetImportedEnvs gets a reference to the given []InstanceEnvsEnvs and assigns it to the ImportedEnvs field.
func (o *InstanceEnvs) SetImportedEnvs(v []InstanceEnvsEnvs) {
	o.ImportedEnvs = v
}

func (o InstanceEnvs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Envs != nil {
		toSerialize["envs"] = o.Envs
	}
	if o.ReadOnlyEnvs != nil {
		toSerialize["readOnlyEnvs"] = o.ReadOnlyEnvs
	}
	if o.ImportedEnvs != nil {
		toSerialize["importedEnvs"] = o.ImportedEnvs
	}
	return json.Marshal(toSerialize)
}

type NullableInstanceEnvs struct {
	value *InstanceEnvs
	isSet bool
}

func (v NullableInstanceEnvs) Get() *InstanceEnvs {
	return v.value
}

func (v *NullableInstanceEnvs) Set(val *InstanceEnvs) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceEnvs) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceEnvs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceEnvs(val *InstanceEnvs) *NullableInstanceEnvs {
	return &NullableInstanceEnvs{value: val, isSet: true}
}

func (v NullableInstanceEnvs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceEnvs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


