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

// InlineResponse20093 struct for InlineResponse20093
type InlineResponse20093 struct {
	NetworkRouterType *NetworkRouterType `json:"networkRouterType,omitempty"`
}

// NewInlineResponse20093 instantiates a new InlineResponse20093 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20093() *InlineResponse20093 {
	this := InlineResponse20093{}
	return &this
}

// NewInlineResponse20093WithDefaults instantiates a new InlineResponse20093 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20093WithDefaults() *InlineResponse20093 {
	this := InlineResponse20093{}
	return &this
}

// GetNetworkRouterType returns the NetworkRouterType field value if set, zero value otherwise.
func (o *InlineResponse20093) GetNetworkRouterType() NetworkRouterType {
	if o == nil || o.NetworkRouterType == nil {
		var ret NetworkRouterType
		return ret
	}
	return *o.NetworkRouterType
}

// GetNetworkRouterTypeOk returns a tuple with the NetworkRouterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20093) GetNetworkRouterTypeOk() (*NetworkRouterType, bool) {
	if o == nil || o.NetworkRouterType == nil {
		return nil, false
	}
	return o.NetworkRouterType, true
}

// HasNetworkRouterType returns a boolean if a field has been set.
func (o *InlineResponse20093) HasNetworkRouterType() bool {
	if o != nil && o.NetworkRouterType != nil {
		return true
	}

	return false
}

// SetNetworkRouterType gets a reference to the given NetworkRouterType and assigns it to the NetworkRouterType field.
func (o *InlineResponse20093) SetNetworkRouterType(v NetworkRouterType) {
	o.NetworkRouterType = &v
}

func (o InlineResponse20093) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkRouterType != nil {
		toSerialize["networkRouterType"] = o.NetworkRouterType
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20093 struct {
	value *InlineResponse20093
	isSet bool
}

func (v NullableInlineResponse20093) Get() *InlineResponse20093 {
	return v.value
}

func (v *NullableInlineResponse20093) Set(val *InlineResponse20093) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20093) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20093) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20093(val *InlineResponse20093) *NullableInlineResponse20093 {
	return &NullableInlineResponse20093{value: val, isSet: true}
}

func (v NullableInlineResponse20093) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20093) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


