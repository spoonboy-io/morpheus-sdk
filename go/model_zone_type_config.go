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

// ZoneTypeConfig struct for ZoneTypeConfig
type ZoneTypeConfig struct {
	CredentialTypes *[]string `json:"credentialTypes,omitempty"`
}

// NewZoneTypeConfig instantiates a new ZoneTypeConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneTypeConfig() *ZoneTypeConfig {
	this := ZoneTypeConfig{}
	return &this
}

// NewZoneTypeConfigWithDefaults instantiates a new ZoneTypeConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneTypeConfigWithDefaults() *ZoneTypeConfig {
	this := ZoneTypeConfig{}
	return &this
}

// GetCredentialTypes returns the CredentialTypes field value if set, zero value otherwise.
func (o *ZoneTypeConfig) GetCredentialTypes() []string {
	if o == nil || o.CredentialTypes == nil {
		var ret []string
		return ret
	}
	return *o.CredentialTypes
}

// GetCredentialTypesOk returns a tuple with the CredentialTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneTypeConfig) GetCredentialTypesOk() (*[]string, bool) {
	if o == nil || o.CredentialTypes == nil {
		return nil, false
	}
	return o.CredentialTypes, true
}

// HasCredentialTypes returns a boolean if a field has been set.
func (o *ZoneTypeConfig) HasCredentialTypes() bool {
	if o != nil && o.CredentialTypes != nil {
		return true
	}

	return false
}

// SetCredentialTypes gets a reference to the given []string and assigns it to the CredentialTypes field.
func (o *ZoneTypeConfig) SetCredentialTypes(v []string) {
	o.CredentialTypes = &v
}

func (o ZoneTypeConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CredentialTypes != nil {
		toSerialize["credentialTypes"] = o.CredentialTypes
	}
	return json.Marshal(toSerialize)
}

type NullableZoneTypeConfig struct {
	value *ZoneTypeConfig
	isSet bool
}

func (v NullableZoneTypeConfig) Get() *ZoneTypeConfig {
	return v.value
}

func (v *NullableZoneTypeConfig) Set(val *ZoneTypeConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneTypeConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneTypeConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneTypeConfig(val *ZoneTypeConfig) *NullableZoneTypeConfig {
	return &NullableZoneTypeConfig{value: val, isSet: true}
}

func (v NullableZoneTypeConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneTypeConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


