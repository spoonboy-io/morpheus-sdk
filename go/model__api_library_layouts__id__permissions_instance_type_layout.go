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

// ApiLibraryLayoutsIdPermissionsInstanceTypeLayout struct for ApiLibraryLayoutsIdPermissionsInstanceTypeLayout
type ApiLibraryLayoutsIdPermissionsInstanceTypeLayout struct {
	Permissions *ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions `json:"permissions,omitempty"`
}

// NewApiLibraryLayoutsIdPermissionsInstanceTypeLayout instantiates a new ApiLibraryLayoutsIdPermissionsInstanceTypeLayout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLibraryLayoutsIdPermissionsInstanceTypeLayout() *ApiLibraryLayoutsIdPermissionsInstanceTypeLayout {
	this := ApiLibraryLayoutsIdPermissionsInstanceTypeLayout{}
	return &this
}

// NewApiLibraryLayoutsIdPermissionsInstanceTypeLayoutWithDefaults instantiates a new ApiLibraryLayoutsIdPermissionsInstanceTypeLayout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLibraryLayoutsIdPermissionsInstanceTypeLayoutWithDefaults() *ApiLibraryLayoutsIdPermissionsInstanceTypeLayout {
	this := ApiLibraryLayoutsIdPermissionsInstanceTypeLayout{}
	return &this
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *ApiLibraryLayoutsIdPermissionsInstanceTypeLayout) GetPermissions() ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions {
	if o == nil || o.Permissions == nil {
		var ret ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLibraryLayoutsIdPermissionsInstanceTypeLayout) GetPermissionsOk() (*ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *ApiLibraryLayoutsIdPermissionsInstanceTypeLayout) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions and assigns it to the Permissions field.
func (o *ApiLibraryLayoutsIdPermissionsInstanceTypeLayout) SetPermissions(v ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions) {
	o.Permissions = &v
}

func (o ApiLibraryLayoutsIdPermissionsInstanceTypeLayout) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableApiLibraryLayoutsIdPermissionsInstanceTypeLayout struct {
	value *ApiLibraryLayoutsIdPermissionsInstanceTypeLayout
	isSet bool
}

func (v NullableApiLibraryLayoutsIdPermissionsInstanceTypeLayout) Get() *ApiLibraryLayoutsIdPermissionsInstanceTypeLayout {
	return v.value
}

func (v *NullableApiLibraryLayoutsIdPermissionsInstanceTypeLayout) Set(val *ApiLibraryLayoutsIdPermissionsInstanceTypeLayout) {
	v.value = val
	v.isSet = true
}

func (v NullableApiLibraryLayoutsIdPermissionsInstanceTypeLayout) IsSet() bool {
	return v.isSet
}

func (v *NullableApiLibraryLayoutsIdPermissionsInstanceTypeLayout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiLibraryLayoutsIdPermissionsInstanceTypeLayout(val *ApiLibraryLayoutsIdPermissionsInstanceTypeLayout) *NullableApiLibraryLayoutsIdPermissionsInstanceTypeLayout {
	return &NullableApiLibraryLayoutsIdPermissionsInstanceTypeLayout{value: val, isSet: true}
}

func (v NullableApiLibraryLayoutsIdPermissionsInstanceTypeLayout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiLibraryLayoutsIdPermissionsInstanceTypeLayout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

