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

// InlineResponse20049 struct for InlineResponse20049
type InlineResponse20049 struct {
	ProcessEvent *ProcessEvent `json:"processEvent,omitempty"`
}

// NewInlineResponse20049 instantiates a new InlineResponse20049 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20049() *InlineResponse20049 {
	this := InlineResponse20049{}
	return &this
}

// NewInlineResponse20049WithDefaults instantiates a new InlineResponse20049 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20049WithDefaults() *InlineResponse20049 {
	this := InlineResponse20049{}
	return &this
}

// GetProcessEvent returns the ProcessEvent field value if set, zero value otherwise.
func (o *InlineResponse20049) GetProcessEvent() ProcessEvent {
	if o == nil || o.ProcessEvent == nil {
		var ret ProcessEvent
		return ret
	}
	return *o.ProcessEvent
}

// GetProcessEventOk returns a tuple with the ProcessEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20049) GetProcessEventOk() (*ProcessEvent, bool) {
	if o == nil || o.ProcessEvent == nil {
		return nil, false
	}
	return o.ProcessEvent, true
}

// HasProcessEvent returns a boolean if a field has been set.
func (o *InlineResponse20049) HasProcessEvent() bool {
	if o != nil && o.ProcessEvent != nil {
		return true
	}

	return false
}

// SetProcessEvent gets a reference to the given ProcessEvent and assigns it to the ProcessEvent field.
func (o *InlineResponse20049) SetProcessEvent(v ProcessEvent) {
	o.ProcessEvent = &v
}

func (o InlineResponse20049) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProcessEvent != nil {
		toSerialize["processEvent"] = o.ProcessEvent
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20049 struct {
	value *InlineResponse20049
	isSet bool
}

func (v NullableInlineResponse20049) Get() *InlineResponse20049 {
	return v.value
}

func (v *NullableInlineResponse20049) Set(val *InlineResponse20049) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20049) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20049) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20049(val *InlineResponse20049) *NullableInlineResponse20049 {
	return &NullableInlineResponse20049{value: val, isSet: true}
}

func (v NullableInlineResponse20049) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20049) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


