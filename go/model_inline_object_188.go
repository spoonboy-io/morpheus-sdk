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

// InlineObject188 struct for InlineObject188
type InlineObject188 struct {
	Policy PolicyCloudCreate `json:"policy"`
}

// NewInlineObject188 instantiates a new InlineObject188 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject188(policy PolicyCloudCreate, ) *InlineObject188 {
	this := InlineObject188{}
	this.Policy = policy
	return &this
}

// NewInlineObject188WithDefaults instantiates a new InlineObject188 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject188WithDefaults() *InlineObject188 {
	this := InlineObject188{}
	return &this
}

// GetPolicy returns the Policy field value
func (o *InlineObject188) GetPolicy() PolicyCloudCreate {
	if o == nil  {
		var ret PolicyCloudCreate
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *InlineObject188) GetPolicyOk() (*PolicyCloudCreate, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Policy, true
}

// SetPolicy sets field value
func (o *InlineObject188) SetPolicy(v PolicyCloudCreate) {
	o.Policy = v
}

func (o InlineObject188) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["policy"] = o.Policy
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject188 struct {
	value *InlineObject188
	isSet bool
}

func (v NullableInlineObject188) Get() *InlineObject188 {
	return v.value
}

func (v *NullableInlineObject188) Set(val *InlineObject188) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject188) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject188) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject188(val *InlineObject188) *NullableInlineObject188 {
	return &NullableInlineObject188{value: val, isSet: true}
}

func (v NullableInlineObject188) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject188) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


