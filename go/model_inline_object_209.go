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

// InlineObject209 struct for InlineObject209
type InlineObject209 struct {
	Role ApiRolesIdRole `json:"role"`
}

// NewInlineObject209 instantiates a new InlineObject209 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject209(role ApiRolesIdRole, ) *InlineObject209 {
	this := InlineObject209{}
	this.Role = role
	return &this
}

// NewInlineObject209WithDefaults instantiates a new InlineObject209 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject209WithDefaults() *InlineObject209 {
	this := InlineObject209{}
	return &this
}

// GetRole returns the Role field value
func (o *InlineObject209) GetRole() ApiRolesIdRole {
	if o == nil  {
		var ret ApiRolesIdRole
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *InlineObject209) GetRoleOk() (*ApiRolesIdRole, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *InlineObject209) SetRole(v ApiRolesIdRole) {
	o.Role = v
}

func (o InlineObject209) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject209 struct {
	value *InlineObject209
	isSet bool
}

func (v NullableInlineObject209) Get() *InlineObject209 {
	return v.value
}

func (v *NullableInlineObject209) Set(val *InlineObject209) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject209) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject209) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject209(val *InlineObject209) *NullableInlineObject209 {
	return &NullableInlineObject209{value: val, isSet: true}
}

func (v NullableInlineObject209) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject209) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


