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

// InlineObject103 struct for InlineObject103
type InlineObject103 struct {
	Job OneOfobjectobjectobject `json:"job"`
}

// NewInlineObject103 instantiates a new InlineObject103 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject103(job OneOfobjectobjectobject, ) *InlineObject103 {
	this := InlineObject103{}
	this.Job = job
	return &this
}

// NewInlineObject103WithDefaults instantiates a new InlineObject103 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject103WithDefaults() *InlineObject103 {
	this := InlineObject103{}
	return &this
}

// GetJob returns the Job field value
func (o *InlineObject103) GetJob() OneOfobjectobjectobject {
	if o == nil  {
		var ret OneOfobjectobjectobject
		return ret
	}

	return o.Job
}

// GetJobOk returns a tuple with the Job field value
// and a boolean to check if the value has been set.
func (o *InlineObject103) GetJobOk() (*OneOfobjectobjectobject, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Job, true
}

// SetJob sets field value
func (o *InlineObject103) SetJob(v OneOfobjectobjectobject) {
	o.Job = v
}

func (o InlineObject103) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["job"] = o.Job
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject103 struct {
	value *InlineObject103
	isSet bool
}

func (v NullableInlineObject103) Get() *InlineObject103 {
	return v.value
}

func (v *NullableInlineObject103) Set(val *InlineObject103) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject103) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject103) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject103(val *InlineObject103) *NullableInlineObject103 {
	return &NullableInlineObject103{value: val, isSet: true}
}

func (v NullableInlineObject103) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject103) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

