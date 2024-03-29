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

// InlineResponse200148 struct for InlineResponse200148
type InlineResponse200148 struct {
	StorageVolumeType *StorageVolumeType `json:"storageVolumeType,omitempty"`
}

// NewInlineResponse200148 instantiates a new InlineResponse200148 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200148() *InlineResponse200148 {
	this := InlineResponse200148{}
	return &this
}

// NewInlineResponse200148WithDefaults instantiates a new InlineResponse200148 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200148WithDefaults() *InlineResponse200148 {
	this := InlineResponse200148{}
	return &this
}

// GetStorageVolumeType returns the StorageVolumeType field value if set, zero value otherwise.
func (o *InlineResponse200148) GetStorageVolumeType() StorageVolumeType {
	if o == nil || o.StorageVolumeType == nil {
		var ret StorageVolumeType
		return ret
	}
	return *o.StorageVolumeType
}

// GetStorageVolumeTypeOk returns a tuple with the StorageVolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200148) GetStorageVolumeTypeOk() (*StorageVolumeType, bool) {
	if o == nil || o.StorageVolumeType == nil {
		return nil, false
	}
	return o.StorageVolumeType, true
}

// HasStorageVolumeType returns a boolean if a field has been set.
func (o *InlineResponse200148) HasStorageVolumeType() bool {
	if o != nil && o.StorageVolumeType != nil {
		return true
	}

	return false
}

// SetStorageVolumeType gets a reference to the given StorageVolumeType and assigns it to the StorageVolumeType field.
func (o *InlineResponse200148) SetStorageVolumeType(v StorageVolumeType) {
	o.StorageVolumeType = &v
}

func (o InlineResponse200148) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StorageVolumeType != nil {
		toSerialize["storageVolumeType"] = o.StorageVolumeType
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200148 struct {
	value *InlineResponse200148
	isSet bool
}

func (v NullableInlineResponse200148) Get() *InlineResponse200148 {
	return v.value
}

func (v *NullableInlineResponse200148) Set(val *InlineResponse200148) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200148) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200148) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200148(val *InlineResponse200148) *NullableInlineResponse200148 {
	return &NullableInlineResponse200148{value: val, isSet: true}
}

func (v NullableInlineResponse200148) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200148) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


