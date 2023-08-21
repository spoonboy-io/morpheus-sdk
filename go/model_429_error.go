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

// Model429Error struct for Model429Error
type Model429Error struct {
	Msg *string `json:"msg,omitempty"`
}

// NewModel429Error instantiates a new Model429Error object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel429Error() *Model429Error {
	this := Model429Error{}
	return &this
}

// NewModel429ErrorWithDefaults instantiates a new Model429Error object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModel429ErrorWithDefaults() *Model429Error {
	this := Model429Error{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *Model429Error) GetMsg() string {
	if o == nil || o.Msg == nil {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model429Error) GetMsgOk() (*string, bool) {
	if o == nil || o.Msg == nil {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *Model429Error) HasMsg() bool {
	if o != nil && o.Msg != nil {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *Model429Error) SetMsg(v string) {
	o.Msg = &v
}

func (o Model429Error) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Msg != nil {
		toSerialize["msg"] = o.Msg
	}
	return json.Marshal(toSerialize)
}

type NullableModel429Error struct {
	value *Model429Error
	isSet bool
}

func (v NullableModel429Error) Get() *Model429Error {
	return v.value
}

func (v *NullableModel429Error) Set(val *Model429Error) {
	v.value = val
	v.isSet = true
}

func (v NullableModel429Error) IsSet() bool {
	return v.isSet
}

func (v *NullableModel429Error) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel429Error(val *Model429Error) *NullableModel429Error {
	return &NullableModel429Error{value: val, isSet: true}
}

func (v NullableModel429Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel429Error) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


