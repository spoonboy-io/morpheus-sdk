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

// InlineResponse20047 struct for InlineResponse20047
type InlineResponse20047 struct {
	GuidanceSettings *GuidanceSettings `json:"guidanceSettings,omitempty"`
}

// NewInlineResponse20047 instantiates a new InlineResponse20047 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20047() *InlineResponse20047 {
	this := InlineResponse20047{}
	return &this
}

// NewInlineResponse20047WithDefaults instantiates a new InlineResponse20047 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20047WithDefaults() *InlineResponse20047 {
	this := InlineResponse20047{}
	return &this
}

// GetGuidanceSettings returns the GuidanceSettings field value if set, zero value otherwise.
func (o *InlineResponse20047) GetGuidanceSettings() GuidanceSettings {
	if o == nil || o.GuidanceSettings == nil {
		var ret GuidanceSettings
		return ret
	}
	return *o.GuidanceSettings
}

// GetGuidanceSettingsOk returns a tuple with the GuidanceSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20047) GetGuidanceSettingsOk() (*GuidanceSettings, bool) {
	if o == nil || o.GuidanceSettings == nil {
		return nil, false
	}
	return o.GuidanceSettings, true
}

// HasGuidanceSettings returns a boolean if a field has been set.
func (o *InlineResponse20047) HasGuidanceSettings() bool {
	if o != nil && o.GuidanceSettings != nil {
		return true
	}

	return false
}

// SetGuidanceSettings gets a reference to the given GuidanceSettings and assigns it to the GuidanceSettings field.
func (o *InlineResponse20047) SetGuidanceSettings(v GuidanceSettings) {
	o.GuidanceSettings = &v
}

func (o InlineResponse20047) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GuidanceSettings != nil {
		toSerialize["guidanceSettings"] = o.GuidanceSettings
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20047 struct {
	value *InlineResponse20047
	isSet bool
}

func (v NullableInlineResponse20047) Get() *InlineResponse20047 {
	return v.value
}

func (v *NullableInlineResponse20047) Set(val *InlineResponse20047) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20047) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20047) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20047(val *InlineResponse20047) *NullableInlineResponse20047 {
	return &NullableInlineResponse20047{value: val, isSet: true}
}

func (v NullableInlineResponse20047) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20047) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

