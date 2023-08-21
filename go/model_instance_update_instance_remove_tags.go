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

// InstanceUpdateInstanceRemoveTags struct for InstanceUpdateInstanceRemoveTags
type InstanceUpdateInstanceRemoveTags struct {
	Name *string `json:"name,omitempty"`
	Value NullableString `json:"value,omitempty"`
}

// NewInstanceUpdateInstanceRemoveTags instantiates a new InstanceUpdateInstanceRemoveTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceUpdateInstanceRemoveTags() *InstanceUpdateInstanceRemoveTags {
	this := InstanceUpdateInstanceRemoveTags{}
	return &this
}

// NewInstanceUpdateInstanceRemoveTagsWithDefaults instantiates a new InstanceUpdateInstanceRemoveTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceUpdateInstanceRemoveTagsWithDefaults() *InstanceUpdateInstanceRemoveTags {
	this := InstanceUpdateInstanceRemoveTags{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InstanceUpdateInstanceRemoveTags) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceUpdateInstanceRemoveTags) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InstanceUpdateInstanceRemoveTags) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InstanceUpdateInstanceRemoveTags) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceUpdateInstanceRemoveTags) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceUpdateInstanceRemoveTags) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *InstanceUpdateInstanceRemoveTags) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *InstanceUpdateInstanceRemoveTags) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *InstanceUpdateInstanceRemoveTags) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *InstanceUpdateInstanceRemoveTags) UnsetValue() {
	o.Value.Unset()
}

func (o InstanceUpdateInstanceRemoveTags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInstanceUpdateInstanceRemoveTags struct {
	value *InstanceUpdateInstanceRemoveTags
	isSet bool
}

func (v NullableInstanceUpdateInstanceRemoveTags) Get() *InstanceUpdateInstanceRemoveTags {
	return v.value
}

func (v *NullableInstanceUpdateInstanceRemoveTags) Set(val *InstanceUpdateInstanceRemoveTags) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceUpdateInstanceRemoveTags) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceUpdateInstanceRemoveTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceUpdateInstanceRemoveTags(val *InstanceUpdateInstanceRemoveTags) *NullableInstanceUpdateInstanceRemoveTags {
	return &NullableInstanceUpdateInstanceRemoveTags{value: val, isSet: true}
}

func (v NullableInstanceUpdateInstanceRemoveTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceUpdateInstanceRemoveTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


