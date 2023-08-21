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

// OptionTypeListCredential struct for OptionTypeListCredential
type OptionTypeListCredential struct {
	Type *string `json:"type,omitempty"`
}

// NewOptionTypeListCredential instantiates a new OptionTypeListCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionTypeListCredential() *OptionTypeListCredential {
	this := OptionTypeListCredential{}
	return &this
}

// NewOptionTypeListCredentialWithDefaults instantiates a new OptionTypeListCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionTypeListCredentialWithDefaults() *OptionTypeListCredential {
	this := OptionTypeListCredential{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OptionTypeListCredential) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionTypeListCredential) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OptionTypeListCredential) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OptionTypeListCredential) SetType(v string) {
	o.Type = &v
}

func (o OptionTypeListCredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableOptionTypeListCredential struct {
	value *OptionTypeListCredential
	isSet bool
}

func (v NullableOptionTypeListCredential) Get() *OptionTypeListCredential {
	return v.value
}

func (v *NullableOptionTypeListCredential) Set(val *OptionTypeListCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionTypeListCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionTypeListCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionTypeListCredential(val *OptionTypeListCredential) *NullableOptionTypeListCredential {
	return &NullableOptionTypeListCredential{value: val, isSet: true}
}

func (v NullableOptionTypeListCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionTypeListCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

