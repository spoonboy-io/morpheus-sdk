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

// InlineObject257 Create VDI App
type InlineObject257 struct {
	VdiApp OneOfobject `json:"vdiApp"`
}

// NewInlineObject257 instantiates a new InlineObject257 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject257(vdiApp OneOfobject, ) *InlineObject257 {
	this := InlineObject257{}
	this.VdiApp = vdiApp
	return &this
}

// NewInlineObject257WithDefaults instantiates a new InlineObject257 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject257WithDefaults() *InlineObject257 {
	this := InlineObject257{}
	return &this
}

// GetVdiApp returns the VdiApp field value
func (o *InlineObject257) GetVdiApp() OneOfobject {
	if o == nil  {
		var ret OneOfobject
		return ret
	}

	return o.VdiApp
}

// GetVdiAppOk returns a tuple with the VdiApp field value
// and a boolean to check if the value has been set.
func (o *InlineObject257) GetVdiAppOk() (*OneOfobject, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VdiApp, true
}

// SetVdiApp sets field value
func (o *InlineObject257) SetVdiApp(v OneOfobject) {
	o.VdiApp = v
}

func (o InlineObject257) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["vdiApp"] = o.VdiApp
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject257 struct {
	value *InlineObject257
	isSet bool
}

func (v NullableInlineObject257) Get() *InlineObject257 {
	return v.value
}

func (v *NullableInlineObject257) Set(val *InlineObject257) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject257) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject257) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject257(val *InlineObject257) *NullableInlineObject257 {
	return &NullableInlineObject257{value: val, isSet: true}
}

func (v NullableInlineObject257) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject257) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

