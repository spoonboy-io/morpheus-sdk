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

// InlineObject164 struct for InlineObject164
type InlineObject164 struct {
	NetworkDomain *NetworkDomainCreate `json:"networkDomain,omitempty"`
}

// NewInlineObject164 instantiates a new InlineObject164 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject164() *InlineObject164 {
	this := InlineObject164{}
	return &this
}

// NewInlineObject164WithDefaults instantiates a new InlineObject164 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject164WithDefaults() *InlineObject164 {
	this := InlineObject164{}
	return &this
}

// GetNetworkDomain returns the NetworkDomain field value if set, zero value otherwise.
func (o *InlineObject164) GetNetworkDomain() NetworkDomainCreate {
	if o == nil || o.NetworkDomain == nil {
		var ret NetworkDomainCreate
		return ret
	}
	return *o.NetworkDomain
}

// GetNetworkDomainOk returns a tuple with the NetworkDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject164) GetNetworkDomainOk() (*NetworkDomainCreate, bool) {
	if o == nil || o.NetworkDomain == nil {
		return nil, false
	}
	return o.NetworkDomain, true
}

// HasNetworkDomain returns a boolean if a field has been set.
func (o *InlineObject164) HasNetworkDomain() bool {
	if o != nil && o.NetworkDomain != nil {
		return true
	}

	return false
}

// SetNetworkDomain gets a reference to the given NetworkDomainCreate and assigns it to the NetworkDomain field.
func (o *InlineObject164) SetNetworkDomain(v NetworkDomainCreate) {
	o.NetworkDomain = &v
}

func (o InlineObject164) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkDomain != nil {
		toSerialize["networkDomain"] = o.NetworkDomain
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject164 struct {
	value *InlineObject164
	isSet bool
}

func (v NullableInlineObject164) Get() *InlineObject164 {
	return v.value
}

func (v *NullableInlineObject164) Set(val *InlineObject164) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject164) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject164) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject164(val *InlineObject164) *NullableInlineObject164 {
	return &NullableInlineObject164{value: val, isSet: true}
}

func (v NullableInlineObject164) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject164) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


