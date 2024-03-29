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

// InlineObject21 struct for InlineObject21
type InlineObject21 struct {
	ResourcePermission *ApiBlueprintsIdUpdatePermissionsResourcePermission `json:"resourcePermission,omitempty"`
}

// NewInlineObject21 instantiates a new InlineObject21 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject21() *InlineObject21 {
	this := InlineObject21{}
	return &this
}

// NewInlineObject21WithDefaults instantiates a new InlineObject21 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject21WithDefaults() *InlineObject21 {
	this := InlineObject21{}
	return &this
}

// GetResourcePermission returns the ResourcePermission field value if set, zero value otherwise.
func (o *InlineObject21) GetResourcePermission() ApiBlueprintsIdUpdatePermissionsResourcePermission {
	if o == nil || o.ResourcePermission == nil {
		var ret ApiBlueprintsIdUpdatePermissionsResourcePermission
		return ret
	}
	return *o.ResourcePermission
}

// GetResourcePermissionOk returns a tuple with the ResourcePermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject21) GetResourcePermissionOk() (*ApiBlueprintsIdUpdatePermissionsResourcePermission, bool) {
	if o == nil || o.ResourcePermission == nil {
		return nil, false
	}
	return o.ResourcePermission, true
}

// HasResourcePermission returns a boolean if a field has been set.
func (o *InlineObject21) HasResourcePermission() bool {
	if o != nil && o.ResourcePermission != nil {
		return true
	}

	return false
}

// SetResourcePermission gets a reference to the given ApiBlueprintsIdUpdatePermissionsResourcePermission and assigns it to the ResourcePermission field.
func (o *InlineObject21) SetResourcePermission(v ApiBlueprintsIdUpdatePermissionsResourcePermission) {
	o.ResourcePermission = &v
}

func (o InlineObject21) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourcePermission != nil {
		toSerialize["resourcePermission"] = o.ResourcePermission
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject21 struct {
	value *InlineObject21
	isSet bool
}

func (v NullableInlineObject21) Get() *InlineObject21 {
	return v.value
}

func (v *NullableInlineObject21) Set(val *InlineObject21) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject21) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject21) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject21(val *InlineObject21) *NullableInlineObject21 {
	return &NullableInlineObject21{value: val, isSet: true}
}

func (v NullableInlineObject21) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject21) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


