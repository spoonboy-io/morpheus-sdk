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

// ZoneTypeProvisionType struct for ZoneTypeProvisionType
type ZoneTypeProvisionType struct {
	Id *int64 `json:"id,omitempty"`
	Code *string `json:"code,omitempty"`
	Name *string `json:"name,omitempty"`
	HasNetworks *bool `json:"hasNetworks,omitempty"`
	HasZonePools *bool `json:"hasZonePools,omitempty"`
}

// NewZoneTypeProvisionType instantiates a new ZoneTypeProvisionType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneTypeProvisionType() *ZoneTypeProvisionType {
	this := ZoneTypeProvisionType{}
	return &this
}

// NewZoneTypeProvisionTypeWithDefaults instantiates a new ZoneTypeProvisionType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneTypeProvisionTypeWithDefaults() *ZoneTypeProvisionType {
	this := ZoneTypeProvisionType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ZoneTypeProvisionType) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneTypeProvisionType) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ZoneTypeProvisionType) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ZoneTypeProvisionType) SetId(v int64) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ZoneTypeProvisionType) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneTypeProvisionType) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ZoneTypeProvisionType) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ZoneTypeProvisionType) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ZoneTypeProvisionType) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneTypeProvisionType) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ZoneTypeProvisionType) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ZoneTypeProvisionType) SetName(v string) {
	o.Name = &v
}

// GetHasNetworks returns the HasNetworks field value if set, zero value otherwise.
func (o *ZoneTypeProvisionType) GetHasNetworks() bool {
	if o == nil || o.HasNetworks == nil {
		var ret bool
		return ret
	}
	return *o.HasNetworks
}

// GetHasNetworksOk returns a tuple with the HasNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneTypeProvisionType) GetHasNetworksOk() (*bool, bool) {
	if o == nil || o.HasNetworks == nil {
		return nil, false
	}
	return o.HasNetworks, true
}

// HasHasNetworks returns a boolean if a field has been set.
func (o *ZoneTypeProvisionType) HasHasNetworks() bool {
	if o != nil && o.HasNetworks != nil {
		return true
	}

	return false
}

// SetHasNetworks gets a reference to the given bool and assigns it to the HasNetworks field.
func (o *ZoneTypeProvisionType) SetHasNetworks(v bool) {
	o.HasNetworks = &v
}

// GetHasZonePools returns the HasZonePools field value if set, zero value otherwise.
func (o *ZoneTypeProvisionType) GetHasZonePools() bool {
	if o == nil || o.HasZonePools == nil {
		var ret bool
		return ret
	}
	return *o.HasZonePools
}

// GetHasZonePoolsOk returns a tuple with the HasZonePools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneTypeProvisionType) GetHasZonePoolsOk() (*bool, bool) {
	if o == nil || o.HasZonePools == nil {
		return nil, false
	}
	return o.HasZonePools, true
}

// HasHasZonePools returns a boolean if a field has been set.
func (o *ZoneTypeProvisionType) HasHasZonePools() bool {
	if o != nil && o.HasZonePools != nil {
		return true
	}

	return false
}

// SetHasZonePools gets a reference to the given bool and assigns it to the HasZonePools field.
func (o *ZoneTypeProvisionType) SetHasZonePools(v bool) {
	o.HasZonePools = &v
}

func (o ZoneTypeProvisionType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.HasNetworks != nil {
		toSerialize["hasNetworks"] = o.HasNetworks
	}
	if o.HasZonePools != nil {
		toSerialize["hasZonePools"] = o.HasZonePools
	}
	return json.Marshal(toSerialize)
}

type NullableZoneTypeProvisionType struct {
	value *ZoneTypeProvisionType
	isSet bool
}

func (v NullableZoneTypeProvisionType) Get() *ZoneTypeProvisionType {
	return v.value
}

func (v *NullableZoneTypeProvisionType) Set(val *ZoneTypeProvisionType) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneTypeProvisionType) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneTypeProvisionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneTypeProvisionType(val *ZoneTypeProvisionType) *NullableZoneTypeProvisionType {
	return &NullableZoneTypeProvisionType{value: val, isSet: true}
}

func (v NullableZoneTypeProvisionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneTypeProvisionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


