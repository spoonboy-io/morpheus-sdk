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

// InlineObject126 struct for InlineObject126
type InlineObject126 struct {
	// License Key. This is a long and unique string of your Morpheus license. License keys can be requested from the [Morpheus Hub](https://morpheushub.com).
	License string `json:"license"`
}

// NewInlineObject126 instantiates a new InlineObject126 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject126(license string, ) *InlineObject126 {
	this := InlineObject126{}
	this.License = license
	return &this
}

// NewInlineObject126WithDefaults instantiates a new InlineObject126 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject126WithDefaults() *InlineObject126 {
	this := InlineObject126{}
	return &this
}

// GetLicense returns the License field value
func (o *InlineObject126) GetLicense() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.License
}

// GetLicenseOk returns a tuple with the License field value
// and a boolean to check if the value has been set.
func (o *InlineObject126) GetLicenseOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.License, true
}

// SetLicense sets field value
func (o *InlineObject126) SetLicense(v string) {
	o.License = v
}

func (o InlineObject126) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["license"] = o.License
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject126 struct {
	value *InlineObject126
	isSet bool
}

func (v NullableInlineObject126) Get() *InlineObject126 {
	return v.value
}

func (v *NullableInlineObject126) Set(val *InlineObject126) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject126) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject126) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject126(val *InlineObject126) *NullableInlineObject126 {
	return &NullableInlineObject126{value: val, isSet: true}
}

func (v NullableInlineObject126) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject126) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


