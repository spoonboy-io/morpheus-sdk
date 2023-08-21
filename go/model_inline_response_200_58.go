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

// InlineResponse20058 struct for InlineResponse20058
type InlineResponse20058 struct {
	Instance *InlineResponse20040AppDeployInstance `json:"instance,omitempty"`
	InstanceThreshold *InstanceThreshold `json:"instanceThreshold,omitempty"`
	InstanceSchedules *[]InstanceSchedule `json:"instanceSchedules,omitempty"`
}

// NewInlineResponse20058 instantiates a new InlineResponse20058 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20058() *InlineResponse20058 {
	this := InlineResponse20058{}
	return &this
}

// NewInlineResponse20058WithDefaults instantiates a new InlineResponse20058 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20058WithDefaults() *InlineResponse20058 {
	this := InlineResponse20058{}
	return &this
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *InlineResponse20058) GetInstance() InlineResponse20040AppDeployInstance {
	if o == nil || o.Instance == nil {
		var ret InlineResponse20040AppDeployInstance
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20058) GetInstanceOk() (*InlineResponse20040AppDeployInstance, bool) {
	if o == nil || o.Instance == nil {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *InlineResponse20058) HasInstance() bool {
	if o != nil && o.Instance != nil {
		return true
	}

	return false
}

// SetInstance gets a reference to the given InlineResponse20040AppDeployInstance and assigns it to the Instance field.
func (o *InlineResponse20058) SetInstance(v InlineResponse20040AppDeployInstance) {
	o.Instance = &v
}

// GetInstanceThreshold returns the InstanceThreshold field value if set, zero value otherwise.
func (o *InlineResponse20058) GetInstanceThreshold() InstanceThreshold {
	if o == nil || o.InstanceThreshold == nil {
		var ret InstanceThreshold
		return ret
	}
	return *o.InstanceThreshold
}

// GetInstanceThresholdOk returns a tuple with the InstanceThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20058) GetInstanceThresholdOk() (*InstanceThreshold, bool) {
	if o == nil || o.InstanceThreshold == nil {
		return nil, false
	}
	return o.InstanceThreshold, true
}

// HasInstanceThreshold returns a boolean if a field has been set.
func (o *InlineResponse20058) HasInstanceThreshold() bool {
	if o != nil && o.InstanceThreshold != nil {
		return true
	}

	return false
}

// SetInstanceThreshold gets a reference to the given InstanceThreshold and assigns it to the InstanceThreshold field.
func (o *InlineResponse20058) SetInstanceThreshold(v InstanceThreshold) {
	o.InstanceThreshold = &v
}

// GetInstanceSchedules returns the InstanceSchedules field value if set, zero value otherwise.
func (o *InlineResponse20058) GetInstanceSchedules() []InstanceSchedule {
	if o == nil || o.InstanceSchedules == nil {
		var ret []InstanceSchedule
		return ret
	}
	return *o.InstanceSchedules
}

// GetInstanceSchedulesOk returns a tuple with the InstanceSchedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20058) GetInstanceSchedulesOk() (*[]InstanceSchedule, bool) {
	if o == nil || o.InstanceSchedules == nil {
		return nil, false
	}
	return o.InstanceSchedules, true
}

// HasInstanceSchedules returns a boolean if a field has been set.
func (o *InlineResponse20058) HasInstanceSchedules() bool {
	if o != nil && o.InstanceSchedules != nil {
		return true
	}

	return false
}

// SetInstanceSchedules gets a reference to the given []InstanceSchedule and assigns it to the InstanceSchedules field.
func (o *InlineResponse20058) SetInstanceSchedules(v []InstanceSchedule) {
	o.InstanceSchedules = &v
}

func (o InlineResponse20058) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Instance != nil {
		toSerialize["instance"] = o.Instance
	}
	if o.InstanceThreshold != nil {
		toSerialize["instanceThreshold"] = o.InstanceThreshold
	}
	if o.InstanceSchedules != nil {
		toSerialize["instanceSchedules"] = o.InstanceSchedules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20058 struct {
	value *InlineResponse20058
	isSet bool
}

func (v NullableInlineResponse20058) Get() *InlineResponse20058 {
	return v.value
}

func (v *NullableInlineResponse20058) Set(val *InlineResponse20058) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20058) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20058) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20058(val *InlineResponse20058) *NullableInlineResponse20058 {
	return &NullableInlineResponse20058{value: val, isSet: true}
}

func (v NullableInlineResponse20058) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20058) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

