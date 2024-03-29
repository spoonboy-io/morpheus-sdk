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

// ClusterPermissionsResourcePool struct for ClusterPermissionsResourcePool
type ClusterPermissionsResourcePool struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
}

// NewClusterPermissionsResourcePool instantiates a new ClusterPermissionsResourcePool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterPermissionsResourcePool() *ClusterPermissionsResourcePool {
	this := ClusterPermissionsResourcePool{}
	return &this
}

// NewClusterPermissionsResourcePoolWithDefaults instantiates a new ClusterPermissionsResourcePool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterPermissionsResourcePoolWithDefaults() *ClusterPermissionsResourcePool {
	this := ClusterPermissionsResourcePool{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClusterPermissionsResourcePool) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterPermissionsResourcePool) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClusterPermissionsResourcePool) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ClusterPermissionsResourcePool) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClusterPermissionsResourcePool) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterPermissionsResourcePool) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClusterPermissionsResourcePool) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClusterPermissionsResourcePool) SetName(v string) {
	o.Name = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *ClusterPermissionsResourcePool) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterPermissionsResourcePool) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *ClusterPermissionsResourcePool) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *ClusterPermissionsResourcePool) SetVisibility(v string) {
	o.Visibility = &v
}

func (o ClusterPermissionsResourcePool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	return json.Marshal(toSerialize)
}

type NullableClusterPermissionsResourcePool struct {
	value *ClusterPermissionsResourcePool
	isSet bool
}

func (v NullableClusterPermissionsResourcePool) Get() *ClusterPermissionsResourcePool {
	return v.value
}

func (v *NullableClusterPermissionsResourcePool) Set(val *ClusterPermissionsResourcePool) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterPermissionsResourcePool) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterPermissionsResourcePool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterPermissionsResourcePool(val *ClusterPermissionsResourcePool) *NullableClusterPermissionsResourcePool {
	return &NullableClusterPermissionsResourcePool{value: val, isSet: true}
}

func (v NullableClusterPermissionsResourcePool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterPermissionsResourcePool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


