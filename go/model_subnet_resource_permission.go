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

// SubnetResourcePermission struct for SubnetResourcePermission
type SubnetResourcePermission struct {
	All *bool `json:"all,omitempty"`
	Sites *[]map[string]interface{} `json:"sites,omitempty"`
	AllPlans *bool `json:"allPlans,omitempty"`
	Plans *[]map[string]interface{} `json:"plans,omitempty"`
}

// NewSubnetResourcePermission instantiates a new SubnetResourcePermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubnetResourcePermission() *SubnetResourcePermission {
	this := SubnetResourcePermission{}
	return &this
}

// NewSubnetResourcePermissionWithDefaults instantiates a new SubnetResourcePermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubnetResourcePermissionWithDefaults() *SubnetResourcePermission {
	this := SubnetResourcePermission{}
	return &this
}

// GetAll returns the All field value if set, zero value otherwise.
func (o *SubnetResourcePermission) GetAll() bool {
	if o == nil || o.All == nil {
		var ret bool
		return ret
	}
	return *o.All
}

// GetAllOk returns a tuple with the All field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubnetResourcePermission) GetAllOk() (*bool, bool) {
	if o == nil || o.All == nil {
		return nil, false
	}
	return o.All, true
}

// HasAll returns a boolean if a field has been set.
func (o *SubnetResourcePermission) HasAll() bool {
	if o != nil && o.All != nil {
		return true
	}

	return false
}

// SetAll gets a reference to the given bool and assigns it to the All field.
func (o *SubnetResourcePermission) SetAll(v bool) {
	o.All = &v
}

// GetSites returns the Sites field value if set, zero value otherwise.
func (o *SubnetResourcePermission) GetSites() []map[string]interface{} {
	if o == nil || o.Sites == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Sites
}

// GetSitesOk returns a tuple with the Sites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubnetResourcePermission) GetSitesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Sites == nil {
		return nil, false
	}
	return o.Sites, true
}

// HasSites returns a boolean if a field has been set.
func (o *SubnetResourcePermission) HasSites() bool {
	if o != nil && o.Sites != nil {
		return true
	}

	return false
}

// SetSites gets a reference to the given []map[string]interface{} and assigns it to the Sites field.
func (o *SubnetResourcePermission) SetSites(v []map[string]interface{}) {
	o.Sites = &v
}

// GetAllPlans returns the AllPlans field value if set, zero value otherwise.
func (o *SubnetResourcePermission) GetAllPlans() bool {
	if o == nil || o.AllPlans == nil {
		var ret bool
		return ret
	}
	return *o.AllPlans
}

// GetAllPlansOk returns a tuple with the AllPlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubnetResourcePermission) GetAllPlansOk() (*bool, bool) {
	if o == nil || o.AllPlans == nil {
		return nil, false
	}
	return o.AllPlans, true
}

// HasAllPlans returns a boolean if a field has been set.
func (o *SubnetResourcePermission) HasAllPlans() bool {
	if o != nil && o.AllPlans != nil {
		return true
	}

	return false
}

// SetAllPlans gets a reference to the given bool and assigns it to the AllPlans field.
func (o *SubnetResourcePermission) SetAllPlans(v bool) {
	o.AllPlans = &v
}

// GetPlans returns the Plans field value if set, zero value otherwise.
func (o *SubnetResourcePermission) GetPlans() []map[string]interface{} {
	if o == nil || o.Plans == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubnetResourcePermission) GetPlansOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Plans == nil {
		return nil, false
	}
	return o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *SubnetResourcePermission) HasPlans() bool {
	if o != nil && o.Plans != nil {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []map[string]interface{} and assigns it to the Plans field.
func (o *SubnetResourcePermission) SetPlans(v []map[string]interface{}) {
	o.Plans = &v
}

func (o SubnetResourcePermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.All != nil {
		toSerialize["all"] = o.All
	}
	if o.Sites != nil {
		toSerialize["sites"] = o.Sites
	}
	if o.AllPlans != nil {
		toSerialize["allPlans"] = o.AllPlans
	}
	if o.Plans != nil {
		toSerialize["plans"] = o.Plans
	}
	return json.Marshal(toSerialize)
}

type NullableSubnetResourcePermission struct {
	value *SubnetResourcePermission
	isSet bool
}

func (v NullableSubnetResourcePermission) Get() *SubnetResourcePermission {
	return v.value
}

func (v *NullableSubnetResourcePermission) Set(val *SubnetResourcePermission) {
	v.value = val
	v.isSet = true
}

func (v NullableSubnetResourcePermission) IsSet() bool {
	return v.isSet
}

func (v *NullableSubnetResourcePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubnetResourcePermission(val *SubnetResourcePermission) *NullableSubnetResourcePermission {
	return &NullableSubnetResourcePermission{value: val, isSet: true}
}

func (v NullableSubnetResourcePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubnetResourcePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


