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

// IntegrationConfig struct for IntegrationConfig
type IntegrationConfig struct {
	Integration IntegrationConfigIntegration `json:"integration"`
}

// NewIntegrationConfig instantiates a new IntegrationConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationConfig(integration IntegrationConfigIntegration, ) *IntegrationConfig {
	this := IntegrationConfig{}
	this.Integration = integration
	return &this
}

// NewIntegrationConfigWithDefaults instantiates a new IntegrationConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationConfigWithDefaults() *IntegrationConfig {
	this := IntegrationConfig{}
	return &this
}

// GetIntegration returns the Integration field value
func (o *IntegrationConfig) GetIntegration() IntegrationConfigIntegration {
	if o == nil  {
		var ret IntegrationConfigIntegration
		return ret
	}

	return o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value
// and a boolean to check if the value has been set.
func (o *IntegrationConfig) GetIntegrationOk() (*IntegrationConfigIntegration, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Integration, true
}

// SetIntegration sets field value
func (o *IntegrationConfig) SetIntegration(v IntegrationConfigIntegration) {
	o.Integration = v
}

func (o IntegrationConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["integration"] = o.Integration
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationConfig struct {
	value *IntegrationConfig
	isSet bool
}

func (v NullableIntegrationConfig) Get() *IntegrationConfig {
	return v.value
}

func (v *NullableIntegrationConfig) Set(val *IntegrationConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationConfig(val *IntegrationConfig) *NullableIntegrationConfig {
	return &NullableIntegrationConfig{value: val, isSet: true}
}

func (v NullableIntegrationConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


