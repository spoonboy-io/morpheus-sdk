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

// ContainerContainerTypeSet struct for ContainerContainerTypeSet
type ContainerContainerTypeSet struct {
	Id *int32 `json:"id,omitempty"`
	Code *string `json:"code,omitempty"`
	Category *string `json:"category,omitempty"`
}

// NewContainerContainerTypeSet instantiates a new ContainerContainerTypeSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerContainerTypeSet() *ContainerContainerTypeSet {
	this := ContainerContainerTypeSet{}
	return &this
}

// NewContainerContainerTypeSetWithDefaults instantiates a new ContainerContainerTypeSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerContainerTypeSetWithDefaults() *ContainerContainerTypeSet {
	this := ContainerContainerTypeSet{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContainerContainerTypeSet) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerTypeSet) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContainerContainerTypeSet) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ContainerContainerTypeSet) SetId(v int32) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ContainerContainerTypeSet) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerTypeSet) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ContainerContainerTypeSet) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ContainerContainerTypeSet) SetCode(v string) {
	o.Code = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ContainerContainerTypeSet) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerTypeSet) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ContainerContainerTypeSet) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *ContainerContainerTypeSet) SetCategory(v string) {
	o.Category = &v
}

func (o ContainerContainerTypeSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	return json.Marshal(toSerialize)
}

type NullableContainerContainerTypeSet struct {
	value *ContainerContainerTypeSet
	isSet bool
}

func (v NullableContainerContainerTypeSet) Get() *ContainerContainerTypeSet {
	return v.value
}

func (v *NullableContainerContainerTypeSet) Set(val *ContainerContainerTypeSet) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerContainerTypeSet) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerContainerTypeSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerContainerTypeSet(val *ContainerContainerTypeSet) *NullableContainerContainerTypeSet {
	return &NullableContainerContainerTypeSet{value: val, isSet: true}
}

func (v NullableContainerContainerTypeSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerContainerTypeSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


