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

// InlineObject243 struct for InlineObject243
type InlineObject243 struct {
	User UserCreate `json:"user"`
}

// NewInlineObject243 instantiates a new InlineObject243 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject243(user UserCreate, ) *InlineObject243 {
	this := InlineObject243{}
	this.User = user
	return &this
}

// NewInlineObject243WithDefaults instantiates a new InlineObject243 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject243WithDefaults() *InlineObject243 {
	this := InlineObject243{}
	return &this
}

// GetUser returns the User field value
func (o *InlineObject243) GetUser() UserCreate {
	if o == nil  {
		var ret UserCreate
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *InlineObject243) GetUserOk() (*UserCreate, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *InlineObject243) SetUser(v UserCreate) {
	o.User = v
}

func (o InlineObject243) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject243 struct {
	value *InlineObject243
	isSet bool
}

func (v NullableInlineObject243) Get() *InlineObject243 {
	return v.value
}

func (v *NullableInlineObject243) Set(val *InlineObject243) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject243) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject243) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject243(val *InlineObject243) *NullableInlineObject243 {
	return &NullableInlineObject243{value: val, isSet: true}
}

func (v NullableInlineObject243) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject243) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


