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

// InlineResponse20020 struct for InlineResponse20020
type InlineResponse20020 struct {
	ZoneType *ZoneType `json:"zoneType,omitempty"`
}

// NewInlineResponse20020 instantiates a new InlineResponse20020 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20020() *InlineResponse20020 {
	this := InlineResponse20020{}
	return &this
}

// NewInlineResponse20020WithDefaults instantiates a new InlineResponse20020 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20020WithDefaults() *InlineResponse20020 {
	this := InlineResponse20020{}
	return &this
}

// GetZoneType returns the ZoneType field value if set, zero value otherwise.
func (o *InlineResponse20020) GetZoneType() ZoneType {
	if o == nil || o.ZoneType == nil {
		var ret ZoneType
		return ret
	}
	return *o.ZoneType
}

// GetZoneTypeOk returns a tuple with the ZoneType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20020) GetZoneTypeOk() (*ZoneType, bool) {
	if o == nil || o.ZoneType == nil {
		return nil, false
	}
	return o.ZoneType, true
}

// HasZoneType returns a boolean if a field has been set.
func (o *InlineResponse20020) HasZoneType() bool {
	if o != nil && o.ZoneType != nil {
		return true
	}

	return false
}

// SetZoneType gets a reference to the given ZoneType and assigns it to the ZoneType field.
func (o *InlineResponse20020) SetZoneType(v ZoneType) {
	o.ZoneType = &v
}

func (o InlineResponse20020) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ZoneType != nil {
		toSerialize["zoneType"] = o.ZoneType
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20020 struct {
	value *InlineResponse20020
	isSet bool
}

func (v NullableInlineResponse20020) Get() *InlineResponse20020 {
	return v.value
}

func (v *NullableInlineResponse20020) Set(val *InlineResponse20020) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20020) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20020) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20020(val *InlineResponse20020) *NullableInlineResponse20020 {
	return &NullableInlineResponse20020{value: val, isSet: true}
}

func (v NullableInlineResponse20020) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20020) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

