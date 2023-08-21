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

// IntegrationGitHubConfigIntegration struct for IntegrationGitHubConfigIntegration
type IntegrationGitHubConfigIntegration struct {
	// Name, a unique identifier for the integration
	Name string `json:"name"`
	// Integration Type Code
	Type string `json:"type"`
	// Username
	ServiceUsername string `json:"serviceUsername"`
	// Password
	ServicePassword *string `json:"servicePassword,omitempty"`
	// Access Token
	ServiceToken *string `json:"serviceToken,omitempty"`
	// Key Pair ID
	ServiceKey *int64 `json:"serviceKey,omitempty"`
	Config *IntegrationGitHubConfigIntegrationConfig `json:"config,omitempty"`
}

// NewIntegrationGitHubConfigIntegration instantiates a new IntegrationGitHubConfigIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationGitHubConfigIntegration(name string, type_ string, serviceUsername string, ) *IntegrationGitHubConfigIntegration {
	this := IntegrationGitHubConfigIntegration{}
	this.Name = name
	this.Type = type_
	this.ServiceUsername = serviceUsername
	return &this
}

// NewIntegrationGitHubConfigIntegrationWithDefaults instantiates a new IntegrationGitHubConfigIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationGitHubConfigIntegrationWithDefaults() *IntegrationGitHubConfigIntegration {
	this := IntegrationGitHubConfigIntegration{}
	return &this
}

// GetName returns the Name field value
func (o *IntegrationGitHubConfigIntegration) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IntegrationGitHubConfigIntegration) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IntegrationGitHubConfigIntegration) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *IntegrationGitHubConfigIntegration) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IntegrationGitHubConfigIntegration) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IntegrationGitHubConfigIntegration) SetType(v string) {
	o.Type = v
}

// GetServiceUsername returns the ServiceUsername field value
func (o *IntegrationGitHubConfigIntegration) GetServiceUsername() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ServiceUsername
}

// GetServiceUsernameOk returns a tuple with the ServiceUsername field value
// and a boolean to check if the value has been set.
func (o *IntegrationGitHubConfigIntegration) GetServiceUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServiceUsername, true
}

// SetServiceUsername sets field value
func (o *IntegrationGitHubConfigIntegration) SetServiceUsername(v string) {
	o.ServiceUsername = v
}

// GetServicePassword returns the ServicePassword field value if set, zero value otherwise.
func (o *IntegrationGitHubConfigIntegration) GetServicePassword() string {
	if o == nil || o.ServicePassword == nil {
		var ret string
		return ret
	}
	return *o.ServicePassword
}

// GetServicePasswordOk returns a tuple with the ServicePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationGitHubConfigIntegration) GetServicePasswordOk() (*string, bool) {
	if o == nil || o.ServicePassword == nil {
		return nil, false
	}
	return o.ServicePassword, true
}

// HasServicePassword returns a boolean if a field has been set.
func (o *IntegrationGitHubConfigIntegration) HasServicePassword() bool {
	if o != nil && o.ServicePassword != nil {
		return true
	}

	return false
}

// SetServicePassword gets a reference to the given string and assigns it to the ServicePassword field.
func (o *IntegrationGitHubConfigIntegration) SetServicePassword(v string) {
	o.ServicePassword = &v
}

// GetServiceToken returns the ServiceToken field value if set, zero value otherwise.
func (o *IntegrationGitHubConfigIntegration) GetServiceToken() string {
	if o == nil || o.ServiceToken == nil {
		var ret string
		return ret
	}
	return *o.ServiceToken
}

// GetServiceTokenOk returns a tuple with the ServiceToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationGitHubConfigIntegration) GetServiceTokenOk() (*string, bool) {
	if o == nil || o.ServiceToken == nil {
		return nil, false
	}
	return o.ServiceToken, true
}

// HasServiceToken returns a boolean if a field has been set.
func (o *IntegrationGitHubConfigIntegration) HasServiceToken() bool {
	if o != nil && o.ServiceToken != nil {
		return true
	}

	return false
}

// SetServiceToken gets a reference to the given string and assigns it to the ServiceToken field.
func (o *IntegrationGitHubConfigIntegration) SetServiceToken(v string) {
	o.ServiceToken = &v
}

// GetServiceKey returns the ServiceKey field value if set, zero value otherwise.
func (o *IntegrationGitHubConfigIntegration) GetServiceKey() int64 {
	if o == nil || o.ServiceKey == nil {
		var ret int64
		return ret
	}
	return *o.ServiceKey
}

// GetServiceKeyOk returns a tuple with the ServiceKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationGitHubConfigIntegration) GetServiceKeyOk() (*int64, bool) {
	if o == nil || o.ServiceKey == nil {
		return nil, false
	}
	return o.ServiceKey, true
}

// HasServiceKey returns a boolean if a field has been set.
func (o *IntegrationGitHubConfigIntegration) HasServiceKey() bool {
	if o != nil && o.ServiceKey != nil {
		return true
	}

	return false
}

// SetServiceKey gets a reference to the given int64 and assigns it to the ServiceKey field.
func (o *IntegrationGitHubConfigIntegration) SetServiceKey(v int64) {
	o.ServiceKey = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *IntegrationGitHubConfigIntegration) GetConfig() IntegrationGitHubConfigIntegrationConfig {
	if o == nil || o.Config == nil {
		var ret IntegrationGitHubConfigIntegrationConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationGitHubConfigIntegration) GetConfigOk() (*IntegrationGitHubConfigIntegrationConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *IntegrationGitHubConfigIntegration) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given IntegrationGitHubConfigIntegrationConfig and assigns it to the Config field.
func (o *IntegrationGitHubConfigIntegration) SetConfig(v IntegrationGitHubConfigIntegrationConfig) {
	o.Config = &v
}

func (o IntegrationGitHubConfigIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["serviceUsername"] = o.ServiceUsername
	}
	if o.ServicePassword != nil {
		toSerialize["servicePassword"] = o.ServicePassword
	}
	if o.ServiceToken != nil {
		toSerialize["serviceToken"] = o.ServiceToken
	}
	if o.ServiceKey != nil {
		toSerialize["serviceKey"] = o.ServiceKey
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationGitHubConfigIntegration struct {
	value *IntegrationGitHubConfigIntegration
	isSet bool
}

func (v NullableIntegrationGitHubConfigIntegration) Get() *IntegrationGitHubConfigIntegration {
	return v.value
}

func (v *NullableIntegrationGitHubConfigIntegration) Set(val *IntegrationGitHubConfigIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationGitHubConfigIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationGitHubConfigIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationGitHubConfigIntegration(val *IntegrationGitHubConfigIntegration) *NullableIntegrationGitHubConfigIntegration {
	return &NullableIntegrationGitHubConfigIntegration{value: val, isSet: true}
}

func (v NullableIntegrationGitHubConfigIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationGitHubConfigIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


