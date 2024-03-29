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

// InlineObject221 struct for InlineObject221
type InlineObject221 struct {
	// Move associated instances to specified Tenant account.
	MoveAssociatedInstances *bool `json:"moveAssociatedInstances,omitempty"`
}

// NewInlineObject221 instantiates a new InlineObject221 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject221() *InlineObject221 {
	this := InlineObject221{}
	var moveAssociatedInstances bool = true
	this.MoveAssociatedInstances = &moveAssociatedInstances
	return &this
}

// NewInlineObject221WithDefaults instantiates a new InlineObject221 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject221WithDefaults() *InlineObject221 {
	this := InlineObject221{}
	var moveAssociatedInstances bool = true
	this.MoveAssociatedInstances = &moveAssociatedInstances
	return &this
}

// GetMoveAssociatedInstances returns the MoveAssociatedInstances field value if set, zero value otherwise.
func (o *InlineObject221) GetMoveAssociatedInstances() bool {
	if o == nil || o.MoveAssociatedInstances == nil {
		var ret bool
		return ret
	}
	return *o.MoveAssociatedInstances
}

// GetMoveAssociatedInstancesOk returns a tuple with the MoveAssociatedInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject221) GetMoveAssociatedInstancesOk() (*bool, bool) {
	if o == nil || o.MoveAssociatedInstances == nil {
		return nil, false
	}
	return o.MoveAssociatedInstances, true
}

// HasMoveAssociatedInstances returns a boolean if a field has been set.
func (o *InlineObject221) HasMoveAssociatedInstances() bool {
	if o != nil && o.MoveAssociatedInstances != nil {
		return true
	}

	return false
}

// SetMoveAssociatedInstances gets a reference to the given bool and assigns it to the MoveAssociatedInstances field.
func (o *InlineObject221) SetMoveAssociatedInstances(v bool) {
	o.MoveAssociatedInstances = &v
}

func (o InlineObject221) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MoveAssociatedInstances != nil {
		toSerialize["moveAssociatedInstances"] = o.MoveAssociatedInstances
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject221 struct {
	value *InlineObject221
	isSet bool
}

func (v NullableInlineObject221) Get() *InlineObject221 {
	return v.value
}

func (v *NullableInlineObject221) Set(val *InlineObject221) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject221) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject221) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject221(val *InlineObject221) *NullableInlineObject221 {
	return &NullableInlineObject221{value: val, isSet: true}
}

func (v NullableInlineObject221) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject221) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


