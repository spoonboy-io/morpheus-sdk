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

// InlineResponse200125 struct for InlineResponse200125
type InlineResponse200125 struct {
	ProvisionType *ProvisionType `json:"provisionType,omitempty"`
}

// NewInlineResponse200125 instantiates a new InlineResponse200125 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200125() *InlineResponse200125 {
	this := InlineResponse200125{}
	return &this
}

// NewInlineResponse200125WithDefaults instantiates a new InlineResponse200125 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200125WithDefaults() *InlineResponse200125 {
	this := InlineResponse200125{}
	return &this
}

// GetProvisionType returns the ProvisionType field value if set, zero value otherwise.
func (o *InlineResponse200125) GetProvisionType() ProvisionType {
	if o == nil || o.ProvisionType == nil {
		var ret ProvisionType
		return ret
	}
	return *o.ProvisionType
}

// GetProvisionTypeOk returns a tuple with the ProvisionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200125) GetProvisionTypeOk() (*ProvisionType, bool) {
	if o == nil || o.ProvisionType == nil {
		return nil, false
	}
	return o.ProvisionType, true
}

// HasProvisionType returns a boolean if a field has been set.
func (o *InlineResponse200125) HasProvisionType() bool {
	if o != nil && o.ProvisionType != nil {
		return true
	}

	return false
}

// SetProvisionType gets a reference to the given ProvisionType and assigns it to the ProvisionType field.
func (o *InlineResponse200125) SetProvisionType(v ProvisionType) {
	o.ProvisionType = &v
}

func (o InlineResponse200125) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProvisionType != nil {
		toSerialize["provisionType"] = o.ProvisionType
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200125 struct {
	value *InlineResponse200125
	isSet bool
}

func (v NullableInlineResponse200125) Get() *InlineResponse200125 {
	return v.value
}

func (v *NullableInlineResponse200125) Set(val *InlineResponse200125) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200125) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200125) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200125(val *InlineResponse200125) *NullableInlineResponse200125 {
	return &NullableInlineResponse200125{value: val, isSet: true}
}

func (v NullableInlineResponse200125) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200125) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

