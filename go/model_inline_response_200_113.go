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

// InlineResponse200113 struct for InlineResponse200113
type InlineResponse200113 struct {
	NetworkDhcpServer *map[string]interface{} `json:"networkDhcpServer,omitempty"`
}

// NewInlineResponse200113 instantiates a new InlineResponse200113 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200113() *InlineResponse200113 {
	this := InlineResponse200113{}
	return &this
}

// NewInlineResponse200113WithDefaults instantiates a new InlineResponse200113 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200113WithDefaults() *InlineResponse200113 {
	this := InlineResponse200113{}
	return &this
}

// GetNetworkDhcpServer returns the NetworkDhcpServer field value if set, zero value otherwise.
func (o *InlineResponse200113) GetNetworkDhcpServer() map[string]interface{} {
	if o == nil || o.NetworkDhcpServer == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.NetworkDhcpServer
}

// GetNetworkDhcpServerOk returns a tuple with the NetworkDhcpServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200113) GetNetworkDhcpServerOk() (*map[string]interface{}, bool) {
	if o == nil || o.NetworkDhcpServer == nil {
		return nil, false
	}
	return o.NetworkDhcpServer, true
}

// HasNetworkDhcpServer returns a boolean if a field has been set.
func (o *InlineResponse200113) HasNetworkDhcpServer() bool {
	if o != nil && o.NetworkDhcpServer != nil {
		return true
	}

	return false
}

// SetNetworkDhcpServer gets a reference to the given map[string]interface{} and assigns it to the NetworkDhcpServer field.
func (o *InlineResponse200113) SetNetworkDhcpServer(v map[string]interface{}) {
	o.NetworkDhcpServer = &v
}

func (o InlineResponse200113) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkDhcpServer != nil {
		toSerialize["networkDhcpServer"] = o.NetworkDhcpServer
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200113 struct {
	value *InlineResponse200113
	isSet bool
}

func (v NullableInlineResponse200113) Get() *InlineResponse200113 {
	return v.value
}

func (v *NullableInlineResponse200113) Set(val *InlineResponse200113) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200113) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200113) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200113(val *InlineResponse200113) *NullableInlineResponse200113 {
	return &NullableInlineResponse200113{value: val, isSet: true}
}

func (v NullableInlineResponse200113) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200113) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


