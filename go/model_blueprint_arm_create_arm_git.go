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

// BlueprintARMCreateArmGit struct for BlueprintARMCreateArmGit
type BlueprintARMCreateArmGit struct {
	// Morpheus SCM Repository ID
	RepoId int64 `json:"repoId"`
	// Path to ARM Files in the Repository
	Path string `json:"path"`
	// Morpheus SCM Integration ID
	IntegrationId int64 `json:"integrationId"`
	// Branch Name
	Branch string `json:"branch"`
}

// NewBlueprintARMCreateArmGit instantiates a new BlueprintARMCreateArmGit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlueprintARMCreateArmGit(repoId int64, path string, integrationId int64, branch string, ) *BlueprintARMCreateArmGit {
	this := BlueprintARMCreateArmGit{}
	this.RepoId = repoId
	this.Path = path
	this.IntegrationId = integrationId
	this.Branch = branch
	return &this
}

// NewBlueprintARMCreateArmGitWithDefaults instantiates a new BlueprintARMCreateArmGit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlueprintARMCreateArmGitWithDefaults() *BlueprintARMCreateArmGit {
	this := BlueprintARMCreateArmGit{}
	return &this
}

// GetRepoId returns the RepoId field value
func (o *BlueprintARMCreateArmGit) GetRepoId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.RepoId
}

// GetRepoIdOk returns a tuple with the RepoId field value
// and a boolean to check if the value has been set.
func (o *BlueprintARMCreateArmGit) GetRepoIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RepoId, true
}

// SetRepoId sets field value
func (o *BlueprintARMCreateArmGit) SetRepoId(v int64) {
	o.RepoId = v
}

// GetPath returns the Path field value
func (o *BlueprintARMCreateArmGit) GetPath() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *BlueprintARMCreateArmGit) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *BlueprintARMCreateArmGit) SetPath(v string) {
	o.Path = v
}

// GetIntegrationId returns the IntegrationId field value
func (o *BlueprintARMCreateArmGit) GetIntegrationId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value
// and a boolean to check if the value has been set.
func (o *BlueprintARMCreateArmGit) GetIntegrationIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IntegrationId, true
}

// SetIntegrationId sets field value
func (o *BlueprintARMCreateArmGit) SetIntegrationId(v int64) {
	o.IntegrationId = v
}

// GetBranch returns the Branch field value
func (o *BlueprintARMCreateArmGit) GetBranch() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Branch
}

// GetBranchOk returns a tuple with the Branch field value
// and a boolean to check if the value has been set.
func (o *BlueprintARMCreateArmGit) GetBranchOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Branch, true
}

// SetBranch sets field value
func (o *BlueprintARMCreateArmGit) SetBranch(v string) {
	o.Branch = v
}

func (o BlueprintARMCreateArmGit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["repoId"] = o.RepoId
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["integrationId"] = o.IntegrationId
	}
	if true {
		toSerialize["branch"] = o.Branch
	}
	return json.Marshal(toSerialize)
}

type NullableBlueprintARMCreateArmGit struct {
	value *BlueprintARMCreateArmGit
	isSet bool
}

func (v NullableBlueprintARMCreateArmGit) Get() *BlueprintARMCreateArmGit {
	return v.value
}

func (v *NullableBlueprintARMCreateArmGit) Set(val *BlueprintARMCreateArmGit) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintARMCreateArmGit) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintARMCreateArmGit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintARMCreateArmGit(val *BlueprintARMCreateArmGit) *NullableBlueprintARMCreateArmGit {
	return &NullableBlueprintARMCreateArmGit{value: val, isSet: true}
}

func (v NullableBlueprintARMCreateArmGit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintARMCreateArmGit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


