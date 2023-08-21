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

// RolePermissionGroupAll struct for RolePermissionGroupAll
type RolePermissionGroupAll struct {
	// Apply to all groups (site)
	AllGroups bool `json:"allGroups"`
	// The new access level.
	Access string `json:"access"`
}

// NewRolePermissionGroupAll instantiates a new RolePermissionGroupAll object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolePermissionGroupAll(allGroups bool, access string, ) *RolePermissionGroupAll {
	this := RolePermissionGroupAll{}
	this.AllGroups = allGroups
	this.Access = access
	return &this
}

// NewRolePermissionGroupAllWithDefaults instantiates a new RolePermissionGroupAll object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolePermissionGroupAllWithDefaults() *RolePermissionGroupAll {
	this := RolePermissionGroupAll{}
	return &this
}

// GetAllGroups returns the AllGroups field value
func (o *RolePermissionGroupAll) GetAllGroups() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.AllGroups
}

// GetAllGroupsOk returns a tuple with the AllGroups field value
// and a boolean to check if the value has been set.
func (o *RolePermissionGroupAll) GetAllGroupsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AllGroups, true
}

// SetAllGroups sets field value
func (o *RolePermissionGroupAll) SetAllGroups(v bool) {
	o.AllGroups = v
}

// GetAccess returns the Access field value
func (o *RolePermissionGroupAll) GetAccess() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *RolePermissionGroupAll) GetAccessOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *RolePermissionGroupAll) SetAccess(v string) {
	o.Access = v
}

func (o RolePermissionGroupAll) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["allGroups"] = o.AllGroups
	}
	if true {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableRolePermissionGroupAll struct {
	value *RolePermissionGroupAll
	isSet bool
}

func (v NullableRolePermissionGroupAll) Get() *RolePermissionGroupAll {
	return v.value
}

func (v *NullableRolePermissionGroupAll) Set(val *RolePermissionGroupAll) {
	v.value = val
	v.isSet = true
}

func (v NullableRolePermissionGroupAll) IsSet() bool {
	return v.isSet
}

func (v *NullableRolePermissionGroupAll) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolePermissionGroupAll(val *RolePermissionGroupAll) *NullableRolePermissionGroupAll {
	return &NullableRolePermissionGroupAll{value: val, isSet: true}
}

func (v NullableRolePermissionGroupAll) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolePermissionGroupAll) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

