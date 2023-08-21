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

// InlineObject158 struct for InlineObject158
type InlineObject158 struct {
	Permissions *NetworkRouterPermissionsUpdate `json:"permissions,omitempty"`
}

// NewInlineObject158 instantiates a new InlineObject158 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject158() *InlineObject158 {
	this := InlineObject158{}
	return &this
}

// NewInlineObject158WithDefaults instantiates a new InlineObject158 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject158WithDefaults() *InlineObject158 {
	this := InlineObject158{}
	return &this
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *InlineObject158) GetPermissions() NetworkRouterPermissionsUpdate {
	if o == nil || o.Permissions == nil {
		var ret NetworkRouterPermissionsUpdate
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject158) GetPermissionsOk() (*NetworkRouterPermissionsUpdate, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *InlineObject158) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given NetworkRouterPermissionsUpdate and assigns it to the Permissions field.
func (o *InlineObject158) SetPermissions(v NetworkRouterPermissionsUpdate) {
	o.Permissions = &v
}

func (o InlineObject158) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject158 struct {
	value *InlineObject158
	isSet bool
}

func (v NullableInlineObject158) Get() *InlineObject158 {
	return v.value
}

func (v *NullableInlineObject158) Set(val *InlineObject158) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject158) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject158) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject158(val *InlineObject158) *NullableInlineObject158 {
	return &NullableInlineObject158{value: val, isSet: true}
}

func (v NullableInlineObject158) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject158) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


