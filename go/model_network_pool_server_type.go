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

// NetworkPoolServerType Network Pool Server Type
type NetworkPoolServerType struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Code *string `json:"code,omitempty"`
}

// NewNetworkPoolServerType instantiates a new NetworkPoolServerType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkPoolServerType() *NetworkPoolServerType {
	this := NetworkPoolServerType{}
	return &this
}

// NewNetworkPoolServerTypeWithDefaults instantiates a new NetworkPoolServerType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkPoolServerTypeWithDefaults() *NetworkPoolServerType {
	this := NetworkPoolServerType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworkPoolServerType) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServerType) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworkPoolServerType) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *NetworkPoolServerType) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkPoolServerType) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServerType) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkPoolServerType) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkPoolServerType) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *NetworkPoolServerType) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServerType) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *NetworkPoolServerType) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *NetworkPoolServerType) SetCode(v string) {
	o.Code = &v
}

func (o NetworkPoolServerType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkPoolServerType struct {
	value *NetworkPoolServerType
	isSet bool
}

func (v NullableNetworkPoolServerType) Get() *NetworkPoolServerType {
	return v.value
}

func (v *NullableNetworkPoolServerType) Set(val *NetworkPoolServerType) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkPoolServerType) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkPoolServerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkPoolServerType(val *NetworkPoolServerType) *NullableNetworkPoolServerType {
	return &NullableNetworkPoolServerType{value: val, isSet: true}
}

func (v NullableNetworkPoolServerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkPoolServerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


