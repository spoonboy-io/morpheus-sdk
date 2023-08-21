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

// InstanceInstancePrice struct for InstanceInstancePrice
type InstanceInstancePrice struct {
	Price *float32 `json:"price,omitempty"`
	Cost *float32 `json:"cost,omitempty"`
	Currency NullableString `json:"currency,omitempty"`
	Unit *string `json:"unit,omitempty"`
}

// NewInstanceInstancePrice instantiates a new InstanceInstancePrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceInstancePrice() *InstanceInstancePrice {
	this := InstanceInstancePrice{}
	return &this
}

// NewInstanceInstancePriceWithDefaults instantiates a new InstanceInstancePrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceInstancePriceWithDefaults() *InstanceInstancePrice {
	this := InstanceInstancePrice{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *InstanceInstancePrice) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceInstancePrice) GetPriceOk() (*float32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *InstanceInstancePrice) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *InstanceInstancePrice) SetPrice(v float32) {
	o.Price = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *InstanceInstancePrice) GetCost() float32 {
	if o == nil || o.Cost == nil {
		var ret float32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceInstancePrice) GetCostOk() (*float32, bool) {
	if o == nil || o.Cost == nil {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *InstanceInstancePrice) HasCost() bool {
	if o != nil && o.Cost != nil {
		return true
	}

	return false
}

// SetCost gets a reference to the given float32 and assigns it to the Cost field.
func (o *InstanceInstancePrice) SetCost(v float32) {
	o.Cost = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceInstancePrice) GetCurrency() string {
	if o == nil || o.Currency.Get() == nil {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceInstancePrice) GetCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *InstanceInstancePrice) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *InstanceInstancePrice) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *InstanceInstancePrice) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *InstanceInstancePrice) UnsetCurrency() {
	o.Currency.Unset()
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *InstanceInstancePrice) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceInstancePrice) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *InstanceInstancePrice) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *InstanceInstancePrice) SetUnit(v string) {
	o.Unit = &v
}

func (o InstanceInstancePrice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.Cost != nil {
		toSerialize["cost"] = o.Cost
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableInstanceInstancePrice struct {
	value *InstanceInstancePrice
	isSet bool
}

func (v NullableInstanceInstancePrice) Get() *InstanceInstancePrice {
	return v.value
}

func (v *NullableInstanceInstancePrice) Set(val *InstanceInstancePrice) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceInstancePrice) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceInstancePrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceInstancePrice(val *InstanceInstancePrice) *NullableInstanceInstancePrice {
	return &NullableInstanceInstancePrice{value: val, isSet: true}
}

func (v NullableInstanceInstancePrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceInstancePrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

