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

// NetworkPoolIpCreate struct for NetworkPoolIpCreate
type NetworkPoolIpCreate struct {
	// IP Address
	IpAddress *string `json:"ipAddress,omitempty"`
	// Hostname
	Hostname *string `json:"hostname,omitempty"`
}

// NewNetworkPoolIpCreate instantiates a new NetworkPoolIpCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkPoolIpCreate() *NetworkPoolIpCreate {
	this := NetworkPoolIpCreate{}
	return &this
}

// NewNetworkPoolIpCreateWithDefaults instantiates a new NetworkPoolIpCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkPoolIpCreateWithDefaults() *NetworkPoolIpCreate {
	this := NetworkPoolIpCreate{}
	return &this
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *NetworkPoolIpCreate) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolIpCreate) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *NetworkPoolIpCreate) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *NetworkPoolIpCreate) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *NetworkPoolIpCreate) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolIpCreate) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *NetworkPoolIpCreate) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *NetworkPoolIpCreate) SetHostname(v string) {
	o.Hostname = &v
}

func (o NetworkPoolIpCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IpAddress != nil {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkPoolIpCreate struct {
	value *NetworkPoolIpCreate
	isSet bool
}

func (v NullableNetworkPoolIpCreate) Get() *NetworkPoolIpCreate {
	return v.value
}

func (v *NullableNetworkPoolIpCreate) Set(val *NetworkPoolIpCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkPoolIpCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkPoolIpCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkPoolIpCreate(val *NetworkPoolIpCreate) *NullableNetworkPoolIpCreate {
	return &NullableNetworkPoolIpCreate{value: val, isSet: true}
}

func (v NullableNetworkPoolIpCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkPoolIpCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


