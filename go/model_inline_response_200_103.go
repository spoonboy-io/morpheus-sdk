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

// InlineResponse200103 struct for InlineResponse200103
type InlineResponse200103 struct {
	NetworkRouterNAT *NetworkRouterNat `json:"networkRouterNAT,omitempty"`
}

// NewInlineResponse200103 instantiates a new InlineResponse200103 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200103() *InlineResponse200103 {
	this := InlineResponse200103{}
	return &this
}

// NewInlineResponse200103WithDefaults instantiates a new InlineResponse200103 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200103WithDefaults() *InlineResponse200103 {
	this := InlineResponse200103{}
	return &this
}

// GetNetworkRouterNAT returns the NetworkRouterNAT field value if set, zero value otherwise.
func (o *InlineResponse200103) GetNetworkRouterNAT() NetworkRouterNat {
	if o == nil || o.NetworkRouterNAT == nil {
		var ret NetworkRouterNat
		return ret
	}
	return *o.NetworkRouterNAT
}

// GetNetworkRouterNATOk returns a tuple with the NetworkRouterNAT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200103) GetNetworkRouterNATOk() (*NetworkRouterNat, bool) {
	if o == nil || o.NetworkRouterNAT == nil {
		return nil, false
	}
	return o.NetworkRouterNAT, true
}

// HasNetworkRouterNAT returns a boolean if a field has been set.
func (o *InlineResponse200103) HasNetworkRouterNAT() bool {
	if o != nil && o.NetworkRouterNAT != nil {
		return true
	}

	return false
}

// SetNetworkRouterNAT gets a reference to the given NetworkRouterNat and assigns it to the NetworkRouterNAT field.
func (o *InlineResponse200103) SetNetworkRouterNAT(v NetworkRouterNat) {
	o.NetworkRouterNAT = &v
}

func (o InlineResponse200103) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkRouterNAT != nil {
		toSerialize["networkRouterNAT"] = o.NetworkRouterNAT
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200103 struct {
	value *InlineResponse200103
	isSet bool
}

func (v NullableInlineResponse200103) Get() *InlineResponse200103 {
	return v.value
}

func (v *NullableInlineResponse200103) Set(val *InlineResponse200103) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200103) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200103) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200103(val *InlineResponse200103) *NullableInlineResponse200103 {
	return &NullableInlineResponse200103{value: val, isSet: true}
}

func (v NullableInlineResponse200103) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200103) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

