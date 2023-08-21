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

// VirtualImageUpdateRemoveTags struct for VirtualImageUpdateRemoveTags
type VirtualImageUpdateRemoveTags struct {
	Name string `json:"name"`
	Value *string `json:"value,omitempty"`
}

// NewVirtualImageUpdateRemoveTags instantiates a new VirtualImageUpdateRemoveTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualImageUpdateRemoveTags(name string, ) *VirtualImageUpdateRemoveTags {
	this := VirtualImageUpdateRemoveTags{}
	this.Name = name
	return &this
}

// NewVirtualImageUpdateRemoveTagsWithDefaults instantiates a new VirtualImageUpdateRemoveTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualImageUpdateRemoveTagsWithDefaults() *VirtualImageUpdateRemoveTags {
	this := VirtualImageUpdateRemoveTags{}
	return &this
}

// GetName returns the Name field value
func (o *VirtualImageUpdateRemoveTags) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VirtualImageUpdateRemoveTags) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VirtualImageUpdateRemoveTags) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *VirtualImageUpdateRemoveTags) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualImageUpdateRemoveTags) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *VirtualImageUpdateRemoveTags) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *VirtualImageUpdateRemoveTags) SetValue(v string) {
	o.Value = &v
}

func (o VirtualImageUpdateRemoveTags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualImageUpdateRemoveTags struct {
	value *VirtualImageUpdateRemoveTags
	isSet bool
}

func (v NullableVirtualImageUpdateRemoveTags) Get() *VirtualImageUpdateRemoveTags {
	return v.value
}

func (v *NullableVirtualImageUpdateRemoveTags) Set(val *VirtualImageUpdateRemoveTags) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualImageUpdateRemoveTags) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualImageUpdateRemoveTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualImageUpdateRemoveTags(val *VirtualImageUpdateRemoveTags) *NullableVirtualImageUpdateRemoveTags {
	return &NullableVirtualImageUpdateRemoveTags{value: val, isSet: true}
}

func (v NullableVirtualImageUpdateRemoveTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualImageUpdateRemoveTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


