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

// InlineObject7 struct for InlineObject7
type InlineObject7 struct {
	ArchiveBucket *ArchiveBucketCreate `json:"archiveBucket,omitempty"`
}

// NewInlineObject7 instantiates a new InlineObject7 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject7() *InlineObject7 {
	this := InlineObject7{}
	return &this
}

// NewInlineObject7WithDefaults instantiates a new InlineObject7 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject7WithDefaults() *InlineObject7 {
	this := InlineObject7{}
	return &this
}

// GetArchiveBucket returns the ArchiveBucket field value if set, zero value otherwise.
func (o *InlineObject7) GetArchiveBucket() ArchiveBucketCreate {
	if o == nil || o.ArchiveBucket == nil {
		var ret ArchiveBucketCreate
		return ret
	}
	return *o.ArchiveBucket
}

// GetArchiveBucketOk returns a tuple with the ArchiveBucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject7) GetArchiveBucketOk() (*ArchiveBucketCreate, bool) {
	if o == nil || o.ArchiveBucket == nil {
		return nil, false
	}
	return o.ArchiveBucket, true
}

// HasArchiveBucket returns a boolean if a field has been set.
func (o *InlineObject7) HasArchiveBucket() bool {
	if o != nil && o.ArchiveBucket != nil {
		return true
	}

	return false
}

// SetArchiveBucket gets a reference to the given ArchiveBucketCreate and assigns it to the ArchiveBucket field.
func (o *InlineObject7) SetArchiveBucket(v ArchiveBucketCreate) {
	o.ArchiveBucket = &v
}

func (o InlineObject7) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArchiveBucket != nil {
		toSerialize["archiveBucket"] = o.ArchiveBucket
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject7 struct {
	value *InlineObject7
	isSet bool
}

func (v NullableInlineObject7) Get() *InlineObject7 {
	return v.value
}

func (v *NullableInlineObject7) Set(val *InlineObject7) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject7) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject7) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject7(val *InlineObject7) *NullableInlineObject7 {
	return &NullableInlineObject7{value: val, isSet: true}
}

func (v NullableInlineObject7) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject7) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


