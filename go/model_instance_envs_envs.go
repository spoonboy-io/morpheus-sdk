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

// InstanceEnvsEnvs struct for InstanceEnvsEnvs
type InstanceEnvsEnvs struct {
	Export bool `json:"export"`
	Masked bool `json:"masked"`
	Name string `json:"name"`
	Value string `json:"value"`
}

// NewInstanceEnvsEnvs instantiates a new InstanceEnvsEnvs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceEnvsEnvs(export bool, masked bool, name string, value string, ) *InstanceEnvsEnvs {
	this := InstanceEnvsEnvs{}
	this.Export = export
	this.Masked = masked
	this.Name = name
	this.Value = value
	return &this
}

// NewInstanceEnvsEnvsWithDefaults instantiates a new InstanceEnvsEnvs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceEnvsEnvsWithDefaults() *InstanceEnvsEnvs {
	this := InstanceEnvsEnvs{}
	return &this
}

// GetExport returns the Export field value
func (o *InstanceEnvsEnvs) GetExport() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Export
}

// GetExportOk returns a tuple with the Export field value
// and a boolean to check if the value has been set.
func (o *InstanceEnvsEnvs) GetExportOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Export, true
}

// SetExport sets field value
func (o *InstanceEnvsEnvs) SetExport(v bool) {
	o.Export = v
}

// GetMasked returns the Masked field value
func (o *InstanceEnvsEnvs) GetMasked() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Masked
}

// GetMaskedOk returns a tuple with the Masked field value
// and a boolean to check if the value has been set.
func (o *InstanceEnvsEnvs) GetMaskedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Masked, true
}

// SetMasked sets field value
func (o *InstanceEnvsEnvs) SetMasked(v bool) {
	o.Masked = v
}

// GetName returns the Name field value
func (o *InstanceEnvsEnvs) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InstanceEnvsEnvs) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InstanceEnvsEnvs) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *InstanceEnvsEnvs) GetValue() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *InstanceEnvsEnvs) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *InstanceEnvsEnvs) SetValue(v string) {
	o.Value = v
}

func (o InstanceEnvsEnvs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["export"] = o.Export
	}
	if true {
		toSerialize["masked"] = o.Masked
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableInstanceEnvsEnvs struct {
	value *InstanceEnvsEnvs
	isSet bool
}

func (v NullableInstanceEnvsEnvs) Get() *InstanceEnvsEnvs {
	return v.value
}

func (v *NullableInstanceEnvsEnvs) Set(val *InstanceEnvsEnvs) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceEnvsEnvs) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceEnvsEnvs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceEnvsEnvs(val *InstanceEnvsEnvs) *NullableInstanceEnvsEnvs {
	return &NullableInstanceEnvsEnvs{value: val, isSet: true}
}

func (v NullableInstanceEnvsEnvs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceEnvsEnvs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


