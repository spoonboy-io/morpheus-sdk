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

// InlineResponse200166 struct for InlineResponse200166
type InlineResponse200166 struct {
	WhitelabelSettings *WhitelabelSettings `json:"whitelabelSettings,omitempty"`
}

// NewInlineResponse200166 instantiates a new InlineResponse200166 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200166() *InlineResponse200166 {
	this := InlineResponse200166{}
	return &this
}

// NewInlineResponse200166WithDefaults instantiates a new InlineResponse200166 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200166WithDefaults() *InlineResponse200166 {
	this := InlineResponse200166{}
	return &this
}

// GetWhitelabelSettings returns the WhitelabelSettings field value if set, zero value otherwise.
func (o *InlineResponse200166) GetWhitelabelSettings() WhitelabelSettings {
	if o == nil || o.WhitelabelSettings == nil {
		var ret WhitelabelSettings
		return ret
	}
	return *o.WhitelabelSettings
}

// GetWhitelabelSettingsOk returns a tuple with the WhitelabelSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200166) GetWhitelabelSettingsOk() (*WhitelabelSettings, bool) {
	if o == nil || o.WhitelabelSettings == nil {
		return nil, false
	}
	return o.WhitelabelSettings, true
}

// HasWhitelabelSettings returns a boolean if a field has been set.
func (o *InlineResponse200166) HasWhitelabelSettings() bool {
	if o != nil && o.WhitelabelSettings != nil {
		return true
	}

	return false
}

// SetWhitelabelSettings gets a reference to the given WhitelabelSettings and assigns it to the WhitelabelSettings field.
func (o *InlineResponse200166) SetWhitelabelSettings(v WhitelabelSettings) {
	o.WhitelabelSettings = &v
}

func (o InlineResponse200166) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WhitelabelSettings != nil {
		toSerialize["whitelabelSettings"] = o.WhitelabelSettings
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200166 struct {
	value *InlineResponse200166
	isSet bool
}

func (v NullableInlineResponse200166) Get() *InlineResponse200166 {
	return v.value
}

func (v *NullableInlineResponse200166) Set(val *InlineResponse200166) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200166) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200166) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200166(val *InlineResponse200166) *NullableInlineResponse200166 {
	return &NullableInlineResponse200166{value: val, isSet: true}
}

func (v NullableInlineResponse200166) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200166) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


