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

// RolePermissionInstanceTypeAll struct for RolePermissionInstanceTypeAll
type RolePermissionInstanceTypeAll struct {
	// Apply to all instance types
	AllInstanceTypes *bool `json:"allInstanceTypes,omitempty"`
	// The new access level.
	Access string `json:"access"`
}

// NewRolePermissionInstanceTypeAll instantiates a new RolePermissionInstanceTypeAll object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolePermissionInstanceTypeAll(access string, ) *RolePermissionInstanceTypeAll {
	this := RolePermissionInstanceTypeAll{}
	this.Access = access
	return &this
}

// NewRolePermissionInstanceTypeAllWithDefaults instantiates a new RolePermissionInstanceTypeAll object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolePermissionInstanceTypeAllWithDefaults() *RolePermissionInstanceTypeAll {
	this := RolePermissionInstanceTypeAll{}
	return &this
}

// GetAllInstanceTypes returns the AllInstanceTypes field value if set, zero value otherwise.
func (o *RolePermissionInstanceTypeAll) GetAllInstanceTypes() bool {
	if o == nil || o.AllInstanceTypes == nil {
		var ret bool
		return ret
	}
	return *o.AllInstanceTypes
}

// GetAllInstanceTypesOk returns a tuple with the AllInstanceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolePermissionInstanceTypeAll) GetAllInstanceTypesOk() (*bool, bool) {
	if o == nil || o.AllInstanceTypes == nil {
		return nil, false
	}
	return o.AllInstanceTypes, true
}

// HasAllInstanceTypes returns a boolean if a field has been set.
func (o *RolePermissionInstanceTypeAll) HasAllInstanceTypes() bool {
	if o != nil && o.AllInstanceTypes != nil {
		return true
	}

	return false
}

// SetAllInstanceTypes gets a reference to the given bool and assigns it to the AllInstanceTypes field.
func (o *RolePermissionInstanceTypeAll) SetAllInstanceTypes(v bool) {
	o.AllInstanceTypes = &v
}

// GetAccess returns the Access field value
func (o *RolePermissionInstanceTypeAll) GetAccess() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *RolePermissionInstanceTypeAll) GetAccessOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *RolePermissionInstanceTypeAll) SetAccess(v string) {
	o.Access = v
}

func (o RolePermissionInstanceTypeAll) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllInstanceTypes != nil {
		toSerialize["allInstanceTypes"] = o.AllInstanceTypes
	}
	if true {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableRolePermissionInstanceTypeAll struct {
	value *RolePermissionInstanceTypeAll
	isSet bool
}

func (v NullableRolePermissionInstanceTypeAll) Get() *RolePermissionInstanceTypeAll {
	return v.value
}

func (v *NullableRolePermissionInstanceTypeAll) Set(val *RolePermissionInstanceTypeAll) {
	v.value = val
	v.isSet = true
}

func (v NullableRolePermissionInstanceTypeAll) IsSet() bool {
	return v.isSet
}

func (v *NullableRolePermissionInstanceTypeAll) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolePermissionInstanceTypeAll(val *RolePermissionInstanceTypeAll) *NullableRolePermissionInstanceTypeAll {
	return &NullableRolePermissionInstanceTypeAll{value: val, isSet: true}
}

func (v NullableRolePermissionInstanceTypeAll) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolePermissionInstanceTypeAll) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

