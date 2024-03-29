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
	"time"
)

// BillingServerApplicablePrices struct for BillingServerApplicablePrices
type BillingServerApplicablePrices struct {
	StartDate *time.Time `json:"startDate,omitempty"`
	EndDate *time.Time `json:"endDate,omitempty"`
	NumUnits *float32 `json:"numUnits,omitempty"`
	Cost *float32 `json:"cost,omitempty"`
	Price *float32 `json:"price,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Prices *[]BillingServerPrices `json:"prices,omitempty"`
}

// NewBillingServerApplicablePrices instantiates a new BillingServerApplicablePrices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingServerApplicablePrices() *BillingServerApplicablePrices {
	this := BillingServerApplicablePrices{}
	return &this
}

// NewBillingServerApplicablePricesWithDefaults instantiates a new BillingServerApplicablePrices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingServerApplicablePricesWithDefaults() *BillingServerApplicablePrices {
	this := BillingServerApplicablePrices{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *BillingServerApplicablePrices) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServerApplicablePrices) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *BillingServerApplicablePrices) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *BillingServerApplicablePrices) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *BillingServerApplicablePrices) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServerApplicablePrices) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *BillingServerApplicablePrices) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *BillingServerApplicablePrices) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetNumUnits returns the NumUnits field value if set, zero value otherwise.
func (o *BillingServerApplicablePrices) GetNumUnits() float32 {
	if o == nil || o.NumUnits == nil {
		var ret float32
		return ret
	}
	return *o.NumUnits
}

// GetNumUnitsOk returns a tuple with the NumUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServerApplicablePrices) GetNumUnitsOk() (*float32, bool) {
	if o == nil || o.NumUnits == nil {
		return nil, false
	}
	return o.NumUnits, true
}

// HasNumUnits returns a boolean if a field has been set.
func (o *BillingServerApplicablePrices) HasNumUnits() bool {
	if o != nil && o.NumUnits != nil {
		return true
	}

	return false
}

// SetNumUnits gets a reference to the given float32 and assigns it to the NumUnits field.
func (o *BillingServerApplicablePrices) SetNumUnits(v float32) {
	o.NumUnits = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *BillingServerApplicablePrices) GetCost() float32 {
	if o == nil || o.Cost == nil {
		var ret float32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServerApplicablePrices) GetCostOk() (*float32, bool) {
	if o == nil || o.Cost == nil {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *BillingServerApplicablePrices) HasCost() bool {
	if o != nil && o.Cost != nil {
		return true
	}

	return false
}

// SetCost gets a reference to the given float32 and assigns it to the Cost field.
func (o *BillingServerApplicablePrices) SetCost(v float32) {
	o.Cost = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *BillingServerApplicablePrices) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServerApplicablePrices) GetPriceOk() (*float32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *BillingServerApplicablePrices) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *BillingServerApplicablePrices) SetPrice(v float32) {
	o.Price = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *BillingServerApplicablePrices) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServerApplicablePrices) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *BillingServerApplicablePrices) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *BillingServerApplicablePrices) SetCurrency(v string) {
	o.Currency = &v
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *BillingServerApplicablePrices) GetPrices() []BillingServerPrices {
	if o == nil || o.Prices == nil {
		var ret []BillingServerPrices
		return ret
	}
	return *o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServerApplicablePrices) GetPricesOk() (*[]BillingServerPrices, bool) {
	if o == nil || o.Prices == nil {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *BillingServerApplicablePrices) HasPrices() bool {
	if o != nil && o.Prices != nil {
		return true
	}

	return false
}

// SetPrices gets a reference to the given []BillingServerPrices and assigns it to the Prices field.
func (o *BillingServerApplicablePrices) SetPrices(v []BillingServerPrices) {
	o.Prices = &v
}

func (o BillingServerApplicablePrices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	if o.EndDate != nil {
		toSerialize["endDate"] = o.EndDate
	}
	if o.NumUnits != nil {
		toSerialize["numUnits"] = o.NumUnits
	}
	if o.Cost != nil {
		toSerialize["cost"] = o.Cost
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.Prices != nil {
		toSerialize["prices"] = o.Prices
	}
	return json.Marshal(toSerialize)
}

type NullableBillingServerApplicablePrices struct {
	value *BillingServerApplicablePrices
	isSet bool
}

func (v NullableBillingServerApplicablePrices) Get() *BillingServerApplicablePrices {
	return v.value
}

func (v *NullableBillingServerApplicablePrices) Set(val *BillingServerApplicablePrices) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingServerApplicablePrices) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingServerApplicablePrices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingServerApplicablePrices(val *BillingServerApplicablePrices) *NullableBillingServerApplicablePrices {
	return &NullableBillingServerApplicablePrices{value: val, isSet: true}
}

func (v NullableBillingServerApplicablePrices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingServerApplicablePrices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


