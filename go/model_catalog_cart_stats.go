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

// CatalogCartStats struct for CatalogCartStats
type CatalogCartStats struct {
	Price *float32 `json:"price,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Unit *string `json:"unit,omitempty"`
}

// NewCatalogCartStats instantiates a new CatalogCartStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogCartStats() *CatalogCartStats {
	this := CatalogCartStats{}
	return &this
}

// NewCatalogCartStatsWithDefaults instantiates a new CatalogCartStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogCartStatsWithDefaults() *CatalogCartStats {
	this := CatalogCartStats{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *CatalogCartStats) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogCartStats) GetPriceOk() (*float32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *CatalogCartStats) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *CatalogCartStats) SetPrice(v float32) {
	o.Price = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *CatalogCartStats) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogCartStats) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *CatalogCartStats) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *CatalogCartStats) SetCurrency(v string) {
	o.Currency = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *CatalogCartStats) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogCartStats) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *CatalogCartStats) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *CatalogCartStats) SetUnit(v string) {
	o.Unit = &v
}

func (o CatalogCartStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableCatalogCartStats struct {
	value *CatalogCartStats
	isSet bool
}

func (v NullableCatalogCartStats) Get() *CatalogCartStats {
	return v.value
}

func (v *NullableCatalogCartStats) Set(val *CatalogCartStats) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogCartStats) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogCartStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogCartStats(val *CatalogCartStats) *NullableCatalogCartStats {
	return &NullableCatalogCartStats{value: val, isSet: true}
}

func (v NullableCatalogCartStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogCartStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


