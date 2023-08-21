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

// ClusterContainersContainerType struct for ClusterContainersContainerType
type ClusterContainersContainerType struct {
	Id *int64 `json:"id,omitempty"`
	Code *string `json:"code,omitempty"`
	Category NullableString `json:"category,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewClusterContainersContainerType instantiates a new ClusterContainersContainerType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterContainersContainerType() *ClusterContainersContainerType {
	this := ClusterContainersContainerType{}
	return &this
}

// NewClusterContainersContainerTypeWithDefaults instantiates a new ClusterContainersContainerType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterContainersContainerTypeWithDefaults() *ClusterContainersContainerType {
	this := ClusterContainersContainerType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClusterContainersContainerType) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterContainersContainerType) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClusterContainersContainerType) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ClusterContainersContainerType) SetId(v int64) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ClusterContainersContainerType) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterContainersContainerType) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ClusterContainersContainerType) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ClusterContainersContainerType) SetCode(v string) {
	o.Code = &v
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterContainersContainerType) GetCategory() string {
	if o == nil || o.Category.Get() == nil {
		var ret string
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterContainersContainerType) GetCategoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// HasCategory returns a boolean if a field has been set.
func (o *ClusterContainersContainerType) HasCategory() bool {
	if o != nil && o.Category.IsSet() {
		return true
	}

	return false
}

// SetCategory gets a reference to the given NullableString and assigns it to the Category field.
func (o *ClusterContainersContainerType) SetCategory(v string) {
	o.Category.Set(&v)
}
// SetCategoryNil sets the value for Category to be an explicit nil
func (o *ClusterContainersContainerType) SetCategoryNil() {
	o.Category.Set(nil)
}

// UnsetCategory ensures that no value is present for Category, not even an explicit nil
func (o *ClusterContainersContainerType) UnsetCategory() {
	o.Category.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClusterContainersContainerType) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterContainersContainerType) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClusterContainersContainerType) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClusterContainersContainerType) SetName(v string) {
	o.Name = &v
}

func (o ClusterContainersContainerType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Category.IsSet() {
		toSerialize["category"] = o.Category.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableClusterContainersContainerType struct {
	value *ClusterContainersContainerType
	isSet bool
}

func (v NullableClusterContainersContainerType) Get() *ClusterContainersContainerType {
	return v.value
}

func (v *NullableClusterContainersContainerType) Set(val *ClusterContainersContainerType) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterContainersContainerType) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterContainersContainerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterContainersContainerType(val *ClusterContainersContainerType) *NullableClusterContainersContainerType {
	return &NullableClusterContainersContainerType{value: val, isSet: true}
}

func (v NullableClusterContainersContainerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterContainersContainerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

