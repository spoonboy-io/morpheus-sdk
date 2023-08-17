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

// checks if the UpdateBlueprintPermissionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateBlueprintPermissionsRequest{}

// UpdateBlueprintPermissionsRequest struct for UpdateBlueprintPermissionsRequest
type UpdateBlueprintPermissionsRequest struct {
	ResourcePermission *UpdateBlueprintPermissionsRequestResourcePermission `json:"resourcePermission,omitempty"`
}

// NewUpdateBlueprintPermissionsRequest instantiates a new UpdateBlueprintPermissionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateBlueprintPermissionsRequest() *UpdateBlueprintPermissionsRequest {
	this := UpdateBlueprintPermissionsRequest{}
	return &this
}

// NewUpdateBlueprintPermissionsRequestWithDefaults instantiates a new UpdateBlueprintPermissionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateBlueprintPermissionsRequestWithDefaults() *UpdateBlueprintPermissionsRequest {
	this := UpdateBlueprintPermissionsRequest{}
	return &this
}

// GetResourcePermission returns the ResourcePermission field value if set, zero value otherwise.
func (o *UpdateBlueprintPermissionsRequest) GetResourcePermission() UpdateBlueprintPermissionsRequestResourcePermission {
	if o == nil || IsNil(o.ResourcePermission) {
		var ret UpdateBlueprintPermissionsRequestResourcePermission
		return ret
	}
	return *o.ResourcePermission
}

// GetResourcePermissionOk returns a tuple with the ResourcePermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBlueprintPermissionsRequest) GetResourcePermissionOk() (*UpdateBlueprintPermissionsRequestResourcePermission, bool) {
	if o == nil || IsNil(o.ResourcePermission) {
		return nil, false
	}
	return o.ResourcePermission, true
}

// HasResourcePermission returns a boolean if a field has been set.
func (o *UpdateBlueprintPermissionsRequest) HasResourcePermission() bool {
	if o != nil && !IsNil(o.ResourcePermission) {
		return true
	}

	return false
}

// SetResourcePermission gets a reference to the given UpdateBlueprintPermissionsRequestResourcePermission and assigns it to the ResourcePermission field.
func (o *UpdateBlueprintPermissionsRequest) SetResourcePermission(v UpdateBlueprintPermissionsRequestResourcePermission) {
	o.ResourcePermission = &v
}

func (o UpdateBlueprintPermissionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateBlueprintPermissionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourcePermission) {
		toSerialize["resourcePermission"] = o.ResourcePermission
	}
	return toSerialize, nil
}

type NullableUpdateBlueprintPermissionsRequest struct {
	value *UpdateBlueprintPermissionsRequest
	isSet bool
}

func (v NullableUpdateBlueprintPermissionsRequest) Get() *UpdateBlueprintPermissionsRequest {
	return v.value
}

func (v *NullableUpdateBlueprintPermissionsRequest) Set(val *UpdateBlueprintPermissionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateBlueprintPermissionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateBlueprintPermissionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateBlueprintPermissionsRequest(val *UpdateBlueprintPermissionsRequest) *NullableUpdateBlueprintPermissionsRequest {
	return &NullableUpdateBlueprintPermissionsRequest{value: val, isSet: true}
}

func (v NullableUpdateBlueprintPermissionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateBlueprintPermissionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


