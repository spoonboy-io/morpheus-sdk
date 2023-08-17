/*
Morpheus API

Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.

API version: 6.1.1
Contact: dev@morpheusdata.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UpdateExecuteSchedulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateExecuteSchedulesRequest{}

// UpdateExecuteSchedulesRequest struct for UpdateExecuteSchedulesRequest
type UpdateExecuteSchedulesRequest struct {
	Schedule UpdateExecuteSchedulesRequestSchedule `json:"schedule"`
}

// NewUpdateExecuteSchedulesRequest instantiates a new UpdateExecuteSchedulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateExecuteSchedulesRequest(schedule UpdateExecuteSchedulesRequestSchedule) *UpdateExecuteSchedulesRequest {
	this := UpdateExecuteSchedulesRequest{}
	this.Schedule = schedule
	return &this
}

// NewUpdateExecuteSchedulesRequestWithDefaults instantiates a new UpdateExecuteSchedulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateExecuteSchedulesRequestWithDefaults() *UpdateExecuteSchedulesRequest {
	this := UpdateExecuteSchedulesRequest{}
	return &this
}

// GetSchedule returns the Schedule field value
func (o *UpdateExecuteSchedulesRequest) GetSchedule() UpdateExecuteSchedulesRequestSchedule {
	if o == nil {
		var ret UpdateExecuteSchedulesRequestSchedule
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *UpdateExecuteSchedulesRequest) GetScheduleOk() (*UpdateExecuteSchedulesRequestSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *UpdateExecuteSchedulesRequest) SetSchedule(v UpdateExecuteSchedulesRequestSchedule) {
	o.Schedule = v
}

func (o UpdateExecuteSchedulesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateExecuteSchedulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schedule"] = o.Schedule
	return toSerialize, nil
}

type NullableUpdateExecuteSchedulesRequest struct {
	value *UpdateExecuteSchedulesRequest
	isSet bool
}

func (v NullableUpdateExecuteSchedulesRequest) Get() *UpdateExecuteSchedulesRequest {
	return v.value
}

func (v *NullableUpdateExecuteSchedulesRequest) Set(val *UpdateExecuteSchedulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateExecuteSchedulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateExecuteSchedulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateExecuteSchedulesRequest(val *UpdateExecuteSchedulesRequest) *NullableUpdateExecuteSchedulesRequest {
	return &NullableUpdateExecuteSchedulesRequest{value: val, isSet: true}
}

func (v NullableUpdateExecuteSchedulesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateExecuteSchedulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


