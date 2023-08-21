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

// IntegrationChefConfigDatabags struct for IntegrationChefConfigDatabags
type IntegrationChefConfigDatabags struct {
	Path *string `json:"path,omitempty"`
	Key *string `json:"key,omitempty"`
}

// NewIntegrationChefConfigDatabags instantiates a new IntegrationChefConfigDatabags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationChefConfigDatabags() *IntegrationChefConfigDatabags {
	this := IntegrationChefConfigDatabags{}
	return &this
}

// NewIntegrationChefConfigDatabagsWithDefaults instantiates a new IntegrationChefConfigDatabags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationChefConfigDatabagsWithDefaults() *IntegrationChefConfigDatabags {
	this := IntegrationChefConfigDatabags{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *IntegrationChefConfigDatabags) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationChefConfigDatabags) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *IntegrationChefConfigDatabags) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *IntegrationChefConfigDatabags) SetPath(v string) {
	o.Path = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *IntegrationChefConfigDatabags) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationChefConfigDatabags) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *IntegrationChefConfigDatabags) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *IntegrationChefConfigDatabags) SetKey(v string) {
	o.Key = &v
}

func (o IntegrationChefConfigDatabags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationChefConfigDatabags struct {
	value *IntegrationChefConfigDatabags
	isSet bool
}

func (v NullableIntegrationChefConfigDatabags) Get() *IntegrationChefConfigDatabags {
	return v.value
}

func (v *NullableIntegrationChefConfigDatabags) Set(val *IntegrationChefConfigDatabags) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationChefConfigDatabags) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationChefConfigDatabags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationChefConfigDatabags(val *IntegrationChefConfigDatabags) *NullableIntegrationChefConfigDatabags {
	return &NullableIntegrationChefConfigDatabags{value: val, isSet: true}
}

func (v NullableIntegrationChefConfigDatabags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationChefConfigDatabags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


