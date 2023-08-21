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

// UserSourceCreate struct for UserSourceCreate
type UserSourceCreate struct {
	UserSource UserSourceCreateUserSource `json:"userSource"`
}

// NewUserSourceCreate instantiates a new UserSourceCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSourceCreate(userSource UserSourceCreateUserSource, ) *UserSourceCreate {
	this := UserSourceCreate{}
	this.UserSource = userSource
	return &this
}

// NewUserSourceCreateWithDefaults instantiates a new UserSourceCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSourceCreateWithDefaults() *UserSourceCreate {
	this := UserSourceCreate{}
	return &this
}

// GetUserSource returns the UserSource field value
func (o *UserSourceCreate) GetUserSource() UserSourceCreateUserSource {
	if o == nil  {
		var ret UserSourceCreateUserSource
		return ret
	}

	return o.UserSource
}

// GetUserSourceOk returns a tuple with the UserSource field value
// and a boolean to check if the value has been set.
func (o *UserSourceCreate) GetUserSourceOk() (*UserSourceCreateUserSource, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserSource, true
}

// SetUserSource sets field value
func (o *UserSourceCreate) SetUserSource(v UserSourceCreateUserSource) {
	o.UserSource = v
}

func (o UserSourceCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["userSource"] = o.UserSource
	}
	return json.Marshal(toSerialize)
}

type NullableUserSourceCreate struct {
	value *UserSourceCreate
	isSet bool
}

func (v NullableUserSourceCreate) Get() *UserSourceCreate {
	return v.value
}

func (v *NullableUserSourceCreate) Set(val *UserSourceCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSourceCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSourceCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSourceCreate(val *UserSourceCreate) *NullableUserSourceCreate {
	return &NullableUserSourceCreate{value: val, isSet: true}
}

func (v NullableUserSourceCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSourceCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


