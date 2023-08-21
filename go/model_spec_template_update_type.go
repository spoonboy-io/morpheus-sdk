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

// SpecTemplateUpdateType struct for SpecTemplateUpdateType
type SpecTemplateUpdateType struct {
	// Spec Template Type. i.e. arm, cloudFormation, helm, kubernetes, oneview, terraform, ucs.
	Code *string `json:"code,omitempty"`
}

// NewSpecTemplateUpdateType instantiates a new SpecTemplateUpdateType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpecTemplateUpdateType() *SpecTemplateUpdateType {
	this := SpecTemplateUpdateType{}
	return &this
}

// NewSpecTemplateUpdateTypeWithDefaults instantiates a new SpecTemplateUpdateType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpecTemplateUpdateTypeWithDefaults() *SpecTemplateUpdateType {
	this := SpecTemplateUpdateType{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SpecTemplateUpdateType) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecTemplateUpdateType) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SpecTemplateUpdateType) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SpecTemplateUpdateType) SetCode(v string) {
	o.Code = &v
}

func (o SpecTemplateUpdateType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	return json.Marshal(toSerialize)
}

type NullableSpecTemplateUpdateType struct {
	value *SpecTemplateUpdateType
	isSet bool
}

func (v NullableSpecTemplateUpdateType) Get() *SpecTemplateUpdateType {
	return v.value
}

func (v *NullableSpecTemplateUpdateType) Set(val *SpecTemplateUpdateType) {
	v.value = val
	v.isSet = true
}

func (v NullableSpecTemplateUpdateType) IsSet() bool {
	return v.isSet
}

func (v *NullableSpecTemplateUpdateType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpecTemplateUpdateType(val *SpecTemplateUpdateType) *NullableSpecTemplateUpdateType {
	return &NullableSpecTemplateUpdateType{value: val, isSet: true}
}

func (v NullableSpecTemplateUpdateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpecTemplateUpdateType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


