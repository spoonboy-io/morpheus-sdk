/*
Morpheus API

Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.

API version: 6.1.1
Contact: dev@morpheusdata.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner{}

// BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner struct for BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner
type BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner struct {
	Type *string `json:"type,omitempty"`
	PricePerUnit *float32 `json:"pricePerUnit,omitempty"`
	CostPerUnit *float32 `json:"costPerUnit,omitempty"`
	Cost *float32 `json:"cost,omitempty"`
	Price *float32 `json:"price,omitempty"`
	Quantity *float32 `json:"quantity,omitempty"`
}

// NewBillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner instantiates a new BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner() *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner {
	this := BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner{}
	return &this
}

// NewBillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInnerWithDefaults instantiates a new BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInnerWithDefaults() *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner {
	this := BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) SetType(v string) {
	o.Type = &v
}

// GetPricePerUnit returns the PricePerUnit field value if set, zero value otherwise.
func (o *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) GetPricePerUnit() float32 {
	if o == nil || IsNil(o.PricePerUnit) {
		var ret float32
		return ret
	}
	return *o.PricePerUnit
}

// GetPricePerUnitOk returns a tuple with the PricePerUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) GetPricePerUnitOk() (*float32, bool) {
	if o == nil || IsNil(o.PricePerUnit) {
		return nil, false
	}
	return o.PricePerUnit, true
}

// HasPricePerUnit returns a boolean if a field has been set.
func (o *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) HasPricePerUnit() bool {
	if o != nil && !IsNil(o.PricePerUnit) {
		return true
	}

	return false
}

// SetPricePerUnit gets a reference to the given float32 and assigns it to the PricePerUnit field.
func (o *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) SetPricePerUnit(v float32) {
	o.PricePerUnit = &v
}

// GetCostPerUnit returns the CostPerUnit field value if set, zero value otherwise.
func (o *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) GetCostPerUnit() float32 {
	if o == nil || IsNil(o.CostPerUnit) {
		var ret float32
		return ret
	}
	return *o.CostPerUnit
}

// GetCostPerUnitOk returns a tuple with the CostPerUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) GetCostPerUnitOk() (*float32, bool) {
	if o == nil || IsNil(o.CostPerUnit) {
		return nil, false
	}
	return o.CostPerUnit, true
}

// HasCostPerUnit returns a boolean if a field has been set.
func (o *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) HasCostPerUnit() bool {
	if o != nil && !IsNil(o.CostPerUnit) {
		return true
	}

	return false
}

// SetCostPerUnit gets a reference to the given float32 and assigns it to the CostPerUnit field.
func (o *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) SetCostPerUnit(v float32) {
	o.CostPerUnit = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) GetCost() float32 {
	if o == nil || IsNil(o.Cost) {
		var ret float32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) GetCostOk() (*float32, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given float32 and assigns it to the Cost field.
func (o *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) SetCost(v float32) {
	o.Cost = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) SetPrice(v float32) {
	o.Price = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) GetQuantity() float32 {
	if o == nil || IsNil(o.Quantity) {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) GetQuantityOk() (*float32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) SetQuantity(v float32) {
	o.Quantity = &v
}

func (o BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.PricePerUnit) {
		toSerialize["pricePerUnit"] = o.PricePerUnit
	}
	if !IsNil(o.CostPerUnit) {
		toSerialize["costPerUnit"] = o.CostPerUnit
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

type NullableBillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner struct {
	value *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner
	isSet bool
}

func (v NullableBillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) Get() *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner {
	return v.value
}

func (v *NullableBillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) Set(val *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner(val *BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) *NullableBillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner {
	return &NullableBillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner{value: val, isSet: true}
}

func (v NullableBillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


