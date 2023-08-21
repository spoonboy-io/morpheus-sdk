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

// DashboardProvisioning struct for DashboardProvisioning
type DashboardProvisioning struct {
	InstanceCount *int64 `json:"instanceCount,omitempty"`
	FavoriteInstances *[]map[string]interface{} `json:"favoriteInstances,omitempty"`
}

// NewDashboardProvisioning instantiates a new DashboardProvisioning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardProvisioning() *DashboardProvisioning {
	this := DashboardProvisioning{}
	return &this
}

// NewDashboardProvisioningWithDefaults instantiates a new DashboardProvisioning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardProvisioningWithDefaults() *DashboardProvisioning {
	this := DashboardProvisioning{}
	return &this
}

// GetInstanceCount returns the InstanceCount field value if set, zero value otherwise.
func (o *DashboardProvisioning) GetInstanceCount() int64 {
	if o == nil || o.InstanceCount == nil {
		var ret int64
		return ret
	}
	return *o.InstanceCount
}

// GetInstanceCountOk returns a tuple with the InstanceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardProvisioning) GetInstanceCountOk() (*int64, bool) {
	if o == nil || o.InstanceCount == nil {
		return nil, false
	}
	return o.InstanceCount, true
}

// HasInstanceCount returns a boolean if a field has been set.
func (o *DashboardProvisioning) HasInstanceCount() bool {
	if o != nil && o.InstanceCount != nil {
		return true
	}

	return false
}

// SetInstanceCount gets a reference to the given int64 and assigns it to the InstanceCount field.
func (o *DashboardProvisioning) SetInstanceCount(v int64) {
	o.InstanceCount = &v
}

// GetFavoriteInstances returns the FavoriteInstances field value if set, zero value otherwise.
func (o *DashboardProvisioning) GetFavoriteInstances() []map[string]interface{} {
	if o == nil || o.FavoriteInstances == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.FavoriteInstances
}

// GetFavoriteInstancesOk returns a tuple with the FavoriteInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardProvisioning) GetFavoriteInstancesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.FavoriteInstances == nil {
		return nil, false
	}
	return o.FavoriteInstances, true
}

// HasFavoriteInstances returns a boolean if a field has been set.
func (o *DashboardProvisioning) HasFavoriteInstances() bool {
	if o != nil && o.FavoriteInstances != nil {
		return true
	}

	return false
}

// SetFavoriteInstances gets a reference to the given []map[string]interface{} and assigns it to the FavoriteInstances field.
func (o *DashboardProvisioning) SetFavoriteInstances(v []map[string]interface{}) {
	o.FavoriteInstances = &v
}

func (o DashboardProvisioning) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InstanceCount != nil {
		toSerialize["instanceCount"] = o.InstanceCount
	}
	if o.FavoriteInstances != nil {
		toSerialize["favoriteInstances"] = o.FavoriteInstances
	}
	return json.Marshal(toSerialize)
}

type NullableDashboardProvisioning struct {
	value *DashboardProvisioning
	isSet bool
}

func (v NullableDashboardProvisioning) Get() *DashboardProvisioning {
	return v.value
}

func (v *NullableDashboardProvisioning) Set(val *DashboardProvisioning) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardProvisioning) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardProvisioning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardProvisioning(val *DashboardProvisioning) *NullableDashboardProvisioning {
	return &NullableDashboardProvisioning{value: val, isSet: true}
}

func (v NullableDashboardProvisioning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardProvisioning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


