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

// InlineObject261 Create Pool Object
type InlineObject261 struct {
	VdiPool OneOfobjectobject `json:"vdiPool"`
}

// NewInlineObject261 instantiates a new InlineObject261 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject261(vdiPool OneOfobjectobject, ) *InlineObject261 {
	this := InlineObject261{}
	this.VdiPool = vdiPool
	return &this
}

// NewInlineObject261WithDefaults instantiates a new InlineObject261 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject261WithDefaults() *InlineObject261 {
	this := InlineObject261{}
	return &this
}

// GetVdiPool returns the VdiPool field value
func (o *InlineObject261) GetVdiPool() OneOfobjectobject {
	if o == nil  {
		var ret OneOfobjectobject
		return ret
	}

	return o.VdiPool
}

// GetVdiPoolOk returns a tuple with the VdiPool field value
// and a boolean to check if the value has been set.
func (o *InlineObject261) GetVdiPoolOk() (*OneOfobjectobject, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VdiPool, true
}

// SetVdiPool sets field value
func (o *InlineObject261) SetVdiPool(v OneOfobjectobject) {
	o.VdiPool = v
}

func (o InlineObject261) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["vdiPool"] = o.VdiPool
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject261 struct {
	value *InlineObject261
	isSet bool
}

func (v NullableInlineObject261) Get() *InlineObject261 {
	return v.value
}

func (v *NullableInlineObject261) Set(val *InlineObject261) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject261) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject261) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject261(val *InlineObject261) *NullableInlineObject261 {
	return &NullableInlineObject261{value: val, isSet: true}
}

func (v NullableInlineObject261) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject261) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


