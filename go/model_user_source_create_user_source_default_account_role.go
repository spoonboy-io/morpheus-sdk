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

// UserSourceCreateUserSourceDefaultAccountRole struct for UserSourceCreateUserSourceDefaultAccountRole
type UserSourceCreateUserSourceDefaultAccountRole struct {
	// Default `Role ID`
	Id int64 `json:"id"`
}

// NewUserSourceCreateUserSourceDefaultAccountRole instantiates a new UserSourceCreateUserSourceDefaultAccountRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSourceCreateUserSourceDefaultAccountRole(id int64, ) *UserSourceCreateUserSourceDefaultAccountRole {
	this := UserSourceCreateUserSourceDefaultAccountRole{}
	this.Id = id
	return &this
}

// NewUserSourceCreateUserSourceDefaultAccountRoleWithDefaults instantiates a new UserSourceCreateUserSourceDefaultAccountRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSourceCreateUserSourceDefaultAccountRoleWithDefaults() *UserSourceCreateUserSourceDefaultAccountRole {
	this := UserSourceCreateUserSourceDefaultAccountRole{}
	return &this
}

// GetId returns the Id field value
func (o *UserSourceCreateUserSourceDefaultAccountRole) GetId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserSourceCreateUserSourceDefaultAccountRole) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserSourceCreateUserSourceDefaultAccountRole) SetId(v int64) {
	o.Id = v
}

func (o UserSourceCreateUserSourceDefaultAccountRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableUserSourceCreateUserSourceDefaultAccountRole struct {
	value *UserSourceCreateUserSourceDefaultAccountRole
	isSet bool
}

func (v NullableUserSourceCreateUserSourceDefaultAccountRole) Get() *UserSourceCreateUserSourceDefaultAccountRole {
	return v.value
}

func (v *NullableUserSourceCreateUserSourceDefaultAccountRole) Set(val *UserSourceCreateUserSourceDefaultAccountRole) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSourceCreateUserSourceDefaultAccountRole) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSourceCreateUserSourceDefaultAccountRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSourceCreateUserSourceDefaultAccountRole(val *UserSourceCreateUserSourceDefaultAccountRole) *NullableUserSourceCreateUserSourceDefaultAccountRole {
	return &NullableUserSourceCreateUserSourceDefaultAccountRole{value: val, isSet: true}
}

func (v NullableUserSourceCreateUserSourceDefaultAccountRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSourceCreateUserSourceDefaultAccountRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


