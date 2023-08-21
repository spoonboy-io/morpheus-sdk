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

// InstanceCreateNetworkNetwork struct for InstanceCreateNetworkNetwork
type InstanceCreateNetworkNetwork struct {
	// id of the network to be used. A network group can be specified instead by prefixing its ID with `networkGroup-`.
	Id string `json:"id"`
}

// NewInstanceCreateNetworkNetwork instantiates a new InstanceCreateNetworkNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceCreateNetworkNetwork(id string, ) *InstanceCreateNetworkNetwork {
	this := InstanceCreateNetworkNetwork{}
	this.Id = id
	return &this
}

// NewInstanceCreateNetworkNetworkWithDefaults instantiates a new InstanceCreateNetworkNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceCreateNetworkNetworkWithDefaults() *InstanceCreateNetworkNetwork {
	this := InstanceCreateNetworkNetwork{}
	return &this
}

// GetId returns the Id field value
func (o *InstanceCreateNetworkNetwork) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InstanceCreateNetworkNetwork) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InstanceCreateNetworkNetwork) SetId(v string) {
	o.Id = v
}

func (o InstanceCreateNetworkNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableInstanceCreateNetworkNetwork struct {
	value *InstanceCreateNetworkNetwork
	isSet bool
}

func (v NullableInstanceCreateNetworkNetwork) Get() *InstanceCreateNetworkNetwork {
	return v.value
}

func (v *NullableInstanceCreateNetworkNetwork) Set(val *InstanceCreateNetworkNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceCreateNetworkNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceCreateNetworkNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceCreateNetworkNetwork(val *InstanceCreateNetworkNetwork) *NullableInstanceCreateNetworkNetwork {
	return &NullableInstanceCreateNetworkNetwork{value: val, isSet: true}
}

func (v NullableInstanceCreateNetworkNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceCreateNetworkNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

