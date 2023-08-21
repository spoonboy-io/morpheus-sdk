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

// NetworkStaticRouteCreate struct for NetworkStaticRouteCreate
type NetworkStaticRouteCreate struct {
	// Source CIDR
	Source string `json:"source"`
	// Destination address
	Destination string `json:"destination"`
}

// NewNetworkStaticRouteCreate instantiates a new NetworkStaticRouteCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkStaticRouteCreate(source string, destination string, ) *NetworkStaticRouteCreate {
	this := NetworkStaticRouteCreate{}
	this.Source = source
	this.Destination = destination
	return &this
}

// NewNetworkStaticRouteCreateWithDefaults instantiates a new NetworkStaticRouteCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkStaticRouteCreateWithDefaults() *NetworkStaticRouteCreate {
	this := NetworkStaticRouteCreate{}
	return &this
}

// GetSource returns the Source field value
func (o *NetworkStaticRouteCreate) GetSource() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *NetworkStaticRouteCreate) GetSourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *NetworkStaticRouteCreate) SetSource(v string) {
	o.Source = v
}

// GetDestination returns the Destination field value
func (o *NetworkStaticRouteCreate) GetDestination() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *NetworkStaticRouteCreate) GetDestinationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *NetworkStaticRouteCreate) SetDestination(v string) {
	o.Destination = v
}

func (o NetworkStaticRouteCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["destination"] = o.Destination
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkStaticRouteCreate struct {
	value *NetworkStaticRouteCreate
	isSet bool
}

func (v NullableNetworkStaticRouteCreate) Get() *NetworkStaticRouteCreate {
	return v.value
}

func (v *NullableNetworkStaticRouteCreate) Set(val *NetworkStaticRouteCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkStaticRouteCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkStaticRouteCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkStaticRouteCreate(val *NetworkStaticRouteCreate) *NullableNetworkStaticRouteCreate {
	return &NullableNetworkStaticRouteCreate{value: val, isSet: true}
}

func (v NullableNetworkStaticRouteCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkStaticRouteCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

