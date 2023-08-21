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

// ApiGroupsIdUpdateZonesGroup struct for ApiGroupsIdUpdateZonesGroup
type ApiGroupsIdUpdateZonesGroup struct {
	// An array of all the zones assigned to this group.
	Zones []map[string]interface{} `json:"zones"`
}

// NewApiGroupsIdUpdateZonesGroup instantiates a new ApiGroupsIdUpdateZonesGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiGroupsIdUpdateZonesGroup(zones []map[string]interface{}, ) *ApiGroupsIdUpdateZonesGroup {
	this := ApiGroupsIdUpdateZonesGroup{}
	this.Zones = zones
	return &this
}

// NewApiGroupsIdUpdateZonesGroupWithDefaults instantiates a new ApiGroupsIdUpdateZonesGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiGroupsIdUpdateZonesGroupWithDefaults() *ApiGroupsIdUpdateZonesGroup {
	this := ApiGroupsIdUpdateZonesGroup{}
	return &this
}

// GetZones returns the Zones field value
func (o *ApiGroupsIdUpdateZonesGroup) GetZones() []map[string]interface{} {
	if o == nil  {
		var ret []map[string]interface{}
		return ret
	}

	return o.Zones
}

// GetZonesOk returns a tuple with the Zones field value
// and a boolean to check if the value has been set.
func (o *ApiGroupsIdUpdateZonesGroup) GetZonesOk() (*[]map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Zones, true
}

// SetZones sets field value
func (o *ApiGroupsIdUpdateZonesGroup) SetZones(v []map[string]interface{}) {
	o.Zones = v
}

func (o ApiGroupsIdUpdateZonesGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["zones"] = o.Zones
	}
	return json.Marshal(toSerialize)
}

type NullableApiGroupsIdUpdateZonesGroup struct {
	value *ApiGroupsIdUpdateZonesGroup
	isSet bool
}

func (v NullableApiGroupsIdUpdateZonesGroup) Get() *ApiGroupsIdUpdateZonesGroup {
	return v.value
}

func (v *NullableApiGroupsIdUpdateZonesGroup) Set(val *ApiGroupsIdUpdateZonesGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableApiGroupsIdUpdateZonesGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableApiGroupsIdUpdateZonesGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiGroupsIdUpdateZonesGroup(val *ApiGroupsIdUpdateZonesGroup) *NullableApiGroupsIdUpdateZonesGroup {
	return &NullableApiGroupsIdUpdateZonesGroup{value: val, isSet: true}
}

func (v NullableApiGroupsIdUpdateZonesGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiGroupsIdUpdateZonesGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


