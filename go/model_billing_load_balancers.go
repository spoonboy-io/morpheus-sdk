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

// BillingLoadBalancers struct for BillingLoadBalancers
type BillingLoadBalancers struct {
	Price *float32 `json:"price,omitempty"`
	Cost *float32 `json:"cost,omitempty"`
	LoadBalancers *[]map[string]interface{} `json:"loadBalancers,omitempty"`
	Count *int64 `json:"count,omitempty"`
}

// NewBillingLoadBalancers instantiates a new BillingLoadBalancers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingLoadBalancers() *BillingLoadBalancers {
	this := BillingLoadBalancers{}
	return &this
}

// NewBillingLoadBalancersWithDefaults instantiates a new BillingLoadBalancers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingLoadBalancersWithDefaults() *BillingLoadBalancers {
	this := BillingLoadBalancers{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *BillingLoadBalancers) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingLoadBalancers) GetPriceOk() (*float32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *BillingLoadBalancers) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *BillingLoadBalancers) SetPrice(v float32) {
	o.Price = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *BillingLoadBalancers) GetCost() float32 {
	if o == nil || o.Cost == nil {
		var ret float32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingLoadBalancers) GetCostOk() (*float32, bool) {
	if o == nil || o.Cost == nil {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *BillingLoadBalancers) HasCost() bool {
	if o != nil && o.Cost != nil {
		return true
	}

	return false
}

// SetCost gets a reference to the given float32 and assigns it to the Cost field.
func (o *BillingLoadBalancers) SetCost(v float32) {
	o.Cost = &v
}

// GetLoadBalancers returns the LoadBalancers field value if set, zero value otherwise.
func (o *BillingLoadBalancers) GetLoadBalancers() []map[string]interface{} {
	if o == nil || o.LoadBalancers == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.LoadBalancers
}

// GetLoadBalancersOk returns a tuple with the LoadBalancers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingLoadBalancers) GetLoadBalancersOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.LoadBalancers == nil {
		return nil, false
	}
	return o.LoadBalancers, true
}

// HasLoadBalancers returns a boolean if a field has been set.
func (o *BillingLoadBalancers) HasLoadBalancers() bool {
	if o != nil && o.LoadBalancers != nil {
		return true
	}

	return false
}

// SetLoadBalancers gets a reference to the given []map[string]interface{} and assigns it to the LoadBalancers field.
func (o *BillingLoadBalancers) SetLoadBalancers(v []map[string]interface{}) {
	o.LoadBalancers = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *BillingLoadBalancers) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingLoadBalancers) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *BillingLoadBalancers) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *BillingLoadBalancers) SetCount(v int64) {
	o.Count = &v
}

func (o BillingLoadBalancers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.Cost != nil {
		toSerialize["cost"] = o.Cost
	}
	if o.LoadBalancers != nil {
		toSerialize["loadBalancers"] = o.LoadBalancers
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableBillingLoadBalancers struct {
	value *BillingLoadBalancers
	isSet bool
}

func (v NullableBillingLoadBalancers) Get() *BillingLoadBalancers {
	return v.value
}

func (v *NullableBillingLoadBalancers) Set(val *BillingLoadBalancers) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingLoadBalancers) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingLoadBalancers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingLoadBalancers(val *BillingLoadBalancers) *NullableBillingLoadBalancers {
	return &NullableBillingLoadBalancers{value: val, isSet: true}
}

func (v NullableBillingLoadBalancers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingLoadBalancers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


