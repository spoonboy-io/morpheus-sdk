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

// AppStateInputProviders struct for AppStateInputProviders
type AppStateInputProviders struct {
	Name *string `json:"name,omitempty"`
}

// NewAppStateInputProviders instantiates a new AppStateInputProviders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStateInputProviders() *AppStateInputProviders {
	this := AppStateInputProviders{}
	return &this
}

// NewAppStateInputProvidersWithDefaults instantiates a new AppStateInputProviders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStateInputProvidersWithDefaults() *AppStateInputProviders {
	this := AppStateInputProviders{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppStateInputProviders) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStateInputProviders) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppStateInputProviders) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppStateInputProviders) SetName(v string) {
	o.Name = &v
}

func (o AppStateInputProviders) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableAppStateInputProviders struct {
	value *AppStateInputProviders
	isSet bool
}

func (v NullableAppStateInputProviders) Get() *AppStateInputProviders {
	return v.value
}

func (v *NullableAppStateInputProviders) Set(val *AppStateInputProviders) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStateInputProviders) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStateInputProviders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStateInputProviders(val *AppStateInputProviders) *NullableAppStateInputProviders {
	return &NullableAppStateInputProviders{value: val, isSet: true}
}

func (v NullableAppStateInputProviders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStateInputProviders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


