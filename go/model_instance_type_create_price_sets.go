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

// InstanceTypeCreatePriceSets struct for InstanceTypeCreatePriceSets
type InstanceTypeCreatePriceSets struct {
	// Price Set ID
	Id int64 `json:"id"`
}

// NewInstanceTypeCreatePriceSets instantiates a new InstanceTypeCreatePriceSets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceTypeCreatePriceSets(id int64, ) *InstanceTypeCreatePriceSets {
	this := InstanceTypeCreatePriceSets{}
	this.Id = id
	return &this
}

// NewInstanceTypeCreatePriceSetsWithDefaults instantiates a new InstanceTypeCreatePriceSets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceTypeCreatePriceSetsWithDefaults() *InstanceTypeCreatePriceSets {
	this := InstanceTypeCreatePriceSets{}
	return &this
}

// GetId returns the Id field value
func (o *InstanceTypeCreatePriceSets) GetId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InstanceTypeCreatePriceSets) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InstanceTypeCreatePriceSets) SetId(v int64) {
	o.Id = v
}

func (o InstanceTypeCreatePriceSets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableInstanceTypeCreatePriceSets struct {
	value *InstanceTypeCreatePriceSets
	isSet bool
}

func (v NullableInstanceTypeCreatePriceSets) Get() *InstanceTypeCreatePriceSets {
	return v.value
}

func (v *NullableInstanceTypeCreatePriceSets) Set(val *InstanceTypeCreatePriceSets) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceTypeCreatePriceSets) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceTypeCreatePriceSets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceTypeCreatePriceSets(val *InstanceTypeCreatePriceSets) *NullableInstanceTypeCreatePriceSets {
	return &NullableInstanceTypeCreatePriceSets{value: val, isSet: true}
}

func (v NullableInstanceTypeCreatePriceSets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceTypeCreatePriceSets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


