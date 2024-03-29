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

// InlineObject14 struct for InlineObject14
type InlineObject14 struct {
	// A script or command to be executed
	Script string `json:"script"`
}

// NewInlineObject14 instantiates a new InlineObject14 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject14(script string, ) *InlineObject14 {
	this := InlineObject14{}
	this.Script = script
	return &this
}

// NewInlineObject14WithDefaults instantiates a new InlineObject14 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject14WithDefaults() *InlineObject14 {
	this := InlineObject14{}
	return &this
}

// GetScript returns the Script field value
func (o *InlineObject14) GetScript() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Script
}

// GetScriptOk returns a tuple with the Script field value
// and a boolean to check if the value has been set.
func (o *InlineObject14) GetScriptOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Script, true
}

// SetScript sets field value
func (o *InlineObject14) SetScript(v string) {
	o.Script = v
}

func (o InlineObject14) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["script"] = o.Script
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject14 struct {
	value *InlineObject14
	isSet bool
}

func (v NullableInlineObject14) Get() *InlineObject14 {
	return v.value
}

func (v *NullableInlineObject14) Set(val *InlineObject14) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject14) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject14) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject14(val *InlineObject14) *NullableInlineObject14 {
	return &NullableInlineObject14{value: val, isSet: true}
}

func (v NullableInlineObject14) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject14) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


