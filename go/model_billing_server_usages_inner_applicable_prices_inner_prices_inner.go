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

// checks if the BillingServerUsagesInnerApplicablePricesInnerPricesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingServerUsagesInnerApplicablePricesInnerPricesInner{}

// BillingServerUsagesInnerApplicablePricesInnerPricesInner struct for BillingServerUsagesInnerApplicablePricesInnerPricesInner
type BillingServerUsagesInnerApplicablePricesInnerPricesInner struct {
	Type *string `json:"type,omitempty"`
	PricePerUnit *float32 `json:"pricePerUnit,omitempty"`
	CostPerUnit *float32 `json:"costPerUnit,omitempty"`
	Cost *float32 `json:"cost,omitempty"`
	Price *float32 `json:"price,omitempty"`
	Quantity *int64 `json:"quantity,omitempty"`
	DatastoreId NullableString `json:"datastoreId,omitempty"`
	VolumeType *string `json:"volumeType,omitempty"`
	Datastore *string `json:"datastore,omitempty"`
}

// NewBillingServerUsagesInnerApplicablePricesInnerPricesInner instantiates a new BillingServerUsagesInnerApplicablePricesInnerPricesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingServerUsagesInnerApplicablePricesInnerPricesInner() *BillingServerUsagesInnerApplicablePricesInnerPricesInner {
	this := BillingServerUsagesInnerApplicablePricesInnerPricesInner{}
	return &this
}

// NewBillingServerUsagesInnerApplicablePricesInnerPricesInnerWithDefaults instantiates a new BillingServerUsagesInnerApplicablePricesInnerPricesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingServerUsagesInnerApplicablePricesInnerPricesInnerWithDefaults() *BillingServerUsagesInnerApplicablePricesInnerPricesInner {
	this := BillingServerUsagesInnerApplicablePricesInnerPricesInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) SetType(v string) {
	o.Type = &v
}

// GetPricePerUnit returns the PricePerUnit field value if set, zero value otherwise.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) GetPricePerUnit() float32 {
	if o == nil || IsNil(o.PricePerUnit) {
		var ret float32
		return ret
	}
	return *o.PricePerUnit
}

// GetPricePerUnitOk returns a tuple with the PricePerUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) GetPricePerUnitOk() (*float32, bool) {
	if o == nil || IsNil(o.PricePerUnit) {
		return nil, false
	}
	return o.PricePerUnit, true
}

// HasPricePerUnit returns a boolean if a field has been set.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) HasPricePerUnit() bool {
	if o != nil && !IsNil(o.PricePerUnit) {
		return true
	}

	return false
}

// SetPricePerUnit gets a reference to the given float32 and assigns it to the PricePerUnit field.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) SetPricePerUnit(v float32) {
	o.PricePerUnit = &v
}

// GetCostPerUnit returns the CostPerUnit field value if set, zero value otherwise.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) GetCostPerUnit() float32 {
	if o == nil || IsNil(o.CostPerUnit) {
		var ret float32
		return ret
	}
	return *o.CostPerUnit
}

// GetCostPerUnitOk returns a tuple with the CostPerUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) GetCostPerUnitOk() (*float32, bool) {
	if o == nil || IsNil(o.CostPerUnit) {
		return nil, false
	}
	return o.CostPerUnit, true
}

// HasCostPerUnit returns a boolean if a field has been set.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) HasCostPerUnit() bool {
	if o != nil && !IsNil(o.CostPerUnit) {
		return true
	}

	return false
}

// SetCostPerUnit gets a reference to the given float32 and assigns it to the CostPerUnit field.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) SetCostPerUnit(v float32) {
	o.CostPerUnit = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) GetCost() float32 {
	if o == nil || IsNil(o.Cost) {
		var ret float32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) GetCostOk() (*float32, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given float32 and assigns it to the Cost field.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) SetCost(v float32) {
	o.Cost = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) SetPrice(v float32) {
	o.Price = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetDatastoreId returns the DatastoreId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) GetDatastoreId() string {
	if o == nil || IsNil(o.DatastoreId.Get()) {
		var ret string
		return ret
	}
	return *o.DatastoreId.Get()
}

// GetDatastoreIdOk returns a tuple with the DatastoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) GetDatastoreIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatastoreId.Get(), o.DatastoreId.IsSet()
}

// HasDatastoreId returns a boolean if a field has been set.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) HasDatastoreId() bool {
	if o != nil && o.DatastoreId.IsSet() {
		return true
	}

	return false
}

// SetDatastoreId gets a reference to the given NullableString and assigns it to the DatastoreId field.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) SetDatastoreId(v string) {
	o.DatastoreId.Set(&v)
}
// SetDatastoreIdNil sets the value for DatastoreId to be an explicit nil
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) SetDatastoreIdNil() {
	o.DatastoreId.Set(nil)
}

// UnsetDatastoreId ensures that no value is present for DatastoreId, not even an explicit nil
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) UnsetDatastoreId() {
	o.DatastoreId.Unset()
}

// GetVolumeType returns the VolumeType field value if set, zero value otherwise.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) GetVolumeType() string {
	if o == nil || IsNil(o.VolumeType) {
		var ret string
		return ret
	}
	return *o.VolumeType
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) GetVolumeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeType) {
		return nil, false
	}
	return o.VolumeType, true
}

// HasVolumeType returns a boolean if a field has been set.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) HasVolumeType() bool {
	if o != nil && !IsNil(o.VolumeType) {
		return true
	}

	return false
}

// SetVolumeType gets a reference to the given string and assigns it to the VolumeType field.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) SetVolumeType(v string) {
	o.VolumeType = &v
}

// GetDatastore returns the Datastore field value if set, zero value otherwise.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) GetDatastore() string {
	if o == nil || IsNil(o.Datastore) {
		var ret string
		return ret
	}
	return *o.Datastore
}

// GetDatastoreOk returns a tuple with the Datastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) GetDatastoreOk() (*string, bool) {
	if o == nil || IsNil(o.Datastore) {
		return nil, false
	}
	return o.Datastore, true
}

// HasDatastore returns a boolean if a field has been set.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) HasDatastore() bool {
	if o != nil && !IsNil(o.Datastore) {
		return true
	}

	return false
}

// SetDatastore gets a reference to the given string and assigns it to the Datastore field.
func (o *BillingServerUsagesInnerApplicablePricesInnerPricesInner) SetDatastore(v string) {
	o.Datastore = &v
}

func (o BillingServerUsagesInnerApplicablePricesInnerPricesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingServerUsagesInnerApplicablePricesInnerPricesInner) ToMap() (map[string]interface{}, error) {
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
	if o.DatastoreId.IsSet() {
		toSerialize["datastoreId"] = o.DatastoreId.Get()
	}
	if !IsNil(o.VolumeType) {
		toSerialize["volumeType"] = o.VolumeType
	}
	if !IsNil(o.Datastore) {
		toSerialize["datastore"] = o.Datastore
	}
	return toSerialize, nil
}

type NullableBillingServerUsagesInnerApplicablePricesInnerPricesInner struct {
	value *BillingServerUsagesInnerApplicablePricesInnerPricesInner
	isSet bool
}

func (v NullableBillingServerUsagesInnerApplicablePricesInnerPricesInner) Get() *BillingServerUsagesInnerApplicablePricesInnerPricesInner {
	return v.value
}

func (v *NullableBillingServerUsagesInnerApplicablePricesInnerPricesInner) Set(val *BillingServerUsagesInnerApplicablePricesInnerPricesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingServerUsagesInnerApplicablePricesInnerPricesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingServerUsagesInnerApplicablePricesInnerPricesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingServerUsagesInnerApplicablePricesInnerPricesInner(val *BillingServerUsagesInnerApplicablePricesInnerPricesInner) *NullableBillingServerUsagesInnerApplicablePricesInnerPricesInner {
	return &NullableBillingServerUsagesInnerApplicablePricesInnerPricesInner{value: val, isSet: true}
}

func (v NullableBillingServerUsagesInnerApplicablePricesInnerPricesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingServerUsagesInnerApplicablePricesInnerPricesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

