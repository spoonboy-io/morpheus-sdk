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

// ApiIntegrationsIdObjectsObject struct for ApiIntegrationsIdObjectsObject
type ApiIntegrationsIdObjectsObject struct {
	// Name to display
	Name *string `json:"name,omitempty"`
	// Integration Object Type Code
	Type string `json:"type"`
	// Catalog Item Type ID
	CatalogItemType int64 `json:"catalogItemType"`
}

// NewApiIntegrationsIdObjectsObject instantiates a new ApiIntegrationsIdObjectsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiIntegrationsIdObjectsObject(type_ string, catalogItemType int64, ) *ApiIntegrationsIdObjectsObject {
	this := ApiIntegrationsIdObjectsObject{}
	this.Type = type_
	this.CatalogItemType = catalogItemType
	return &this
}

// NewApiIntegrationsIdObjectsObjectWithDefaults instantiates a new ApiIntegrationsIdObjectsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiIntegrationsIdObjectsObjectWithDefaults() *ApiIntegrationsIdObjectsObject {
	this := ApiIntegrationsIdObjectsObject{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiIntegrationsIdObjectsObject) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiIntegrationsIdObjectsObject) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiIntegrationsIdObjectsObject) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiIntegrationsIdObjectsObject) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value
func (o *ApiIntegrationsIdObjectsObject) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApiIntegrationsIdObjectsObject) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApiIntegrationsIdObjectsObject) SetType(v string) {
	o.Type = v
}

// GetCatalogItemType returns the CatalogItemType field value
func (o *ApiIntegrationsIdObjectsObject) GetCatalogItemType() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.CatalogItemType
}

// GetCatalogItemTypeOk returns a tuple with the CatalogItemType field value
// and a boolean to check if the value has been set.
func (o *ApiIntegrationsIdObjectsObject) GetCatalogItemTypeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CatalogItemType, true
}

// SetCatalogItemType sets field value
func (o *ApiIntegrationsIdObjectsObject) SetCatalogItemType(v int64) {
	o.CatalogItemType = v
}

func (o ApiIntegrationsIdObjectsObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["catalogItemType"] = o.CatalogItemType
	}
	return json.Marshal(toSerialize)
}

type NullableApiIntegrationsIdObjectsObject struct {
	value *ApiIntegrationsIdObjectsObject
	isSet bool
}

func (v NullableApiIntegrationsIdObjectsObject) Get() *ApiIntegrationsIdObjectsObject {
	return v.value
}

func (v *NullableApiIntegrationsIdObjectsObject) Set(val *ApiIntegrationsIdObjectsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableApiIntegrationsIdObjectsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableApiIntegrationsIdObjectsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiIntegrationsIdObjectsObject(val *ApiIntegrationsIdObjectsObject) *NullableApiIntegrationsIdObjectsObject {
	return &NullableApiIntegrationsIdObjectsObject{value: val, isSet: true}
}

func (v NullableApiIntegrationsIdObjectsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiIntegrationsIdObjectsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


