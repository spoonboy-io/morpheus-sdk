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

// RolePermissionDefaultVDIPool struct for RolePermissionDefaultVDIPool
type RolePermissionDefaultVDIPool struct {
	// `VdiPools` is the code for Default VDI Pool Access
	PermissionCode string `json:"permissionCode"`
	// The new access level.
	Access string `json:"access"`
}

// NewRolePermissionDefaultVDIPool instantiates a new RolePermissionDefaultVDIPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolePermissionDefaultVDIPool(permissionCode string, access string, ) *RolePermissionDefaultVDIPool {
	this := RolePermissionDefaultVDIPool{}
	this.PermissionCode = permissionCode
	this.Access = access
	return &this
}

// NewRolePermissionDefaultVDIPoolWithDefaults instantiates a new RolePermissionDefaultVDIPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolePermissionDefaultVDIPoolWithDefaults() *RolePermissionDefaultVDIPool {
	this := RolePermissionDefaultVDIPool{}
	return &this
}

// GetPermissionCode returns the PermissionCode field value
func (o *RolePermissionDefaultVDIPool) GetPermissionCode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.PermissionCode
}

// GetPermissionCodeOk returns a tuple with the PermissionCode field value
// and a boolean to check if the value has been set.
func (o *RolePermissionDefaultVDIPool) GetPermissionCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PermissionCode, true
}

// SetPermissionCode sets field value
func (o *RolePermissionDefaultVDIPool) SetPermissionCode(v string) {
	o.PermissionCode = v
}

// GetAccess returns the Access field value
func (o *RolePermissionDefaultVDIPool) GetAccess() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *RolePermissionDefaultVDIPool) GetAccessOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *RolePermissionDefaultVDIPool) SetAccess(v string) {
	o.Access = v
}

func (o RolePermissionDefaultVDIPool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["permissionCode"] = o.PermissionCode
	}
	if true {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableRolePermissionDefaultVDIPool struct {
	value *RolePermissionDefaultVDIPool
	isSet bool
}

func (v NullableRolePermissionDefaultVDIPool) Get() *RolePermissionDefaultVDIPool {
	return v.value
}

func (v *NullableRolePermissionDefaultVDIPool) Set(val *RolePermissionDefaultVDIPool) {
	v.value = val
	v.isSet = true
}

func (v NullableRolePermissionDefaultVDIPool) IsSet() bool {
	return v.isSet
}

func (v *NullableRolePermissionDefaultVDIPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolePermissionDefaultVDIPool(val *RolePermissionDefaultVDIPool) *NullableRolePermissionDefaultVDIPool {
	return &NullableRolePermissionDefaultVDIPool{value: val, isSet: true}
}

func (v NullableRolePermissionDefaultVDIPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolePermissionDefaultVDIPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


