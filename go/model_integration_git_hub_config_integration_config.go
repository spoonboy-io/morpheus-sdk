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

// IntegrationGitHubConfigIntegrationConfig struct for IntegrationGitHubConfigIntegrationConfig
type IntegrationGitHubConfigIntegrationConfig struct {
	// Enable Git Repository Caching
	CacheEnabled *bool `json:"cacheEnabled,omitempty"`
}

// NewIntegrationGitHubConfigIntegrationConfig instantiates a new IntegrationGitHubConfigIntegrationConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationGitHubConfigIntegrationConfig() *IntegrationGitHubConfigIntegrationConfig {
	this := IntegrationGitHubConfigIntegrationConfig{}
	return &this
}

// NewIntegrationGitHubConfigIntegrationConfigWithDefaults instantiates a new IntegrationGitHubConfigIntegrationConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationGitHubConfigIntegrationConfigWithDefaults() *IntegrationGitHubConfigIntegrationConfig {
	this := IntegrationGitHubConfigIntegrationConfig{}
	return &this
}

// GetCacheEnabled returns the CacheEnabled field value if set, zero value otherwise.
func (o *IntegrationGitHubConfigIntegrationConfig) GetCacheEnabled() bool {
	if o == nil || o.CacheEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CacheEnabled
}

// GetCacheEnabledOk returns a tuple with the CacheEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationGitHubConfigIntegrationConfig) GetCacheEnabledOk() (*bool, bool) {
	if o == nil || o.CacheEnabled == nil {
		return nil, false
	}
	return o.CacheEnabled, true
}

// HasCacheEnabled returns a boolean if a field has been set.
func (o *IntegrationGitHubConfigIntegrationConfig) HasCacheEnabled() bool {
	if o != nil && o.CacheEnabled != nil {
		return true
	}

	return false
}

// SetCacheEnabled gets a reference to the given bool and assigns it to the CacheEnabled field.
func (o *IntegrationGitHubConfigIntegrationConfig) SetCacheEnabled(v bool) {
	o.CacheEnabled = &v
}

func (o IntegrationGitHubConfigIntegrationConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CacheEnabled != nil {
		toSerialize["cacheEnabled"] = o.CacheEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationGitHubConfigIntegrationConfig struct {
	value *IntegrationGitHubConfigIntegrationConfig
	isSet bool
}

func (v NullableIntegrationGitHubConfigIntegrationConfig) Get() *IntegrationGitHubConfigIntegrationConfig {
	return v.value
}

func (v *NullableIntegrationGitHubConfigIntegrationConfig) Set(val *IntegrationGitHubConfigIntegrationConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationGitHubConfigIntegrationConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationGitHubConfigIntegrationConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationGitHubConfigIntegrationConfig(val *IntegrationGitHubConfigIntegrationConfig) *NullableIntegrationGitHubConfigIntegrationConfig {
	return &NullableIntegrationGitHubConfigIntegrationConfig{value: val, isSet: true}
}

func (v NullableIntegrationGitHubConfigIntegrationConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationGitHubConfigIntegrationConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


