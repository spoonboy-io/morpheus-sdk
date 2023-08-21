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

// VirtualImageCreateTags struct for VirtualImageCreateTags
type VirtualImageCreateTags struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

// NewVirtualImageCreateTags instantiates a new VirtualImageCreateTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualImageCreateTags(name string, value string, ) *VirtualImageCreateTags {
	this := VirtualImageCreateTags{}
	this.Name = name
	this.Value = value
	return &this
}

// NewVirtualImageCreateTagsWithDefaults instantiates a new VirtualImageCreateTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualImageCreateTagsWithDefaults() *VirtualImageCreateTags {
	this := VirtualImageCreateTags{}
	return &this
}

// GetName returns the Name field value
func (o *VirtualImageCreateTags) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VirtualImageCreateTags) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VirtualImageCreateTags) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *VirtualImageCreateTags) GetValue() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VirtualImageCreateTags) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VirtualImageCreateTags) SetValue(v string) {
	o.Value = v
}

func (o VirtualImageCreateTags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualImageCreateTags struct {
	value *VirtualImageCreateTags
	isSet bool
}

func (v NullableVirtualImageCreateTags) Get() *VirtualImageCreateTags {
	return v.value
}

func (v *NullableVirtualImageCreateTags) Set(val *VirtualImageCreateTags) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualImageCreateTags) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualImageCreateTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualImageCreateTags(val *VirtualImageCreateTags) *NullableVirtualImageCreateTags {
	return &NullableVirtualImageCreateTags{value: val, isSet: true}
}

func (v NullableVirtualImageCreateTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualImageCreateTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


