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

// InlineObject201 struct for InlineObject201
type InlineObject201 struct {
	Price ApiPricesIdPrice `json:"price"`
}

// NewInlineObject201 instantiates a new InlineObject201 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject201(price ApiPricesIdPrice, ) *InlineObject201 {
	this := InlineObject201{}
	this.Price = price
	return &this
}

// NewInlineObject201WithDefaults instantiates a new InlineObject201 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject201WithDefaults() *InlineObject201 {
	this := InlineObject201{}
	return &this
}

// GetPrice returns the Price field value
func (o *InlineObject201) GetPrice() ApiPricesIdPrice {
	if o == nil  {
		var ret ApiPricesIdPrice
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *InlineObject201) GetPriceOk() (*ApiPricesIdPrice, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *InlineObject201) SetPrice(v ApiPricesIdPrice) {
	o.Price = v
}

func (o InlineObject201) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["price"] = o.Price
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject201 struct {
	value *InlineObject201
	isSet bool
}

func (v NullableInlineObject201) Get() *InlineObject201 {
	return v.value
}

func (v *NullableInlineObject201) Set(val *InlineObject201) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject201) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject201) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject201(val *InlineObject201) *NullableInlineObject201 {
	return &NullableInlineObject201{value: val, isSet: true}
}

func (v NullableInlineObject201) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject201) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


