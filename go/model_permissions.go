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

// Permissions struct for Permissions
type Permissions struct {
	ResourcePool *ClusterPermissionsResourcePool `json:"resourcePool,omitempty"`
	ResourcePermissions *ResourcePermissions `json:"resourcePermissions,omitempty"`
	TenantPermissions *ApiZonesZoneIdFoldersIdFolderTenantPermissions `json:"tenantPermissions,omitempty"`
}

// NewPermissions instantiates a new Permissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissions() *Permissions {
	this := Permissions{}
	return &this
}

// NewPermissionsWithDefaults instantiates a new Permissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionsWithDefaults() *Permissions {
	this := Permissions{}
	return &this
}

// GetResourcePool returns the ResourcePool field value if set, zero value otherwise.
func (o *Permissions) GetResourcePool() ClusterPermissionsResourcePool {
	if o == nil || o.ResourcePool == nil {
		var ret ClusterPermissionsResourcePool
		return ret
	}
	return *o.ResourcePool
}

// GetResourcePoolOk returns a tuple with the ResourcePool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetResourcePoolOk() (*ClusterPermissionsResourcePool, bool) {
	if o == nil || o.ResourcePool == nil {
		return nil, false
	}
	return o.ResourcePool, true
}

// HasResourcePool returns a boolean if a field has been set.
func (o *Permissions) HasResourcePool() bool {
	if o != nil && o.ResourcePool != nil {
		return true
	}

	return false
}

// SetResourcePool gets a reference to the given ClusterPermissionsResourcePool and assigns it to the ResourcePool field.
func (o *Permissions) SetResourcePool(v ClusterPermissionsResourcePool) {
	o.ResourcePool = &v
}

// GetResourcePermissions returns the ResourcePermissions field value if set, zero value otherwise.
func (o *Permissions) GetResourcePermissions() ResourcePermissions {
	if o == nil || o.ResourcePermissions == nil {
		var ret ResourcePermissions
		return ret
	}
	return *o.ResourcePermissions
}

// GetResourcePermissionsOk returns a tuple with the ResourcePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetResourcePermissionsOk() (*ResourcePermissions, bool) {
	if o == nil || o.ResourcePermissions == nil {
		return nil, false
	}
	return o.ResourcePermissions, true
}

// HasResourcePermissions returns a boolean if a field has been set.
func (o *Permissions) HasResourcePermissions() bool {
	if o != nil && o.ResourcePermissions != nil {
		return true
	}

	return false
}

// SetResourcePermissions gets a reference to the given ResourcePermissions and assigns it to the ResourcePermissions field.
func (o *Permissions) SetResourcePermissions(v ResourcePermissions) {
	o.ResourcePermissions = &v
}

// GetTenantPermissions returns the TenantPermissions field value if set, zero value otherwise.
func (o *Permissions) GetTenantPermissions() ApiZonesZoneIdFoldersIdFolderTenantPermissions {
	if o == nil || o.TenantPermissions == nil {
		var ret ApiZonesZoneIdFoldersIdFolderTenantPermissions
		return ret
	}
	return *o.TenantPermissions
}

// GetTenantPermissionsOk returns a tuple with the TenantPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetTenantPermissionsOk() (*ApiZonesZoneIdFoldersIdFolderTenantPermissions, bool) {
	if o == nil || o.TenantPermissions == nil {
		return nil, false
	}
	return o.TenantPermissions, true
}

// HasTenantPermissions returns a boolean if a field has been set.
func (o *Permissions) HasTenantPermissions() bool {
	if o != nil && o.TenantPermissions != nil {
		return true
	}

	return false
}

// SetTenantPermissions gets a reference to the given ApiZonesZoneIdFoldersIdFolderTenantPermissions and assigns it to the TenantPermissions field.
func (o *Permissions) SetTenantPermissions(v ApiZonesZoneIdFoldersIdFolderTenantPermissions) {
	o.TenantPermissions = &v
}

func (o Permissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourcePool != nil {
		toSerialize["resourcePool"] = o.ResourcePool
	}
	if o.ResourcePermissions != nil {
		toSerialize["resourcePermissions"] = o.ResourcePermissions
	}
	if o.TenantPermissions != nil {
		toSerialize["tenantPermissions"] = o.TenantPermissions
	}
	return json.Marshal(toSerialize)
}

type NullablePermissions struct {
	value *Permissions
	isSet bool
}

func (v NullablePermissions) Get() *Permissions {
	return v.value
}

func (v *NullablePermissions) Set(val *Permissions) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissions) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissions(val *Permissions) *NullablePermissions {
	return &NullablePermissions{value: val, isSet: true}
}

func (v NullablePermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


