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

// checks if the AddClientRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddClientRequest{}

// AddClientRequest struct for AddClientRequest
type AddClientRequest struct {
	Client AddClientRequestClient `json:"client"`
}

// NewAddClientRequest instantiates a new AddClientRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddClientRequest(client AddClientRequestClient) *AddClientRequest {
	this := AddClientRequest{}
	this.Client = client
	return &this
}

// NewAddClientRequestWithDefaults instantiates a new AddClientRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddClientRequestWithDefaults() *AddClientRequest {
	this := AddClientRequest{}
	return &this
}

// GetClient returns the Client field value
func (o *AddClientRequest) GetClient() AddClientRequestClient {
	if o == nil {
		var ret AddClientRequestClient
		return ret
	}

	return o.Client
}

// GetClientOk returns a tuple with the Client field value
// and a boolean to check if the value has been set.
func (o *AddClientRequest) GetClientOk() (*AddClientRequestClient, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Client, true
}

// SetClient sets field value
func (o *AddClientRequest) SetClient(v AddClientRequestClient) {
	o.Client = v
}

func (o AddClientRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddClientRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["client"] = o.Client
	return toSerialize, nil
}

type NullableAddClientRequest struct {
	value *AddClientRequest
	isSet bool
}

func (v NullableAddClientRequest) Get() *AddClientRequest {
	return v.value
}

func (v *NullableAddClientRequest) Set(val *AddClientRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddClientRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddClientRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddClientRequest(val *AddClientRequest) *NullableAddClientRequest {
	return &NullableAddClientRequest{value: val, isSet: true}
}

func (v NullableAddClientRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddClientRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


