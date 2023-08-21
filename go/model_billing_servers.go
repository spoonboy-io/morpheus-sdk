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

// BillingServers struct for BillingServers
type BillingServers struct {
	Price *float32 `json:"price,omitempty"`
	Cost *float32 `json:"cost,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
	EndDate *time.Time `json:"endDate,omitempty"`
	Servers *[]BillingServersServers `json:"servers,omitempty"`
}

// NewBillingServers instantiates a new BillingServers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingServers() *BillingServers {
	this := BillingServers{}
	return &this
}

// NewBillingServersWithDefaults instantiates a new BillingServers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingServersWithDefaults() *BillingServers {
	this := BillingServers{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *BillingServers) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServers) GetPriceOk() (*float32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *BillingServers) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *BillingServers) SetPrice(v float32) {
	o.Price = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *BillingServers) GetCost() float32 {
	if o == nil || o.Cost == nil {
		var ret float32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServers) GetCostOk() (*float32, bool) {
	if o == nil || o.Cost == nil {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *BillingServers) HasCost() bool {
	if o != nil && o.Cost != nil {
		return true
	}

	return false
}

// SetCost gets a reference to the given float32 and assigns it to the Cost field.
func (o *BillingServers) SetCost(v float32) {
	o.Cost = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *BillingServers) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServers) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *BillingServers) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *BillingServers) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *BillingServers) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServers) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *BillingServers) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *BillingServers) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *BillingServers) GetServers() []BillingServersServers {
	if o == nil || o.Servers == nil {
		var ret []BillingServersServers
		return ret
	}
	return *o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServers) GetServersOk() (*[]BillingServersServers, bool) {
	if o == nil || o.Servers == nil {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *BillingServers) HasServers() bool {
	if o != nil && o.Servers != nil {
		return true
	}

	return false
}

// SetServers gets a reference to the given []BillingServersServers and assigns it to the Servers field.
func (o *BillingServers) SetServers(v []BillingServersServers) {
	o.Servers = &v
}

func (o BillingServers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.Cost != nil {
		toSerialize["cost"] = o.Cost
	}
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	if o.EndDate != nil {
		toSerialize["endDate"] = o.EndDate
	}
	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}
	return json.Marshal(toSerialize)
}

type NullableBillingServers struct {
	value *BillingServers
	isSet bool
}

func (v NullableBillingServers) Get() *BillingServers {
	return v.value
}

func (v *NullableBillingServers) Set(val *BillingServers) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingServers) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingServers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingServers(val *BillingServers) *NullableBillingServers {
	return &NullableBillingServers{value: val, isSet: true}
}

func (v NullableBillingServers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingServers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


