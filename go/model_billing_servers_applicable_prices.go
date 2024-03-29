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

// BillingServersApplicablePrices struct for BillingServersApplicablePrices
type BillingServersApplicablePrices struct {
	StartDate *time.Time `json:"startDate,omitempty"`
	EndDate *time.Time `json:"endDate,omitempty"`
	NumUnits *float32 `json:"numUnits,omitempty"`
	Cost *float32 `json:"cost,omitempty"`
	Price *float32 `json:"price,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Prices *[]BillingServersPrices `json:"prices,omitempty"`
}

// NewBillingServersApplicablePrices instantiates a new BillingServersApplicablePrices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingServersApplicablePrices() *BillingServersApplicablePrices {
	this := BillingServersApplicablePrices{}
	return &this
}

// NewBillingServersApplicablePricesWithDefaults instantiates a new BillingServersApplicablePrices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingServersApplicablePricesWithDefaults() *BillingServersApplicablePrices {
	this := BillingServersApplicablePrices{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *BillingServersApplicablePrices) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServersApplicablePrices) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *BillingServersApplicablePrices) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *BillingServersApplicablePrices) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *BillingServersApplicablePrices) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServersApplicablePrices) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *BillingServersApplicablePrices) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *BillingServersApplicablePrices) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetNumUnits returns the NumUnits field value if set, zero value otherwise.
func (o *BillingServersApplicablePrices) GetNumUnits() float32 {
	if o == nil || o.NumUnits == nil {
		var ret float32
		return ret
	}
	return *o.NumUnits
}

// GetNumUnitsOk returns a tuple with the NumUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServersApplicablePrices) GetNumUnitsOk() (*float32, bool) {
	if o == nil || o.NumUnits == nil {
		return nil, false
	}
	return o.NumUnits, true
}

// HasNumUnits returns a boolean if a field has been set.
func (o *BillingServersApplicablePrices) HasNumUnits() bool {
	if o != nil && o.NumUnits != nil {
		return true
	}

	return false
}

// SetNumUnits gets a reference to the given float32 and assigns it to the NumUnits field.
func (o *BillingServersApplicablePrices) SetNumUnits(v float32) {
	o.NumUnits = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *BillingServersApplicablePrices) GetCost() float32 {
	if o == nil || o.Cost == nil {
		var ret float32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServersApplicablePrices) GetCostOk() (*float32, bool) {
	if o == nil || o.Cost == nil {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *BillingServersApplicablePrices) HasCost() bool {
	if o != nil && o.Cost != nil {
		return true
	}

	return false
}

// SetCost gets a reference to the given float32 and assigns it to the Cost field.
func (o *BillingServersApplicablePrices) SetCost(v float32) {
	o.Cost = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *BillingServersApplicablePrices) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServersApplicablePrices) GetPriceOk() (*float32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *BillingServersApplicablePrices) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *BillingServersApplicablePrices) SetPrice(v float32) {
	o.Price = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *BillingServersApplicablePrices) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServersApplicablePrices) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *BillingServersApplicablePrices) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *BillingServersApplicablePrices) SetCurrency(v string) {
	o.Currency = &v
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *BillingServersApplicablePrices) GetPrices() []BillingServersPrices {
	if o == nil || o.Prices == nil {
		var ret []BillingServersPrices
		return ret
	}
	return *o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServersApplicablePrices) GetPricesOk() (*[]BillingServersPrices, bool) {
	if o == nil || o.Prices == nil {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *BillingServersApplicablePrices) HasPrices() bool {
	if o != nil && o.Prices != nil {
		return true
	}

	return false
}

// SetPrices gets a reference to the given []BillingServersPrices and assigns it to the Prices field.
func (o *BillingServersApplicablePrices) SetPrices(v []BillingServersPrices) {
	o.Prices = &v
}

func (o BillingServersApplicablePrices) MarshalJSON() ([]byte, error) {
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

type NullableBillingServersApplicablePrices struct {
	value *BillingServersApplicablePrices
	isSet bool
}

func (v NullableBillingServersApplicablePrices) Get() *BillingServersApplicablePrices {
	return v.value
}

func (v *NullableBillingServersApplicablePrices) Set(val *BillingServersApplicablePrices) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingServersApplicablePrices) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingServersApplicablePrices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingServersApplicablePrices(val *BillingServersApplicablePrices) *NullableBillingServersApplicablePrices {
	return &NullableBillingServersApplicablePrices{value: val, isSet: true}
}

func (v NullableBillingServersApplicablePrices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingServersApplicablePrices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


