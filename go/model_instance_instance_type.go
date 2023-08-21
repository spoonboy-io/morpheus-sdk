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

// InstanceInstanceType struct for InstanceInstanceType
type InstanceInstanceType struct {
	Id *int64 `json:"id,omitempty"`
	Code *string `json:"code,omitempty"`
	Category *string `json:"category,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewInstanceInstanceType instantiates a new InstanceInstanceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceInstanceType() *InstanceInstanceType {
	this := InstanceInstanceType{}
	return &this
}

// NewInstanceInstanceTypeWithDefaults instantiates a new InstanceInstanceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceInstanceTypeWithDefaults() *InstanceInstanceType {
	this := InstanceInstanceType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InstanceInstanceType) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceInstanceType) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InstanceInstanceType) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *InstanceInstanceType) SetId(v int64) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InstanceInstanceType) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceInstanceType) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InstanceInstanceType) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *InstanceInstanceType) SetCode(v string) {
	o.Code = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *InstanceInstanceType) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceInstanceType) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *InstanceInstanceType) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *InstanceInstanceType) SetCategory(v string) {
	o.Category = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InstanceInstanceType) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceInstanceType) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InstanceInstanceType) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InstanceInstanceType) SetName(v string) {
	o.Name = &v
}

func (o InstanceInstanceType) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableInstanceInstanceType struct {
	value *InstanceInstanceType
	isSet bool
}

func (v NullableInstanceInstanceType) Get() *InstanceInstanceType {
	return v.value
}

func (v *NullableInstanceInstanceType) Set(val *InstanceInstanceType) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceInstanceType) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceInstanceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceInstanceType(val *InstanceInstanceType) *NullableInstanceInstanceType {
	return &NullableInstanceInstanceType{value: val, isSet: true}
}

func (v NullableInstanceInstanceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceInstanceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


