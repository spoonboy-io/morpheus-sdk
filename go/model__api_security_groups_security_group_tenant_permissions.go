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

// ApiSecurityGroupsSecurityGroupTenantPermissions struct for ApiSecurityGroupsSecurityGroupTenantPermissions
type ApiSecurityGroupsSecurityGroupTenantPermissions struct {
	// Array of tenant account ids that are allowed access
	Accounts *[]int64 `json:"accounts,omitempty"`
	// Array of tenant account ids that can manage
	CanManageAccounts *[]int64 `json:"canManageAccounts,omitempty"`
}

// NewApiSecurityGroupsSecurityGroupTenantPermissions instantiates a new ApiSecurityGroupsSecurityGroupTenantPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSecurityGroupsSecurityGroupTenantPermissions() *ApiSecurityGroupsSecurityGroupTenantPermissions {
	this := ApiSecurityGroupsSecurityGroupTenantPermissions{}
	return &this
}

// NewApiSecurityGroupsSecurityGroupTenantPermissionsWithDefaults instantiates a new ApiSecurityGroupsSecurityGroupTenantPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSecurityGroupsSecurityGroupTenantPermissionsWithDefaults() *ApiSecurityGroupsSecurityGroupTenantPermissions {
	this := ApiSecurityGroupsSecurityGroupTenantPermissions{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *ApiSecurityGroupsSecurityGroupTenantPermissions) GetAccounts() []int64 {
	if o == nil || o.Accounts == nil {
		var ret []int64
		return ret
	}
	return *o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSecurityGroupsSecurityGroupTenantPermissions) GetAccountsOk() (*[]int64, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *ApiSecurityGroupsSecurityGroupTenantPermissions) HasAccounts() bool {
	if o != nil && o.Accounts != nil {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []int64 and assigns it to the Accounts field.
func (o *ApiSecurityGroupsSecurityGroupTenantPermissions) SetAccounts(v []int64) {
	o.Accounts = &v
}

// GetCanManageAccounts returns the CanManageAccounts field value if set, zero value otherwise.
func (o *ApiSecurityGroupsSecurityGroupTenantPermissions) GetCanManageAccounts() []int64 {
	if o == nil || o.CanManageAccounts == nil {
		var ret []int64
		return ret
	}
	return *o.CanManageAccounts
}

// GetCanManageAccountsOk returns a tuple with the CanManageAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSecurityGroupsSecurityGroupTenantPermissions) GetCanManageAccountsOk() (*[]int64, bool) {
	if o == nil || o.CanManageAccounts == nil {
		return nil, false
	}
	return o.CanManageAccounts, true
}

// HasCanManageAccounts returns a boolean if a field has been set.
func (o *ApiSecurityGroupsSecurityGroupTenantPermissions) HasCanManageAccounts() bool {
	if o != nil && o.CanManageAccounts != nil {
		return true
	}

	return false
}

// SetCanManageAccounts gets a reference to the given []int64 and assigns it to the CanManageAccounts field.
func (o *ApiSecurityGroupsSecurityGroupTenantPermissions) SetCanManageAccounts(v []int64) {
	o.CanManageAccounts = &v
}

func (o ApiSecurityGroupsSecurityGroupTenantPermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}
	if o.CanManageAccounts != nil {
		toSerialize["canManageAccounts"] = o.CanManageAccounts
	}
	return json.Marshal(toSerialize)
}

type NullableApiSecurityGroupsSecurityGroupTenantPermissions struct {
	value *ApiSecurityGroupsSecurityGroupTenantPermissions
	isSet bool
}

func (v NullableApiSecurityGroupsSecurityGroupTenantPermissions) Get() *ApiSecurityGroupsSecurityGroupTenantPermissions {
	return v.value
}

func (v *NullableApiSecurityGroupsSecurityGroupTenantPermissions) Set(val *ApiSecurityGroupsSecurityGroupTenantPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSecurityGroupsSecurityGroupTenantPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSecurityGroupsSecurityGroupTenantPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSecurityGroupsSecurityGroupTenantPermissions(val *ApiSecurityGroupsSecurityGroupTenantPermissions) *NullableApiSecurityGroupsSecurityGroupTenantPermissions {
	return &NullableApiSecurityGroupsSecurityGroupTenantPermissions{value: val, isSet: true}
}

func (v NullableApiSecurityGroupsSecurityGroupTenantPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSecurityGroupsSecurityGroupTenantPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


