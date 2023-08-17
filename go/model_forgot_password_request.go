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

// checks if the ForgotPasswordRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForgotPasswordRequest{}

// ForgotPasswordRequest struct for ForgotPasswordRequest
type ForgotPasswordRequest struct {
	// Specified as \"username\" or \"tenantId\\username\" with proper HTML encoding and used in conjunction with `password`. 
	Username string `json:"username"`
}

// NewForgotPasswordRequest instantiates a new ForgotPasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForgotPasswordRequest(username string) *ForgotPasswordRequest {
	this := ForgotPasswordRequest{}
	this.Username = username
	return &this
}

// NewForgotPasswordRequestWithDefaults instantiates a new ForgotPasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForgotPasswordRequestWithDefaults() *ForgotPasswordRequest {
	this := ForgotPasswordRequest{}
	return &this
}

// GetUsername returns the Username field value
func (o *ForgotPasswordRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *ForgotPasswordRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *ForgotPasswordRequest) SetUsername(v string) {
	o.Username = v
}

func (o ForgotPasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForgotPasswordRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

type NullableForgotPasswordRequest struct {
	value *ForgotPasswordRequest
	isSet bool
}

func (v NullableForgotPasswordRequest) Get() *ForgotPasswordRequest {
	return v.value
}

func (v *NullableForgotPasswordRequest) Set(val *ForgotPasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableForgotPasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableForgotPasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForgotPasswordRequest(val *ForgotPasswordRequest) *NullableForgotPasswordRequest {
	return &NullableForgotPasswordRequest{value: val, isSet: true}
}

func (v NullableForgotPasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForgotPasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

