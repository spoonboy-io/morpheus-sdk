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

// ClusterUpdatePermissionsResourcePermissions struct for ClusterUpdatePermissionsResourcePermissions
type ClusterUpdatePermissionsResourcePermissions struct {
	// Pass true to allow access to all groups
	All *bool `json:"all,omitempty"`
	// Array of groups that are allowed access
	Sites *[]ClusterUpdatePermissionsResourcePermissionsSites `json:"sites,omitempty"`
	// Pass true to allow access to all plans
	AllPlans *bool `json:"allPlans,omitempty"`
	// Array of plans that are allowed access
	Plans *[]ClusterUpdatePermissionsResourcePermissionsSites `json:"plans,omitempty"`
}

// NewClusterUpdatePermissionsResourcePermissions instantiates a new ClusterUpdatePermissionsResourcePermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterUpdatePermissionsResourcePermissions() *ClusterUpdatePermissionsResourcePermissions {
	this := ClusterUpdatePermissionsResourcePermissions{}
	return &this
}

// NewClusterUpdatePermissionsResourcePermissionsWithDefaults instantiates a new ClusterUpdatePermissionsResourcePermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterUpdatePermissionsResourcePermissionsWithDefaults() *ClusterUpdatePermissionsResourcePermissions {
	this := ClusterUpdatePermissionsResourcePermissions{}
	return &this
}

// GetAll returns the All field value if set, zero value otherwise.
func (o *ClusterUpdatePermissionsResourcePermissions) GetAll() bool {
	if o == nil || o.All == nil {
		var ret bool
		return ret
	}
	return *o.All
}

// GetAllOk returns a tuple with the All field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdatePermissionsResourcePermissions) GetAllOk() (*bool, bool) {
	if o == nil || o.All == nil {
		return nil, false
	}
	return o.All, true
}

// HasAll returns a boolean if a field has been set.
func (o *ClusterUpdatePermissionsResourcePermissions) HasAll() bool {
	if o != nil && o.All != nil {
		return true
	}

	return false
}

// SetAll gets a reference to the given bool and assigns it to the All field.
func (o *ClusterUpdatePermissionsResourcePermissions) SetAll(v bool) {
	o.All = &v
}

// GetSites returns the Sites field value if set, zero value otherwise.
func (o *ClusterUpdatePermissionsResourcePermissions) GetSites() []ClusterUpdatePermissionsResourcePermissionsSites {
	if o == nil || o.Sites == nil {
		var ret []ClusterUpdatePermissionsResourcePermissionsSites
		return ret
	}
	return *o.Sites
}

// GetSitesOk returns a tuple with the Sites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdatePermissionsResourcePermissions) GetSitesOk() (*[]ClusterUpdatePermissionsResourcePermissionsSites, bool) {
	if o == nil || o.Sites == nil {
		return nil, false
	}
	return o.Sites, true
}

// HasSites returns a boolean if a field has been set.
func (o *ClusterUpdatePermissionsResourcePermissions) HasSites() bool {
	if o != nil && o.Sites != nil {
		return true
	}

	return false
}

// SetSites gets a reference to the given []ClusterUpdatePermissionsResourcePermissionsSites and assigns it to the Sites field.
func (o *ClusterUpdatePermissionsResourcePermissions) SetSites(v []ClusterUpdatePermissionsResourcePermissionsSites) {
	o.Sites = &v
}

// GetAllPlans returns the AllPlans field value if set, zero value otherwise.
func (o *ClusterUpdatePermissionsResourcePermissions) GetAllPlans() bool {
	if o == nil || o.AllPlans == nil {
		var ret bool
		return ret
	}
	return *o.AllPlans
}

// GetAllPlansOk returns a tuple with the AllPlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdatePermissionsResourcePermissions) GetAllPlansOk() (*bool, bool) {
	if o == nil || o.AllPlans == nil {
		return nil, false
	}
	return o.AllPlans, true
}

// HasAllPlans returns a boolean if a field has been set.
func (o *ClusterUpdatePermissionsResourcePermissions) HasAllPlans() bool {
	if o != nil && o.AllPlans != nil {
		return true
	}

	return false
}

// SetAllPlans gets a reference to the given bool and assigns it to the AllPlans field.
func (o *ClusterUpdatePermissionsResourcePermissions) SetAllPlans(v bool) {
	o.AllPlans = &v
}

// GetPlans returns the Plans field value if set, zero value otherwise.
func (o *ClusterUpdatePermissionsResourcePermissions) GetPlans() []ClusterUpdatePermissionsResourcePermissionsSites {
	if o == nil || o.Plans == nil {
		var ret []ClusterUpdatePermissionsResourcePermissionsSites
		return ret
	}
	return *o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdatePermissionsResourcePermissions) GetPlansOk() (*[]ClusterUpdatePermissionsResourcePermissionsSites, bool) {
	if o == nil || o.Plans == nil {
		return nil, false
	}
	return o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *ClusterUpdatePermissionsResourcePermissions) HasPlans() bool {
	if o != nil && o.Plans != nil {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []ClusterUpdatePermissionsResourcePermissionsSites and assigns it to the Plans field.
func (o *ClusterUpdatePermissionsResourcePermissions) SetPlans(v []ClusterUpdatePermissionsResourcePermissionsSites) {
	o.Plans = &v
}

func (o ClusterUpdatePermissionsResourcePermissions) MarshalJSON() ([]byte, error) {
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

type NullableClusterUpdatePermissionsResourcePermissions struct {
	value *ClusterUpdatePermissionsResourcePermissions
	isSet bool
}

func (v NullableClusterUpdatePermissionsResourcePermissions) Get() *ClusterUpdatePermissionsResourcePermissions {
	return v.value
}

func (v *NullableClusterUpdatePermissionsResourcePermissions) Set(val *ClusterUpdatePermissionsResourcePermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterUpdatePermissionsResourcePermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterUpdatePermissionsResourcePermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterUpdatePermissionsResourcePermissions(val *ClusterUpdatePermissionsResourcePermissions) *NullableClusterUpdatePermissionsResourcePermissions {
	return &NullableClusterUpdatePermissionsResourcePermissions{value: val, isSet: true}
}

func (v NullableClusterUpdatePermissionsResourcePermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterUpdatePermissionsResourcePermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


