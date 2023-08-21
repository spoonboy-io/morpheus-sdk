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

// InlineObject2 struct for InlineObject2
type InlineObject2 struct {
	ApplianceSettings *ApplianceSettingsUpdate `json:"applianceSettings,omitempty"`
}

// NewInlineObject2 instantiates a new InlineObject2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject2() *InlineObject2 {
	this := InlineObject2{}
	return &this
}

// NewInlineObject2WithDefaults instantiates a new InlineObject2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject2WithDefaults() *InlineObject2 {
	this := InlineObject2{}
	return &this
}

// GetApplianceSettings returns the ApplianceSettings field value if set, zero value otherwise.
func (o *InlineObject2) GetApplianceSettings() ApplianceSettingsUpdate {
	if o == nil || o.ApplianceSettings == nil {
		var ret ApplianceSettingsUpdate
		return ret
	}
	return *o.ApplianceSettings
}

// GetApplianceSettingsOk returns a tuple with the ApplianceSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetApplianceSettingsOk() (*ApplianceSettingsUpdate, bool) {
	if o == nil || o.ApplianceSettings == nil {
		return nil, false
	}
	return o.ApplianceSettings, true
}

// HasApplianceSettings returns a boolean if a field has been set.
func (o *InlineObject2) HasApplianceSettings() bool {
	if o != nil && o.ApplianceSettings != nil {
		return true
	}

	return false
}

// SetApplianceSettings gets a reference to the given ApplianceSettingsUpdate and assigns it to the ApplianceSettings field.
func (o *InlineObject2) SetApplianceSettings(v ApplianceSettingsUpdate) {
	o.ApplianceSettings = &v
}

func (o InlineObject2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplianceSettings != nil {
		toSerialize["applianceSettings"] = o.ApplianceSettings
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject2 struct {
	value *InlineObject2
	isSet bool
}

func (v NullableInlineObject2) Get() *InlineObject2 {
	return v.value
}

func (v *NullableInlineObject2) Set(val *InlineObject2) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject2) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject2(val *InlineObject2) *NullableInlineObject2 {
	return &NullableInlineObject2{value: val, isSet: true}
}

func (v NullableInlineObject2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


