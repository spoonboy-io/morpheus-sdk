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

// InlineObject122 struct for InlineObject122
type InlineObject122 struct {
	OptionType *OptionTypeUpdate `json:"optionType,omitempty"`
}

// NewInlineObject122 instantiates a new InlineObject122 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject122() *InlineObject122 {
	this := InlineObject122{}
	return &this
}

// NewInlineObject122WithDefaults instantiates a new InlineObject122 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject122WithDefaults() *InlineObject122 {
	this := InlineObject122{}
	return &this
}

// GetOptionType returns the OptionType field value if set, zero value otherwise.
func (o *InlineObject122) GetOptionType() OptionTypeUpdate {
	if o == nil || o.OptionType == nil {
		var ret OptionTypeUpdate
		return ret
	}
	return *o.OptionType
}

// GetOptionTypeOk returns a tuple with the OptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject122) GetOptionTypeOk() (*OptionTypeUpdate, bool) {
	if o == nil || o.OptionType == nil {
		return nil, false
	}
	return o.OptionType, true
}

// HasOptionType returns a boolean if a field has been set.
func (o *InlineObject122) HasOptionType() bool {
	if o != nil && o.OptionType != nil {
		return true
	}

	return false
}

// SetOptionType gets a reference to the given OptionTypeUpdate and assigns it to the OptionType field.
func (o *InlineObject122) SetOptionType(v OptionTypeUpdate) {
	o.OptionType = &v
}

func (o InlineObject122) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OptionType != nil {
		toSerialize["optionType"] = o.OptionType
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject122 struct {
	value *InlineObject122
	isSet bool
}

func (v NullableInlineObject122) Get() *InlineObject122 {
	return v.value
}

func (v *NullableInlineObject122) Set(val *InlineObject122) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject122) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject122) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject122(val *InlineObject122) *NullableInlineObject122 {
	return &NullableInlineObject122{value: val, isSet: true}
}

func (v NullableInlineObject122) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject122) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


