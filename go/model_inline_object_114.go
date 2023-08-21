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

// InlineObject114 struct for InlineObject114
type InlineObject114 struct {
	InstanceType *InstanceTypeUpdate `json:"instanceType,omitempty"`
}

// NewInlineObject114 instantiates a new InlineObject114 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject114() *InlineObject114 {
	this := InlineObject114{}
	return &this
}

// NewInlineObject114WithDefaults instantiates a new InlineObject114 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject114WithDefaults() *InlineObject114 {
	this := InlineObject114{}
	return &this
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *InlineObject114) GetInstanceType() InstanceTypeUpdate {
	if o == nil || o.InstanceType == nil {
		var ret InstanceTypeUpdate
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject114) GetInstanceTypeOk() (*InstanceTypeUpdate, bool) {
	if o == nil || o.InstanceType == nil {
		return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *InlineObject114) HasInstanceType() bool {
	if o != nil && o.InstanceType != nil {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given InstanceTypeUpdate and assigns it to the InstanceType field.
func (o *InlineObject114) SetInstanceType(v InstanceTypeUpdate) {
	o.InstanceType = &v
}

func (o InlineObject114) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InstanceType != nil {
		toSerialize["instanceType"] = o.InstanceType
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject114 struct {
	value *InlineObject114
	isSet bool
}

func (v NullableInlineObject114) Get() *InlineObject114 {
	return v.value
}

func (v *NullableInlineObject114) Set(val *InlineObject114) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject114) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject114) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject114(val *InlineObject114) *NullableInlineObject114 {
	return &NullableInlineObject114{value: val, isSet: true}
}

func (v NullableInlineObject114) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject114) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


