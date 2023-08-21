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

// RolePermissionDefaultPersona struct for RolePermissionDefaultPersona
type RolePermissionDefaultPersona struct {
	// `Persona` is the code for Default Persona Access
	PermissionCode string `json:"permissionCode"`
	// The new access level.
	Access string `json:"access"`
}

// NewRolePermissionDefaultPersona instantiates a new RolePermissionDefaultPersona object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolePermissionDefaultPersona(permissionCode string, access string, ) *RolePermissionDefaultPersona {
	this := RolePermissionDefaultPersona{}
	this.PermissionCode = permissionCode
	this.Access = access
	return &this
}

// NewRolePermissionDefaultPersonaWithDefaults instantiates a new RolePermissionDefaultPersona object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolePermissionDefaultPersonaWithDefaults() *RolePermissionDefaultPersona {
	this := RolePermissionDefaultPersona{}
	return &this
}

// GetPermissionCode returns the PermissionCode field value
func (o *RolePermissionDefaultPersona) GetPermissionCode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.PermissionCode
}

// GetPermissionCodeOk returns a tuple with the PermissionCode field value
// and a boolean to check if the value has been set.
func (o *RolePermissionDefaultPersona) GetPermissionCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PermissionCode, true
}

// SetPermissionCode sets field value
func (o *RolePermissionDefaultPersona) SetPermissionCode(v string) {
	o.PermissionCode = v
}

// GetAccess returns the Access field value
func (o *RolePermissionDefaultPersona) GetAccess() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *RolePermissionDefaultPersona) GetAccessOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *RolePermissionDefaultPersona) SetAccess(v string) {
	o.Access = v
}

func (o RolePermissionDefaultPersona) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["permissionCode"] = o.PermissionCode
	}
	if true {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableRolePermissionDefaultPersona struct {
	value *RolePermissionDefaultPersona
	isSet bool
}

func (v NullableRolePermissionDefaultPersona) Get() *RolePermissionDefaultPersona {
	return v.value
}

func (v *NullableRolePermissionDefaultPersona) Set(val *RolePermissionDefaultPersona) {
	v.value = val
	v.isSet = true
}

func (v NullableRolePermissionDefaultPersona) IsSet() bool {
	return v.isSet
}

func (v *NullableRolePermissionDefaultPersona) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolePermissionDefaultPersona(val *RolePermissionDefaultPersona) *NullableRolePermissionDefaultPersona {
	return &NullableRolePermissionDefaultPersona{value: val, isSet: true}
}

func (v NullableRolePermissionDefaultPersona) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolePermissionDefaultPersona) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


