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

// ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT For a full list of available NAT options, see natOptionTypes in the specific Network Router Type
type ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT struct {
	// Sets name of NAT
	Name interface{} `json:"name,omitempty"`
}

// NewApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT instantiates a new ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT() *ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT {
	this := ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT{}
	return &this
}

// NewApiNetworksRoutersRouterIdNatsIdNetworkRouterNATWithDefaults instantiates a new ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiNetworksRoutersRouterIdNatsIdNetworkRouterNATWithDefaults() *ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT {
	this := ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT) GetName() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT) GetNameOk() (*interface{}, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT) SetName(v interface{}) {
	o.Name = v
}

func (o ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT struct {
	value *ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT
	isSet bool
}

func (v NullableApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT) Get() *ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT {
	return v.value
}

func (v *NullableApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT) Set(val *ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT) {
	v.value = val
	v.isSet = true
}

func (v NullableApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT) IsSet() bool {
	return v.isSet
}

func (v *NullableApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT(val *ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT) *NullableApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT {
	return &NullableApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT{value: val, isSet: true}
}

func (v NullableApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


