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

// NetworkPoolCreateIpRanges struct for NetworkPoolCreateIpRanges
type NetworkPoolCreateIpRanges struct {
	// Starting IP Address
	StartAddress *string `json:"startAddress,omitempty"`
	// Ending IP Address
	EndAddress *string `json:"endAddress,omitempty"`
	// IPv6 Network CIDR
	CidrIPv6 *string `json:"cidrIPv6,omitempty"`
}

// NewNetworkPoolCreateIpRanges instantiates a new NetworkPoolCreateIpRanges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkPoolCreateIpRanges() *NetworkPoolCreateIpRanges {
	this := NetworkPoolCreateIpRanges{}
	return &this
}

// NewNetworkPoolCreateIpRangesWithDefaults instantiates a new NetworkPoolCreateIpRanges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkPoolCreateIpRangesWithDefaults() *NetworkPoolCreateIpRanges {
	this := NetworkPoolCreateIpRanges{}
	return &this
}

// GetStartAddress returns the StartAddress field value if set, zero value otherwise.
func (o *NetworkPoolCreateIpRanges) GetStartAddress() string {
	if o == nil || o.StartAddress == nil {
		var ret string
		return ret
	}
	return *o.StartAddress
}

// GetStartAddressOk returns a tuple with the StartAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolCreateIpRanges) GetStartAddressOk() (*string, bool) {
	if o == nil || o.StartAddress == nil {
		return nil, false
	}
	return o.StartAddress, true
}

// HasStartAddress returns a boolean if a field has been set.
func (o *NetworkPoolCreateIpRanges) HasStartAddress() bool {
	if o != nil && o.StartAddress != nil {
		return true
	}

	return false
}

// SetStartAddress gets a reference to the given string and assigns it to the StartAddress field.
func (o *NetworkPoolCreateIpRanges) SetStartAddress(v string) {
	o.StartAddress = &v
}

// GetEndAddress returns the EndAddress field value if set, zero value otherwise.
func (o *NetworkPoolCreateIpRanges) GetEndAddress() string {
	if o == nil || o.EndAddress == nil {
		var ret string
		return ret
	}
	return *o.EndAddress
}

// GetEndAddressOk returns a tuple with the EndAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolCreateIpRanges) GetEndAddressOk() (*string, bool) {
	if o == nil || o.EndAddress == nil {
		return nil, false
	}
	return o.EndAddress, true
}

// HasEndAddress returns a boolean if a field has been set.
func (o *NetworkPoolCreateIpRanges) HasEndAddress() bool {
	if o != nil && o.EndAddress != nil {
		return true
	}

	return false
}

// SetEndAddress gets a reference to the given string and assigns it to the EndAddress field.
func (o *NetworkPoolCreateIpRanges) SetEndAddress(v string) {
	o.EndAddress = &v
}

// GetCidrIPv6 returns the CidrIPv6 field value if set, zero value otherwise.
func (o *NetworkPoolCreateIpRanges) GetCidrIPv6() string {
	if o == nil || o.CidrIPv6 == nil {
		var ret string
		return ret
	}
	return *o.CidrIPv6
}

// GetCidrIPv6Ok returns a tuple with the CidrIPv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolCreateIpRanges) GetCidrIPv6Ok() (*string, bool) {
	if o == nil || o.CidrIPv6 == nil {
		return nil, false
	}
	return o.CidrIPv6, true
}

// HasCidrIPv6 returns a boolean if a field has been set.
func (o *NetworkPoolCreateIpRanges) HasCidrIPv6() bool {
	if o != nil && o.CidrIPv6 != nil {
		return true
	}

	return false
}

// SetCidrIPv6 gets a reference to the given string and assigns it to the CidrIPv6 field.
func (o *NetworkPoolCreateIpRanges) SetCidrIPv6(v string) {
	o.CidrIPv6 = &v
}

func (o NetworkPoolCreateIpRanges) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StartAddress != nil {
		toSerialize["startAddress"] = o.StartAddress
	}
	if o.EndAddress != nil {
		toSerialize["endAddress"] = o.EndAddress
	}
	if o.CidrIPv6 != nil {
		toSerialize["cidrIPv6"] = o.CidrIPv6
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkPoolCreateIpRanges struct {
	value *NetworkPoolCreateIpRanges
	isSet bool
}

func (v NullableNetworkPoolCreateIpRanges) Get() *NetworkPoolCreateIpRanges {
	return v.value
}

func (v *NullableNetworkPoolCreateIpRanges) Set(val *NetworkPoolCreateIpRanges) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkPoolCreateIpRanges) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkPoolCreateIpRanges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkPoolCreateIpRanges(val *NetworkPoolCreateIpRanges) *NullableNetworkPoolCreateIpRanges {
	return &NullableNetworkPoolCreateIpRanges{value: val, isSet: true}
}

func (v NullableNetworkPoolCreateIpRanges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkPoolCreateIpRanges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


