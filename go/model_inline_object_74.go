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

// InlineObject74 struct for InlineObject74
type InlineObject74 struct {
	Environment ApiEnvironmentsEnvironment `json:"environment"`
}

// NewInlineObject74 instantiates a new InlineObject74 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject74(environment ApiEnvironmentsEnvironment, ) *InlineObject74 {
	this := InlineObject74{}
	this.Environment = environment
	return &this
}

// NewInlineObject74WithDefaults instantiates a new InlineObject74 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject74WithDefaults() *InlineObject74 {
	this := InlineObject74{}
	return &this
}

// GetEnvironment returns the Environment field value
func (o *InlineObject74) GetEnvironment() ApiEnvironmentsEnvironment {
	if o == nil  {
		var ret ApiEnvironmentsEnvironment
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *InlineObject74) GetEnvironmentOk() (*ApiEnvironmentsEnvironment, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *InlineObject74) SetEnvironment(v ApiEnvironmentsEnvironment) {
	o.Environment = v
}

func (o InlineObject74) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["environment"] = o.Environment
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject74 struct {
	value *InlineObject74
	isSet bool
}

func (v NullableInlineObject74) Get() *InlineObject74 {
	return v.value
}

func (v *NullableInlineObject74) Set(val *InlineObject74) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject74) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject74) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject74(val *InlineObject74) *NullableInlineObject74 {
	return &NullableInlineObject74{value: val, isSet: true}
}

func (v NullableInlineObject74) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject74) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


