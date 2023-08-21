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

// IntegrationAnsibleConfigIntegrationConfig struct for IntegrationAnsibleConfigIntegrationConfig
type IntegrationAnsibleConfigIntegrationConfig struct {
	// default branch
	DefaultBranch *string `json:"defaultBranch,omitempty"`
	// Playbooks path
	AnsiblePlaybooks *string `json:"ansiblePlaybooks,omitempty"`
	// Roles path
	AnsibleRoles *string `json:"ansibleRoles,omitempty"`
	// Group variables path
	AnsibleGroupVars *string `json:"ansibleGroupVars,omitempty"`
	// Host variables path
	AnsibleHostVars *string `json:"ansibleHostVars,omitempty"`
	// Use Ansible Galaxy
	AnsibleGalaxyEnabled *bool `json:"ansibleGalaxyEnabled,omitempty"`
	// Use verbose logging
	AnsibleVerbose *bool `json:"ansibleVerbose,omitempty"`
	// Use Morpheus Agent Command Bus
	AnsibleCommandBus *bool `json:"ansibleCommandBus,omitempty"`
	// Enable Git repository caching
	CacheEnabled *bool `json:"cacheEnabled,omitempty"`
}

// NewIntegrationAnsibleConfigIntegrationConfig instantiates a new IntegrationAnsibleConfigIntegrationConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationAnsibleConfigIntegrationConfig() *IntegrationAnsibleConfigIntegrationConfig {
	this := IntegrationAnsibleConfigIntegrationConfig{}
	return &this
}

// NewIntegrationAnsibleConfigIntegrationConfigWithDefaults instantiates a new IntegrationAnsibleConfigIntegrationConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationAnsibleConfigIntegrationConfigWithDefaults() *IntegrationAnsibleConfigIntegrationConfig {
	this := IntegrationAnsibleConfigIntegrationConfig{}
	return &this
}

// GetDefaultBranch returns the DefaultBranch field value if set, zero value otherwise.
func (o *IntegrationAnsibleConfigIntegrationConfig) GetDefaultBranch() string {
	if o == nil || o.DefaultBranch == nil {
		var ret string
		return ret
	}
	return *o.DefaultBranch
}

// GetDefaultBranchOk returns a tuple with the DefaultBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationAnsibleConfigIntegrationConfig) GetDefaultBranchOk() (*string, bool) {
	if o == nil || o.DefaultBranch == nil {
		return nil, false
	}
	return o.DefaultBranch, true
}

// HasDefaultBranch returns a boolean if a field has been set.
func (o *IntegrationAnsibleConfigIntegrationConfig) HasDefaultBranch() bool {
	if o != nil && o.DefaultBranch != nil {
		return true
	}

	return false
}

// SetDefaultBranch gets a reference to the given string and assigns it to the DefaultBranch field.
func (o *IntegrationAnsibleConfigIntegrationConfig) SetDefaultBranch(v string) {
	o.DefaultBranch = &v
}

// GetAnsiblePlaybooks returns the AnsiblePlaybooks field value if set, zero value otherwise.
func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsiblePlaybooks() string {
	if o == nil || o.AnsiblePlaybooks == nil {
		var ret string
		return ret
	}
	return *o.AnsiblePlaybooks
}

// GetAnsiblePlaybooksOk returns a tuple with the AnsiblePlaybooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsiblePlaybooksOk() (*string, bool) {
	if o == nil || o.AnsiblePlaybooks == nil {
		return nil, false
	}
	return o.AnsiblePlaybooks, true
}

// HasAnsiblePlaybooks returns a boolean if a field has been set.
func (o *IntegrationAnsibleConfigIntegrationConfig) HasAnsiblePlaybooks() bool {
	if o != nil && o.AnsiblePlaybooks != nil {
		return true
	}

	return false
}

// SetAnsiblePlaybooks gets a reference to the given string and assigns it to the AnsiblePlaybooks field.
func (o *IntegrationAnsibleConfigIntegrationConfig) SetAnsiblePlaybooks(v string) {
	o.AnsiblePlaybooks = &v
}

// GetAnsibleRoles returns the AnsibleRoles field value if set, zero value otherwise.
func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsibleRoles() string {
	if o == nil || o.AnsibleRoles == nil {
		var ret string
		return ret
	}
	return *o.AnsibleRoles
}

// GetAnsibleRolesOk returns a tuple with the AnsibleRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsibleRolesOk() (*string, bool) {
	if o == nil || o.AnsibleRoles == nil {
		return nil, false
	}
	return o.AnsibleRoles, true
}

// HasAnsibleRoles returns a boolean if a field has been set.
func (o *IntegrationAnsibleConfigIntegrationConfig) HasAnsibleRoles() bool {
	if o != nil && o.AnsibleRoles != nil {
		return true
	}

	return false
}

// SetAnsibleRoles gets a reference to the given string and assigns it to the AnsibleRoles field.
func (o *IntegrationAnsibleConfigIntegrationConfig) SetAnsibleRoles(v string) {
	o.AnsibleRoles = &v
}

// GetAnsibleGroupVars returns the AnsibleGroupVars field value if set, zero value otherwise.
func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsibleGroupVars() string {
	if o == nil || o.AnsibleGroupVars == nil {
		var ret string
		return ret
	}
	return *o.AnsibleGroupVars
}

// GetAnsibleGroupVarsOk returns a tuple with the AnsibleGroupVars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsibleGroupVarsOk() (*string, bool) {
	if o == nil || o.AnsibleGroupVars == nil {
		return nil, false
	}
	return o.AnsibleGroupVars, true
}

// HasAnsibleGroupVars returns a boolean if a field has been set.
func (o *IntegrationAnsibleConfigIntegrationConfig) HasAnsibleGroupVars() bool {
	if o != nil && o.AnsibleGroupVars != nil {
		return true
	}

	return false
}

// SetAnsibleGroupVars gets a reference to the given string and assigns it to the AnsibleGroupVars field.
func (o *IntegrationAnsibleConfigIntegrationConfig) SetAnsibleGroupVars(v string) {
	o.AnsibleGroupVars = &v
}

// GetAnsibleHostVars returns the AnsibleHostVars field value if set, zero value otherwise.
func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsibleHostVars() string {
	if o == nil || o.AnsibleHostVars == nil {
		var ret string
		return ret
	}
	return *o.AnsibleHostVars
}

// GetAnsibleHostVarsOk returns a tuple with the AnsibleHostVars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsibleHostVarsOk() (*string, bool) {
	if o == nil || o.AnsibleHostVars == nil {
		return nil, false
	}
	return o.AnsibleHostVars, true
}

// HasAnsibleHostVars returns a boolean if a field has been set.
func (o *IntegrationAnsibleConfigIntegrationConfig) HasAnsibleHostVars() bool {
	if o != nil && o.AnsibleHostVars != nil {
		return true
	}

	return false
}

// SetAnsibleHostVars gets a reference to the given string and assigns it to the AnsibleHostVars field.
func (o *IntegrationAnsibleConfigIntegrationConfig) SetAnsibleHostVars(v string) {
	o.AnsibleHostVars = &v
}

// GetAnsibleGalaxyEnabled returns the AnsibleGalaxyEnabled field value if set, zero value otherwise.
func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsibleGalaxyEnabled() bool {
	if o == nil || o.AnsibleGalaxyEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AnsibleGalaxyEnabled
}

// GetAnsibleGalaxyEnabledOk returns a tuple with the AnsibleGalaxyEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsibleGalaxyEnabledOk() (*bool, bool) {
	if o == nil || o.AnsibleGalaxyEnabled == nil {
		return nil, false
	}
	return o.AnsibleGalaxyEnabled, true
}

// HasAnsibleGalaxyEnabled returns a boolean if a field has been set.
func (o *IntegrationAnsibleConfigIntegrationConfig) HasAnsibleGalaxyEnabled() bool {
	if o != nil && o.AnsibleGalaxyEnabled != nil {
		return true
	}

	return false
}

// SetAnsibleGalaxyEnabled gets a reference to the given bool and assigns it to the AnsibleGalaxyEnabled field.
func (o *IntegrationAnsibleConfigIntegrationConfig) SetAnsibleGalaxyEnabled(v bool) {
	o.AnsibleGalaxyEnabled = &v
}

// GetAnsibleVerbose returns the AnsibleVerbose field value if set, zero value otherwise.
func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsibleVerbose() bool {
	if o == nil || o.AnsibleVerbose == nil {
		var ret bool
		return ret
	}
	return *o.AnsibleVerbose
}

// GetAnsibleVerboseOk returns a tuple with the AnsibleVerbose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsibleVerboseOk() (*bool, bool) {
	if o == nil || o.AnsibleVerbose == nil {
		return nil, false
	}
	return o.AnsibleVerbose, true
}

// HasAnsibleVerbose returns a boolean if a field has been set.
func (o *IntegrationAnsibleConfigIntegrationConfig) HasAnsibleVerbose() bool {
	if o != nil && o.AnsibleVerbose != nil {
		return true
	}

	return false
}

// SetAnsibleVerbose gets a reference to the given bool and assigns it to the AnsibleVerbose field.
func (o *IntegrationAnsibleConfigIntegrationConfig) SetAnsibleVerbose(v bool) {
	o.AnsibleVerbose = &v
}

// GetAnsibleCommandBus returns the AnsibleCommandBus field value if set, zero value otherwise.
func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsibleCommandBus() bool {
	if o == nil || o.AnsibleCommandBus == nil {
		var ret bool
		return ret
	}
	return *o.AnsibleCommandBus
}

// GetAnsibleCommandBusOk returns a tuple with the AnsibleCommandBus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsibleCommandBusOk() (*bool, bool) {
	if o == nil || o.AnsibleCommandBus == nil {
		return nil, false
	}
	return o.AnsibleCommandBus, true
}

// HasAnsibleCommandBus returns a boolean if a field has been set.
func (o *IntegrationAnsibleConfigIntegrationConfig) HasAnsibleCommandBus() bool {
	if o != nil && o.AnsibleCommandBus != nil {
		return true
	}

	return false
}

// SetAnsibleCommandBus gets a reference to the given bool and assigns it to the AnsibleCommandBus field.
func (o *IntegrationAnsibleConfigIntegrationConfig) SetAnsibleCommandBus(v bool) {
	o.AnsibleCommandBus = &v
}

// GetCacheEnabled returns the CacheEnabled field value if set, zero value otherwise.
func (o *IntegrationAnsibleConfigIntegrationConfig) GetCacheEnabled() bool {
	if o == nil || o.CacheEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CacheEnabled
}

// GetCacheEnabledOk returns a tuple with the CacheEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationAnsibleConfigIntegrationConfig) GetCacheEnabledOk() (*bool, bool) {
	if o == nil || o.CacheEnabled == nil {
		return nil, false
	}
	return o.CacheEnabled, true
}

// HasCacheEnabled returns a boolean if a field has been set.
func (o *IntegrationAnsibleConfigIntegrationConfig) HasCacheEnabled() bool {
	if o != nil && o.CacheEnabled != nil {
		return true
	}

	return false
}

// SetCacheEnabled gets a reference to the given bool and assigns it to the CacheEnabled field.
func (o *IntegrationAnsibleConfigIntegrationConfig) SetCacheEnabled(v bool) {
	o.CacheEnabled = &v
}

func (o IntegrationAnsibleConfigIntegrationConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultBranch != nil {
		toSerialize["defaultBranch"] = o.DefaultBranch
	}
	if o.AnsiblePlaybooks != nil {
		toSerialize["ansiblePlaybooks"] = o.AnsiblePlaybooks
	}
	if o.AnsibleRoles != nil {
		toSerialize["ansibleRoles"] = o.AnsibleRoles
	}
	if o.AnsibleGroupVars != nil {
		toSerialize["ansibleGroupVars"] = o.AnsibleGroupVars
	}
	if o.AnsibleHostVars != nil {
		toSerialize["ansibleHostVars"] = o.AnsibleHostVars
	}
	if o.AnsibleGalaxyEnabled != nil {
		toSerialize["ansibleGalaxyEnabled"] = o.AnsibleGalaxyEnabled
	}
	if o.AnsibleVerbose != nil {
		toSerialize["ansibleVerbose"] = o.AnsibleVerbose
	}
	if o.AnsibleCommandBus != nil {
		toSerialize["ansibleCommandBus"] = o.AnsibleCommandBus
	}
	if o.CacheEnabled != nil {
		toSerialize["cacheEnabled"] = o.CacheEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationAnsibleConfigIntegrationConfig struct {
	value *IntegrationAnsibleConfigIntegrationConfig
	isSet bool
}

func (v NullableIntegrationAnsibleConfigIntegrationConfig) Get() *IntegrationAnsibleConfigIntegrationConfig {
	return v.value
}

func (v *NullableIntegrationAnsibleConfigIntegrationConfig) Set(val *IntegrationAnsibleConfigIntegrationConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationAnsibleConfigIntegrationConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationAnsibleConfigIntegrationConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationAnsibleConfigIntegrationConfig(val *IntegrationAnsibleConfigIntegrationConfig) *NullableIntegrationAnsibleConfigIntegrationConfig {
	return &NullableIntegrationAnsibleConfigIntegrationConfig{value: val, isSet: true}
}

func (v NullableIntegrationAnsibleConfigIntegrationConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationAnsibleConfigIntegrationConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


