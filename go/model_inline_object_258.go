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

// InlineObject258 Updates VDI App
type InlineObject258 struct {
	VdiApp OneOfobject `json:"vdiApp"`
}

// NewInlineObject258 instantiates a new InlineObject258 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject258(vdiApp OneOfobject, ) *InlineObject258 {
	this := InlineObject258{}
	this.VdiApp = vdiApp
	return &this
}

// NewInlineObject258WithDefaults instantiates a new InlineObject258 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject258WithDefaults() *InlineObject258 {
	this := InlineObject258{}
	return &this
}

// GetVdiApp returns the VdiApp field value
func (o *InlineObject258) GetVdiApp() OneOfobject {
	if o == nil  {
		var ret OneOfobject
		return ret
	}

	return o.VdiApp
}

// GetVdiAppOk returns a tuple with the VdiApp field value
// and a boolean to check if the value has been set.
func (o *InlineObject258) GetVdiAppOk() (*OneOfobject, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VdiApp, true
}

// SetVdiApp sets field value
func (o *InlineObject258) SetVdiApp(v OneOfobject) {
	o.VdiApp = v
}

func (o InlineObject258) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["vdiApp"] = o.VdiApp
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject258 struct {
	value *InlineObject258
	isSet bool
}

func (v NullableInlineObject258) Get() *InlineObject258 {
	return v.value
}

func (v *NullableInlineObject258) Set(val *InlineObject258) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject258) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject258) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject258(val *InlineObject258) *NullableInlineObject258 {
	return &NullableInlineObject258{value: val, isSet: true}
}

func (v NullableInlineObject258) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject258) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


