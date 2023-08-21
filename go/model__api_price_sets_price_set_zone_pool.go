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

// ApiPriceSetsPriceSetZonePool struct for ApiPriceSetsPriceSetZonePool
type ApiPriceSetsPriceSetZonePool struct {
	// Resource Pool ID
	Id *int64 `json:"id,omitempty"`
}

// NewApiPriceSetsPriceSetZonePool instantiates a new ApiPriceSetsPriceSetZonePool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPriceSetsPriceSetZonePool() *ApiPriceSetsPriceSetZonePool {
	this := ApiPriceSetsPriceSetZonePool{}
	return &this
}

// NewApiPriceSetsPriceSetZonePoolWithDefaults instantiates a new ApiPriceSetsPriceSetZonePool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPriceSetsPriceSetZonePoolWithDefaults() *ApiPriceSetsPriceSetZonePool {
	this := ApiPriceSetsPriceSetZonePool{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiPriceSetsPriceSetZonePool) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPriceSetsPriceSetZonePool) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiPriceSetsPriceSetZonePool) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ApiPriceSetsPriceSetZonePool) SetId(v int64) {
	o.Id = &v
}

func (o ApiPriceSetsPriceSetZonePool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableApiPriceSetsPriceSetZonePool struct {
	value *ApiPriceSetsPriceSetZonePool
	isSet bool
}

func (v NullableApiPriceSetsPriceSetZonePool) Get() *ApiPriceSetsPriceSetZonePool {
	return v.value
}

func (v *NullableApiPriceSetsPriceSetZonePool) Set(val *ApiPriceSetsPriceSetZonePool) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPriceSetsPriceSetZonePool) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPriceSetsPriceSetZonePool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPriceSetsPriceSetZonePool(val *ApiPriceSetsPriceSetZonePool) *NullableApiPriceSetsPriceSetZonePool {
	return &NullableApiPriceSetsPriceSetZonePool{value: val, isSet: true}
}

func (v NullableApiPriceSetsPriceSetZonePool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPriceSetsPriceSetZonePool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

