/*
Morpheus API

Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.

API version: 6.1.1
Contact: dev@morpheusdata.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AppStateInputProvidersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStateInputProvidersInner{}

// AppStateInputProvidersInner struct for AppStateInputProvidersInner
type AppStateInputProvidersInner struct {
	Name *string `json:"name,omitempty"`
}

// NewAppStateInputProvidersInner instantiates a new AppStateInputProvidersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStateInputProvidersInner() *AppStateInputProvidersInner {
	this := AppStateInputProvidersInner{}
	return &this
}

// NewAppStateInputProvidersInnerWithDefaults instantiates a new AppStateInputProvidersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStateInputProvidersInnerWithDefaults() *AppStateInputProvidersInner {
	this := AppStateInputProvidersInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppStateInputProvidersInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStateInputProvidersInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppStateInputProvidersInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppStateInputProvidersInner) SetName(v string) {
	o.Name = &v
}

func (o AppStateInputProvidersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStateInputProvidersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableAppStateInputProvidersInner struct {
	value *AppStateInputProvidersInner
	isSet bool
}

func (v NullableAppStateInputProvidersInner) Get() *AppStateInputProvidersInner {
	return v.value
}

func (v *NullableAppStateInputProvidersInner) Set(val *AppStateInputProvidersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStateInputProvidersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStateInputProvidersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStateInputProvidersInner(val *AppStateInputProvidersInner) *NullableAppStateInputProvidersInner {
	return &NullableAppStateInputProvidersInner{value: val, isSet: true}
}

func (v NullableAppStateInputProvidersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStateInputProvidersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


