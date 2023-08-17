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

// checks if the UpdateCatalogItemTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCatalogItemTypeRequest{}

// UpdateCatalogItemTypeRequest struct for UpdateCatalogItemTypeRequest
type UpdateCatalogItemTypeRequest struct {
	CatalogItemType *UpdateCatalogItemTypeRequestCatalogItemType `json:"catalogItemType,omitempty"`
}

// NewUpdateCatalogItemTypeRequest instantiates a new UpdateCatalogItemTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCatalogItemTypeRequest() *UpdateCatalogItemTypeRequest {
	this := UpdateCatalogItemTypeRequest{}
	return &this
}

// NewUpdateCatalogItemTypeRequestWithDefaults instantiates a new UpdateCatalogItemTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCatalogItemTypeRequestWithDefaults() *UpdateCatalogItemTypeRequest {
	this := UpdateCatalogItemTypeRequest{}
	return &this
}

// GetCatalogItemType returns the CatalogItemType field value if set, zero value otherwise.
func (o *UpdateCatalogItemTypeRequest) GetCatalogItemType() UpdateCatalogItemTypeRequestCatalogItemType {
	if o == nil || IsNil(o.CatalogItemType) {
		var ret UpdateCatalogItemTypeRequestCatalogItemType
		return ret
	}
	return *o.CatalogItemType
}

// GetCatalogItemTypeOk returns a tuple with the CatalogItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCatalogItemTypeRequest) GetCatalogItemTypeOk() (*UpdateCatalogItemTypeRequestCatalogItemType, bool) {
	if o == nil || IsNil(o.CatalogItemType) {
		return nil, false
	}
	return o.CatalogItemType, true
}

// HasCatalogItemType returns a boolean if a field has been set.
func (o *UpdateCatalogItemTypeRequest) HasCatalogItemType() bool {
	if o != nil && !IsNil(o.CatalogItemType) {
		return true
	}

	return false
}

// SetCatalogItemType gets a reference to the given UpdateCatalogItemTypeRequestCatalogItemType and assigns it to the CatalogItemType field.
func (o *UpdateCatalogItemTypeRequest) SetCatalogItemType(v UpdateCatalogItemTypeRequestCatalogItemType) {
	o.CatalogItemType = &v
}

func (o UpdateCatalogItemTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCatalogItemTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CatalogItemType) {
		toSerialize["catalogItemType"] = o.CatalogItemType
	}
	return toSerialize, nil
}

type NullableUpdateCatalogItemTypeRequest struct {
	value *UpdateCatalogItemTypeRequest
	isSet bool
}

func (v NullableUpdateCatalogItemTypeRequest) Get() *UpdateCatalogItemTypeRequest {
	return v.value
}

func (v *NullableUpdateCatalogItemTypeRequest) Set(val *UpdateCatalogItemTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCatalogItemTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCatalogItemTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCatalogItemTypeRequest(val *UpdateCatalogItemTypeRequest) *NullableUpdateCatalogItemTypeRequest {
	return &NullableUpdateCatalogItemTypeRequest{value: val, isSet: true}
}

func (v NullableUpdateCatalogItemTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCatalogItemTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

