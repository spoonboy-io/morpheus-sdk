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

// BlueprintARMCreateArm struct for BlueprintARMCreateArm
type BlueprintARMCreateArm struct {
	// Configuration Type
	ConfigType string `json:"configType"`
	// ARM Template in JSON
	Json *string `json:"json,omitempty"`
	// ARM Template in YAML
	Yaml *string `json:"yaml,omitempty"`
	Git *BlueprintARMCreateArmGit `json:"git,omitempty"`
	// OS Type
	OsType *string `json:"osType,omitempty"`
	// Install Morpheus Agent
	InstallAgent *OneOfbooleanstring `json:"installAgent,omitempty"`
	// Cloud Init Enabled
	CloudInitEnabled *OneOfbooleanstring `json:"cloudInitEnabled,omitempty"`
}

// NewBlueprintARMCreateArm instantiates a new BlueprintARMCreateArm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlueprintARMCreateArm(configType string, ) *BlueprintARMCreateArm {
	this := BlueprintARMCreateArm{}
	this.ConfigType = configType
	return &this
}

// NewBlueprintARMCreateArmWithDefaults instantiates a new BlueprintARMCreateArm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlueprintARMCreateArmWithDefaults() *BlueprintARMCreateArm {
	this := BlueprintARMCreateArm{}
	return &this
}

// GetConfigType returns the ConfigType field value
func (o *BlueprintARMCreateArm) GetConfigType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ConfigType
}

// GetConfigTypeOk returns a tuple with the ConfigType field value
// and a boolean to check if the value has been set.
func (o *BlueprintARMCreateArm) GetConfigTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfigType, true
}

// SetConfigType sets field value
func (o *BlueprintARMCreateArm) SetConfigType(v string) {
	o.ConfigType = v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *BlueprintARMCreateArm) GetJson() string {
	if o == nil || o.Json == nil {
		var ret string
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintARMCreateArm) GetJsonOk() (*string, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *BlueprintARMCreateArm) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given string and assigns it to the Json field.
func (o *BlueprintARMCreateArm) SetJson(v string) {
	o.Json = &v
}

// GetYaml returns the Yaml field value if set, zero value otherwise.
func (o *BlueprintARMCreateArm) GetYaml() string {
	if o == nil || o.Yaml == nil {
		var ret string
		return ret
	}
	return *o.Yaml
}

// GetYamlOk returns a tuple with the Yaml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintARMCreateArm) GetYamlOk() (*string, bool) {
	if o == nil || o.Yaml == nil {
		return nil, false
	}
	return o.Yaml, true
}

// HasYaml returns a boolean if a field has been set.
func (o *BlueprintARMCreateArm) HasYaml() bool {
	if o != nil && o.Yaml != nil {
		return true
	}

	return false
}

// SetYaml gets a reference to the given string and assigns it to the Yaml field.
func (o *BlueprintARMCreateArm) SetYaml(v string) {
	o.Yaml = &v
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *BlueprintARMCreateArm) GetGit() BlueprintARMCreateArmGit {
	if o == nil || o.Git == nil {
		var ret BlueprintARMCreateArmGit
		return ret
	}
	return *o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintARMCreateArm) GetGitOk() (*BlueprintARMCreateArmGit, bool) {
	if o == nil || o.Git == nil {
		return nil, false
	}
	return o.Git, true
}

// HasGit returns a boolean if a field has been set.
func (o *BlueprintARMCreateArm) HasGit() bool {
	if o != nil && o.Git != nil {
		return true
	}

	return false
}

// SetGit gets a reference to the given BlueprintARMCreateArmGit and assigns it to the Git field.
func (o *BlueprintARMCreateArm) SetGit(v BlueprintARMCreateArmGit) {
	o.Git = &v
}

// GetOsType returns the OsType field value if set, zero value otherwise.
func (o *BlueprintARMCreateArm) GetOsType() string {
	if o == nil || o.OsType == nil {
		var ret string
		return ret
	}
	return *o.OsType
}

// GetOsTypeOk returns a tuple with the OsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintARMCreateArm) GetOsTypeOk() (*string, bool) {
	if o == nil || o.OsType == nil {
		return nil, false
	}
	return o.OsType, true
}

// HasOsType returns a boolean if a field has been set.
func (o *BlueprintARMCreateArm) HasOsType() bool {
	if o != nil && o.OsType != nil {
		return true
	}

	return false
}

// SetOsType gets a reference to the given string and assigns it to the OsType field.
func (o *BlueprintARMCreateArm) SetOsType(v string) {
	o.OsType = &v
}

// GetInstallAgent returns the InstallAgent field value if set, zero value otherwise.
func (o *BlueprintARMCreateArm) GetInstallAgent() OneOfbooleanstring {
	if o == nil || o.InstallAgent == nil {
		var ret OneOfbooleanstring
		return ret
	}
	return *o.InstallAgent
}

// GetInstallAgentOk returns a tuple with the InstallAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintARMCreateArm) GetInstallAgentOk() (*OneOfbooleanstring, bool) {
	if o == nil || o.InstallAgent == nil {
		return nil, false
	}
	return o.InstallAgent, true
}

// HasInstallAgent returns a boolean if a field has been set.
func (o *BlueprintARMCreateArm) HasInstallAgent() bool {
	if o != nil && o.InstallAgent != nil {
		return true
	}

	return false
}

// SetInstallAgent gets a reference to the given OneOfbooleanstring and assigns it to the InstallAgent field.
func (o *BlueprintARMCreateArm) SetInstallAgent(v OneOfbooleanstring) {
	o.InstallAgent = &v
}

// GetCloudInitEnabled returns the CloudInitEnabled field value if set, zero value otherwise.
func (o *BlueprintARMCreateArm) GetCloudInitEnabled() OneOfbooleanstring {
	if o == nil || o.CloudInitEnabled == nil {
		var ret OneOfbooleanstring
		return ret
	}
	return *o.CloudInitEnabled
}

// GetCloudInitEnabledOk returns a tuple with the CloudInitEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintARMCreateArm) GetCloudInitEnabledOk() (*OneOfbooleanstring, bool) {
	if o == nil || o.CloudInitEnabled == nil {
		return nil, false
	}
	return o.CloudInitEnabled, true
}

// HasCloudInitEnabled returns a boolean if a field has been set.
func (o *BlueprintARMCreateArm) HasCloudInitEnabled() bool {
	if o != nil && o.CloudInitEnabled != nil {
		return true
	}

	return false
}

// SetCloudInitEnabled gets a reference to the given OneOfbooleanstring and assigns it to the CloudInitEnabled field.
func (o *BlueprintARMCreateArm) SetCloudInitEnabled(v OneOfbooleanstring) {
	o.CloudInitEnabled = &v
}

func (o BlueprintARMCreateArm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["configType"] = o.ConfigType
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.Yaml != nil {
		toSerialize["yaml"] = o.Yaml
	}
	if o.Git != nil {
		toSerialize["git"] = o.Git
	}
	if o.OsType != nil {
		toSerialize["osType"] = o.OsType
	}
	if o.InstallAgent != nil {
		toSerialize["installAgent"] = o.InstallAgent
	}
	if o.CloudInitEnabled != nil {
		toSerialize["cloudInitEnabled"] = o.CloudInitEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableBlueprintARMCreateArm struct {
	value *BlueprintARMCreateArm
	isSet bool
}

func (v NullableBlueprintARMCreateArm) Get() *BlueprintARMCreateArm {
	return v.value
}

func (v *NullableBlueprintARMCreateArm) Set(val *BlueprintARMCreateArm) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintARMCreateArm) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintARMCreateArm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintARMCreateArm(val *BlueprintARMCreateArm) *NullableBlueprintARMCreateArm {
	return &NullableBlueprintARMCreateArm{value: val, isSet: true}
}

func (v NullableBlueprintARMCreateArm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintARMCreateArm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

