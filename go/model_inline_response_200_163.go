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

// InlineResponse200163 struct for InlineResponse200163
type InlineResponse200163 struct {
	VdiAllocation *VdiAllocation `json:"vdiAllocation,omitempty"`
}

// NewInlineResponse200163 instantiates a new InlineResponse200163 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200163() *InlineResponse200163 {
	this := InlineResponse200163{}
	return &this
}

// NewInlineResponse200163WithDefaults instantiates a new InlineResponse200163 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200163WithDefaults() *InlineResponse200163 {
	this := InlineResponse200163{}
	return &this
}

// GetVdiAllocation returns the VdiAllocation field value if set, zero value otherwise.
func (o *InlineResponse200163) GetVdiAllocation() VdiAllocation {
	if o == nil || o.VdiAllocation == nil {
		var ret VdiAllocation
		return ret
	}
	return *o.VdiAllocation
}

// GetVdiAllocationOk returns a tuple with the VdiAllocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200163) GetVdiAllocationOk() (*VdiAllocation, bool) {
	if o == nil || o.VdiAllocation == nil {
		return nil, false
	}
	return o.VdiAllocation, true
}

// HasVdiAllocation returns a boolean if a field has been set.
func (o *InlineResponse200163) HasVdiAllocation() bool {
	if o != nil && o.VdiAllocation != nil {
		return true
	}

	return false
}

// SetVdiAllocation gets a reference to the given VdiAllocation and assigns it to the VdiAllocation field.
func (o *InlineResponse200163) SetVdiAllocation(v VdiAllocation) {
	o.VdiAllocation = &v
}

func (o InlineResponse200163) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VdiAllocation != nil {
		toSerialize["vdiAllocation"] = o.VdiAllocation
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200163 struct {
	value *InlineResponse200163
	isSet bool
}

func (v NullableInlineResponse200163) Get() *InlineResponse200163 {
	return v.value
}

func (v *NullableInlineResponse200163) Set(val *InlineResponse200163) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200163) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200163) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200163(val *InlineResponse200163) *NullableInlineResponse200163 {
	return &NullableInlineResponse200163{value: val, isSet: true}
}

func (v NullableInlineResponse200163) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200163) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

