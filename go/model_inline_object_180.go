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

// InlineObject180 struct for InlineObject180
type InlineObject180 struct {
	// Payload for creating a new Network Pool Server
	NetworkPoolServer *AnyOfnetworkPoolServerCreateBluecatnetworkPoolServerCreateInfobloxnetworkPoolServerCreatePhpIpamnetworkPoolServerCreateSolarWinds `json:"networkPoolServer,omitempty"`
}

// NewInlineObject180 instantiates a new InlineObject180 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject180() *InlineObject180 {
	this := InlineObject180{}
	return &this
}

// NewInlineObject180WithDefaults instantiates a new InlineObject180 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject180WithDefaults() *InlineObject180 {
	this := InlineObject180{}
	return &this
}

// GetNetworkPoolServer returns the NetworkPoolServer field value if set, zero value otherwise.
func (o *InlineObject180) GetNetworkPoolServer() AnyOfnetworkPoolServerCreateBluecatnetworkPoolServerCreateInfobloxnetworkPoolServerCreatePhpIpamnetworkPoolServerCreateSolarWinds {
	if o == nil || o.NetworkPoolServer == nil {
		var ret AnyOfnetworkPoolServerCreateBluecatnetworkPoolServerCreateInfobloxnetworkPoolServerCreatePhpIpamnetworkPoolServerCreateSolarWinds
		return ret
	}
	return *o.NetworkPoolServer
}

// GetNetworkPoolServerOk returns a tuple with the NetworkPoolServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject180) GetNetworkPoolServerOk() (*AnyOfnetworkPoolServerCreateBluecatnetworkPoolServerCreateInfobloxnetworkPoolServerCreatePhpIpamnetworkPoolServerCreateSolarWinds, bool) {
	if o == nil || o.NetworkPoolServer == nil {
		return nil, false
	}
	return o.NetworkPoolServer, true
}

// HasNetworkPoolServer returns a boolean if a field has been set.
func (o *InlineObject180) HasNetworkPoolServer() bool {
	if o != nil && o.NetworkPoolServer != nil {
		return true
	}

	return false
}

// SetNetworkPoolServer gets a reference to the given AnyOfnetworkPoolServerCreateBluecatnetworkPoolServerCreateInfobloxnetworkPoolServerCreatePhpIpamnetworkPoolServerCreateSolarWinds and assigns it to the NetworkPoolServer field.
func (o *InlineObject180) SetNetworkPoolServer(v AnyOfnetworkPoolServerCreateBluecatnetworkPoolServerCreateInfobloxnetworkPoolServerCreatePhpIpamnetworkPoolServerCreateSolarWinds) {
	o.NetworkPoolServer = &v
}

func (o InlineObject180) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkPoolServer != nil {
		toSerialize["networkPoolServer"] = o.NetworkPoolServer
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject180 struct {
	value *InlineObject180
	isSet bool
}

func (v NullableInlineObject180) Get() *InlineObject180 {
	return v.value
}

func (v *NullableInlineObject180) Set(val *InlineObject180) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject180) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject180) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject180(val *InlineObject180) *NullableInlineObject180 {
	return &NullableInlineObject180{value: val, isSet: true}
}

func (v NullableInlineObject180) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject180) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

