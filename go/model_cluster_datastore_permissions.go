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

// ClusterDatastorePermissions struct for ClusterDatastorePermissions
type ClusterDatastorePermissions struct {
	ResourcePermissions *ClusterDatastorePermissionsResourcePermissions `json:"resourcePermissions,omitempty"`
	TenantPermissions *InlineResponse20095NetworkRouterPermissionsTenantPermissions `json:"tenantPermissions,omitempty"`
}

// NewClusterDatastorePermissions instantiates a new ClusterDatastorePermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterDatastorePermissions() *ClusterDatastorePermissions {
	this := ClusterDatastorePermissions{}
	return &this
}

// NewClusterDatastorePermissionsWithDefaults instantiates a new ClusterDatastorePermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterDatastorePermissionsWithDefaults() *ClusterDatastorePermissions {
	this := ClusterDatastorePermissions{}
	return &this
}

// GetResourcePermissions returns the ResourcePermissions field value if set, zero value otherwise.
func (o *ClusterDatastorePermissions) GetResourcePermissions() ClusterDatastorePermissionsResourcePermissions {
	if o == nil || o.ResourcePermissions == nil {
		var ret ClusterDatastorePermissionsResourcePermissions
		return ret
	}
	return *o.ResourcePermissions
}

// GetResourcePermissionsOk returns a tuple with the ResourcePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDatastorePermissions) GetResourcePermissionsOk() (*ClusterDatastorePermissionsResourcePermissions, bool) {
	if o == nil || o.ResourcePermissions == nil {
		return nil, false
	}
	return o.ResourcePermissions, true
}

// HasResourcePermissions returns a boolean if a field has been set.
func (o *ClusterDatastorePermissions) HasResourcePermissions() bool {
	if o != nil && o.ResourcePermissions != nil {
		return true
	}

	return false
}

// SetResourcePermissions gets a reference to the given ClusterDatastorePermissionsResourcePermissions and assigns it to the ResourcePermissions field.
func (o *ClusterDatastorePermissions) SetResourcePermissions(v ClusterDatastorePermissionsResourcePermissions) {
	o.ResourcePermissions = &v
}

// GetTenantPermissions returns the TenantPermissions field value if set, zero value otherwise.
func (o *ClusterDatastorePermissions) GetTenantPermissions() InlineResponse20095NetworkRouterPermissionsTenantPermissions {
	if o == nil || o.TenantPermissions == nil {
		var ret InlineResponse20095NetworkRouterPermissionsTenantPermissions
		return ret
	}
	return *o.TenantPermissions
}

// GetTenantPermissionsOk returns a tuple with the TenantPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDatastorePermissions) GetTenantPermissionsOk() (*InlineResponse20095NetworkRouterPermissionsTenantPermissions, bool) {
	if o == nil || o.TenantPermissions == nil {
		return nil, false
	}
	return o.TenantPermissions, true
}

// HasTenantPermissions returns a boolean if a field has been set.
func (o *ClusterDatastorePermissions) HasTenantPermissions() bool {
	if o != nil && o.TenantPermissions != nil {
		return true
	}

	return false
}

// SetTenantPermissions gets a reference to the given InlineResponse20095NetworkRouterPermissionsTenantPermissions and assigns it to the TenantPermissions field.
func (o *ClusterDatastorePermissions) SetTenantPermissions(v InlineResponse20095NetworkRouterPermissionsTenantPermissions) {
	o.TenantPermissions = &v
}

func (o ClusterDatastorePermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourcePermissions != nil {
		toSerialize["resourcePermissions"] = o.ResourcePermissions
	}
	if o.TenantPermissions != nil {
		toSerialize["tenantPermissions"] = o.TenantPermissions
	}
	return json.Marshal(toSerialize)
}

type NullableClusterDatastorePermissions struct {
	value *ClusterDatastorePermissions
	isSet bool
}

func (v NullableClusterDatastorePermissions) Get() *ClusterDatastorePermissions {
	return v.value
}

func (v *NullableClusterDatastorePermissions) Set(val *ClusterDatastorePermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterDatastorePermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterDatastorePermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterDatastorePermissions(val *ClusterDatastorePermissions) *NullableClusterDatastorePermissions {
	return &NullableClusterDatastorePermissions{value: val, isSet: true}
}

func (v NullableClusterDatastorePermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterDatastorePermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


