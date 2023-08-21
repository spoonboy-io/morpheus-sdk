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

// InlineResponse200119 struct for InlineResponse200119
type InlineResponse200119 struct {
	NetworkServices *[]NetworkService `json:"networkServices,omitempty"`
}

// NewInlineResponse200119 instantiates a new InlineResponse200119 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200119() *InlineResponse200119 {
	this := InlineResponse200119{}
	return &this
}

// NewInlineResponse200119WithDefaults instantiates a new InlineResponse200119 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200119WithDefaults() *InlineResponse200119 {
	this := InlineResponse200119{}
	return &this
}

// GetNetworkServices returns the NetworkServices field value if set, zero value otherwise.
func (o *InlineResponse200119) GetNetworkServices() []NetworkService {
	if o == nil || o.NetworkServices == nil {
		var ret []NetworkService
		return ret
	}
	return *o.NetworkServices
}

// GetNetworkServicesOk returns a tuple with the NetworkServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200119) GetNetworkServicesOk() (*[]NetworkService, bool) {
	if o == nil || o.NetworkServices == nil {
		return nil, false
	}
	return o.NetworkServices, true
}

// HasNetworkServices returns a boolean if a field has been set.
func (o *InlineResponse200119) HasNetworkServices() bool {
	if o != nil && o.NetworkServices != nil {
		return true
	}

	return false
}

// SetNetworkServices gets a reference to the given []NetworkService and assigns it to the NetworkServices field.
func (o *InlineResponse200119) SetNetworkServices(v []NetworkService) {
	o.NetworkServices = &v
}

func (o InlineResponse200119) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkServices != nil {
		toSerialize["networkServices"] = o.NetworkServices
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200119 struct {
	value *InlineResponse200119
	isSet bool
}

func (v NullableInlineResponse200119) Get() *InlineResponse200119 {
	return v.value
}

func (v *NullableInlineResponse200119) Set(val *InlineResponse200119) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200119) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200119) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200119(val *InlineResponse200119) *NullableInlineResponse200119 {
	return &NullableInlineResponse200119{value: val, isSet: true}
}

func (v NullableInlineResponse200119) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200119) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

