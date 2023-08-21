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

// InlineObject143 struct for InlineObject143
type InlineObject143 struct {
	Network *NetworkUpdate `json:"network,omitempty"`
}

// NewInlineObject143 instantiates a new InlineObject143 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject143() *InlineObject143 {
	this := InlineObject143{}
	return &this
}

// NewInlineObject143WithDefaults instantiates a new InlineObject143 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject143WithDefaults() *InlineObject143 {
	this := InlineObject143{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineObject143) GetNetwork() NetworkUpdate {
	if o == nil || o.Network == nil {
		var ret NetworkUpdate
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject143) GetNetworkOk() (*NetworkUpdate, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineObject143) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given NetworkUpdate and assigns it to the Network field.
func (o *InlineObject143) SetNetwork(v NetworkUpdate) {
	o.Network = &v
}

func (o InlineObject143) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject143 struct {
	value *InlineObject143
	isSet bool
}

func (v NullableInlineObject143) Get() *InlineObject143 {
	return v.value
}

func (v *NullableInlineObject143) Set(val *InlineObject143) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject143) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject143) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject143(val *InlineObject143) *NullableInlineObject143 {
	return &NullableInlineObject143{value: val, isSet: true}
}

func (v NullableInlineObject143) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject143) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


