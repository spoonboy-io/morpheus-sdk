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

// RolePermissionReportTypeAll struct for RolePermissionReportTypeAll
type RolePermissionReportTypeAll struct {
	// Apply to all report types
	AllReportTypes bool `json:"allReportTypes"`
	// The new access level.
	Access string `json:"access"`
}

// NewRolePermissionReportTypeAll instantiates a new RolePermissionReportTypeAll object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolePermissionReportTypeAll(allReportTypes bool, access string, ) *RolePermissionReportTypeAll {
	this := RolePermissionReportTypeAll{}
	this.AllReportTypes = allReportTypes
	this.Access = access
	return &this
}

// NewRolePermissionReportTypeAllWithDefaults instantiates a new RolePermissionReportTypeAll object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolePermissionReportTypeAllWithDefaults() *RolePermissionReportTypeAll {
	this := RolePermissionReportTypeAll{}
	return &this
}

// GetAllReportTypes returns the AllReportTypes field value
func (o *RolePermissionReportTypeAll) GetAllReportTypes() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.AllReportTypes
}

// GetAllReportTypesOk returns a tuple with the AllReportTypes field value
// and a boolean to check if the value has been set.
func (o *RolePermissionReportTypeAll) GetAllReportTypesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AllReportTypes, true
}

// SetAllReportTypes sets field value
func (o *RolePermissionReportTypeAll) SetAllReportTypes(v bool) {
	o.AllReportTypes = v
}

// GetAccess returns the Access field value
func (o *RolePermissionReportTypeAll) GetAccess() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *RolePermissionReportTypeAll) GetAccessOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *RolePermissionReportTypeAll) SetAccess(v string) {
	o.Access = v
}

func (o RolePermissionReportTypeAll) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["allReportTypes"] = o.AllReportTypes
	}
	if true {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableRolePermissionReportTypeAll struct {
	value *RolePermissionReportTypeAll
	isSet bool
}

func (v NullableRolePermissionReportTypeAll) Get() *RolePermissionReportTypeAll {
	return v.value
}

func (v *NullableRolePermissionReportTypeAll) Set(val *RolePermissionReportTypeAll) {
	v.value = val
	v.isSet = true
}

func (v NullableRolePermissionReportTypeAll) IsSet() bool {
	return v.isSet
}

func (v *NullableRolePermissionReportTypeAll) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolePermissionReportTypeAll(val *RolePermissionReportTypeAll) *NullableRolePermissionReportTypeAll {
	return &NullableRolePermissionReportTypeAll{value: val, isSet: true}
}

func (v NullableRolePermissionReportTypeAll) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolePermissionReportTypeAll) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

