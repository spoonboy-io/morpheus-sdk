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

// CredentialAccessSecretKeyConfig struct for CredentialAccessSecretKeyConfig
type CredentialAccessSecretKeyConfig struct {
	// Credential Type Code
	Type string `json:"type"`
	// A unique name scoped to your account for the credential
	Name string `json:"name"`
	// Optional Description
	Description *string `json:"description,omitempty"`
	// Credential enabled
	Enabled *bool `json:"enabled,omitempty"`
	Integration NullableCredentialAccessSecretKeyConfigIntegration `json:"integration,omitempty"`
	// Access Key
	Username string `json:"username"`
	// Secret Key
	Password string `json:"password"`
}

// NewCredentialAccessSecretKeyConfig instantiates a new CredentialAccessSecretKeyConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialAccessSecretKeyConfig(type_ string, name string, username string, password string, ) *CredentialAccessSecretKeyConfig {
	this := CredentialAccessSecretKeyConfig{}
	this.Type = type_
	this.Name = name
	var enabled bool = true
	this.Enabled = &enabled
	this.Username = username
	this.Password = password
	return &this
}

// NewCredentialAccessSecretKeyConfigWithDefaults instantiates a new CredentialAccessSecretKeyConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialAccessSecretKeyConfigWithDefaults() *CredentialAccessSecretKeyConfig {
	this := CredentialAccessSecretKeyConfig{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetType returns the Type field value
func (o *CredentialAccessSecretKeyConfig) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CredentialAccessSecretKeyConfig) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CredentialAccessSecretKeyConfig) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *CredentialAccessSecretKeyConfig) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CredentialAccessSecretKeyConfig) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CredentialAccessSecretKeyConfig) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CredentialAccessSecretKeyConfig) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialAccessSecretKeyConfig) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CredentialAccessSecretKeyConfig) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CredentialAccessSecretKeyConfig) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CredentialAccessSecretKeyConfig) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialAccessSecretKeyConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CredentialAccessSecretKeyConfig) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CredentialAccessSecretKeyConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIntegration returns the Integration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialAccessSecretKeyConfig) GetIntegration() CredentialAccessSecretKeyConfigIntegration {
	if o == nil || o.Integration.Get() == nil {
		var ret CredentialAccessSecretKeyConfigIntegration
		return ret
	}
	return *o.Integration.Get()
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialAccessSecretKeyConfig) GetIntegrationOk() (*CredentialAccessSecretKeyConfigIntegration, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Integration.Get(), o.Integration.IsSet()
}

// HasIntegration returns a boolean if a field has been set.
func (o *CredentialAccessSecretKeyConfig) HasIntegration() bool {
	if o != nil && o.Integration.IsSet() {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given NullableCredentialAccessSecretKeyConfigIntegration and assigns it to the Integration field.
func (o *CredentialAccessSecretKeyConfig) SetIntegration(v CredentialAccessSecretKeyConfigIntegration) {
	o.Integration.Set(&v)
}
// SetIntegrationNil sets the value for Integration to be an explicit nil
func (o *CredentialAccessSecretKeyConfig) SetIntegrationNil() {
	o.Integration.Set(nil)
}

// UnsetIntegration ensures that no value is present for Integration, not even an explicit nil
func (o *CredentialAccessSecretKeyConfig) UnsetIntegration() {
	o.Integration.Unset()
}

// GetUsername returns the Username field value
func (o *CredentialAccessSecretKeyConfig) GetUsername() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CredentialAccessSecretKeyConfig) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CredentialAccessSecretKeyConfig) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *CredentialAccessSecretKeyConfig) GetPassword() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CredentialAccessSecretKeyConfig) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CredentialAccessSecretKeyConfig) SetPassword(v string) {
	o.Password = v
}

func (o CredentialAccessSecretKeyConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Integration.IsSet() {
		toSerialize["integration"] = o.Integration.Get()
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableCredentialAccessSecretKeyConfig struct {
	value *CredentialAccessSecretKeyConfig
	isSet bool
}

func (v NullableCredentialAccessSecretKeyConfig) Get() *CredentialAccessSecretKeyConfig {
	return v.value
}

func (v *NullableCredentialAccessSecretKeyConfig) Set(val *CredentialAccessSecretKeyConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialAccessSecretKeyConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialAccessSecretKeyConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialAccessSecretKeyConfig(val *CredentialAccessSecretKeyConfig) *NullableCredentialAccessSecretKeyConfig {
	return &NullableCredentialAccessSecretKeyConfig{value: val, isSet: true}
}

func (v NullableCredentialAccessSecretKeyConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialAccessSecretKeyConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


