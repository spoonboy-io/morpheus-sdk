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

// InlineObject178 The parameters for creating a network transport zone is type dependent. The following lists the common parameters. See get a specific type to list available options for the network server type.
type InlineObject178 struct {
	NetworkScope *NetworkTransportZoneCreate `json:"networkScope,omitempty"`
}

// NewInlineObject178 instantiates a new InlineObject178 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject178() *InlineObject178 {
	this := InlineObject178{}
	return &this
}

// NewInlineObject178WithDefaults instantiates a new InlineObject178 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject178WithDefaults() *InlineObject178 {
	this := InlineObject178{}
	return &this
}

// GetNetworkScope returns the NetworkScope field value if set, zero value otherwise.
func (o *InlineObject178) GetNetworkScope() NetworkTransportZoneCreate {
	if o == nil || o.NetworkScope == nil {
		var ret NetworkTransportZoneCreate
		return ret
	}
	return *o.NetworkScope
}

// GetNetworkScopeOk returns a tuple with the NetworkScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject178) GetNetworkScopeOk() (*NetworkTransportZoneCreate, bool) {
	if o == nil || o.NetworkScope == nil {
		return nil, false
	}
	return o.NetworkScope, true
}

// HasNetworkScope returns a boolean if a field has been set.
func (o *InlineObject178) HasNetworkScope() bool {
	if o != nil && o.NetworkScope != nil {
		return true
	}

	return false
}

// SetNetworkScope gets a reference to the given NetworkTransportZoneCreate and assigns it to the NetworkScope field.
func (o *InlineObject178) SetNetworkScope(v NetworkTransportZoneCreate) {
	o.NetworkScope = &v
}

func (o InlineObject178) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkScope != nil {
		toSerialize["networkScope"] = o.NetworkScope
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject178 struct {
	value *InlineObject178
	isSet bool
}

func (v NullableInlineObject178) Get() *InlineObject178 {
	return v.value
}

func (v *NullableInlineObject178) Set(val *InlineObject178) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject178) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject178) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject178(val *InlineObject178) *NullableInlineObject178 {
	return &NullableInlineObject178{value: val, isSet: true}
}

func (v NullableInlineObject178) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject178) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


