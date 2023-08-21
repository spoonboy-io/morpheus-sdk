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

// InlineObject199 struct for InlineObject199
type InlineObject199 struct {
	PriceSet ApiPriceSetsIdPriceSet `json:"priceSet"`
}

// NewInlineObject199 instantiates a new InlineObject199 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject199(priceSet ApiPriceSetsIdPriceSet, ) *InlineObject199 {
	this := InlineObject199{}
	this.PriceSet = priceSet
	return &this
}

// NewInlineObject199WithDefaults instantiates a new InlineObject199 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject199WithDefaults() *InlineObject199 {
	this := InlineObject199{}
	return &this
}

// GetPriceSet returns the PriceSet field value
func (o *InlineObject199) GetPriceSet() ApiPriceSetsIdPriceSet {
	if o == nil  {
		var ret ApiPriceSetsIdPriceSet
		return ret
	}

	return o.PriceSet
}

// GetPriceSetOk returns a tuple with the PriceSet field value
// and a boolean to check if the value has been set.
func (o *InlineObject199) GetPriceSetOk() (*ApiPriceSetsIdPriceSet, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PriceSet, true
}

// SetPriceSet sets field value
func (o *InlineObject199) SetPriceSet(v ApiPriceSetsIdPriceSet) {
	o.PriceSet = v
}

func (o InlineObject199) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["priceSet"] = o.PriceSet
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject199 struct {
	value *InlineObject199
	isSet bool
}

func (v NullableInlineObject199) Get() *InlineObject199 {
	return v.value
}

func (v *NullableInlineObject199) Set(val *InlineObject199) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject199) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject199) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject199(val *InlineObject199) *NullableInlineObject199 {
	return &NullableInlineObject199{value: val, isSet: true}
}

func (v NullableInlineObject199) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject199) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


