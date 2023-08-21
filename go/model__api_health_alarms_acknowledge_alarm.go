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

// ApiHealthAlarmsAcknowledgeAlarm struct for ApiHealthAlarmsAcknowledgeAlarm
type ApiHealthAlarmsAcknowledgeAlarm struct {
	// Pass `true` to ackowledge an alarm, or pass `false` to unacknowledge it.
	Acknowledged bool `json:"acknowledged"`
	// Array of Alarm ID(s)to be updated.
	Ids *[]int64 `json:"ids,omitempty"`
	// Pass `true` to update all alarms instead of passing ids. This will update any active alarm that is not already acknowledged. 
	All *bool `json:"all,omitempty"`
}

// NewApiHealthAlarmsAcknowledgeAlarm instantiates a new ApiHealthAlarmsAcknowledgeAlarm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiHealthAlarmsAcknowledgeAlarm(acknowledged bool, ) *ApiHealthAlarmsAcknowledgeAlarm {
	this := ApiHealthAlarmsAcknowledgeAlarm{}
	this.Acknowledged = acknowledged
	var all bool = false
	this.All = &all
	return &this
}

// NewApiHealthAlarmsAcknowledgeAlarmWithDefaults instantiates a new ApiHealthAlarmsAcknowledgeAlarm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiHealthAlarmsAcknowledgeAlarmWithDefaults() *ApiHealthAlarmsAcknowledgeAlarm {
	this := ApiHealthAlarmsAcknowledgeAlarm{}
	var all bool = false
	this.All = &all
	return &this
}

// GetAcknowledged returns the Acknowledged field value
func (o *ApiHealthAlarmsAcknowledgeAlarm) GetAcknowledged() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Acknowledged
}

// GetAcknowledgedOk returns a tuple with the Acknowledged field value
// and a boolean to check if the value has been set.
func (o *ApiHealthAlarmsAcknowledgeAlarm) GetAcknowledgedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Acknowledged, true
}

// SetAcknowledged sets field value
func (o *ApiHealthAlarmsAcknowledgeAlarm) SetAcknowledged(v bool) {
	o.Acknowledged = v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *ApiHealthAlarmsAcknowledgeAlarm) GetIds() []int64 {
	if o == nil || o.Ids == nil {
		var ret []int64
		return ret
	}
	return *o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHealthAlarmsAcknowledgeAlarm) GetIdsOk() (*[]int64, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *ApiHealthAlarmsAcknowledgeAlarm) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given []int64 and assigns it to the Ids field.
func (o *ApiHealthAlarmsAcknowledgeAlarm) SetIds(v []int64) {
	o.Ids = &v
}

// GetAll returns the All field value if set, zero value otherwise.
func (o *ApiHealthAlarmsAcknowledgeAlarm) GetAll() bool {
	if o == nil || o.All == nil {
		var ret bool
		return ret
	}
	return *o.All
}

// GetAllOk returns a tuple with the All field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHealthAlarmsAcknowledgeAlarm) GetAllOk() (*bool, bool) {
	if o == nil || o.All == nil {
		return nil, false
	}
	return o.All, true
}

// HasAll returns a boolean if a field has been set.
func (o *ApiHealthAlarmsAcknowledgeAlarm) HasAll() bool {
	if o != nil && o.All != nil {
		return true
	}

	return false
}

// SetAll gets a reference to the given bool and assigns it to the All field.
func (o *ApiHealthAlarmsAcknowledgeAlarm) SetAll(v bool) {
	o.All = &v
}

func (o ApiHealthAlarmsAcknowledgeAlarm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["acknowledged"] = o.Acknowledged
	}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.All != nil {
		toSerialize["all"] = o.All
	}
	return json.Marshal(toSerialize)
}

type NullableApiHealthAlarmsAcknowledgeAlarm struct {
	value *ApiHealthAlarmsAcknowledgeAlarm
	isSet bool
}

func (v NullableApiHealthAlarmsAcknowledgeAlarm) Get() *ApiHealthAlarmsAcknowledgeAlarm {
	return v.value
}

func (v *NullableApiHealthAlarmsAcknowledgeAlarm) Set(val *ApiHealthAlarmsAcknowledgeAlarm) {
	v.value = val
	v.isSet = true
}

func (v NullableApiHealthAlarmsAcknowledgeAlarm) IsSet() bool {
	return v.isSet
}

func (v *NullableApiHealthAlarmsAcknowledgeAlarm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiHealthAlarmsAcknowledgeAlarm(val *ApiHealthAlarmsAcknowledgeAlarm) *NullableApiHealthAlarmsAcknowledgeAlarm {
	return &NullableApiHealthAlarmsAcknowledgeAlarm{value: val, isSet: true}
}

func (v NullableApiHealthAlarmsAcknowledgeAlarm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiHealthAlarmsAcknowledgeAlarm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


