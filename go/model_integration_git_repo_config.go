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

// IntegrationGitRepoConfig struct for IntegrationGitRepoConfig
type IntegrationGitRepoConfig struct {
	DefaultBranch *string `json:"defaultBranch,omitempty"`
	CacheEnabled *bool `json:"cacheEnabled,omitempty"`
}

// NewIntegrationGitRepoConfig instantiates a new IntegrationGitRepoConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationGitRepoConfig() *IntegrationGitRepoConfig {
	this := IntegrationGitRepoConfig{}
	return &this
}

// NewIntegrationGitRepoConfigWithDefaults instantiates a new IntegrationGitRepoConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationGitRepoConfigWithDefaults() *IntegrationGitRepoConfig {
	this := IntegrationGitRepoConfig{}
	return &this
}

// GetDefaultBranch returns the DefaultBranch field value if set, zero value otherwise.
func (o *IntegrationGitRepoConfig) GetDefaultBranch() string {
	if o == nil || o.DefaultBranch == nil {
		var ret string
		return ret
	}
	return *o.DefaultBranch
}

// GetDefaultBranchOk returns a tuple with the DefaultBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationGitRepoConfig) GetDefaultBranchOk() (*string, bool) {
	if o == nil || o.DefaultBranch == nil {
		return nil, false
	}
	return o.DefaultBranch, true
}

// HasDefaultBranch returns a boolean if a field has been set.
func (o *IntegrationGitRepoConfig) HasDefaultBranch() bool {
	if o != nil && o.DefaultBranch != nil {
		return true
	}

	return false
}

// SetDefaultBranch gets a reference to the given string and assigns it to the DefaultBranch field.
func (o *IntegrationGitRepoConfig) SetDefaultBranch(v string) {
	o.DefaultBranch = &v
}

// GetCacheEnabled returns the CacheEnabled field value if set, zero value otherwise.
func (o *IntegrationGitRepoConfig) GetCacheEnabled() bool {
	if o == nil || o.CacheEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CacheEnabled
}

// GetCacheEnabledOk returns a tuple with the CacheEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationGitRepoConfig) GetCacheEnabledOk() (*bool, bool) {
	if o == nil || o.CacheEnabled == nil {
		return nil, false
	}
	return o.CacheEnabled, true
}

// HasCacheEnabled returns a boolean if a field has been set.
func (o *IntegrationGitRepoConfig) HasCacheEnabled() bool {
	if o != nil && o.CacheEnabled != nil {
		return true
	}

	return false
}

// SetCacheEnabled gets a reference to the given bool and assigns it to the CacheEnabled field.
func (o *IntegrationGitRepoConfig) SetCacheEnabled(v bool) {
	o.CacheEnabled = &v
}

func (o IntegrationGitRepoConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultBranch != nil {
		toSerialize["defaultBranch"] = o.DefaultBranch
	}
	if o.CacheEnabled != nil {
		toSerialize["cacheEnabled"] = o.CacheEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationGitRepoConfig struct {
	value *IntegrationGitRepoConfig
	isSet bool
}

func (v NullableIntegrationGitRepoConfig) Get() *IntegrationGitRepoConfig {
	return v.value
}

func (v *NullableIntegrationGitRepoConfig) Set(val *IntegrationGitRepoConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationGitRepoConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationGitRepoConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationGitRepoConfig(val *IntegrationGitRepoConfig) *NullableIntegrationGitRepoConfig {
	return &NullableIntegrationGitRepoConfig{value: val, isSet: true}
}

func (v NullableIntegrationGitRepoConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationGitRepoConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

