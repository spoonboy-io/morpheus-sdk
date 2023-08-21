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

// InlineObject176 struct for InlineObject176
type InlineObject176 struct {
	Group *NetworkServerGroupCreate `json:"group,omitempty"`
}

// NewInlineObject176 instantiates a new InlineObject176 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject176() *InlineObject176 {
	this := InlineObject176{}
	return &this
}

// NewInlineObject176WithDefaults instantiates a new InlineObject176 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject176WithDefaults() *InlineObject176 {
	this := InlineObject176{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *InlineObject176) GetGroup() NetworkServerGroupCreate {
	if o == nil || o.Group == nil {
		var ret NetworkServerGroupCreate
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject176) GetGroupOk() (*NetworkServerGroupCreate, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *InlineObject176) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NetworkServerGroupCreate and assigns it to the Group field.
func (o *InlineObject176) SetGroup(v NetworkServerGroupCreate) {
	o.Group = &v
}

func (o InlineObject176) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject176 struct {
	value *InlineObject176
	isSet bool
}

func (v NullableInlineObject176) Get() *InlineObject176 {
	return v.value
}

func (v *NullableInlineObject176) Set(val *InlineObject176) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject176) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject176) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject176(val *InlineObject176) *NullableInlineObject176 {
	return &NullableInlineObject176{value: val, isSet: true}
}

func (v NullableInlineObject176) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject176) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

