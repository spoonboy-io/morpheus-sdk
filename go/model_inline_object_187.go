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

// InlineObject187 struct for InlineObject187
type InlineObject187 struct {
	Policy PolicyGroupUpdate `json:"policy"`
}

// NewInlineObject187 instantiates a new InlineObject187 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject187(policy PolicyGroupUpdate, ) *InlineObject187 {
	this := InlineObject187{}
	this.Policy = policy
	return &this
}

// NewInlineObject187WithDefaults instantiates a new InlineObject187 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject187WithDefaults() *InlineObject187 {
	this := InlineObject187{}
	return &this
}

// GetPolicy returns the Policy field value
func (o *InlineObject187) GetPolicy() PolicyGroupUpdate {
	if o == nil  {
		var ret PolicyGroupUpdate
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *InlineObject187) GetPolicyOk() (*PolicyGroupUpdate, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Policy, true
}

// SetPolicy sets field value
func (o *InlineObject187) SetPolicy(v PolicyGroupUpdate) {
	o.Policy = v
}

func (o InlineObject187) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["policy"] = o.Policy
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject187 struct {
	value *InlineObject187
	isSet bool
}

func (v NullableInlineObject187) Get() *InlineObject187 {
	return v.value
}

func (v *NullableInlineObject187) Set(val *InlineObject187) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject187) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject187) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject187(val *InlineObject187) *NullableInlineObject187 {
	return &NullableInlineObject187{value: val, isSet: true}
}

func (v NullableInlineObject187) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject187) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

