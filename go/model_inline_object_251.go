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

// InlineObject251 struct for InlineObject251
type InlineObject251 struct {
	Job ApiTasksIdExecuteJob `json:"job"`
}

// NewInlineObject251 instantiates a new InlineObject251 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject251(job ApiTasksIdExecuteJob, ) *InlineObject251 {
	this := InlineObject251{}
	this.Job = job
	return &this
}

// NewInlineObject251WithDefaults instantiates a new InlineObject251 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject251WithDefaults() *InlineObject251 {
	this := InlineObject251{}
	return &this
}

// GetJob returns the Job field value
func (o *InlineObject251) GetJob() ApiTasksIdExecuteJob {
	if o == nil  {
		var ret ApiTasksIdExecuteJob
		return ret
	}

	return o.Job
}

// GetJobOk returns a tuple with the Job field value
// and a boolean to check if the value has been set.
func (o *InlineObject251) GetJobOk() (*ApiTasksIdExecuteJob, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Job, true
}

// SetJob sets field value
func (o *InlineObject251) SetJob(v ApiTasksIdExecuteJob) {
	o.Job = v
}

func (o InlineObject251) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["job"] = o.Job
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject251 struct {
	value *InlineObject251
	isSet bool
}

func (v NullableInlineObject251) Get() *InlineObject251 {
	return v.value
}

func (v *NullableInlineObject251) Set(val *InlineObject251) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject251) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject251) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject251(val *InlineObject251) *NullableInlineObject251 {
	return &NullableInlineObject251{value: val, isSet: true}
}

func (v NullableInlineObject251) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject251) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


