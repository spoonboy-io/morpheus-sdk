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

// GroupStats struct for GroupStats
type GroupStats struct {
	InstanceCounts *GroupStatsInstanceCounts `json:"instanceCounts,omitempty"`
	ServerCounts *ZoneStatsServerCounts `json:"serverCounts,omitempty"`
}

// NewGroupStats instantiates a new GroupStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupStats() *GroupStats {
	this := GroupStats{}
	return &this
}

// NewGroupStatsWithDefaults instantiates a new GroupStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupStatsWithDefaults() *GroupStats {
	this := GroupStats{}
	return &this
}

// GetInstanceCounts returns the InstanceCounts field value if set, zero value otherwise.
func (o *GroupStats) GetInstanceCounts() GroupStatsInstanceCounts {
	if o == nil || o.InstanceCounts == nil {
		var ret GroupStatsInstanceCounts
		return ret
	}
	return *o.InstanceCounts
}

// GetInstanceCountsOk returns a tuple with the InstanceCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupStats) GetInstanceCountsOk() (*GroupStatsInstanceCounts, bool) {
	if o == nil || o.InstanceCounts == nil {
		return nil, false
	}
	return o.InstanceCounts, true
}

// HasInstanceCounts returns a boolean if a field has been set.
func (o *GroupStats) HasInstanceCounts() bool {
	if o != nil && o.InstanceCounts != nil {
		return true
	}

	return false
}

// SetInstanceCounts gets a reference to the given GroupStatsInstanceCounts and assigns it to the InstanceCounts field.
func (o *GroupStats) SetInstanceCounts(v GroupStatsInstanceCounts) {
	o.InstanceCounts = &v
}

// GetServerCounts returns the ServerCounts field value if set, zero value otherwise.
func (o *GroupStats) GetServerCounts() ZoneStatsServerCounts {
	if o == nil || o.ServerCounts == nil {
		var ret ZoneStatsServerCounts
		return ret
	}
	return *o.ServerCounts
}

// GetServerCountsOk returns a tuple with the ServerCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupStats) GetServerCountsOk() (*ZoneStatsServerCounts, bool) {
	if o == nil || o.ServerCounts == nil {
		return nil, false
	}
	return o.ServerCounts, true
}

// HasServerCounts returns a boolean if a field has been set.
func (o *GroupStats) HasServerCounts() bool {
	if o != nil && o.ServerCounts != nil {
		return true
	}

	return false
}

// SetServerCounts gets a reference to the given ZoneStatsServerCounts and assigns it to the ServerCounts field.
func (o *GroupStats) SetServerCounts(v ZoneStatsServerCounts) {
	o.ServerCounts = &v
}

func (o GroupStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InstanceCounts != nil {
		toSerialize["instanceCounts"] = o.InstanceCounts
	}
	if o.ServerCounts != nil {
		toSerialize["serverCounts"] = o.ServerCounts
	}
	return json.Marshal(toSerialize)
}

type NullableGroupStats struct {
	value *GroupStats
	isSet bool
}

func (v NullableGroupStats) Get() *GroupStats {
	return v.value
}

func (v *NullableGroupStats) Set(val *GroupStats) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupStats) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupStats(val *GroupStats) *NullableGroupStats {
	return &NullableGroupStats{value: val, isSet: true}
}

func (v NullableGroupStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


