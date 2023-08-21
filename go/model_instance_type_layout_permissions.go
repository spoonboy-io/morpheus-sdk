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

// InstanceTypeLayoutPermissions struct for InstanceTypeLayoutPermissions
type InstanceTypeLayoutPermissions struct {
	ResourcePermissions *InstanceTypeLayoutPermissionsResourcePermissions `json:"resourcePermissions,omitempty"`
}

// NewInstanceTypeLayoutPermissions instantiates a new InstanceTypeLayoutPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceTypeLayoutPermissions() *InstanceTypeLayoutPermissions {
	this := InstanceTypeLayoutPermissions{}
	return &this
}

// NewInstanceTypeLayoutPermissionsWithDefaults instantiates a new InstanceTypeLayoutPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceTypeLayoutPermissionsWithDefaults() *InstanceTypeLayoutPermissions {
	this := InstanceTypeLayoutPermissions{}
	return &this
}

// GetResourcePermissions returns the ResourcePermissions field value if set, zero value otherwise.
func (o *InstanceTypeLayoutPermissions) GetResourcePermissions() InstanceTypeLayoutPermissionsResourcePermissions {
	if o == nil || o.ResourcePermissions == nil {
		var ret InstanceTypeLayoutPermissionsResourcePermissions
		return ret
	}
	return *o.ResourcePermissions
}

// GetResourcePermissionsOk returns a tuple with the ResourcePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeLayoutPermissions) GetResourcePermissionsOk() (*InstanceTypeLayoutPermissionsResourcePermissions, bool) {
	if o == nil || o.ResourcePermissions == nil {
		return nil, false
	}
	return o.ResourcePermissions, true
}

// HasResourcePermissions returns a boolean if a field has been set.
func (o *InstanceTypeLayoutPermissions) HasResourcePermissions() bool {
	if o != nil && o.ResourcePermissions != nil {
		return true
	}

	return false
}

// SetResourcePermissions gets a reference to the given InstanceTypeLayoutPermissionsResourcePermissions and assigns it to the ResourcePermissions field.
func (o *InstanceTypeLayoutPermissions) SetResourcePermissions(v InstanceTypeLayoutPermissionsResourcePermissions) {
	o.ResourcePermissions = &v
}

func (o InstanceTypeLayoutPermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourcePermissions != nil {
		toSerialize["resourcePermissions"] = o.ResourcePermissions
	}
	return json.Marshal(toSerialize)
}

type NullableInstanceTypeLayoutPermissions struct {
	value *InstanceTypeLayoutPermissions
	isSet bool
}

func (v NullableInstanceTypeLayoutPermissions) Get() *InstanceTypeLayoutPermissions {
	return v.value
}

func (v *NullableInstanceTypeLayoutPermissions) Set(val *InstanceTypeLayoutPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceTypeLayoutPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceTypeLayoutPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceTypeLayoutPermissions(val *InstanceTypeLayoutPermissions) *NullableInstanceTypeLayoutPermissions {
	return &NullableInstanceTypeLayoutPermissions{value: val, isSet: true}
}

func (v NullableInstanceTypeLayoutPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceTypeLayoutPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

