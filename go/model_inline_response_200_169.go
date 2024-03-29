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

// InlineResponse200169 struct for InlineResponse200169
type InlineResponse200169 struct {
	Categories *[]map[string]interface{} `json:"categories,omitempty"`
}

// NewInlineResponse200169 instantiates a new InlineResponse200169 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200169() *InlineResponse200169 {
	this := InlineResponse200169{}
	return &this
}

// NewInlineResponse200169WithDefaults instantiates a new InlineResponse200169 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200169WithDefaults() *InlineResponse200169 {
	this := InlineResponse200169{}
	return &this
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *InlineResponse200169) GetCategories() []map[string]interface{} {
	if o == nil || o.Categories == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200169) GetCategoriesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Categories == nil {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *InlineResponse200169) HasCategories() bool {
	if o != nil && o.Categories != nil {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []map[string]interface{} and assigns it to the Categories field.
func (o *InlineResponse200169) SetCategories(v []map[string]interface{}) {
	o.Categories = &v
}

func (o InlineResponse200169) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Categories != nil {
		toSerialize["categories"] = o.Categories
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200169 struct {
	value *InlineResponse200169
	isSet bool
}

func (v NullableInlineResponse200169) Get() *InlineResponse200169 {
	return v.value
}

func (v *NullableInlineResponse200169) Set(val *InlineResponse200169) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200169) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200169) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200169(val *InlineResponse200169) *NullableInlineResponse200169 {
	return &NullableInlineResponse200169{value: val, isSet: true}
}

func (v NullableInlineResponse200169) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200169) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


