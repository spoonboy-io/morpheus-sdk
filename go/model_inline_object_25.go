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

// InlineObject25 struct for InlineObject25
type InlineObject25 struct {
	CatalogItemType *AnyOfcatalogItemTypeInstanceUpdatecatalogItemTypeBlueprintUpdatecatalogItemTypeWorkflowUpdate `json:"catalogItemType,omitempty"`
}

// NewInlineObject25 instantiates a new InlineObject25 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject25() *InlineObject25 {
	this := InlineObject25{}
	return &this
}

// NewInlineObject25WithDefaults instantiates a new InlineObject25 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject25WithDefaults() *InlineObject25 {
	this := InlineObject25{}
	return &this
}

// GetCatalogItemType returns the CatalogItemType field value if set, zero value otherwise.
func (o *InlineObject25) GetCatalogItemType() AnyOfcatalogItemTypeInstanceUpdatecatalogItemTypeBlueprintUpdatecatalogItemTypeWorkflowUpdate {
	if o == nil || o.CatalogItemType == nil {
		var ret AnyOfcatalogItemTypeInstanceUpdatecatalogItemTypeBlueprintUpdatecatalogItemTypeWorkflowUpdate
		return ret
	}
	return *o.CatalogItemType
}

// GetCatalogItemTypeOk returns a tuple with the CatalogItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject25) GetCatalogItemTypeOk() (*AnyOfcatalogItemTypeInstanceUpdatecatalogItemTypeBlueprintUpdatecatalogItemTypeWorkflowUpdate, bool) {
	if o == nil || o.CatalogItemType == nil {
		return nil, false
	}
	return o.CatalogItemType, true
}

// HasCatalogItemType returns a boolean if a field has been set.
func (o *InlineObject25) HasCatalogItemType() bool {
	if o != nil && o.CatalogItemType != nil {
		return true
	}

	return false
}

// SetCatalogItemType gets a reference to the given AnyOfcatalogItemTypeInstanceUpdatecatalogItemTypeBlueprintUpdatecatalogItemTypeWorkflowUpdate and assigns it to the CatalogItemType field.
func (o *InlineObject25) SetCatalogItemType(v AnyOfcatalogItemTypeInstanceUpdatecatalogItemTypeBlueprintUpdatecatalogItemTypeWorkflowUpdate) {
	o.CatalogItemType = &v
}

func (o InlineObject25) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CatalogItemType != nil {
		toSerialize["catalogItemType"] = o.CatalogItemType
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject25 struct {
	value *InlineObject25
	isSet bool
}

func (v NullableInlineObject25) Get() *InlineObject25 {
	return v.value
}

func (v *NullableInlineObject25) Set(val *InlineObject25) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject25) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject25) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject25(val *InlineObject25) *NullableInlineObject25 {
	return &NullableInlineObject25{value: val, isSet: true}
}

func (v NullableInlineObject25) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject25) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

