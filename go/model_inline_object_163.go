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

// InlineObject163 struct for InlineObject163
type InlineObject163 struct {
	NetworkDomain *NetworkDomainCreate `json:"networkDomain,omitempty"`
}

// NewInlineObject163 instantiates a new InlineObject163 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject163() *InlineObject163 {
	this := InlineObject163{}
	return &this
}

// NewInlineObject163WithDefaults instantiates a new InlineObject163 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject163WithDefaults() *InlineObject163 {
	this := InlineObject163{}
	return &this
}

// GetNetworkDomain returns the NetworkDomain field value if set, zero value otherwise.
func (o *InlineObject163) GetNetworkDomain() NetworkDomainCreate {
	if o == nil || o.NetworkDomain == nil {
		var ret NetworkDomainCreate
		return ret
	}
	return *o.NetworkDomain
}

// GetNetworkDomainOk returns a tuple with the NetworkDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject163) GetNetworkDomainOk() (*NetworkDomainCreate, bool) {
	if o == nil || o.NetworkDomain == nil {
		return nil, false
	}
	return o.NetworkDomain, true
}

// HasNetworkDomain returns a boolean if a field has been set.
func (o *InlineObject163) HasNetworkDomain() bool {
	if o != nil && o.NetworkDomain != nil {
		return true
	}

	return false
}

// SetNetworkDomain gets a reference to the given NetworkDomainCreate and assigns it to the NetworkDomain field.
func (o *InlineObject163) SetNetworkDomain(v NetworkDomainCreate) {
	o.NetworkDomain = &v
}

func (o InlineObject163) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkDomain != nil {
		toSerialize["networkDomain"] = o.NetworkDomain
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject163 struct {
	value *InlineObject163
	isSet bool
}

func (v NullableInlineObject163) Get() *InlineObject163 {
	return v.value
}

func (v *NullableInlineObject163) Set(val *InlineObject163) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject163) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject163) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject163(val *InlineObject163) *NullableInlineObject163 {
	return &NullableInlineObject163{value: val, isSet: true}
}

func (v NullableInlineObject163) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject163) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

