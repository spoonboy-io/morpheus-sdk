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

// HealthRabbitQueues struct for HealthRabbitQueues
type HealthRabbitQueues struct {
	Name *string `json:"name,omitempty"`
	Count *int64 `json:"count,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewHealthRabbitQueues instantiates a new HealthRabbitQueues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthRabbitQueues() *HealthRabbitQueues {
	this := HealthRabbitQueues{}
	return &this
}

// NewHealthRabbitQueuesWithDefaults instantiates a new HealthRabbitQueues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthRabbitQueuesWithDefaults() *HealthRabbitQueues {
	this := HealthRabbitQueues{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HealthRabbitQueues) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthRabbitQueues) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HealthRabbitQueues) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HealthRabbitQueues) SetName(v string) {
	o.Name = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *HealthRabbitQueues) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthRabbitQueues) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *HealthRabbitQueues) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *HealthRabbitQueues) SetCount(v int64) {
	o.Count = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HealthRabbitQueues) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthRabbitQueues) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HealthRabbitQueues) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HealthRabbitQueues) SetStatus(v string) {
	o.Status = &v
}

func (o HealthRabbitQueues) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableHealthRabbitQueues struct {
	value *HealthRabbitQueues
	isSet bool
}

func (v NullableHealthRabbitQueues) Get() *HealthRabbitQueues {
	return v.value
}

func (v *NullableHealthRabbitQueues) Set(val *HealthRabbitQueues) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthRabbitQueues) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthRabbitQueues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthRabbitQueues(val *HealthRabbitQueues) *NullableHealthRabbitQueues {
	return &NullableHealthRabbitQueues{value: val, isSet: true}
}

func (v NullableHealthRabbitQueues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthRabbitQueues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


