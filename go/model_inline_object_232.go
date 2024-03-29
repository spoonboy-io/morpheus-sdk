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

// InlineObject232 struct for InlineObject232
type InlineObject232 struct {
	StorageBucket ApiStorageBucketsStorageBucket `json:"storageBucket"`
}

// NewInlineObject232 instantiates a new InlineObject232 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject232(storageBucket ApiStorageBucketsStorageBucket, ) *InlineObject232 {
	this := InlineObject232{}
	this.StorageBucket = storageBucket
	return &this
}

// NewInlineObject232WithDefaults instantiates a new InlineObject232 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject232WithDefaults() *InlineObject232 {
	this := InlineObject232{}
	return &this
}

// GetStorageBucket returns the StorageBucket field value
func (o *InlineObject232) GetStorageBucket() ApiStorageBucketsStorageBucket {
	if o == nil  {
		var ret ApiStorageBucketsStorageBucket
		return ret
	}

	return o.StorageBucket
}

// GetStorageBucketOk returns a tuple with the StorageBucket field value
// and a boolean to check if the value has been set.
func (o *InlineObject232) GetStorageBucketOk() (*ApiStorageBucketsStorageBucket, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StorageBucket, true
}

// SetStorageBucket sets field value
func (o *InlineObject232) SetStorageBucket(v ApiStorageBucketsStorageBucket) {
	o.StorageBucket = v
}

func (o InlineObject232) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["storageBucket"] = o.StorageBucket
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject232 struct {
	value *InlineObject232
	isSet bool
}

func (v NullableInlineObject232) Get() *InlineObject232 {
	return v.value
}

func (v *NullableInlineObject232) Set(val *InlineObject232) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject232) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject232) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject232(val *InlineObject232) *NullableInlineObject232 {
	return &NullableInlineObject232{value: val, isSet: true}
}

func (v NullableInlineObject232) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject232) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


