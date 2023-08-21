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

// InlineResponse20059 struct for InlineResponse20059
type InlineResponse20059 struct {
	InstanceSchedule *InstanceSchedule `json:"instanceSchedule,omitempty"`
}

// NewInlineResponse20059 instantiates a new InlineResponse20059 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20059() *InlineResponse20059 {
	this := InlineResponse20059{}
	return &this
}

// NewInlineResponse20059WithDefaults instantiates a new InlineResponse20059 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20059WithDefaults() *InlineResponse20059 {
	this := InlineResponse20059{}
	return &this
}

// GetInstanceSchedule returns the InstanceSchedule field value if set, zero value otherwise.
func (o *InlineResponse20059) GetInstanceSchedule() InstanceSchedule {
	if o == nil || o.InstanceSchedule == nil {
		var ret InstanceSchedule
		return ret
	}
	return *o.InstanceSchedule
}

// GetInstanceScheduleOk returns a tuple with the InstanceSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20059) GetInstanceScheduleOk() (*InstanceSchedule, bool) {
	if o == nil || o.InstanceSchedule == nil {
		return nil, false
	}
	return o.InstanceSchedule, true
}

// HasInstanceSchedule returns a boolean if a field has been set.
func (o *InlineResponse20059) HasInstanceSchedule() bool {
	if o != nil && o.InstanceSchedule != nil {
		return true
	}

	return false
}

// SetInstanceSchedule gets a reference to the given InstanceSchedule and assigns it to the InstanceSchedule field.
func (o *InlineResponse20059) SetInstanceSchedule(v InstanceSchedule) {
	o.InstanceSchedule = &v
}

func (o InlineResponse20059) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InstanceSchedule != nil {
		toSerialize["instanceSchedule"] = o.InstanceSchedule
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20059 struct {
	value *InlineResponse20059
	isSet bool
}

func (v NullableInlineResponse20059) Get() *InlineResponse20059 {
	return v.value
}

func (v *NullableInlineResponse20059) Set(val *InlineResponse20059) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20059) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20059) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20059(val *InlineResponse20059) *NullableInlineResponse20059 {
	return &NullableInlineResponse20059{value: val, isSet: true}
}

func (v NullableInlineResponse20059) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20059) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

