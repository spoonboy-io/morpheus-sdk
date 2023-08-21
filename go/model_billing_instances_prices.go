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

// BillingInstancesPrices struct for BillingInstancesPrices
type BillingInstancesPrices struct {
	Type *string `json:"type,omitempty"`
	PricePerUnit *float32 `json:"pricePerUnit,omitempty"`
	CostPerUnit *float32 `json:"costPerUnit,omitempty"`
	Cost *float32 `json:"cost,omitempty"`
	Price *float32 `json:"price,omitempty"`
	Quantity *int64 `json:"quantity,omitempty"`
}

// NewBillingInstancesPrices instantiates a new BillingInstancesPrices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInstancesPrices() *BillingInstancesPrices {
	this := BillingInstancesPrices{}
	return &this
}

// NewBillingInstancesPricesWithDefaults instantiates a new BillingInstancesPrices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInstancesPricesWithDefaults() *BillingInstancesPrices {
	this := BillingInstancesPrices{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BillingInstancesPrices) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesPrices) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BillingInstancesPrices) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BillingInstancesPrices) SetType(v string) {
	o.Type = &v
}

// GetPricePerUnit returns the PricePerUnit field value if set, zero value otherwise.
func (o *BillingInstancesPrices) GetPricePerUnit() float32 {
	if o == nil || o.PricePerUnit == nil {
		var ret float32
		return ret
	}
	return *o.PricePerUnit
}

// GetPricePerUnitOk returns a tuple with the PricePerUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesPrices) GetPricePerUnitOk() (*float32, bool) {
	if o == nil || o.PricePerUnit == nil {
		return nil, false
	}
	return o.PricePerUnit, true
}

// HasPricePerUnit returns a boolean if a field has been set.
func (o *BillingInstancesPrices) HasPricePerUnit() bool {
	if o != nil && o.PricePerUnit != nil {
		return true
	}

	return false
}

// SetPricePerUnit gets a reference to the given float32 and assigns it to the PricePerUnit field.
func (o *BillingInstancesPrices) SetPricePerUnit(v float32) {
	o.PricePerUnit = &v
}

// GetCostPerUnit returns the CostPerUnit field value if set, zero value otherwise.
func (o *BillingInstancesPrices) GetCostPerUnit() float32 {
	if o == nil || o.CostPerUnit == nil {
		var ret float32
		return ret
	}
	return *o.CostPerUnit
}

// GetCostPerUnitOk returns a tuple with the CostPerUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesPrices) GetCostPerUnitOk() (*float32, bool) {
	if o == nil || o.CostPerUnit == nil {
		return nil, false
	}
	return o.CostPerUnit, true
}

// HasCostPerUnit returns a boolean if a field has been set.
func (o *BillingInstancesPrices) HasCostPerUnit() bool {
	if o != nil && o.CostPerUnit != nil {
		return true
	}

	return false
}

// SetCostPerUnit gets a reference to the given float32 and assigns it to the CostPerUnit field.
func (o *BillingInstancesPrices) SetCostPerUnit(v float32) {
	o.CostPerUnit = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *BillingInstancesPrices) GetCost() float32 {
	if o == nil || o.Cost == nil {
		var ret float32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesPrices) GetCostOk() (*float32, bool) {
	if o == nil || o.Cost == nil {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *BillingInstancesPrices) HasCost() bool {
	if o != nil && o.Cost != nil {
		return true
	}

	return false
}

// SetCost gets a reference to the given float32 and assigns it to the Cost field.
func (o *BillingInstancesPrices) SetCost(v float32) {
	o.Cost = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *BillingInstancesPrices) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesPrices) GetPriceOk() (*float32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *BillingInstancesPrices) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *BillingInstancesPrices) SetPrice(v float32) {
	o.Price = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *BillingInstancesPrices) GetQuantity() int64 {
	if o == nil || o.Quantity == nil {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesPrices) GetQuantityOk() (*int64, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *BillingInstancesPrices) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *BillingInstancesPrices) SetQuantity(v int64) {
	o.Quantity = &v
}

func (o BillingInstancesPrices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.PricePerUnit != nil {
		toSerialize["pricePerUnit"] = o.PricePerUnit
	}
	if o.CostPerUnit != nil {
		toSerialize["costPerUnit"] = o.CostPerUnit
	}
	if o.Cost != nil {
		toSerialize["cost"] = o.Cost
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	return json.Marshal(toSerialize)
}

type NullableBillingInstancesPrices struct {
	value *BillingInstancesPrices
	isSet bool
}

func (v NullableBillingInstancesPrices) Get() *BillingInstancesPrices {
	return v.value
}

func (v *NullableBillingInstancesPrices) Set(val *BillingInstancesPrices) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInstancesPrices) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInstancesPrices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInstancesPrices(val *BillingInstancesPrices) *NullableBillingInstancesPrices {
	return &NullableBillingInstancesPrices{value: val, isSet: true}
}

func (v NullableBillingInstancesPrices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInstancesPrices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

