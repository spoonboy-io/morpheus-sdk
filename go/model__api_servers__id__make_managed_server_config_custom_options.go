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

// ApiServersIdMakeManagedServerConfigCustomOptions struct for ApiServersIdMakeManagedServerConfigCustomOptions
type ApiServersIdMakeManagedServerConfigCustomOptions struct {
	Dbfoldername *string `json:"dbfoldername,omitempty"`
}

// NewApiServersIdMakeManagedServerConfigCustomOptions instantiates a new ApiServersIdMakeManagedServerConfigCustomOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiServersIdMakeManagedServerConfigCustomOptions() *ApiServersIdMakeManagedServerConfigCustomOptions {
	this := ApiServersIdMakeManagedServerConfigCustomOptions{}
	return &this
}

// NewApiServersIdMakeManagedServerConfigCustomOptionsWithDefaults instantiates a new ApiServersIdMakeManagedServerConfigCustomOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiServersIdMakeManagedServerConfigCustomOptionsWithDefaults() *ApiServersIdMakeManagedServerConfigCustomOptions {
	this := ApiServersIdMakeManagedServerConfigCustomOptions{}
	return &this
}

// GetDbfoldername returns the Dbfoldername field value if set, zero value otherwise.
func (o *ApiServersIdMakeManagedServerConfigCustomOptions) GetDbfoldername() string {
	if o == nil || o.Dbfoldername == nil {
		var ret string
		return ret
	}
	return *o.Dbfoldername
}

// GetDbfoldernameOk returns a tuple with the Dbfoldername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiServersIdMakeManagedServerConfigCustomOptions) GetDbfoldernameOk() (*string, bool) {
	if o == nil || o.Dbfoldername == nil {
		return nil, false
	}
	return o.Dbfoldername, true
}

// HasDbfoldername returns a boolean if a field has been set.
func (o *ApiServersIdMakeManagedServerConfigCustomOptions) HasDbfoldername() bool {
	if o != nil && o.Dbfoldername != nil {
		return true
	}

	return false
}

// SetDbfoldername gets a reference to the given string and assigns it to the Dbfoldername field.
func (o *ApiServersIdMakeManagedServerConfigCustomOptions) SetDbfoldername(v string) {
	o.Dbfoldername = &v
}

func (o ApiServersIdMakeManagedServerConfigCustomOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dbfoldername != nil {
		toSerialize["dbfoldername"] = o.Dbfoldername
	}
	return json.Marshal(toSerialize)
}

type NullableApiServersIdMakeManagedServerConfigCustomOptions struct {
	value *ApiServersIdMakeManagedServerConfigCustomOptions
	isSet bool
}

func (v NullableApiServersIdMakeManagedServerConfigCustomOptions) Get() *ApiServersIdMakeManagedServerConfigCustomOptions {
	return v.value
}

func (v *NullableApiServersIdMakeManagedServerConfigCustomOptions) Set(val *ApiServersIdMakeManagedServerConfigCustomOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableApiServersIdMakeManagedServerConfigCustomOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableApiServersIdMakeManagedServerConfigCustomOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiServersIdMakeManagedServerConfigCustomOptions(val *ApiServersIdMakeManagedServerConfigCustomOptions) *NullableApiServersIdMakeManagedServerConfigCustomOptions {
	return &NullableApiServersIdMakeManagedServerConfigCustomOptions{value: val, isSet: true}
}

func (v NullableApiServersIdMakeManagedServerConfigCustomOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiServersIdMakeManagedServerConfigCustomOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

