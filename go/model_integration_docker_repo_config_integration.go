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

// IntegrationDockerRepoConfigIntegration struct for IntegrationDockerRepoConfigIntegration
type IntegrationDockerRepoConfigIntegration struct {
	// Name, a unique identifier for the integration
	Name string `json:"name"`
	// Integration Type Code
	Type string `json:"type"`
	// Repository URL
	ServiceUrl string `json:"serviceUrl"`
	// Username
	ServiceUsername string `json:"serviceUsername"`
	// Password
	ServicePassword string `json:"servicePassword"`
}

// NewIntegrationDockerRepoConfigIntegration instantiates a new IntegrationDockerRepoConfigIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationDockerRepoConfigIntegration(name string, type_ string, serviceUrl string, serviceUsername string, servicePassword string, ) *IntegrationDockerRepoConfigIntegration {
	this := IntegrationDockerRepoConfigIntegration{}
	this.Name = name
	this.Type = type_
	this.ServiceUrl = serviceUrl
	this.ServiceUsername = serviceUsername
	this.ServicePassword = servicePassword
	return &this
}

// NewIntegrationDockerRepoConfigIntegrationWithDefaults instantiates a new IntegrationDockerRepoConfigIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationDockerRepoConfigIntegrationWithDefaults() *IntegrationDockerRepoConfigIntegration {
	this := IntegrationDockerRepoConfigIntegration{}
	return &this
}

// GetName returns the Name field value
func (o *IntegrationDockerRepoConfigIntegration) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IntegrationDockerRepoConfigIntegration) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IntegrationDockerRepoConfigIntegration) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *IntegrationDockerRepoConfigIntegration) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IntegrationDockerRepoConfigIntegration) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IntegrationDockerRepoConfigIntegration) SetType(v string) {
	o.Type = v
}

// GetServiceUrl returns the ServiceUrl field value
func (o *IntegrationDockerRepoConfigIntegration) GetServiceUrl() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ServiceUrl
}

// GetServiceUrlOk returns a tuple with the ServiceUrl field value
// and a boolean to check if the value has been set.
func (o *IntegrationDockerRepoConfigIntegration) GetServiceUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServiceUrl, true
}

// SetServiceUrl sets field value
func (o *IntegrationDockerRepoConfigIntegration) SetServiceUrl(v string) {
	o.ServiceUrl = v
}

// GetServiceUsername returns the ServiceUsername field value
func (o *IntegrationDockerRepoConfigIntegration) GetServiceUsername() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ServiceUsername
}

// GetServiceUsernameOk returns a tuple with the ServiceUsername field value
// and a boolean to check if the value has been set.
func (o *IntegrationDockerRepoConfigIntegration) GetServiceUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServiceUsername, true
}

// SetServiceUsername sets field value
func (o *IntegrationDockerRepoConfigIntegration) SetServiceUsername(v string) {
	o.ServiceUsername = v
}

// GetServicePassword returns the ServicePassword field value
func (o *IntegrationDockerRepoConfigIntegration) GetServicePassword() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ServicePassword
}

// GetServicePasswordOk returns a tuple with the ServicePassword field value
// and a boolean to check if the value has been set.
func (o *IntegrationDockerRepoConfigIntegration) GetServicePasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServicePassword, true
}

// SetServicePassword sets field value
func (o *IntegrationDockerRepoConfigIntegration) SetServicePassword(v string) {
	o.ServicePassword = v
}

func (o IntegrationDockerRepoConfigIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["serviceUrl"] = o.ServiceUrl
	}
	if true {
		toSerialize["serviceUsername"] = o.ServiceUsername
	}
	if true {
		toSerialize["servicePassword"] = o.ServicePassword
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationDockerRepoConfigIntegration struct {
	value *IntegrationDockerRepoConfigIntegration
	isSet bool
}

func (v NullableIntegrationDockerRepoConfigIntegration) Get() *IntegrationDockerRepoConfigIntegration {
	return v.value
}

func (v *NullableIntegrationDockerRepoConfigIntegration) Set(val *IntegrationDockerRepoConfigIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationDockerRepoConfigIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationDockerRepoConfigIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationDockerRepoConfigIntegration(val *IntegrationDockerRepoConfigIntegration) *NullableIntegrationDockerRepoConfigIntegration {
	return &NullableIntegrationDockerRepoConfigIntegration{value: val, isSet: true}
}

func (v NullableIntegrationDockerRepoConfigIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationDockerRepoConfigIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


