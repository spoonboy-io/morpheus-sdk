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

// InlineObject79 struct for InlineObject79
type InlineObject79 struct {
	GuidanceSettings *GuidanceSettings `json:"guidanceSettings,omitempty"`
}

// NewInlineObject79 instantiates a new InlineObject79 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject79() *InlineObject79 {
	this := InlineObject79{}
	return &this
}

// NewInlineObject79WithDefaults instantiates a new InlineObject79 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject79WithDefaults() *InlineObject79 {
	this := InlineObject79{}
	return &this
}

// GetGuidanceSettings returns the GuidanceSettings field value if set, zero value otherwise.
func (o *InlineObject79) GetGuidanceSettings() GuidanceSettings {
	if o == nil || o.GuidanceSettings == nil {
		var ret GuidanceSettings
		return ret
	}
	return *o.GuidanceSettings
}

// GetGuidanceSettingsOk returns a tuple with the GuidanceSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject79) GetGuidanceSettingsOk() (*GuidanceSettings, bool) {
	if o == nil || o.GuidanceSettings == nil {
		return nil, false
	}
	return o.GuidanceSettings, true
}

// HasGuidanceSettings returns a boolean if a field has been set.
func (o *InlineObject79) HasGuidanceSettings() bool {
	if o != nil && o.GuidanceSettings != nil {
		return true
	}

	return false
}

// SetGuidanceSettings gets a reference to the given GuidanceSettings and assigns it to the GuidanceSettings field.
func (o *InlineObject79) SetGuidanceSettings(v GuidanceSettings) {
	o.GuidanceSettings = &v
}

func (o InlineObject79) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GuidanceSettings != nil {
		toSerialize["guidanceSettings"] = o.GuidanceSettings
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject79 struct {
	value *InlineObject79
	isSet bool
}

func (v NullableInlineObject79) Get() *InlineObject79 {
	return v.value
}

func (v *NullableInlineObject79) Set(val *InlineObject79) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject79) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject79) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject79(val *InlineObject79) *NullableInlineObject79 {
	return &NullableInlineObject79{value: val, isSet: true}
}

func (v NullableInlineObject79) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject79) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


