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

// checks if the BudgetStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetStats{}

// BudgetStats struct for BudgetStats
type BudgetStats struct {
	Currency *string `json:"currency,omitempty"`
	ConversionRate *int64 `json:"conversionRate,omitempty"`
	Intervals []BudgetStatsIntervalsInner `json:"intervals,omitempty"`
	Current *BudgetStatsCurrent `json:"current,omitempty"`
}

// NewBudgetStats instantiates a new BudgetStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetStats() *BudgetStats {
	this := BudgetStats{}
	return &this
}

// NewBudgetStatsWithDefaults instantiates a new BudgetStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetStatsWithDefaults() *BudgetStats {
	this := BudgetStats{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *BudgetStats) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetStats) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *BudgetStats) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *BudgetStats) SetCurrency(v string) {
	o.Currency = &v
}

// GetConversionRate returns the ConversionRate field value if set, zero value otherwise.
func (o *BudgetStats) GetConversionRate() int64 {
	if o == nil || IsNil(o.ConversionRate) {
		var ret int64
		return ret
	}
	return *o.ConversionRate
}

// GetConversionRateOk returns a tuple with the ConversionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetStats) GetConversionRateOk() (*int64, bool) {
	if o == nil || IsNil(o.ConversionRate) {
		return nil, false
	}
	return o.ConversionRate, true
}

// HasConversionRate returns a boolean if a field has been set.
func (o *BudgetStats) HasConversionRate() bool {
	if o != nil && !IsNil(o.ConversionRate) {
		return true
	}

	return false
}

// SetConversionRate gets a reference to the given int64 and assigns it to the ConversionRate field.
func (o *BudgetStats) SetConversionRate(v int64) {
	o.ConversionRate = &v
}

// GetIntervals returns the Intervals field value if set, zero value otherwise.
func (o *BudgetStats) GetIntervals() []BudgetStatsIntervalsInner {
	if o == nil || IsNil(o.Intervals) {
		var ret []BudgetStatsIntervalsInner
		return ret
	}
	return o.Intervals
}

// GetIntervalsOk returns a tuple with the Intervals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetStats) GetIntervalsOk() ([]BudgetStatsIntervalsInner, bool) {
	if o == nil || IsNil(o.Intervals) {
		return nil, false
	}
	return o.Intervals, true
}

// HasIntervals returns a boolean if a field has been set.
func (o *BudgetStats) HasIntervals() bool {
	if o != nil && !IsNil(o.Intervals) {
		return true
	}

	return false
}

// SetIntervals gets a reference to the given []BudgetStatsIntervalsInner and assigns it to the Intervals field.
func (o *BudgetStats) SetIntervals(v []BudgetStatsIntervalsInner) {
	o.Intervals = v
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *BudgetStats) GetCurrent() BudgetStatsCurrent {
	if o == nil || IsNil(o.Current) {
		var ret BudgetStatsCurrent
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetStats) GetCurrentOk() (*BudgetStatsCurrent, bool) {
	if o == nil || IsNil(o.Current) {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *BudgetStats) HasCurrent() bool {
	if o != nil && !IsNil(o.Current) {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given BudgetStatsCurrent and assigns it to the Current field.
func (o *BudgetStats) SetCurrent(v BudgetStatsCurrent) {
	o.Current = &v
}

func (o BudgetStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BudgetStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.ConversionRate) {
		toSerialize["conversionRate"] = o.ConversionRate
	}
	if !IsNil(o.Intervals) {
		toSerialize["intervals"] = o.Intervals
	}
	if !IsNil(o.Current) {
		toSerialize["current"] = o.Current
	}
	return toSerialize, nil
}

type NullableBudgetStats struct {
	value *BudgetStats
	isSet bool
}

func (v NullableBudgetStats) Get() *BudgetStats {
	return v.value
}

func (v *NullableBudgetStats) Set(val *BudgetStats) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetStats) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetStats(val *BudgetStats) *NullableBudgetStats {
	return &NullableBudgetStats{value: val, isSet: true}
}

func (v NullableBudgetStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBudgetStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


