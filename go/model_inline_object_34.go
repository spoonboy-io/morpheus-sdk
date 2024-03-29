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

// InlineObject34 struct for InlineObject34
type InlineObject34 struct {
	// Set to false to unmute
	Muted bool `json:"muted"`
}

// NewInlineObject34 instantiates a new InlineObject34 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject34(muted bool, ) *InlineObject34 {
	this := InlineObject34{}
	this.Muted = muted
	return &this
}

// NewInlineObject34WithDefaults instantiates a new InlineObject34 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject34WithDefaults() *InlineObject34 {
	this := InlineObject34{}
	var muted bool = true
	this.Muted = muted
	return &this
}

// GetMuted returns the Muted field value
func (o *InlineObject34) GetMuted() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Muted
}

// GetMutedOk returns a tuple with the Muted field value
// and a boolean to check if the value has been set.
func (o *InlineObject34) GetMutedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Muted, true
}

// SetMuted sets field value
func (o *InlineObject34) SetMuted(v bool) {
	o.Muted = v
}

func (o InlineObject34) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["muted"] = o.Muted
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject34 struct {
	value *InlineObject34
	isSet bool
}

func (v NullableInlineObject34) Get() *InlineObject34 {
	return v.value
}

func (v *NullableInlineObject34) Set(val *InlineObject34) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject34) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject34) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject34(val *InlineObject34) *NullableInlineObject34 {
	return &NullableInlineObject34{value: val, isSet: true}
}

func (v NullableInlineObject34) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject34) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


