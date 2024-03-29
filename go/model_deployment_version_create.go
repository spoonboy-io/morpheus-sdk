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

// DeploymentVersionCreate struct for DeploymentVersionCreate
type DeploymentVersionCreate struct {
	// Version number (userVersion), a unique version identifier for the deployment version.
	Version *string `json:"version,omitempty"`
	// Alias for version
	UserVersion *string `json:"userVersion,omitempty"`
	// Deploy Type, eg. file, git, fetch
	DeployType *string `json:"deployType,omitempty"`
	GitUrl NullableString `json:"gitUrl,omitempty"`
	GitRef NullableString `json:"gitRef,omitempty"`
	FetchUrl NullableString `json:"fetchUrl,omitempty"`
}

// NewDeploymentVersionCreate instantiates a new DeploymentVersionCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentVersionCreate() *DeploymentVersionCreate {
	this := DeploymentVersionCreate{}
	return &this
}

// NewDeploymentVersionCreateWithDefaults instantiates a new DeploymentVersionCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentVersionCreateWithDefaults() *DeploymentVersionCreate {
	this := DeploymentVersionCreate{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DeploymentVersionCreate) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentVersionCreate) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DeploymentVersionCreate) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *DeploymentVersionCreate) SetVersion(v string) {
	o.Version = &v
}

// GetUserVersion returns the UserVersion field value if set, zero value otherwise.
func (o *DeploymentVersionCreate) GetUserVersion() string {
	if o == nil || o.UserVersion == nil {
		var ret string
		return ret
	}
	return *o.UserVersion
}

// GetUserVersionOk returns a tuple with the UserVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentVersionCreate) GetUserVersionOk() (*string, bool) {
	if o == nil || o.UserVersion == nil {
		return nil, false
	}
	return o.UserVersion, true
}

// HasUserVersion returns a boolean if a field has been set.
func (o *DeploymentVersionCreate) HasUserVersion() bool {
	if o != nil && o.UserVersion != nil {
		return true
	}

	return false
}

// SetUserVersion gets a reference to the given string and assigns it to the UserVersion field.
func (o *DeploymentVersionCreate) SetUserVersion(v string) {
	o.UserVersion = &v
}

// GetDeployType returns the DeployType field value if set, zero value otherwise.
func (o *DeploymentVersionCreate) GetDeployType() string {
	if o == nil || o.DeployType == nil {
		var ret string
		return ret
	}
	return *o.DeployType
}

// GetDeployTypeOk returns a tuple with the DeployType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentVersionCreate) GetDeployTypeOk() (*string, bool) {
	if o == nil || o.DeployType == nil {
		return nil, false
	}
	return o.DeployType, true
}

// HasDeployType returns a boolean if a field has been set.
func (o *DeploymentVersionCreate) HasDeployType() bool {
	if o != nil && o.DeployType != nil {
		return true
	}

	return false
}

// SetDeployType gets a reference to the given string and assigns it to the DeployType field.
func (o *DeploymentVersionCreate) SetDeployType(v string) {
	o.DeployType = &v
}

// GetGitUrl returns the GitUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentVersionCreate) GetGitUrl() string {
	if o == nil || o.GitUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.GitUrl.Get()
}

// GetGitUrlOk returns a tuple with the GitUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentVersionCreate) GetGitUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GitUrl.Get(), o.GitUrl.IsSet()
}

// HasGitUrl returns a boolean if a field has been set.
func (o *DeploymentVersionCreate) HasGitUrl() bool {
	if o != nil && o.GitUrl.IsSet() {
		return true
	}

	return false
}

// SetGitUrl gets a reference to the given NullableString and assigns it to the GitUrl field.
func (o *DeploymentVersionCreate) SetGitUrl(v string) {
	o.GitUrl.Set(&v)
}
// SetGitUrlNil sets the value for GitUrl to be an explicit nil
func (o *DeploymentVersionCreate) SetGitUrlNil() {
	o.GitUrl.Set(nil)
}

// UnsetGitUrl ensures that no value is present for GitUrl, not even an explicit nil
func (o *DeploymentVersionCreate) UnsetGitUrl() {
	o.GitUrl.Unset()
}

// GetGitRef returns the GitRef field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentVersionCreate) GetGitRef() string {
	if o == nil || o.GitRef.Get() == nil {
		var ret string
		return ret
	}
	return *o.GitRef.Get()
}

// GetGitRefOk returns a tuple with the GitRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentVersionCreate) GetGitRefOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GitRef.Get(), o.GitRef.IsSet()
}

// HasGitRef returns a boolean if a field has been set.
func (o *DeploymentVersionCreate) HasGitRef() bool {
	if o != nil && o.GitRef.IsSet() {
		return true
	}

	return false
}

// SetGitRef gets a reference to the given NullableString and assigns it to the GitRef field.
func (o *DeploymentVersionCreate) SetGitRef(v string) {
	o.GitRef.Set(&v)
}
// SetGitRefNil sets the value for GitRef to be an explicit nil
func (o *DeploymentVersionCreate) SetGitRefNil() {
	o.GitRef.Set(nil)
}

// UnsetGitRef ensures that no value is present for GitRef, not even an explicit nil
func (o *DeploymentVersionCreate) UnsetGitRef() {
	o.GitRef.Unset()
}

// GetFetchUrl returns the FetchUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentVersionCreate) GetFetchUrl() string {
	if o == nil || o.FetchUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.FetchUrl.Get()
}

// GetFetchUrlOk returns a tuple with the FetchUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentVersionCreate) GetFetchUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FetchUrl.Get(), o.FetchUrl.IsSet()
}

// HasFetchUrl returns a boolean if a field has been set.
func (o *DeploymentVersionCreate) HasFetchUrl() bool {
	if o != nil && o.FetchUrl.IsSet() {
		return true
	}

	return false
}

// SetFetchUrl gets a reference to the given NullableString and assigns it to the FetchUrl field.
func (o *DeploymentVersionCreate) SetFetchUrl(v string) {
	o.FetchUrl.Set(&v)
}
// SetFetchUrlNil sets the value for FetchUrl to be an explicit nil
func (o *DeploymentVersionCreate) SetFetchUrlNil() {
	o.FetchUrl.Set(nil)
}

// UnsetFetchUrl ensures that no value is present for FetchUrl, not even an explicit nil
func (o *DeploymentVersionCreate) UnsetFetchUrl() {
	o.FetchUrl.Unset()
}

func (o DeploymentVersionCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.UserVersion != nil {
		toSerialize["userVersion"] = o.UserVersion
	}
	if o.DeployType != nil {
		toSerialize["deployType"] = o.DeployType
	}
	if o.GitUrl.IsSet() {
		toSerialize["gitUrl"] = o.GitUrl.Get()
	}
	if o.GitRef.IsSet() {
		toSerialize["gitRef"] = o.GitRef.Get()
	}
	if o.FetchUrl.IsSet() {
		toSerialize["fetchUrl"] = o.FetchUrl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentVersionCreate struct {
	value *DeploymentVersionCreate
	isSet bool
}

func (v NullableDeploymentVersionCreate) Get() *DeploymentVersionCreate {
	return v.value
}

func (v *NullableDeploymentVersionCreate) Set(val *DeploymentVersionCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentVersionCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentVersionCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentVersionCreate(val *DeploymentVersionCreate) *NullableDeploymentVersionCreate {
	return &NullableDeploymentVersionCreate{value: val, isSet: true}
}

func (v NullableDeploymentVersionCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentVersionCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


