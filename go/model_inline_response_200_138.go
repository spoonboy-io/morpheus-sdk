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

// InlineResponse200138 struct for InlineResponse200138
type InlineResponse200138 struct {
	Snapshots *[]Snapshot `json:"snapshots,omitempty"`
}

// NewInlineResponse200138 instantiates a new InlineResponse200138 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200138() *InlineResponse200138 {
	this := InlineResponse200138{}
	return &this
}

// NewInlineResponse200138WithDefaults instantiates a new InlineResponse200138 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200138WithDefaults() *InlineResponse200138 {
	this := InlineResponse200138{}
	return &this
}

// GetSnapshots returns the Snapshots field value if set, zero value otherwise.
func (o *InlineResponse200138) GetSnapshots() []Snapshot {
	if o == nil || o.Snapshots == nil {
		var ret []Snapshot
		return ret
	}
	return *o.Snapshots
}

// GetSnapshotsOk returns a tuple with the Snapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200138) GetSnapshotsOk() (*[]Snapshot, bool) {
	if o == nil || o.Snapshots == nil {
		return nil, false
	}
	return o.Snapshots, true
}

// HasSnapshots returns a boolean if a field has been set.
func (o *InlineResponse200138) HasSnapshots() bool {
	if o != nil && o.Snapshots != nil {
		return true
	}

	return false
}

// SetSnapshots gets a reference to the given []Snapshot and assigns it to the Snapshots field.
func (o *InlineResponse200138) SetSnapshots(v []Snapshot) {
	o.Snapshots = &v
}

func (o InlineResponse200138) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Snapshots != nil {
		toSerialize["snapshots"] = o.Snapshots
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200138 struct {
	value *InlineResponse200138
	isSet bool
}

func (v NullableInlineResponse200138) Get() *InlineResponse200138 {
	return v.value
}

func (v *NullableInlineResponse200138) Set(val *InlineResponse200138) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200138) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200138) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200138(val *InlineResponse200138) *NullableInlineResponse200138 {
	return &NullableInlineResponse200138{value: val, isSet: true}
}

func (v NullableInlineResponse200138) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200138) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

