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

// InlineResponse200109 struct for InlineResponse200109
type InlineResponse200109 struct {
	NetworkDomain *NetworkDomain `json:"networkDomain,omitempty"`
}

// NewInlineResponse200109 instantiates a new InlineResponse200109 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200109() *InlineResponse200109 {
	this := InlineResponse200109{}
	return &this
}

// NewInlineResponse200109WithDefaults instantiates a new InlineResponse200109 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200109WithDefaults() *InlineResponse200109 {
	this := InlineResponse200109{}
	return &this
}

// GetNetworkDomain returns the NetworkDomain field value if set, zero value otherwise.
func (o *InlineResponse200109) GetNetworkDomain() NetworkDomain {
	if o == nil || o.NetworkDomain == nil {
		var ret NetworkDomain
		return ret
	}
	return *o.NetworkDomain
}

// GetNetworkDomainOk returns a tuple with the NetworkDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200109) GetNetworkDomainOk() (*NetworkDomain, bool) {
	if o == nil || o.NetworkDomain == nil {
		return nil, false
	}
	return o.NetworkDomain, true
}

// HasNetworkDomain returns a boolean if a field has been set.
func (o *InlineResponse200109) HasNetworkDomain() bool {
	if o != nil && o.NetworkDomain != nil {
		return true
	}

	return false
}

// SetNetworkDomain gets a reference to the given NetworkDomain and assigns it to the NetworkDomain field.
func (o *InlineResponse200109) SetNetworkDomain(v NetworkDomain) {
	o.NetworkDomain = &v
}

func (o InlineResponse200109) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkDomain != nil {
		toSerialize["networkDomain"] = o.NetworkDomain
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200109 struct {
	value *InlineResponse200109
	isSet bool
}

func (v NullableInlineResponse200109) Get() *InlineResponse200109 {
	return v.value
}

func (v *NullableInlineResponse200109) Set(val *InlineResponse200109) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200109) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200109) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200109(val *InlineResponse200109) *NullableInlineResponse200109 {
	return &NullableInlineResponse200109{value: val, isSet: true}
}

func (v NullableInlineResponse200109) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200109) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


