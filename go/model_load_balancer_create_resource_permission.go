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

// LoadBalancerCreateResourcePermission struct for LoadBalancerCreateResourcePermission
type LoadBalancerCreateResourcePermission struct {
	// Pass true to allow access to all groups
	All *bool `json:"all,omitempty"`
	// Array of groups that are allowed access
	Sites *[]int64 `json:"sites,omitempty"`
}

// NewLoadBalancerCreateResourcePermission instantiates a new LoadBalancerCreateResourcePermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerCreateResourcePermission() *LoadBalancerCreateResourcePermission {
	this := LoadBalancerCreateResourcePermission{}
	return &this
}

// NewLoadBalancerCreateResourcePermissionWithDefaults instantiates a new LoadBalancerCreateResourcePermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerCreateResourcePermissionWithDefaults() *LoadBalancerCreateResourcePermission {
	this := LoadBalancerCreateResourcePermission{}
	return &this
}

// GetAll returns the All field value if set, zero value otherwise.
func (o *LoadBalancerCreateResourcePermission) GetAll() bool {
	if o == nil || o.All == nil {
		var ret bool
		return ret
	}
	return *o.All
}

// GetAllOk returns a tuple with the All field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerCreateResourcePermission) GetAllOk() (*bool, bool) {
	if o == nil || o.All == nil {
		return nil, false
	}
	return o.All, true
}

// HasAll returns a boolean if a field has been set.
func (o *LoadBalancerCreateResourcePermission) HasAll() bool {
	if o != nil && o.All != nil {
		return true
	}

	return false
}

// SetAll gets a reference to the given bool and assigns it to the All field.
func (o *LoadBalancerCreateResourcePermission) SetAll(v bool) {
	o.All = &v
}

// GetSites returns the Sites field value if set, zero value otherwise.
func (o *LoadBalancerCreateResourcePermission) GetSites() []int64 {
	if o == nil || o.Sites == nil {
		var ret []int64
		return ret
	}
	return *o.Sites
}

// GetSitesOk returns a tuple with the Sites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerCreateResourcePermission) GetSitesOk() (*[]int64, bool) {
	if o == nil || o.Sites == nil {
		return nil, false
	}
	return o.Sites, true
}

// HasSites returns a boolean if a field has been set.
func (o *LoadBalancerCreateResourcePermission) HasSites() bool {
	if o != nil && o.Sites != nil {
		return true
	}

	return false
}

// SetSites gets a reference to the given []int64 and assigns it to the Sites field.
func (o *LoadBalancerCreateResourcePermission) SetSites(v []int64) {
	o.Sites = &v
}

func (o LoadBalancerCreateResourcePermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.All != nil {
		toSerialize["all"] = o.All
	}
	if o.Sites != nil {
		toSerialize["sites"] = o.Sites
	}
	return json.Marshal(toSerialize)
}

type NullableLoadBalancerCreateResourcePermission struct {
	value *LoadBalancerCreateResourcePermission
	isSet bool
}

func (v NullableLoadBalancerCreateResourcePermission) Get() *LoadBalancerCreateResourcePermission {
	return v.value
}

func (v *NullableLoadBalancerCreateResourcePermission) Set(val *LoadBalancerCreateResourcePermission) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerCreateResourcePermission) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerCreateResourcePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerCreateResourcePermission(val *LoadBalancerCreateResourcePermission) *NullableLoadBalancerCreateResourcePermission {
	return &NullableLoadBalancerCreateResourcePermission{value: val, isSet: true}
}

func (v NullableLoadBalancerCreateResourcePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerCreateResourcePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


