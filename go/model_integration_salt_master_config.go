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

// IntegrationSaltMasterConfig struct for IntegrationSaltMasterConfig
type IntegrationSaltMasterConfig struct {
	SaltApplyOnMinion *bool `json:"saltApplyOnMinion,omitempty"`
}

// NewIntegrationSaltMasterConfig instantiates a new IntegrationSaltMasterConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationSaltMasterConfig() *IntegrationSaltMasterConfig {
	this := IntegrationSaltMasterConfig{}
	return &this
}

// NewIntegrationSaltMasterConfigWithDefaults instantiates a new IntegrationSaltMasterConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationSaltMasterConfigWithDefaults() *IntegrationSaltMasterConfig {
	this := IntegrationSaltMasterConfig{}
	return &this
}

// GetSaltApplyOnMinion returns the SaltApplyOnMinion field value if set, zero value otherwise.
func (o *IntegrationSaltMasterConfig) GetSaltApplyOnMinion() bool {
	if o == nil || o.SaltApplyOnMinion == nil {
		var ret bool
		return ret
	}
	return *o.SaltApplyOnMinion
}

// GetSaltApplyOnMinionOk returns a tuple with the SaltApplyOnMinion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationSaltMasterConfig) GetSaltApplyOnMinionOk() (*bool, bool) {
	if o == nil || o.SaltApplyOnMinion == nil {
		return nil, false
	}
	return o.SaltApplyOnMinion, true
}

// HasSaltApplyOnMinion returns a boolean if a field has been set.
func (o *IntegrationSaltMasterConfig) HasSaltApplyOnMinion() bool {
	if o != nil && o.SaltApplyOnMinion != nil {
		return true
	}

	return false
}

// SetSaltApplyOnMinion gets a reference to the given bool and assigns it to the SaltApplyOnMinion field.
func (o *IntegrationSaltMasterConfig) SetSaltApplyOnMinion(v bool) {
	o.SaltApplyOnMinion = &v
}

func (o IntegrationSaltMasterConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SaltApplyOnMinion != nil {
		toSerialize["saltApplyOnMinion"] = o.SaltApplyOnMinion
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationSaltMasterConfig struct {
	value *IntegrationSaltMasterConfig
	isSet bool
}

func (v NullableIntegrationSaltMasterConfig) Get() *IntegrationSaltMasterConfig {
	return v.value
}

func (v *NullableIntegrationSaltMasterConfig) Set(val *IntegrationSaltMasterConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationSaltMasterConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationSaltMasterConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationSaltMasterConfig(val *IntegrationSaltMasterConfig) *NullableIntegrationSaltMasterConfig {
	return &NullableIntegrationSaltMasterConfig{value: val, isSet: true}
}

func (v NullableIntegrationSaltMasterConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationSaltMasterConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


