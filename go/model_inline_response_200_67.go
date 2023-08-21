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

// InlineResponse20067 struct for InlineResponse20067
type InlineResponse20067 struct {
	Account *KeyPair `json:"account,omitempty"`
}

// NewInlineResponse20067 instantiates a new InlineResponse20067 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20067() *InlineResponse20067 {
	this := InlineResponse20067{}
	return &this
}

// NewInlineResponse20067WithDefaults instantiates a new InlineResponse20067 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20067WithDefaults() *InlineResponse20067 {
	this := InlineResponse20067{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *InlineResponse20067) GetAccount() KeyPair {
	if o == nil || o.Account == nil {
		var ret KeyPair
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20067) GetAccountOk() (*KeyPair, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *InlineResponse20067) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given KeyPair and assigns it to the Account field.
func (o *InlineResponse20067) SetAccount(v KeyPair) {
	o.Account = &v
}

func (o InlineResponse20067) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20067 struct {
	value *InlineResponse20067
	isSet bool
}

func (v NullableInlineResponse20067) Get() *InlineResponse20067 {
	return v.value
}

func (v *NullableInlineResponse20067) Set(val *InlineResponse20067) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20067) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20067) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20067(val *InlineResponse20067) *NullableInlineResponse20067 {
	return &NullableInlineResponse20067{value: val, isSet: true}
}

func (v NullableInlineResponse20067) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20067) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

