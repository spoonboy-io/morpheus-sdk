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

// InlineResponse200147 struct for InlineResponse200147
type InlineResponse200147 struct {
	StorageServer *StorageServer `json:"storageServer,omitempty"`
}

// NewInlineResponse200147 instantiates a new InlineResponse200147 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200147() *InlineResponse200147 {
	this := InlineResponse200147{}
	return &this
}

// NewInlineResponse200147WithDefaults instantiates a new InlineResponse200147 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200147WithDefaults() *InlineResponse200147 {
	this := InlineResponse200147{}
	return &this
}

// GetStorageServer returns the StorageServer field value if set, zero value otherwise.
func (o *InlineResponse200147) GetStorageServer() StorageServer {
	if o == nil || o.StorageServer == nil {
		var ret StorageServer
		return ret
	}
	return *o.StorageServer
}

// GetStorageServerOk returns a tuple with the StorageServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200147) GetStorageServerOk() (*StorageServer, bool) {
	if o == nil || o.StorageServer == nil {
		return nil, false
	}
	return o.StorageServer, true
}

// HasStorageServer returns a boolean if a field has been set.
func (o *InlineResponse200147) HasStorageServer() bool {
	if o != nil && o.StorageServer != nil {
		return true
	}

	return false
}

// SetStorageServer gets a reference to the given StorageServer and assigns it to the StorageServer field.
func (o *InlineResponse200147) SetStorageServer(v StorageServer) {
	o.StorageServer = &v
}

func (o InlineResponse200147) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StorageServer != nil {
		toSerialize["storageServer"] = o.StorageServer
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200147 struct {
	value *InlineResponse200147
	isSet bool
}

func (v NullableInlineResponse200147) Get() *InlineResponse200147 {
	return v.value
}

func (v *NullableInlineResponse200147) Set(val *InlineResponse200147) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200147) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200147) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200147(val *InlineResponse200147) *NullableInlineResponse200147 {
	return &NullableInlineResponse200147{value: val, isSet: true}
}

func (v NullableInlineResponse200147) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200147) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


