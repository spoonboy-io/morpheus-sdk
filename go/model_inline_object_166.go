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

// InlineObject166 struct for InlineObject166
type InlineObject166 struct {
	NetworkProxy *ApiNetworksProxiesNetworkProxy `json:"networkProxy,omitempty"`
}

// NewInlineObject166 instantiates a new InlineObject166 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject166() *InlineObject166 {
	this := InlineObject166{}
	return &this
}

// NewInlineObject166WithDefaults instantiates a new InlineObject166 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject166WithDefaults() *InlineObject166 {
	this := InlineObject166{}
	return &this
}

// GetNetworkProxy returns the NetworkProxy field value if set, zero value otherwise.
func (o *InlineObject166) GetNetworkProxy() ApiNetworksProxiesNetworkProxy {
	if o == nil || o.NetworkProxy == nil {
		var ret ApiNetworksProxiesNetworkProxy
		return ret
	}
	return *o.NetworkProxy
}

// GetNetworkProxyOk returns a tuple with the NetworkProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject166) GetNetworkProxyOk() (*ApiNetworksProxiesNetworkProxy, bool) {
	if o == nil || o.NetworkProxy == nil {
		return nil, false
	}
	return o.NetworkProxy, true
}

// HasNetworkProxy returns a boolean if a field has been set.
func (o *InlineObject166) HasNetworkProxy() bool {
	if o != nil && o.NetworkProxy != nil {
		return true
	}

	return false
}

// SetNetworkProxy gets a reference to the given ApiNetworksProxiesNetworkProxy and assigns it to the NetworkProxy field.
func (o *InlineObject166) SetNetworkProxy(v ApiNetworksProxiesNetworkProxy) {
	o.NetworkProxy = &v
}

func (o InlineObject166) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkProxy != nil {
		toSerialize["networkProxy"] = o.NetworkProxy
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject166 struct {
	value *InlineObject166
	isSet bool
}

func (v NullableInlineObject166) Get() *InlineObject166 {
	return v.value
}

func (v *NullableInlineObject166) Set(val *InlineObject166) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject166) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject166) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject166(val *InlineObject166) *NullableInlineObject166 {
	return &NullableInlineObject166{value: val, isSet: true}
}

func (v NullableInlineObject166) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject166) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

