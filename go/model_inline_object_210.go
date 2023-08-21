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

// InlineObject210 struct for InlineObject210
type InlineObject210 struct {
	// `id` of the VDI Pool
	VdiPoolId int32 `json:"vdiPoolId"`
	// The new access level.
	Access string `json:"access"`
}

// NewInlineObject210 instantiates a new InlineObject210 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject210(vdiPoolId int32, access string, ) *InlineObject210 {
	this := InlineObject210{}
	this.VdiPoolId = vdiPoolId
	this.Access = access
	return &this
}

// NewInlineObject210WithDefaults instantiates a new InlineObject210 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject210WithDefaults() *InlineObject210 {
	this := InlineObject210{}
	return &this
}

// GetVdiPoolId returns the VdiPoolId field value
func (o *InlineObject210) GetVdiPoolId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.VdiPoolId
}

// GetVdiPoolIdOk returns a tuple with the VdiPoolId field value
// and a boolean to check if the value has been set.
func (o *InlineObject210) GetVdiPoolIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VdiPoolId, true
}

// SetVdiPoolId sets field value
func (o *InlineObject210) SetVdiPoolId(v int32) {
	o.VdiPoolId = v
}

// GetAccess returns the Access field value
func (o *InlineObject210) GetAccess() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *InlineObject210) GetAccessOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *InlineObject210) SetAccess(v string) {
	o.Access = v
}

func (o InlineObject210) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["vdiPoolId"] = o.VdiPoolId
	}
	if true {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject210 struct {
	value *InlineObject210
	isSet bool
}

func (v NullableInlineObject210) Get() *InlineObject210 {
	return v.value
}

func (v *NullableInlineObject210) Set(val *InlineObject210) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject210) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject210) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject210(val *InlineObject210) *NullableInlineObject210 {
	return &NullableInlineObject210{value: val, isSet: true}
}

func (v NullableInlineObject210) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject210) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

