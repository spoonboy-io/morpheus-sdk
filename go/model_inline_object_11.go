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

// InlineObject11 struct for InlineObject11
type InlineObject11 struct {
	// Specified as \"username\" or \"tenantId\\username\" with proper HTML encoding and used in conjunction with `password`. 
	Username string `json:"username"`
}

// NewInlineObject11 instantiates a new InlineObject11 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject11(username string, ) *InlineObject11 {
	this := InlineObject11{}
	this.Username = username
	return &this
}

// NewInlineObject11WithDefaults instantiates a new InlineObject11 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject11WithDefaults() *InlineObject11 {
	this := InlineObject11{}
	return &this
}

// GetUsername returns the Username field value
func (o *InlineObject11) GetUsername() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *InlineObject11) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *InlineObject11) SetUsername(v string) {
	o.Username = v
}

func (o InlineObject11) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject11 struct {
	value *InlineObject11
	isSet bool
}

func (v NullableInlineObject11) Get() *InlineObject11 {
	return v.value
}

func (v *NullableInlineObject11) Set(val *InlineObject11) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject11) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject11) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject11(val *InlineObject11) *NullableInlineObject11 {
	return &NullableInlineObject11{value: val, isSet: true}
}

func (v NullableInlineObject11) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject11) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


