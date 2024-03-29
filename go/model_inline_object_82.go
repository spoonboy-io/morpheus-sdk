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

// InlineObject82 struct for InlineObject82
type InlineObject82 struct {
	// New Subdomain for account
	Subdomain string `json:"subdomain"`
}

// NewInlineObject82 instantiates a new InlineObject82 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject82(subdomain string, ) *InlineObject82 {
	this := InlineObject82{}
	this.Subdomain = subdomain
	return &this
}

// NewInlineObject82WithDefaults instantiates a new InlineObject82 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject82WithDefaults() *InlineObject82 {
	this := InlineObject82{}
	return &this
}

// GetSubdomain returns the Subdomain field value
func (o *InlineObject82) GetSubdomain() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value
// and a boolean to check if the value has been set.
func (o *InlineObject82) GetSubdomainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Subdomain, true
}

// SetSubdomain sets field value
func (o *InlineObject82) SetSubdomain(v string) {
	o.Subdomain = v
}

func (o InlineObject82) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subdomain"] = o.Subdomain
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject82 struct {
	value *InlineObject82
	isSet bool
}

func (v NullableInlineObject82) Get() *InlineObject82 {
	return v.value
}

func (v *NullableInlineObject82) Set(val *InlineObject82) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject82) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject82) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject82(val *InlineObject82) *NullableInlineObject82 {
	return &NullableInlineObject82{value: val, isSet: true}
}

func (v NullableInlineObject82) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject82) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


