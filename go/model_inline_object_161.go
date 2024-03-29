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

// InlineObject161 struct for InlineObject161
type InlineObject161 struct {
	NetworkPool *NetworkPoolCreate `json:"networkPool,omitempty"`
}

// NewInlineObject161 instantiates a new InlineObject161 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject161() *InlineObject161 {
	this := InlineObject161{}
	return &this
}

// NewInlineObject161WithDefaults instantiates a new InlineObject161 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject161WithDefaults() *InlineObject161 {
	this := InlineObject161{}
	return &this
}

// GetNetworkPool returns the NetworkPool field value if set, zero value otherwise.
func (o *InlineObject161) GetNetworkPool() NetworkPoolCreate {
	if o == nil || o.NetworkPool == nil {
		var ret NetworkPoolCreate
		return ret
	}
	return *o.NetworkPool
}

// GetNetworkPoolOk returns a tuple with the NetworkPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject161) GetNetworkPoolOk() (*NetworkPoolCreate, bool) {
	if o == nil || o.NetworkPool == nil {
		return nil, false
	}
	return o.NetworkPool, true
}

// HasNetworkPool returns a boolean if a field has been set.
func (o *InlineObject161) HasNetworkPool() bool {
	if o != nil && o.NetworkPool != nil {
		return true
	}

	return false
}

// SetNetworkPool gets a reference to the given NetworkPoolCreate and assigns it to the NetworkPool field.
func (o *InlineObject161) SetNetworkPool(v NetworkPoolCreate) {
	o.NetworkPool = &v
}

func (o InlineObject161) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkPool != nil {
		toSerialize["networkPool"] = o.NetworkPool
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject161 struct {
	value *InlineObject161
	isSet bool
}

func (v NullableInlineObject161) Get() *InlineObject161 {
	return v.value
}

func (v *NullableInlineObject161) Set(val *InlineObject161) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject161) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject161) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject161(val *InlineObject161) *NullableInlineObject161 {
	return &NullableInlineObject161{value: val, isSet: true}
}

func (v NullableInlineObject161) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject161) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


