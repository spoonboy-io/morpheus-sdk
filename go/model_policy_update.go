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

// PolicyUpdate struct for PolicyUpdate
type PolicyUpdate struct {
	// A name for the policy
	Name *string `json:"name,omitempty"`
	// A description for the policy
	Description *string `json:"description,omitempty"`
	// A map of config values. The expected values vary by policy type. See `Retrieves all Policy Types` endpoint for `fieldName`(s) of required options.
	Config *AnyOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject `json:"config,omitempty"`
	// Set to false to disable
	Enabled *bool `json:"enabled,omitempty"`
	// Scope object type
	RefType *string `json:"refType,omitempty"`
	// Scope object ID (`group`,`cloud`,`user`, etc)
	RefId *int64 `json:"refId,omitempty"`
	// Array of tenants to scope the policy to
	Accounts *[]int64 `json:"accounts,omitempty"`
	// Apply individually to each user in role.  Only when `refType` equals `Role`
	EachUser *bool `json:"eachUser,omitempty"`
}

// NewPolicyUpdate instantiates a new PolicyUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyUpdate() *PolicyUpdate {
	this := PolicyUpdate{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// NewPolicyUpdateWithDefaults instantiates a new PolicyUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyUpdateWithDefaults() *PolicyUpdate {
	this := PolicyUpdate{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PolicyUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PolicyUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PolicyUpdate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PolicyUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PolicyUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PolicyUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *PolicyUpdate) GetConfig() AnyOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject {
	if o == nil || o.Config == nil {
		var ret AnyOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyUpdate) GetConfigOk() (*AnyOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *PolicyUpdate) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given AnyOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject and assigns it to the Config field.
func (o *PolicyUpdate) SetConfig(v AnyOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject) {
	o.Config = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PolicyUpdate) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyUpdate) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PolicyUpdate) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PolicyUpdate) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRefType returns the RefType field value if set, zero value otherwise.
func (o *PolicyUpdate) GetRefType() string {
	if o == nil || o.RefType == nil {
		var ret string
		return ret
	}
	return *o.RefType
}

// GetRefTypeOk returns a tuple with the RefType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyUpdate) GetRefTypeOk() (*string, bool) {
	if o == nil || o.RefType == nil {
		return nil, false
	}
	return o.RefType, true
}

// HasRefType returns a boolean if a field has been set.
func (o *PolicyUpdate) HasRefType() bool {
	if o != nil && o.RefType != nil {
		return true
	}

	return false
}

// SetRefType gets a reference to the given string and assigns it to the RefType field.
func (o *PolicyUpdate) SetRefType(v string) {
	o.RefType = &v
}

// GetRefId returns the RefId field value if set, zero value otherwise.
func (o *PolicyUpdate) GetRefId() int64 {
	if o == nil || o.RefId == nil {
		var ret int64
		return ret
	}
	return *o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyUpdate) GetRefIdOk() (*int64, bool) {
	if o == nil || o.RefId == nil {
		return nil, false
	}
	return o.RefId, true
}

// HasRefId returns a boolean if a field has been set.
func (o *PolicyUpdate) HasRefId() bool {
	if o != nil && o.RefId != nil {
		return true
	}

	return false
}

// SetRefId gets a reference to the given int64 and assigns it to the RefId field.
func (o *PolicyUpdate) SetRefId(v int64) {
	o.RefId = &v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *PolicyUpdate) GetAccounts() []int64 {
	if o == nil || o.Accounts == nil {
		var ret []int64
		return ret
	}
	return *o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyUpdate) GetAccountsOk() (*[]int64, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *PolicyUpdate) HasAccounts() bool {
	if o != nil && o.Accounts != nil {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []int64 and assigns it to the Accounts field.
func (o *PolicyUpdate) SetAccounts(v []int64) {
	o.Accounts = &v
}

// GetEachUser returns the EachUser field value if set, zero value otherwise.
func (o *PolicyUpdate) GetEachUser() bool {
	if o == nil || o.EachUser == nil {
		var ret bool
		return ret
	}
	return *o.EachUser
}

// GetEachUserOk returns a tuple with the EachUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyUpdate) GetEachUserOk() (*bool, bool) {
	if o == nil || o.EachUser == nil {
		return nil, false
	}
	return o.EachUser, true
}

// HasEachUser returns a boolean if a field has been set.
func (o *PolicyUpdate) HasEachUser() bool {
	if o != nil && o.EachUser != nil {
		return true
	}

	return false
}

// SetEachUser gets a reference to the given bool and assigns it to the EachUser field.
func (o *PolicyUpdate) SetEachUser(v bool) {
	o.EachUser = &v
}

func (o PolicyUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.RefType != nil {
		toSerialize["refType"] = o.RefType
	}
	if o.RefId != nil {
		toSerialize["refId"] = o.RefId
	}
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}
	if o.EachUser != nil {
		toSerialize["eachUser"] = o.EachUser
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyUpdate struct {
	value *PolicyUpdate
	isSet bool
}

func (v NullablePolicyUpdate) Get() *PolicyUpdate {
	return v.value
}

func (v *NullablePolicyUpdate) Set(val *PolicyUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyUpdate(val *PolicyUpdate) *NullablePolicyUpdate {
	return &NullablePolicyUpdate{value: val, isSet: true}
}

func (v NullablePolicyUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


