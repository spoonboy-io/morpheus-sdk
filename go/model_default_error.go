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

// DefaultError struct for DefaultError
type DefaultError struct {
	Msg *string `json:"msg,omitempty"`
}

// NewDefaultError instantiates a new DefaultError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultError() *DefaultError {
	this := DefaultError{}
	return &this
}

// NewDefaultErrorWithDefaults instantiates a new DefaultError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultErrorWithDefaults() *DefaultError {
	this := DefaultError{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *DefaultError) GetMsg() string {
	if o == nil || o.Msg == nil {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultError) GetMsgOk() (*string, bool) {
	if o == nil || o.Msg == nil {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *DefaultError) HasMsg() bool {
	if o != nil && o.Msg != nil {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *DefaultError) SetMsg(v string) {
	o.Msg = &v
}

func (o DefaultError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Msg != nil {
		toSerialize["msg"] = o.Msg
	}
	return json.Marshal(toSerialize)
}

type NullableDefaultError struct {
	value *DefaultError
	isSet bool
}

func (v NullableDefaultError) Get() *DefaultError {
	return v.value
}

func (v *NullableDefaultError) Set(val *DefaultError) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultError) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultError(val *DefaultError) *NullableDefaultError {
	return &NullableDefaultError{value: val, isSet: true}
}

func (v NullableDefaultError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


