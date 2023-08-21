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

// InlineObject170 The parameters for update a Network DHCP Server is type dependent. The following lists the common parameters. Get a specific network type to list available options for the network server type. 
type InlineObject170 struct {
	NetworkDhcpServer *map[string]interface{} `json:"networkDhcpServer,omitempty"`
}

// NewInlineObject170 instantiates a new InlineObject170 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject170() *InlineObject170 {
	this := InlineObject170{}
	return &this
}

// NewInlineObject170WithDefaults instantiates a new InlineObject170 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject170WithDefaults() *InlineObject170 {
	this := InlineObject170{}
	return &this
}

// GetNetworkDhcpServer returns the NetworkDhcpServer field value if set, zero value otherwise.
func (o *InlineObject170) GetNetworkDhcpServer() map[string]interface{} {
	if o == nil || o.NetworkDhcpServer == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.NetworkDhcpServer
}

// GetNetworkDhcpServerOk returns a tuple with the NetworkDhcpServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject170) GetNetworkDhcpServerOk() (*map[string]interface{}, bool) {
	if o == nil || o.NetworkDhcpServer == nil {
		return nil, false
	}
	return o.NetworkDhcpServer, true
}

// HasNetworkDhcpServer returns a boolean if a field has been set.
func (o *InlineObject170) HasNetworkDhcpServer() bool {
	if o != nil && o.NetworkDhcpServer != nil {
		return true
	}

	return false
}

// SetNetworkDhcpServer gets a reference to the given map[string]interface{} and assigns it to the NetworkDhcpServer field.
func (o *InlineObject170) SetNetworkDhcpServer(v map[string]interface{}) {
	o.NetworkDhcpServer = &v
}

func (o InlineObject170) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkDhcpServer != nil {
		toSerialize["networkDhcpServer"] = o.NetworkDhcpServer
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject170 struct {
	value *InlineObject170
	isSet bool
}

func (v NullableInlineObject170) Get() *InlineObject170 {
	return v.value
}

func (v *NullableInlineObject170) Set(val *InlineObject170) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject170) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject170) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject170(val *InlineObject170) *NullableInlineObject170 {
	return &NullableInlineObject170{value: val, isSet: true}
}

func (v NullableInlineObject170) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject170) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

