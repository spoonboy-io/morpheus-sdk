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

// Model200Success struct for Model200Success
type Model200Success struct {
	Success *bool `json:"success,omitempty"`
}

// NewModel200Success instantiates a new Model200Success object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel200Success() *Model200Success {
	this := Model200Success{}
	return &this
}

// NewModel200SuccessWithDefaults instantiates a new Model200Success object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModel200SuccessWithDefaults() *Model200Success {
	this := Model200Success{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *Model200Success) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model200Success) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *Model200Success) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *Model200Success) SetSuccess(v bool) {
	o.Success = &v
}

func (o Model200Success) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableModel200Success struct {
	value *Model200Success
	isSet bool
}

func (v NullableModel200Success) Get() *Model200Success {
	return v.value
}

func (v *NullableModel200Success) Set(val *Model200Success) {
	v.value = val
	v.isSet = true
}

func (v NullableModel200Success) IsSet() bool {
	return v.isSet
}

func (v *NullableModel200Success) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel200Success(val *Model200Success) *NullableModel200Success {
	return &NullableModel200Success{value: val, isSet: true}
}

func (v NullableModel200Success) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel200Success) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


