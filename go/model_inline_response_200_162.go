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

// InlineResponse200162 struct for InlineResponse200162
type InlineResponse200162 struct {
	VdiPool *VdiPool `json:"vdiPool,omitempty"`
}

// NewInlineResponse200162 instantiates a new InlineResponse200162 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200162() *InlineResponse200162 {
	this := InlineResponse200162{}
	return &this
}

// NewInlineResponse200162WithDefaults instantiates a new InlineResponse200162 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200162WithDefaults() *InlineResponse200162 {
	this := InlineResponse200162{}
	return &this
}

// GetVdiPool returns the VdiPool field value if set, zero value otherwise.
func (o *InlineResponse200162) GetVdiPool() VdiPool {
	if o == nil || o.VdiPool == nil {
		var ret VdiPool
		return ret
	}
	return *o.VdiPool
}

// GetVdiPoolOk returns a tuple with the VdiPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200162) GetVdiPoolOk() (*VdiPool, bool) {
	if o == nil || o.VdiPool == nil {
		return nil, false
	}
	return o.VdiPool, true
}

// HasVdiPool returns a boolean if a field has been set.
func (o *InlineResponse200162) HasVdiPool() bool {
	if o != nil && o.VdiPool != nil {
		return true
	}

	return false
}

// SetVdiPool gets a reference to the given VdiPool and assigns it to the VdiPool field.
func (o *InlineResponse200162) SetVdiPool(v VdiPool) {
	o.VdiPool = &v
}

func (o InlineResponse200162) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VdiPool != nil {
		toSerialize["vdiPool"] = o.VdiPool
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200162 struct {
	value *InlineResponse200162
	isSet bool
}

func (v NullableInlineResponse200162) Get() *InlineResponse200162 {
	return v.value
}

func (v *NullableInlineResponse200162) Set(val *InlineResponse200162) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200162) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200162) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200162(val *InlineResponse200162) *NullableInlineResponse200162 {
	return &NullableInlineResponse200162{value: val, isSet: true}
}

func (v NullableInlineResponse200162) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200162) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

