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

// InlineResponse20096 struct for InlineResponse20096
type InlineResponse20096 struct {
	NetworkRouterBgpNeighbors *[]InlineResponse20096NetworkRouterBgpNeighbors `json:"networkRouterBgpNeighbors,omitempty"`
}

// NewInlineResponse20096 instantiates a new InlineResponse20096 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20096() *InlineResponse20096 {
	this := InlineResponse20096{}
	return &this
}

// NewInlineResponse20096WithDefaults instantiates a new InlineResponse20096 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20096WithDefaults() *InlineResponse20096 {
	this := InlineResponse20096{}
	return &this
}

// GetNetworkRouterBgpNeighbors returns the NetworkRouterBgpNeighbors field value if set, zero value otherwise.
func (o *InlineResponse20096) GetNetworkRouterBgpNeighbors() []InlineResponse20096NetworkRouterBgpNeighbors {
	if o == nil || o.NetworkRouterBgpNeighbors == nil {
		var ret []InlineResponse20096NetworkRouterBgpNeighbors
		return ret
	}
	return *o.NetworkRouterBgpNeighbors
}

// GetNetworkRouterBgpNeighborsOk returns a tuple with the NetworkRouterBgpNeighbors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20096) GetNetworkRouterBgpNeighborsOk() (*[]InlineResponse20096NetworkRouterBgpNeighbors, bool) {
	if o == nil || o.NetworkRouterBgpNeighbors == nil {
		return nil, false
	}
	return o.NetworkRouterBgpNeighbors, true
}

// HasNetworkRouterBgpNeighbors returns a boolean if a field has been set.
func (o *InlineResponse20096) HasNetworkRouterBgpNeighbors() bool {
	if o != nil && o.NetworkRouterBgpNeighbors != nil {
		return true
	}

	return false
}

// SetNetworkRouterBgpNeighbors gets a reference to the given []InlineResponse20096NetworkRouterBgpNeighbors and assigns it to the NetworkRouterBgpNeighbors field.
func (o *InlineResponse20096) SetNetworkRouterBgpNeighbors(v []InlineResponse20096NetworkRouterBgpNeighbors) {
	o.NetworkRouterBgpNeighbors = &v
}

func (o InlineResponse20096) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkRouterBgpNeighbors != nil {
		toSerialize["networkRouterBgpNeighbors"] = o.NetworkRouterBgpNeighbors
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20096 struct {
	value *InlineResponse20096
	isSet bool
}

func (v NullableInlineResponse20096) Get() *InlineResponse20096 {
	return v.value
}

func (v *NullableInlineResponse20096) Set(val *InlineResponse20096) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20096) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20096) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20096(val *InlineResponse20096) *NullableInlineResponse20096 {
	return &NullableInlineResponse20096{value: val, isSet: true}
}

func (v NullableInlineResponse20096) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20096) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


