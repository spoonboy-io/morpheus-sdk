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

// InlineResponse20032 struct for InlineResponse20032
type InlineResponse20032 struct {
	Versions *[]string `json:"versions,omitempty"`
	CurrentVersion *string `json:"currentVersion,omitempty"`
}

// NewInlineResponse20032 instantiates a new InlineResponse20032 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20032() *InlineResponse20032 {
	this := InlineResponse20032{}
	return &this
}

// NewInlineResponse20032WithDefaults instantiates a new InlineResponse20032 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20032WithDefaults() *InlineResponse20032 {
	this := InlineResponse20032{}
	return &this
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *InlineResponse20032) GetVersions() []string {
	if o == nil || o.Versions == nil {
		var ret []string
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20032) GetVersionsOk() (*[]string, bool) {
	if o == nil || o.Versions == nil {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *InlineResponse20032) HasVersions() bool {
	if o != nil && o.Versions != nil {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []string and assigns it to the Versions field.
func (o *InlineResponse20032) SetVersions(v []string) {
	o.Versions = &v
}

// GetCurrentVersion returns the CurrentVersion field value if set, zero value otherwise.
func (o *InlineResponse20032) GetCurrentVersion() string {
	if o == nil || o.CurrentVersion == nil {
		var ret string
		return ret
	}
	return *o.CurrentVersion
}

// GetCurrentVersionOk returns a tuple with the CurrentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20032) GetCurrentVersionOk() (*string, bool) {
	if o == nil || o.CurrentVersion == nil {
		return nil, false
	}
	return o.CurrentVersion, true
}

// HasCurrentVersion returns a boolean if a field has been set.
func (o *InlineResponse20032) HasCurrentVersion() bool {
	if o != nil && o.CurrentVersion != nil {
		return true
	}

	return false
}

// SetCurrentVersion gets a reference to the given string and assigns it to the CurrentVersion field.
func (o *InlineResponse20032) SetCurrentVersion(v string) {
	o.CurrentVersion = &v
}

func (o InlineResponse20032) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Versions != nil {
		toSerialize["versions"] = o.Versions
	}
	if o.CurrentVersion != nil {
		toSerialize["currentVersion"] = o.CurrentVersion
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20032 struct {
	value *InlineResponse20032
	isSet bool
}

func (v NullableInlineResponse20032) Get() *InlineResponse20032 {
	return v.value
}

func (v *NullableInlineResponse20032) Set(val *InlineResponse20032) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20032) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20032) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20032(val *InlineResponse20032) *NullableInlineResponse20032 {
	return &NullableInlineResponse20032{value: val, isSet: true}
}

func (v NullableInlineResponse20032) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20032) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

