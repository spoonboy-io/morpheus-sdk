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

// InlineObject23 struct for InlineObject23
type InlineObject23 struct {
	Budget ApiBudgetsBudget `json:"budget"`
}

// NewInlineObject23 instantiates a new InlineObject23 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject23(budget ApiBudgetsBudget, ) *InlineObject23 {
	this := InlineObject23{}
	this.Budget = budget
	return &this
}

// NewInlineObject23WithDefaults instantiates a new InlineObject23 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject23WithDefaults() *InlineObject23 {
	this := InlineObject23{}
	return &this
}

// GetBudget returns the Budget field value
func (o *InlineObject23) GetBudget() ApiBudgetsBudget {
	if o == nil  {
		var ret ApiBudgetsBudget
		return ret
	}

	return o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value
// and a boolean to check if the value has been set.
func (o *InlineObject23) GetBudgetOk() (*ApiBudgetsBudget, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Budget, true
}

// SetBudget sets field value
func (o *InlineObject23) SetBudget(v ApiBudgetsBudget) {
	o.Budget = v
}

func (o InlineObject23) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["budget"] = o.Budget
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject23 struct {
	value *InlineObject23
	isSet bool
}

func (v NullableInlineObject23) Get() *InlineObject23 {
	return v.value
}

func (v *NullableInlineObject23) Set(val *InlineObject23) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject23) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject23) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject23(val *InlineObject23) *NullableInlineObject23 {
	return &NullableInlineObject23{value: val, isSet: true}
}

func (v NullableInlineObject23) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject23) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


