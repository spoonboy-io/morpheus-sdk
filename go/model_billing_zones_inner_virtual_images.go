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

// checks if the BillingZonesInnerVirtualImages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingZonesInnerVirtualImages{}

// BillingZonesInnerVirtualImages struct for BillingZonesInnerVirtualImages
type BillingZonesInnerVirtualImages struct {
	Price *float32 `json:"price,omitempty"`
	Cost *float32 `json:"cost,omitempty"`
	VirtualImages []map[string]interface{} `json:"virtualImages,omitempty"`
	Count *int64 `json:"count,omitempty"`
}

// NewBillingZonesInnerVirtualImages instantiates a new BillingZonesInnerVirtualImages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingZonesInnerVirtualImages() *BillingZonesInnerVirtualImages {
	this := BillingZonesInnerVirtualImages{}
	return &this
}

// NewBillingZonesInnerVirtualImagesWithDefaults instantiates a new BillingZonesInnerVirtualImages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingZonesInnerVirtualImagesWithDefaults() *BillingZonesInnerVirtualImages {
	this := BillingZonesInnerVirtualImages{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *BillingZonesInnerVirtualImages) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingZonesInnerVirtualImages) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *BillingZonesInnerVirtualImages) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *BillingZonesInnerVirtualImages) SetPrice(v float32) {
	o.Price = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *BillingZonesInnerVirtualImages) GetCost() float32 {
	if o == nil || IsNil(o.Cost) {
		var ret float32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingZonesInnerVirtualImages) GetCostOk() (*float32, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *BillingZonesInnerVirtualImages) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given float32 and assigns it to the Cost field.
func (o *BillingZonesInnerVirtualImages) SetCost(v float32) {
	o.Cost = &v
}

// GetVirtualImages returns the VirtualImages field value if set, zero value otherwise.
func (o *BillingZonesInnerVirtualImages) GetVirtualImages() []map[string]interface{} {
	if o == nil || IsNil(o.VirtualImages) {
		var ret []map[string]interface{}
		return ret
	}
	return o.VirtualImages
}

// GetVirtualImagesOk returns a tuple with the VirtualImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingZonesInnerVirtualImages) GetVirtualImagesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.VirtualImages) {
		return nil, false
	}
	return o.VirtualImages, true
}

// HasVirtualImages returns a boolean if a field has been set.
func (o *BillingZonesInnerVirtualImages) HasVirtualImages() bool {
	if o != nil && !IsNil(o.VirtualImages) {
		return true
	}

	return false
}

// SetVirtualImages gets a reference to the given []map[string]interface{} and assigns it to the VirtualImages field.
func (o *BillingZonesInnerVirtualImages) SetVirtualImages(v []map[string]interface{}) {
	o.VirtualImages = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *BillingZonesInnerVirtualImages) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingZonesInnerVirtualImages) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *BillingZonesInnerVirtualImages) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *BillingZonesInnerVirtualImages) SetCount(v int64) {
	o.Count = &v
}

func (o BillingZonesInnerVirtualImages) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingZonesInnerVirtualImages) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.VirtualImages) {
		toSerialize["virtualImages"] = o.VirtualImages
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableBillingZonesInnerVirtualImages struct {
	value *BillingZonesInnerVirtualImages
	isSet bool
}

func (v NullableBillingZonesInnerVirtualImages) Get() *BillingZonesInnerVirtualImages {
	return v.value
}

func (v *NullableBillingZonesInnerVirtualImages) Set(val *BillingZonesInnerVirtualImages) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingZonesInnerVirtualImages) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingZonesInnerVirtualImages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingZonesInnerVirtualImages(val *BillingZonesInnerVirtualImages) *NullableBillingZonesInnerVirtualImages {
	return &NullableBillingZonesInnerVirtualImages{value: val, isSet: true}
}

func (v NullableBillingZonesInnerVirtualImages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingZonesInnerVirtualImages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

