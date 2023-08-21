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

// ZoneStats struct for ZoneStats
type ZoneStats struct {
	ServerCounts *ZoneStatsServerCounts `json:"serverCounts,omitempty"`
}

// NewZoneStats instantiates a new ZoneStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneStats() *ZoneStats {
	this := ZoneStats{}
	return &this
}

// NewZoneStatsWithDefaults instantiates a new ZoneStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneStatsWithDefaults() *ZoneStats {
	this := ZoneStats{}
	return &this
}

// GetServerCounts returns the ServerCounts field value if set, zero value otherwise.
func (o *ZoneStats) GetServerCounts() ZoneStatsServerCounts {
	if o == nil || o.ServerCounts == nil {
		var ret ZoneStatsServerCounts
		return ret
	}
	return *o.ServerCounts
}

// GetServerCountsOk returns a tuple with the ServerCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneStats) GetServerCountsOk() (*ZoneStatsServerCounts, bool) {
	if o == nil || o.ServerCounts == nil {
		return nil, false
	}
	return o.ServerCounts, true
}

// HasServerCounts returns a boolean if a field has been set.
func (o *ZoneStats) HasServerCounts() bool {
	if o != nil && o.ServerCounts != nil {
		return true
	}

	return false
}

// SetServerCounts gets a reference to the given ZoneStatsServerCounts and assigns it to the ServerCounts field.
func (o *ZoneStats) SetServerCounts(v ZoneStatsServerCounts) {
	o.ServerCounts = &v
}

func (o ZoneStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServerCounts != nil {
		toSerialize["serverCounts"] = o.ServerCounts
	}
	return json.Marshal(toSerialize)
}

type NullableZoneStats struct {
	value *ZoneStats
	isSet bool
}

func (v NullableZoneStats) Get() *ZoneStats {
	return v.value
}

func (v *NullableZoneStats) Set(val *ZoneStats) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneStats) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneStats(val *ZoneStats) *NullableZoneStats {
	return &NullableZoneStats{value: val, isSet: true}
}

func (v NullableZoneStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


