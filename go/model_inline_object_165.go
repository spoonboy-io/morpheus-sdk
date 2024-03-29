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

// InlineObject165 struct for InlineObject165
type InlineObject165 struct {
	NetworkProxy *ApiNetworksProxiesNetworkProxy `json:"networkProxy,omitempty"`
}

// NewInlineObject165 instantiates a new InlineObject165 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject165() *InlineObject165 {
	this := InlineObject165{}
	return &this
}

// NewInlineObject165WithDefaults instantiates a new InlineObject165 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject165WithDefaults() *InlineObject165 {
	this := InlineObject165{}
	return &this
}

// GetNetworkProxy returns the NetworkProxy field value if set, zero value otherwise.
func (o *InlineObject165) GetNetworkProxy() ApiNetworksProxiesNetworkProxy {
	if o == nil || o.NetworkProxy == nil {
		var ret ApiNetworksProxiesNetworkProxy
		return ret
	}
	return *o.NetworkProxy
}

// GetNetworkProxyOk returns a tuple with the NetworkProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject165) GetNetworkProxyOk() (*ApiNetworksProxiesNetworkProxy, bool) {
	if o == nil || o.NetworkProxy == nil {
		return nil, false
	}
	return o.NetworkProxy, true
}

// HasNetworkProxy returns a boolean if a field has been set.
func (o *InlineObject165) HasNetworkProxy() bool {
	if o != nil && o.NetworkProxy != nil {
		return true
	}

	return false
}

// SetNetworkProxy gets a reference to the given ApiNetworksProxiesNetworkProxy and assigns it to the NetworkProxy field.
func (o *InlineObject165) SetNetworkProxy(v ApiNetworksProxiesNetworkProxy) {
	o.NetworkProxy = &v
}

func (o InlineObject165) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkProxy != nil {
		toSerialize["networkProxy"] = o.NetworkProxy
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject165 struct {
	value *InlineObject165
	isSet bool
}

func (v NullableInlineObject165) Get() *InlineObject165 {
	return v.value
}

func (v *NullableInlineObject165) Set(val *InlineObject165) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject165) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject165) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject165(val *InlineObject165) *NullableInlineObject165 {
	return &NullableInlineObject165{value: val, isSet: true}
}

func (v NullableInlineObject165) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject165) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


