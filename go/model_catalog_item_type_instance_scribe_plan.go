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

// CatalogItemTypeInstanceScribePlan struct for CatalogItemTypeInstanceScribePlan
type CatalogItemTypeInstanceScribePlan struct {
	Id OneOfstringlong `json:"id"`
}

// NewCatalogItemTypeInstanceScribePlan instantiates a new CatalogItemTypeInstanceScribePlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogItemTypeInstanceScribePlan(id OneOfstringlong, ) *CatalogItemTypeInstanceScribePlan {
	this := CatalogItemTypeInstanceScribePlan{}
	this.Id = id
	return &this
}

// NewCatalogItemTypeInstanceScribePlanWithDefaults instantiates a new CatalogItemTypeInstanceScribePlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogItemTypeInstanceScribePlanWithDefaults() *CatalogItemTypeInstanceScribePlan {
	this := CatalogItemTypeInstanceScribePlan{}
	return &this
}

// GetId returns the Id field value
func (o *CatalogItemTypeInstanceScribePlan) GetId() OneOfstringlong {
	if o == nil  {
		var ret OneOfstringlong
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CatalogItemTypeInstanceScribePlan) GetIdOk() (*OneOfstringlong, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CatalogItemTypeInstanceScribePlan) SetId(v OneOfstringlong) {
	o.Id = v
}

func (o CatalogItemTypeInstanceScribePlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableCatalogItemTypeInstanceScribePlan struct {
	value *CatalogItemTypeInstanceScribePlan
	isSet bool
}

func (v NullableCatalogItemTypeInstanceScribePlan) Get() *CatalogItemTypeInstanceScribePlan {
	return v.value
}

func (v *NullableCatalogItemTypeInstanceScribePlan) Set(val *CatalogItemTypeInstanceScribePlan) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogItemTypeInstanceScribePlan) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogItemTypeInstanceScribePlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogItemTypeInstanceScribePlan(val *CatalogItemTypeInstanceScribePlan) *NullableCatalogItemTypeInstanceScribePlan {
	return &NullableCatalogItemTypeInstanceScribePlan{value: val, isSet: true}
}

func (v NullableCatalogItemTypeInstanceScribePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogItemTypeInstanceScribePlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


