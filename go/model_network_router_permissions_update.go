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

// NetworkRouterPermissionsUpdate struct for NetworkRouterPermissionsUpdate
type NetworkRouterPermissionsUpdate struct {
	// Sets visibility - public, private
	Visibility *string `json:"visibility,omitempty"`
	TenantPermissions *NetworkRouterPermissionsUpdateTenantPermissions `json:"tenantPermissions,omitempty"`
}

// NewNetworkRouterPermissionsUpdate instantiates a new NetworkRouterPermissionsUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkRouterPermissionsUpdate() *NetworkRouterPermissionsUpdate {
	this := NetworkRouterPermissionsUpdate{}
	return &this
}

// NewNetworkRouterPermissionsUpdateWithDefaults instantiates a new NetworkRouterPermissionsUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkRouterPermissionsUpdateWithDefaults() *NetworkRouterPermissionsUpdate {
	this := NetworkRouterPermissionsUpdate{}
	return &this
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *NetworkRouterPermissionsUpdate) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRouterPermissionsUpdate) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *NetworkRouterPermissionsUpdate) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *NetworkRouterPermissionsUpdate) SetVisibility(v string) {
	o.Visibility = &v
}

// GetTenantPermissions returns the TenantPermissions field value if set, zero value otherwise.
func (o *NetworkRouterPermissionsUpdate) GetTenantPermissions() NetworkRouterPermissionsUpdateTenantPermissions {
	if o == nil || o.TenantPermissions == nil {
		var ret NetworkRouterPermissionsUpdateTenantPermissions
		return ret
	}
	return *o.TenantPermissions
}

// GetTenantPermissionsOk returns a tuple with the TenantPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRouterPermissionsUpdate) GetTenantPermissionsOk() (*NetworkRouterPermissionsUpdateTenantPermissions, bool) {
	if o == nil || o.TenantPermissions == nil {
		return nil, false
	}
	return o.TenantPermissions, true
}

// HasTenantPermissions returns a boolean if a field has been set.
func (o *NetworkRouterPermissionsUpdate) HasTenantPermissions() bool {
	if o != nil && o.TenantPermissions != nil {
		return true
	}

	return false
}

// SetTenantPermissions gets a reference to the given NetworkRouterPermissionsUpdateTenantPermissions and assigns it to the TenantPermissions field.
func (o *NetworkRouterPermissionsUpdate) SetTenantPermissions(v NetworkRouterPermissionsUpdateTenantPermissions) {
	o.TenantPermissions = &v
}

func (o NetworkRouterPermissionsUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.TenantPermissions != nil {
		toSerialize["tenantPermissions"] = o.TenantPermissions
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkRouterPermissionsUpdate struct {
	value *NetworkRouterPermissionsUpdate
	isSet bool
}

func (v NullableNetworkRouterPermissionsUpdate) Get() *NetworkRouterPermissionsUpdate {
	return v.value
}

func (v *NullableNetworkRouterPermissionsUpdate) Set(val *NetworkRouterPermissionsUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkRouterPermissionsUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkRouterPermissionsUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkRouterPermissionsUpdate(val *NetworkRouterPermissionsUpdate) *NullableNetworkRouterPermissionsUpdate {
	return &NullableNetworkRouterPermissionsUpdate{value: val, isSet: true}
}

func (v NullableNetworkRouterPermissionsUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkRouterPermissionsUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

