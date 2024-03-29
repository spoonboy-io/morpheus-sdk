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

// BlueprintHelmCreateHelm struct for BlueprintHelmCreateHelm
type BlueprintHelmCreateHelm struct {
	// Configuration Type
	ConfigType string `json:"configType"`
	Git *BlueprintHelmCreateHelmGit `json:"git,omitempty"`
}

// NewBlueprintHelmCreateHelm instantiates a new BlueprintHelmCreateHelm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlueprintHelmCreateHelm(configType string, ) *BlueprintHelmCreateHelm {
	this := BlueprintHelmCreateHelm{}
	this.ConfigType = configType
	return &this
}

// NewBlueprintHelmCreateHelmWithDefaults instantiates a new BlueprintHelmCreateHelm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlueprintHelmCreateHelmWithDefaults() *BlueprintHelmCreateHelm {
	this := BlueprintHelmCreateHelm{}
	return &this
}

// GetConfigType returns the ConfigType field value
func (o *BlueprintHelmCreateHelm) GetConfigType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ConfigType
}

// GetConfigTypeOk returns a tuple with the ConfigType field value
// and a boolean to check if the value has been set.
func (o *BlueprintHelmCreateHelm) GetConfigTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfigType, true
}

// SetConfigType sets field value
func (o *BlueprintHelmCreateHelm) SetConfigType(v string) {
	o.ConfigType = v
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *BlueprintHelmCreateHelm) GetGit() BlueprintHelmCreateHelmGit {
	if o == nil || o.Git == nil {
		var ret BlueprintHelmCreateHelmGit
		return ret
	}
	return *o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintHelmCreateHelm) GetGitOk() (*BlueprintHelmCreateHelmGit, bool) {
	if o == nil || o.Git == nil {
		return nil, false
	}
	return o.Git, true
}

// HasGit returns a boolean if a field has been set.
func (o *BlueprintHelmCreateHelm) HasGit() bool {
	if o != nil && o.Git != nil {
		return true
	}

	return false
}

// SetGit gets a reference to the given BlueprintHelmCreateHelmGit and assigns it to the Git field.
func (o *BlueprintHelmCreateHelm) SetGit(v BlueprintHelmCreateHelmGit) {
	o.Git = &v
}

func (o BlueprintHelmCreateHelm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["configType"] = o.ConfigType
	}
	if o.Git != nil {
		toSerialize["git"] = o.Git
	}
	return json.Marshal(toSerialize)
}

type NullableBlueprintHelmCreateHelm struct {
	value *BlueprintHelmCreateHelm
	isSet bool
}

func (v NullableBlueprintHelmCreateHelm) Get() *BlueprintHelmCreateHelm {
	return v.value
}

func (v *NullableBlueprintHelmCreateHelm) Set(val *BlueprintHelmCreateHelm) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintHelmCreateHelm) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintHelmCreateHelm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintHelmCreateHelm(val *BlueprintHelmCreateHelm) *NullableBlueprintHelmCreateHelm {
	return &NullableBlueprintHelmCreateHelm{value: val, isSet: true}
}

func (v NullableBlueprintHelmCreateHelm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintHelmCreateHelm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


