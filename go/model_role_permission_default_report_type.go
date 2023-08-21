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

// RolePermissionDefaultReportType struct for RolePermissionDefaultReportType
type RolePermissionDefaultReportType struct {
	// `ReportTypes` is the code for Default Report Type Access
	PermissionCode string `json:"permissionCode"`
	// The new access level.
	Access string `json:"access"`
}

// NewRolePermissionDefaultReportType instantiates a new RolePermissionDefaultReportType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolePermissionDefaultReportType(permissionCode string, access string, ) *RolePermissionDefaultReportType {
	this := RolePermissionDefaultReportType{}
	this.PermissionCode = permissionCode
	this.Access = access
	return &this
}

// NewRolePermissionDefaultReportTypeWithDefaults instantiates a new RolePermissionDefaultReportType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolePermissionDefaultReportTypeWithDefaults() *RolePermissionDefaultReportType {
	this := RolePermissionDefaultReportType{}
	return &this
}

// GetPermissionCode returns the PermissionCode field value
func (o *RolePermissionDefaultReportType) GetPermissionCode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.PermissionCode
}

// GetPermissionCodeOk returns a tuple with the PermissionCode field value
// and a boolean to check if the value has been set.
func (o *RolePermissionDefaultReportType) GetPermissionCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PermissionCode, true
}

// SetPermissionCode sets field value
func (o *RolePermissionDefaultReportType) SetPermissionCode(v string) {
	o.PermissionCode = v
}

// GetAccess returns the Access field value
func (o *RolePermissionDefaultReportType) GetAccess() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *RolePermissionDefaultReportType) GetAccessOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *RolePermissionDefaultReportType) SetAccess(v string) {
	o.Access = v
}

func (o RolePermissionDefaultReportType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["permissionCode"] = o.PermissionCode
	}
	if true {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableRolePermissionDefaultReportType struct {
	value *RolePermissionDefaultReportType
	isSet bool
}

func (v NullableRolePermissionDefaultReportType) Get() *RolePermissionDefaultReportType {
	return v.value
}

func (v *NullableRolePermissionDefaultReportType) Set(val *RolePermissionDefaultReportType) {
	v.value = val
	v.isSet = true
}

func (v NullableRolePermissionDefaultReportType) IsSet() bool {
	return v.isSet
}

func (v *NullableRolePermissionDefaultReportType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolePermissionDefaultReportType(val *RolePermissionDefaultReportType) *NullableRolePermissionDefaultReportType {
	return &NullableRolePermissionDefaultReportType{value: val, isSet: true}
}

func (v NullableRolePermissionDefaultReportType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolePermissionDefaultReportType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

