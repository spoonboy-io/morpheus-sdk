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

// IntegrationGitRepoConfigIntegrationConfig struct for IntegrationGitRepoConfigIntegrationConfig
type IntegrationGitRepoConfigIntegrationConfig struct {
	// Default Branch
	DefaultBranch *string `json:"defaultBranch,omitempty"`
	// Enable Git Repository Caching
	CacheEnabled *bool `json:"cacheEnabled,omitempty"`
}

// NewIntegrationGitRepoConfigIntegrationConfig instantiates a new IntegrationGitRepoConfigIntegrationConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationGitRepoConfigIntegrationConfig() *IntegrationGitRepoConfigIntegrationConfig {
	this := IntegrationGitRepoConfigIntegrationConfig{}
	return &this
}

// NewIntegrationGitRepoConfigIntegrationConfigWithDefaults instantiates a new IntegrationGitRepoConfigIntegrationConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationGitRepoConfigIntegrationConfigWithDefaults() *IntegrationGitRepoConfigIntegrationConfig {
	this := IntegrationGitRepoConfigIntegrationConfig{}
	return &this
}

// GetDefaultBranch returns the DefaultBranch field value if set, zero value otherwise.
func (o *IntegrationGitRepoConfigIntegrationConfig) GetDefaultBranch() string {
	if o == nil || o.DefaultBranch == nil {
		var ret string
		return ret
	}
	return *o.DefaultBranch
}

// GetDefaultBranchOk returns a tuple with the DefaultBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationGitRepoConfigIntegrationConfig) GetDefaultBranchOk() (*string, bool) {
	if o == nil || o.DefaultBranch == nil {
		return nil, false
	}
	return o.DefaultBranch, true
}

// HasDefaultBranch returns a boolean if a field has been set.
func (o *IntegrationGitRepoConfigIntegrationConfig) HasDefaultBranch() bool {
	if o != nil && o.DefaultBranch != nil {
		return true
	}

	return false
}

// SetDefaultBranch gets a reference to the given string and assigns it to the DefaultBranch field.
func (o *IntegrationGitRepoConfigIntegrationConfig) SetDefaultBranch(v string) {
	o.DefaultBranch = &v
}

// GetCacheEnabled returns the CacheEnabled field value if set, zero value otherwise.
func (o *IntegrationGitRepoConfigIntegrationConfig) GetCacheEnabled() bool {
	if o == nil || o.CacheEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CacheEnabled
}

// GetCacheEnabledOk returns a tuple with the CacheEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationGitRepoConfigIntegrationConfig) GetCacheEnabledOk() (*bool, bool) {
	if o == nil || o.CacheEnabled == nil {
		return nil, false
	}
	return o.CacheEnabled, true
}

// HasCacheEnabled returns a boolean if a field has been set.
func (o *IntegrationGitRepoConfigIntegrationConfig) HasCacheEnabled() bool {
	if o != nil && o.CacheEnabled != nil {
		return true
	}

	return false
}

// SetCacheEnabled gets a reference to the given bool and assigns it to the CacheEnabled field.
func (o *IntegrationGitRepoConfigIntegrationConfig) SetCacheEnabled(v bool) {
	o.CacheEnabled = &v
}

func (o IntegrationGitRepoConfigIntegrationConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultBranch != nil {
		toSerialize["defaultBranch"] = o.DefaultBranch
	}
	if o.CacheEnabled != nil {
		toSerialize["cacheEnabled"] = o.CacheEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationGitRepoConfigIntegrationConfig struct {
	value *IntegrationGitRepoConfigIntegrationConfig
	isSet bool
}

func (v NullableIntegrationGitRepoConfigIntegrationConfig) Get() *IntegrationGitRepoConfigIntegrationConfig {
	return v.value
}

func (v *NullableIntegrationGitRepoConfigIntegrationConfig) Set(val *IntegrationGitRepoConfigIntegrationConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationGitRepoConfigIntegrationConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationGitRepoConfigIntegrationConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationGitRepoConfigIntegrationConfig(val *IntegrationGitRepoConfigIntegrationConfig) *NullableIntegrationGitRepoConfigIntegrationConfig {
	return &NullableIntegrationGitRepoConfigIntegrationConfig{value: val, isSet: true}
}

func (v NullableIntegrationGitRepoConfigIntegrationConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationGitRepoConfigIntegrationConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

