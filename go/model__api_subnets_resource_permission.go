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

// ApiSubnetsResourcePermission struct for ApiSubnetsResourcePermission
type ApiSubnetsResourcePermission struct {
	// Pass true to allow access all groups
	All *bool `json:"all,omitempty"`
	// Array of groups ID objects that are allowed access
	Sites *[]ApiBlueprintsIdUpdatePermissionsResourcePermissionSites `json:"sites,omitempty"`
	AllPlans *bool `json:"allPlans,omitempty"`
	Plans *[]map[string]interface{} `json:"plans,omitempty"`
}

// NewApiSubnetsResourcePermission instantiates a new ApiSubnetsResourcePermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSubnetsResourcePermission() *ApiSubnetsResourcePermission {
	this := ApiSubnetsResourcePermission{}
	return &this
}

// NewApiSubnetsResourcePermissionWithDefaults instantiates a new ApiSubnetsResourcePermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSubnetsResourcePermissionWithDefaults() *ApiSubnetsResourcePermission {
	this := ApiSubnetsResourcePermission{}
	return &this
}

// GetAll returns the All field value if set, zero value otherwise.
func (o *ApiSubnetsResourcePermission) GetAll() bool {
	if o == nil || o.All == nil {
		var ret bool
		return ret
	}
	return *o.All
}

// GetAllOk returns a tuple with the All field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubnetsResourcePermission) GetAllOk() (*bool, bool) {
	if o == nil || o.All == nil {
		return nil, false
	}
	return o.All, true
}

// HasAll returns a boolean if a field has been set.
func (o *ApiSubnetsResourcePermission) HasAll() bool {
	if o != nil && o.All != nil {
		return true
	}

	return false
}

// SetAll gets a reference to the given bool and assigns it to the All field.
func (o *ApiSubnetsResourcePermission) SetAll(v bool) {
	o.All = &v
}

// GetSites returns the Sites field value if set, zero value otherwise.
func (o *ApiSubnetsResourcePermission) GetSites() []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites {
	if o == nil || o.Sites == nil {
		var ret []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites
		return ret
	}
	return *o.Sites
}

// GetSitesOk returns a tuple with the Sites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubnetsResourcePermission) GetSitesOk() (*[]ApiBlueprintsIdUpdatePermissionsResourcePermissionSites, bool) {
	if o == nil || o.Sites == nil {
		return nil, false
	}
	return o.Sites, true
}

// HasSites returns a boolean if a field has been set.
func (o *ApiSubnetsResourcePermission) HasSites() bool {
	if o != nil && o.Sites != nil {
		return true
	}

	return false
}

// SetSites gets a reference to the given []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites and assigns it to the Sites field.
func (o *ApiSubnetsResourcePermission) SetSites(v []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites) {
	o.Sites = &v
}

// GetAllPlans returns the AllPlans field value if set, zero value otherwise.
func (o *ApiSubnetsResourcePermission) GetAllPlans() bool {
	if o == nil || o.AllPlans == nil {
		var ret bool
		return ret
	}
	return *o.AllPlans
}

// GetAllPlansOk returns a tuple with the AllPlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubnetsResourcePermission) GetAllPlansOk() (*bool, bool) {
	if o == nil || o.AllPlans == nil {
		return nil, false
	}
	return o.AllPlans, true
}

// HasAllPlans returns a boolean if a field has been set.
func (o *ApiSubnetsResourcePermission) HasAllPlans() bool {
	if o != nil && o.AllPlans != nil {
		return true
	}

	return false
}

// SetAllPlans gets a reference to the given bool and assigns it to the AllPlans field.
func (o *ApiSubnetsResourcePermission) SetAllPlans(v bool) {
	o.AllPlans = &v
}

// GetPlans returns the Plans field value if set, zero value otherwise.
func (o *ApiSubnetsResourcePermission) GetPlans() []map[string]interface{} {
	if o == nil || o.Plans == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubnetsResourcePermission) GetPlansOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Plans == nil {
		return nil, false
	}
	return o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *ApiSubnetsResourcePermission) HasPlans() bool {
	if o != nil && o.Plans != nil {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []map[string]interface{} and assigns it to the Plans field.
func (o *ApiSubnetsResourcePermission) SetPlans(v []map[string]interface{}) {
	o.Plans = &v
}

func (o ApiSubnetsResourcePermission) MarshalJSON() ([]byte, error) {
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

type NullableApiSubnetsResourcePermission struct {
	value *ApiSubnetsResourcePermission
	isSet bool
}

func (v NullableApiSubnetsResourcePermission) Get() *ApiSubnetsResourcePermission {
	return v.value
}

func (v *NullableApiSubnetsResourcePermission) Set(val *ApiSubnetsResourcePermission) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSubnetsResourcePermission) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSubnetsResourcePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSubnetsResourcePermission(val *ApiSubnetsResourcePermission) *NullableApiSubnetsResourcePermission {
	return &NullableApiSubnetsResourcePermission{value: val, isSet: true}
}

func (v NullableApiSubnetsResourcePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSubnetsResourcePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

