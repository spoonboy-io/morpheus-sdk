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

// ClusterServerCreateUserGroup struct for ClusterServerCreateUserGroup
type ClusterServerCreateUserGroup struct {
	// User Group ID for server host
	Id *int64 `json:"id,omitempty"`
}

// NewClusterServerCreateUserGroup instantiates a new ClusterServerCreateUserGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterServerCreateUserGroup() *ClusterServerCreateUserGroup {
	this := ClusterServerCreateUserGroup{}
	return &this
}

// NewClusterServerCreateUserGroupWithDefaults instantiates a new ClusterServerCreateUserGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterServerCreateUserGroupWithDefaults() *ClusterServerCreateUserGroup {
	this := ClusterServerCreateUserGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClusterServerCreateUserGroup) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterServerCreateUserGroup) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClusterServerCreateUserGroup) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ClusterServerCreateUserGroup) SetId(v int64) {
	o.Id = &v
}

func (o ClusterServerCreateUserGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableClusterServerCreateUserGroup struct {
	value *ClusterServerCreateUserGroup
	isSet bool
}

func (v NullableClusterServerCreateUserGroup) Get() *ClusterServerCreateUserGroup {
	return v.value
}

func (v *NullableClusterServerCreateUserGroup) Set(val *ClusterServerCreateUserGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterServerCreateUserGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterServerCreateUserGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterServerCreateUserGroup(val *ClusterServerCreateUserGroup) *NullableClusterServerCreateUserGroup {
	return &NullableClusterServerCreateUserGroup{value: val, isSet: true}
}

func (v NullableClusterServerCreateUserGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterServerCreateUserGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


