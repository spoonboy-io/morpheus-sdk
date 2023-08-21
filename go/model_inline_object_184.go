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

// InlineObject184 struct for InlineObject184
type InlineObject184 struct {
	Policy PolicyCreate `json:"policy"`
}

// NewInlineObject184 instantiates a new InlineObject184 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject184(policy PolicyCreate, ) *InlineObject184 {
	this := InlineObject184{}
	this.Policy = policy
	return &this
}

// NewInlineObject184WithDefaults instantiates a new InlineObject184 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject184WithDefaults() *InlineObject184 {
	this := InlineObject184{}
	return &this
}

// GetPolicy returns the Policy field value
func (o *InlineObject184) GetPolicy() PolicyCreate {
	if o == nil  {
		var ret PolicyCreate
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *InlineObject184) GetPolicyOk() (*PolicyCreate, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Policy, true
}

// SetPolicy sets field value
func (o *InlineObject184) SetPolicy(v PolicyCreate) {
	o.Policy = v
}

func (o InlineObject184) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["policy"] = o.Policy
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject184 struct {
	value *InlineObject184
	isSet bool
}

func (v NullableInlineObject184) Get() *InlineObject184 {
	return v.value
}

func (v *NullableInlineObject184) Set(val *InlineObject184) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject184) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject184) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject184(val *InlineObject184) *NullableInlineObject184 {
	return &NullableInlineObject184{value: val, isSet: true}
}

func (v NullableInlineObject184) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject184) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


