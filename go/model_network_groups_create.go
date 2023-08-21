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

// NetworkGroupsCreate struct for NetworkGroupsCreate
type NetworkGroupsCreate struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Networks *[]int64 `json:"networks,omitempty"`
	Subnets *[]map[string]interface{} `json:"subnets,omitempty"`
}

// NewNetworkGroupsCreate instantiates a new NetworkGroupsCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkGroupsCreate() *NetworkGroupsCreate {
	this := NetworkGroupsCreate{}
	return &this
}

// NewNetworkGroupsCreateWithDefaults instantiates a new NetworkGroupsCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkGroupsCreateWithDefaults() *NetworkGroupsCreate {
	this := NetworkGroupsCreate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkGroupsCreate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkGroupsCreate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkGroupsCreate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkGroupsCreate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NetworkGroupsCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkGroupsCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NetworkGroupsCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NetworkGroupsCreate) SetDescription(v string) {
	o.Description = &v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *NetworkGroupsCreate) GetNetworks() []int64 {
	if o == nil || o.Networks == nil {
		var ret []int64
		return ret
	}
	return *o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkGroupsCreate) GetNetworksOk() (*[]int64, bool) {
	if o == nil || o.Networks == nil {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *NetworkGroupsCreate) HasNetworks() bool {
	if o != nil && o.Networks != nil {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []int64 and assigns it to the Networks field.
func (o *NetworkGroupsCreate) SetNetworks(v []int64) {
	o.Networks = &v
}

// GetSubnets returns the Subnets field value if set, zero value otherwise.
func (o *NetworkGroupsCreate) GetSubnets() []map[string]interface{} {
	if o == nil || o.Subnets == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Subnets
}

// GetSubnetsOk returns a tuple with the Subnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkGroupsCreate) GetSubnetsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Subnets == nil {
		return nil, false
	}
	return o.Subnets, true
}

// HasSubnets returns a boolean if a field has been set.
func (o *NetworkGroupsCreate) HasSubnets() bool {
	if o != nil && o.Subnets != nil {
		return true
	}

	return false
}

// SetSubnets gets a reference to the given []map[string]interface{} and assigns it to the Subnets field.
func (o *NetworkGroupsCreate) SetSubnets(v []map[string]interface{}) {
	o.Subnets = &v
}

func (o NetworkGroupsCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Networks != nil {
		toSerialize["networks"] = o.Networks
	}
	if o.Subnets != nil {
		toSerialize["subnets"] = o.Subnets
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkGroupsCreate struct {
	value *NetworkGroupsCreate
	isSet bool
}

func (v NullableNetworkGroupsCreate) Get() *NetworkGroupsCreate {
	return v.value
}

func (v *NullableNetworkGroupsCreate) Set(val *NetworkGroupsCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkGroupsCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkGroupsCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkGroupsCreate(val *NetworkGroupsCreate) *NullableNetworkGroupsCreate {
	return &NullableNetworkGroupsCreate{value: val, isSet: true}
}

func (v NullableNetworkGroupsCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkGroupsCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

