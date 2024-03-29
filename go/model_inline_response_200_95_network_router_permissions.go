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

// InlineResponse20095NetworkRouterPermissions struct for InlineResponse20095NetworkRouterPermissions
type InlineResponse20095NetworkRouterPermissions struct {
	Visibility *string `json:"visibility,omitempty"`
	TenantPermissions *InlineResponse20095NetworkRouterPermissionsTenantPermissions `json:"tenantPermissions,omitempty"`
}

// NewInlineResponse20095NetworkRouterPermissions instantiates a new InlineResponse20095NetworkRouterPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20095NetworkRouterPermissions() *InlineResponse20095NetworkRouterPermissions {
	this := InlineResponse20095NetworkRouterPermissions{}
	return &this
}

// NewInlineResponse20095NetworkRouterPermissionsWithDefaults instantiates a new InlineResponse20095NetworkRouterPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20095NetworkRouterPermissionsWithDefaults() *InlineResponse20095NetworkRouterPermissions {
	this := InlineResponse20095NetworkRouterPermissions{}
	return &this
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouterPermissions) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouterPermissions) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouterPermissions) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *InlineResponse20095NetworkRouterPermissions) SetVisibility(v string) {
	o.Visibility = &v
}

// GetTenantPermissions returns the TenantPermissions field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouterPermissions) GetTenantPermissions() InlineResponse20095NetworkRouterPermissionsTenantPermissions {
	if o == nil || o.TenantPermissions == nil {
		var ret InlineResponse20095NetworkRouterPermissionsTenantPermissions
		return ret
	}
	return *o.TenantPermissions
}

// GetTenantPermissionsOk returns a tuple with the TenantPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouterPermissions) GetTenantPermissionsOk() (*InlineResponse20095NetworkRouterPermissionsTenantPermissions, bool) {
	if o == nil || o.TenantPermissions == nil {
		return nil, false
	}
	return o.TenantPermissions, true
}

// HasTenantPermissions returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouterPermissions) HasTenantPermissions() bool {
	if o != nil && o.TenantPermissions != nil {
		return true
	}

	return false
}

// SetTenantPermissions gets a reference to the given InlineResponse20095NetworkRouterPermissionsTenantPermissions and assigns it to the TenantPermissions field.
func (o *InlineResponse20095NetworkRouterPermissions) SetTenantPermissions(v InlineResponse20095NetworkRouterPermissionsTenantPermissions) {
	o.TenantPermissions = &v
}

func (o InlineResponse20095NetworkRouterPermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.TenantPermissions != nil {
		toSerialize["tenantPermissions"] = o.TenantPermissions
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20095NetworkRouterPermissions struct {
	value *InlineResponse20095NetworkRouterPermissions
	isSet bool
}

func (v NullableInlineResponse20095NetworkRouterPermissions) Get() *InlineResponse20095NetworkRouterPermissions {
	return v.value
}

func (v *NullableInlineResponse20095NetworkRouterPermissions) Set(val *InlineResponse20095NetworkRouterPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20095NetworkRouterPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20095NetworkRouterPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20095NetworkRouterPermissions(val *InlineResponse20095NetworkRouterPermissions) *NullableInlineResponse20095NetworkRouterPermissions {
	return &NullableInlineResponse20095NetworkRouterPermissions{value: val, isSet: true}
}

func (v NullableInlineResponse20095NetworkRouterPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20095NetworkRouterPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


