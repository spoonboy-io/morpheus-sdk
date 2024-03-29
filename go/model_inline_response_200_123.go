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

// InlineResponse200123 struct for InlineResponse200123
type InlineResponse200123 struct {
	PriceSet *PriceSet `json:"priceSet,omitempty"`
}

// NewInlineResponse200123 instantiates a new InlineResponse200123 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200123() *InlineResponse200123 {
	this := InlineResponse200123{}
	return &this
}

// NewInlineResponse200123WithDefaults instantiates a new InlineResponse200123 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200123WithDefaults() *InlineResponse200123 {
	this := InlineResponse200123{}
	return &this
}

// GetPriceSet returns the PriceSet field value if set, zero value otherwise.
func (o *InlineResponse200123) GetPriceSet() PriceSet {
	if o == nil || o.PriceSet == nil {
		var ret PriceSet
		return ret
	}
	return *o.PriceSet
}

// GetPriceSetOk returns a tuple with the PriceSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200123) GetPriceSetOk() (*PriceSet, bool) {
	if o == nil || o.PriceSet == nil {
		return nil, false
	}
	return o.PriceSet, true
}

// HasPriceSet returns a boolean if a field has been set.
func (o *InlineResponse200123) HasPriceSet() bool {
	if o != nil && o.PriceSet != nil {
		return true
	}

	return false
}

// SetPriceSet gets a reference to the given PriceSet and assigns it to the PriceSet field.
func (o *InlineResponse200123) SetPriceSet(v PriceSet) {
	o.PriceSet = &v
}

func (o InlineResponse200123) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PriceSet != nil {
		toSerialize["priceSet"] = o.PriceSet
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200123 struct {
	value *InlineResponse200123
	isSet bool
}

func (v NullableInlineResponse200123) Get() *InlineResponse200123 {
	return v.value
}

func (v *NullableInlineResponse200123) Set(val *InlineResponse200123) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200123) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200123) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200123(val *InlineResponse200123) *NullableInlineResponse200123 {
	return &NullableInlineResponse200123{value: val, isSet: true}
}

func (v NullableInlineResponse200123) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200123) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


