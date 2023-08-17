/*
Morpheus API

Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.

API version: 6.1.1
Contact: dev@morpheusdata.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UpdateBlueprintPermissionsRequestResourcePermission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateBlueprintPermissionsRequestResourcePermission{}

// UpdateBlueprintPermissionsRequestResourcePermission struct for UpdateBlueprintPermissionsRequestResourcePermission
type UpdateBlueprintPermissionsRequestResourcePermission struct {
	// Set to true to grant access to all groups
	All *bool `json:"all,omitempty"`
	// Array of objects identifying groups with access
	Sites []UpdateBlueprintPermissionsRequestResourcePermissionSitesInner `json:"sites,omitempty"`
	// User ID, can be used to change blueprint owner.
	OwnerId *int64 `json:"ownerId,omitempty"`
}

// NewUpdateBlueprintPermissionsRequestResourcePermission instantiates a new UpdateBlueprintPermissionsRequestResourcePermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateBlueprintPermissionsRequestResourcePermission() *UpdateBlueprintPermissionsRequestResourcePermission {
	this := UpdateBlueprintPermissionsRequestResourcePermission{}
	return &this
}

// NewUpdateBlueprintPermissionsRequestResourcePermissionWithDefaults instantiates a new UpdateBlueprintPermissionsRequestResourcePermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateBlueprintPermissionsRequestResourcePermissionWithDefaults() *UpdateBlueprintPermissionsRequestResourcePermission {
	this := UpdateBlueprintPermissionsRequestResourcePermission{}
	return &this
}

// GetAll returns the All field value if set, zero value otherwise.
func (o *UpdateBlueprintPermissionsRequestResourcePermission) GetAll() bool {
	if o == nil || IsNil(o.All) {
		var ret bool
		return ret
	}
	return *o.All
}

// GetAllOk returns a tuple with the All field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBlueprintPermissionsRequestResourcePermission) GetAllOk() (*bool, bool) {
	if o == nil || IsNil(o.All) {
		return nil, false
	}
	return o.All, true
}

// HasAll returns a boolean if a field has been set.
func (o *UpdateBlueprintPermissionsRequestResourcePermission) HasAll() bool {
	if o != nil && !IsNil(o.All) {
		return true
	}

	return false
}

// SetAll gets a reference to the given bool and assigns it to the All field.
func (o *UpdateBlueprintPermissionsRequestResourcePermission) SetAll(v bool) {
	o.All = &v
}

// GetSites returns the Sites field value if set, zero value otherwise.
func (o *UpdateBlueprintPermissionsRequestResourcePermission) GetSites() []UpdateBlueprintPermissionsRequestResourcePermissionSitesInner {
	if o == nil || IsNil(o.Sites) {
		var ret []UpdateBlueprintPermissionsRequestResourcePermissionSitesInner
		return ret
	}
	return o.Sites
}

// GetSitesOk returns a tuple with the Sites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBlueprintPermissionsRequestResourcePermission) GetSitesOk() ([]UpdateBlueprintPermissionsRequestResourcePermissionSitesInner, bool) {
	if o == nil || IsNil(o.Sites) {
		return nil, false
	}
	return o.Sites, true
}

// HasSites returns a boolean if a field has been set.
func (o *UpdateBlueprintPermissionsRequestResourcePermission) HasSites() bool {
	if o != nil && !IsNil(o.Sites) {
		return true
	}

	return false
}

// SetSites gets a reference to the given []UpdateBlueprintPermissionsRequestResourcePermissionSitesInner and assigns it to the Sites field.
func (o *UpdateBlueprintPermissionsRequestResourcePermission) SetSites(v []UpdateBlueprintPermissionsRequestResourcePermissionSitesInner) {
	o.Sites = v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *UpdateBlueprintPermissionsRequestResourcePermission) GetOwnerId() int64 {
	if o == nil || IsNil(o.OwnerId) {
		var ret int64
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBlueprintPermissionsRequestResourcePermission) GetOwnerIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *UpdateBlueprintPermissionsRequestResourcePermission) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given int64 and assigns it to the OwnerId field.
func (o *UpdateBlueprintPermissionsRequestResourcePermission) SetOwnerId(v int64) {
	o.OwnerId = &v
}

func (o UpdateBlueprintPermissionsRequestResourcePermission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateBlueprintPermissionsRequestResourcePermission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.All) {
		toSerialize["all"] = o.All
	}
	if !IsNil(o.Sites) {
		toSerialize["sites"] = o.Sites
	}
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	return toSerialize, nil
}

type NullableUpdateBlueprintPermissionsRequestResourcePermission struct {
	value *UpdateBlueprintPermissionsRequestResourcePermission
	isSet bool
}

func (v NullableUpdateBlueprintPermissionsRequestResourcePermission) Get() *UpdateBlueprintPermissionsRequestResourcePermission {
	return v.value
}

func (v *NullableUpdateBlueprintPermissionsRequestResourcePermission) Set(val *UpdateBlueprintPermissionsRequestResourcePermission) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateBlueprintPermissionsRequestResourcePermission) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateBlueprintPermissionsRequestResourcePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateBlueprintPermissionsRequestResourcePermission(val *UpdateBlueprintPermissionsRequestResourcePermission) *NullableUpdateBlueprintPermissionsRequestResourcePermission {
	return &NullableUpdateBlueprintPermissionsRequestResourcePermission{value: val, isSet: true}
}

func (v NullableUpdateBlueprintPermissionsRequestResourcePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateBlueprintPermissionsRequestResourcePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

