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

// NetworkRoutersCreateType struct for NetworkRoutersCreateType
type NetworkRoutersCreateType struct {
	// Network router type ID
	Id int64 `json:"id"`
}

// NewNetworkRoutersCreateType instantiates a new NetworkRoutersCreateType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkRoutersCreateType(id int64, ) *NetworkRoutersCreateType {
	this := NetworkRoutersCreateType{}
	this.Id = id
	return &this
}

// NewNetworkRoutersCreateTypeWithDefaults instantiates a new NetworkRoutersCreateType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkRoutersCreateTypeWithDefaults() *NetworkRoutersCreateType {
	this := NetworkRoutersCreateType{}
	return &this
}

// GetId returns the Id field value
func (o *NetworkRoutersCreateType) GetId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NetworkRoutersCreateType) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NetworkRoutersCreateType) SetId(v int64) {
	o.Id = v
}

func (o NetworkRoutersCreateType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkRoutersCreateType struct {
	value *NetworkRoutersCreateType
	isSet bool
}

func (v NullableNetworkRoutersCreateType) Get() *NetworkRoutersCreateType {
	return v.value
}

func (v *NullableNetworkRoutersCreateType) Set(val *NetworkRoutersCreateType) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkRoutersCreateType) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkRoutersCreateType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkRoutersCreateType(val *NetworkRoutersCreateType) *NullableNetworkRoutersCreateType {
	return &NullableNetworkRoutersCreateType{value: val, isSet: true}
}

func (v NullableNetworkRoutersCreateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkRoutersCreateType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


