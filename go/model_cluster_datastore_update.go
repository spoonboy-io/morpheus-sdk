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

// ClusterDatastoreUpdate struct for ClusterDatastoreUpdate
type ClusterDatastoreUpdate struct {
	// Datastore active
	Active *bool `json:"active,omitempty"`
	Permissions *ClusterUpdatePermissions `json:"permissions,omitempty"`
	// Visibility for datastore
	Visibility *string `json:"visibility,omitempty"`
}

// NewClusterDatastoreUpdate instantiates a new ClusterDatastoreUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterDatastoreUpdate() *ClusterDatastoreUpdate {
	this := ClusterDatastoreUpdate{}
	var active bool = true
	this.Active = &active
	var visibility string = "private"
	this.Visibility = &visibility
	return &this
}

// NewClusterDatastoreUpdateWithDefaults instantiates a new ClusterDatastoreUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterDatastoreUpdateWithDefaults() *ClusterDatastoreUpdate {
	this := ClusterDatastoreUpdate{}
	var active bool = true
	this.Active = &active
	var visibility string = "private"
	this.Visibility = &visibility
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ClusterDatastoreUpdate) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDatastoreUpdate) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ClusterDatastoreUpdate) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ClusterDatastoreUpdate) SetActive(v bool) {
	o.Active = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *ClusterDatastoreUpdate) GetPermissions() ClusterUpdatePermissions {
	if o == nil || o.Permissions == nil {
		var ret ClusterUpdatePermissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDatastoreUpdate) GetPermissionsOk() (*ClusterUpdatePermissions, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *ClusterDatastoreUpdate) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given ClusterUpdatePermissions and assigns it to the Permissions field.
func (o *ClusterDatastoreUpdate) SetPermissions(v ClusterUpdatePermissions) {
	o.Permissions = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *ClusterDatastoreUpdate) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDatastoreUpdate) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *ClusterDatastoreUpdate) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *ClusterDatastoreUpdate) SetVisibility(v string) {
	o.Visibility = &v
}

func (o ClusterDatastoreUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	return json.Marshal(toSerialize)
}

type NullableClusterDatastoreUpdate struct {
	value *ClusterDatastoreUpdate
	isSet bool
}

func (v NullableClusterDatastoreUpdate) Get() *ClusterDatastoreUpdate {
	return v.value
}

func (v *NullableClusterDatastoreUpdate) Set(val *ClusterDatastoreUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterDatastoreUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterDatastoreUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterDatastoreUpdate(val *ClusterDatastoreUpdate) *NullableClusterDatastoreUpdate {
	return &NullableClusterDatastoreUpdate{value: val, isSet: true}
}

func (v NullableClusterDatastoreUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterDatastoreUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


