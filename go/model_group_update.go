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

// GroupUpdate struct for GroupUpdate
type GroupUpdate struct {
	// A unique name scoped to your account for the group
	Name string `json:"name"`
	// Optional code for use with policies
	Code *string `json:"code,omitempty"`
	// Optional location argument for your group
	Location *string `json:"location,omitempty"`
	Config *GroupCreateConfig `json:"config,omitempty"`
}

// NewGroupUpdate instantiates a new GroupUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupUpdate(name string, ) *GroupUpdate {
	this := GroupUpdate{}
	this.Name = name
	return &this
}

// NewGroupUpdateWithDefaults instantiates a new GroupUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupUpdateWithDefaults() *GroupUpdate {
	this := GroupUpdate{}
	return &this
}

// GetName returns the Name field value
func (o *GroupUpdate) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GroupUpdate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GroupUpdate) SetName(v string) {
	o.Name = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GroupUpdate) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupUpdate) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GroupUpdate) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GroupUpdate) SetCode(v string) {
	o.Code = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *GroupUpdate) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupUpdate) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *GroupUpdate) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *GroupUpdate) SetLocation(v string) {
	o.Location = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *GroupUpdate) GetConfig() GroupCreateConfig {
	if o == nil || o.Config == nil {
		var ret GroupCreateConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupUpdate) GetConfigOk() (*GroupCreateConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *GroupUpdate) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given GroupCreateConfig and assigns it to the Config field.
func (o *GroupUpdate) SetConfig(v GroupCreateConfig) {
	o.Config = &v
}

func (o GroupUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableGroupUpdate struct {
	value *GroupUpdate
	isSet bool
}

func (v NullableGroupUpdate) Get() *GroupUpdate {
	return v.value
}

func (v *NullableGroupUpdate) Set(val *GroupUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupUpdate(val *GroupUpdate) *NullableGroupUpdate {
	return &NullableGroupUpdate{value: val, isSet: true}
}

func (v NullableGroupUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


