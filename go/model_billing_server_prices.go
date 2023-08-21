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

// BillingServerPrices struct for BillingServerPrices
type BillingServerPrices struct {
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

// NewBillingServerPrices instantiates a new BillingServerPrices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingServerPrices() *BillingServerPrices {
	this := BillingServerPrices{}
	return &this
}

// NewBillingServerPricesWithDefaults instantiates a new BillingServerPrices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingServerPricesWithDefaults() *BillingServerPrices {
	this := BillingServerPrices{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BillingServerPrices) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServerPrices) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BillingServerPrices) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BillingServerPrices) SetType(v string) {
	o.Type = &v
}

// GetPricePerUnit returns the PricePerUnit field value if set, zero value otherwise.
func (o *BillingServerPrices) GetPricePerUnit() float32 {
	if o == nil || o.PricePerUnit == nil {
		var ret float32
		return ret
	}
	return *o.PricePerUnit
}

// GetPricePerUnitOk returns a tuple with the PricePerUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServerPrices) GetPricePerUnitOk() (*float32, bool) {
	if o == nil || o.PricePerUnit == nil {
		return nil, false
	}
	return o.PricePerUnit, true
}

// HasPricePerUnit returns a boolean if a field has been set.
func (o *BillingServerPrices) HasPricePerUnit() bool {
	if o != nil && o.PricePerUnit != nil {
		return true
	}

	return false
}

// SetPricePerUnit gets a reference to the given float32 and assigns it to the PricePerUnit field.
func (o *BillingServerPrices) SetPricePerUnit(v float32) {
	o.PricePerUnit = &v
}

// GetCostPerUnit returns the CostPerUnit field value if set, zero value otherwise.
func (o *BillingServerPrices) GetCostPerUnit() float32 {
	if o == nil || o.CostPerUnit == nil {
		var ret float32
		return ret
	}
	return *o.CostPerUnit
}

// GetCostPerUnitOk returns a tuple with the CostPerUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServerPrices) GetCostPerUnitOk() (*float32, bool) {
	if o == nil || o.CostPerUnit == nil {
		return nil, false
	}
	return o.CostPerUnit, true
}

// HasCostPerUnit returns a boolean if a field has been set.
func (o *BillingServerPrices) HasCostPerUnit() bool {
	if o != nil && o.CostPerUnit != nil {
		return true
	}

	return false
}

// SetCostPerUnit gets a reference to the given float32 and assigns it to the CostPerUnit field.
func (o *BillingServerPrices) SetCostPerUnit(v float32) {
	o.CostPerUnit = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *BillingServerPrices) GetCost() float32 {
	if o == nil || o.Cost == nil {
		var ret float32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServerPrices) GetCostOk() (*float32, bool) {
	if o == nil || o.Cost == nil {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *BillingServerPrices) HasCost() bool {
	if o != nil && o.Cost != nil {
		return true
	}

	return false
}

// SetCost gets a reference to the given float32 and assigns it to the Cost field.
func (o *BillingServerPrices) SetCost(v float32) {
	o.Cost = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *BillingServerPrices) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServerPrices) GetPriceOk() (*float32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *BillingServerPrices) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *BillingServerPrices) SetPrice(v float32) {
	o.Price = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *BillingServerPrices) GetQuantity() int64 {
	if o == nil || o.Quantity == nil {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServerPrices) GetQuantityOk() (*int64, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *BillingServerPrices) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *BillingServerPrices) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetDatastoreId returns the DatastoreId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingServerPrices) GetDatastoreId() string {
	if o == nil || o.DatastoreId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DatastoreId.Get()
}

// GetDatastoreIdOk returns a tuple with the DatastoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingServerPrices) GetDatastoreIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DatastoreId.Get(), o.DatastoreId.IsSet()
}

// HasDatastoreId returns a boolean if a field has been set.
func (o *BillingServerPrices) HasDatastoreId() bool {
	if o != nil && o.DatastoreId.IsSet() {
		return true
	}

	return false
}

// SetDatastoreId gets a reference to the given NullableString and assigns it to the DatastoreId field.
func (o *BillingServerPrices) SetDatastoreId(v string) {
	o.DatastoreId.Set(&v)
}
// SetDatastoreIdNil sets the value for DatastoreId to be an explicit nil
func (o *BillingServerPrices) SetDatastoreIdNil() {
	o.DatastoreId.Set(nil)
}

// UnsetDatastoreId ensures that no value is present for DatastoreId, not even an explicit nil
func (o *BillingServerPrices) UnsetDatastoreId() {
	o.DatastoreId.Unset()
}

// GetVolumeType returns the VolumeType field value if set, zero value otherwise.
func (o *BillingServerPrices) GetVolumeType() string {
	if o == nil || o.VolumeType == nil {
		var ret string
		return ret
	}
	return *o.VolumeType
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServerPrices) GetVolumeTypeOk() (*string, bool) {
	if o == nil || o.VolumeType == nil {
		return nil, false
	}
	return o.VolumeType, true
}

// HasVolumeType returns a boolean if a field has been set.
func (o *BillingServerPrices) HasVolumeType() bool {
	if o != nil && o.VolumeType != nil {
		return true
	}

	return false
}

// SetVolumeType gets a reference to the given string and assigns it to the VolumeType field.
func (o *BillingServerPrices) SetVolumeType(v string) {
	o.VolumeType = &v
}

// GetDatastore returns the Datastore field value if set, zero value otherwise.
func (o *BillingServerPrices) GetDatastore() string {
	if o == nil || o.Datastore == nil {
		var ret string
		return ret
	}
	return *o.Datastore
}

// GetDatastoreOk returns a tuple with the Datastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServerPrices) GetDatastoreOk() (*string, bool) {
	if o == nil || o.Datastore == nil {
		return nil, false
	}
	return o.Datastore, true
}

// HasDatastore returns a boolean if a field has been set.
func (o *BillingServerPrices) HasDatastore() bool {
	if o != nil && o.Datastore != nil {
		return true
	}

	return false
}

// SetDatastore gets a reference to the given string and assigns it to the Datastore field.
func (o *BillingServerPrices) SetDatastore(v string) {
	o.Datastore = &v
}

func (o BillingServerPrices) MarshalJSON() ([]byte, error) {
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
	if o.DatastoreId.IsSet() {
		toSerialize["datastoreId"] = o.DatastoreId.Get()
	}
	if o.VolumeType != nil {
		toSerialize["volumeType"] = o.VolumeType
	}
	if o.Datastore != nil {
		toSerialize["datastore"] = o.Datastore
	}
	return json.Marshal(toSerialize)
}

type NullableBillingServerPrices struct {
	value *BillingServerPrices
	isSet bool
}

func (v NullableBillingServerPrices) Get() *BillingServerPrices {
	return v.value
}

func (v *NullableBillingServerPrices) Set(val *BillingServerPrices) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingServerPrices) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingServerPrices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingServerPrices(val *BillingServerPrices) *NullableBillingServerPrices {
	return &NullableBillingServerPrices{value: val, isSet: true}
}

func (v NullableBillingServerPrices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingServerPrices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

