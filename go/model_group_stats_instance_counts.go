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

// GroupStatsInstanceCounts struct for GroupStatsInstanceCounts
type GroupStatsInstanceCounts struct {
	All *int64 `json:"all,omitempty"`
}

// NewGroupStatsInstanceCounts instantiates a new GroupStatsInstanceCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupStatsInstanceCounts() *GroupStatsInstanceCounts {
	this := GroupStatsInstanceCounts{}
	return &this
}

// NewGroupStatsInstanceCountsWithDefaults instantiates a new GroupStatsInstanceCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupStatsInstanceCountsWithDefaults() *GroupStatsInstanceCounts {
	this := GroupStatsInstanceCounts{}
	return &this
}

// GetAll returns the All field value if set, zero value otherwise.
func (o *GroupStatsInstanceCounts) GetAll() int64 {
	if o == nil || o.All == nil {
		var ret int64
		return ret
	}
	return *o.All
}

// GetAllOk returns a tuple with the All field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupStatsInstanceCounts) GetAllOk() (*int64, bool) {
	if o == nil || o.All == nil {
		return nil, false
	}
	return o.All, true
}

// HasAll returns a boolean if a field has been set.
func (o *GroupStatsInstanceCounts) HasAll() bool {
	if o != nil && o.All != nil {
		return true
	}

	return false
}

// SetAll gets a reference to the given int64 and assigns it to the All field.
func (o *GroupStatsInstanceCounts) SetAll(v int64) {
	o.All = &v
}

func (o GroupStatsInstanceCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.All != nil {
		toSerialize["all"] = o.All
	}
	return json.Marshal(toSerialize)
}

type NullableGroupStatsInstanceCounts struct {
	value *GroupStatsInstanceCounts
	isSet bool
}

func (v NullableGroupStatsInstanceCounts) Get() *GroupStatsInstanceCounts {
	return v.value
}

func (v *NullableGroupStatsInstanceCounts) Set(val *GroupStatsInstanceCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupStatsInstanceCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupStatsInstanceCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupStatsInstanceCounts(val *GroupStatsInstanceCounts) *NullableGroupStatsInstanceCounts {
	return &NullableGroupStatsInstanceCounts{value: val, isSet: true}
}

func (v NullableGroupStatsInstanceCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupStatsInstanceCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


