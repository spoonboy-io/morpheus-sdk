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

// ClusterNamespacePermissions struct for ClusterNamespacePermissions
type ClusterNamespacePermissions struct {
	ResourcePermissions *ClusterNamespacePermissionsResourcePermissions `json:"resourcePermissions,omitempty"`
}

// NewClusterNamespacePermissions instantiates a new ClusterNamespacePermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterNamespacePermissions() *ClusterNamespacePermissions {
	this := ClusterNamespacePermissions{}
	return &this
}

// NewClusterNamespacePermissionsWithDefaults instantiates a new ClusterNamespacePermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterNamespacePermissionsWithDefaults() *ClusterNamespacePermissions {
	this := ClusterNamespacePermissions{}
	return &this
}

// GetResourcePermissions returns the ResourcePermissions field value if set, zero value otherwise.
func (o *ClusterNamespacePermissions) GetResourcePermissions() ClusterNamespacePermissionsResourcePermissions {
	if o == nil || o.ResourcePermissions == nil {
		var ret ClusterNamespacePermissionsResourcePermissions
		return ret
	}
	return *o.ResourcePermissions
}

// GetResourcePermissionsOk returns a tuple with the ResourcePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNamespacePermissions) GetResourcePermissionsOk() (*ClusterNamespacePermissionsResourcePermissions, bool) {
	if o == nil || o.ResourcePermissions == nil {
		return nil, false
	}
	return o.ResourcePermissions, true
}

// HasResourcePermissions returns a boolean if a field has been set.
func (o *ClusterNamespacePermissions) HasResourcePermissions() bool {
	if o != nil && o.ResourcePermissions != nil {
		return true
	}

	return false
}

// SetResourcePermissions gets a reference to the given ClusterNamespacePermissionsResourcePermissions and assigns it to the ResourcePermissions field.
func (o *ClusterNamespacePermissions) SetResourcePermissions(v ClusterNamespacePermissionsResourcePermissions) {
	o.ResourcePermissions = &v
}

func (o ClusterNamespacePermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourcePermissions != nil {
		toSerialize["resourcePermissions"] = o.ResourcePermissions
	}
	return json.Marshal(toSerialize)
}

type NullableClusterNamespacePermissions struct {
	value *ClusterNamespacePermissions
	isSet bool
}

func (v NullableClusterNamespacePermissions) Get() *ClusterNamespacePermissions {
	return v.value
}

func (v *NullableClusterNamespacePermissions) Set(val *ClusterNamespacePermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterNamespacePermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterNamespacePermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterNamespacePermissions(val *ClusterNamespacePermissions) *NullableClusterNamespacePermissions {
	return &NullableClusterNamespacePermissions{value: val, isSet: true}
}

func (v NullableClusterNamespacePermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterNamespacePermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


