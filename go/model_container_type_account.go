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

// ContainerTypeAccount struct for ContainerTypeAccount
type ContainerTypeAccount struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewContainerTypeAccount instantiates a new ContainerTypeAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerTypeAccount() *ContainerTypeAccount {
	this := ContainerTypeAccount{}
	return &this
}

// NewContainerTypeAccountWithDefaults instantiates a new ContainerTypeAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerTypeAccountWithDefaults() *ContainerTypeAccount {
	this := ContainerTypeAccount{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContainerTypeAccount) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypeAccount) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContainerTypeAccount) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ContainerTypeAccount) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContainerTypeAccount) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypeAccount) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContainerTypeAccount) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContainerTypeAccount) SetName(v string) {
	o.Name = &v
}

func (o ContainerTypeAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableContainerTypeAccount struct {
	value *ContainerTypeAccount
	isSet bool
}

func (v NullableContainerTypeAccount) Get() *ContainerTypeAccount {
	return v.value
}

func (v *NullableContainerTypeAccount) Set(val *ContainerTypeAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerTypeAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerTypeAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerTypeAccount(val *ContainerTypeAccount) *NullableContainerTypeAccount {
	return &NullableContainerTypeAccount{value: val, isSet: true}
}

func (v NullableContainerTypeAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerTypeAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


