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

// ApiKeyPairsGenerateKeyPair struct for ApiKeyPairsGenerateKeyPair
type ApiKeyPairsGenerateKeyPair struct {
	Name string `json:"name"`
}

// NewApiKeyPairsGenerateKeyPair instantiates a new ApiKeyPairsGenerateKeyPair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeyPairsGenerateKeyPair(name string, ) *ApiKeyPairsGenerateKeyPair {
	this := ApiKeyPairsGenerateKeyPair{}
	this.Name = name
	return &this
}

// NewApiKeyPairsGenerateKeyPairWithDefaults instantiates a new ApiKeyPairsGenerateKeyPair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyPairsGenerateKeyPairWithDefaults() *ApiKeyPairsGenerateKeyPair {
	this := ApiKeyPairsGenerateKeyPair{}
	return &this
}

// GetName returns the Name field value
func (o *ApiKeyPairsGenerateKeyPair) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiKeyPairsGenerateKeyPair) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiKeyPairsGenerateKeyPair) SetName(v string) {
	o.Name = v
}

func (o ApiKeyPairsGenerateKeyPair) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableApiKeyPairsGenerateKeyPair struct {
	value *ApiKeyPairsGenerateKeyPair
	isSet bool
}

func (v NullableApiKeyPairsGenerateKeyPair) Get() *ApiKeyPairsGenerateKeyPair {
	return v.value
}

func (v *NullableApiKeyPairsGenerateKeyPair) Set(val *ApiKeyPairsGenerateKeyPair) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKeyPairsGenerateKeyPair) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKeyPairsGenerateKeyPair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKeyPairsGenerateKeyPair(val *ApiKeyPairsGenerateKeyPair) *NullableApiKeyPairsGenerateKeyPair {
	return &NullableApiKeyPairsGenerateKeyPair{value: val, isSet: true}
}

func (v NullableApiKeyPairsGenerateKeyPair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKeyPairsGenerateKeyPair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

