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

// InlineObject8 struct for InlineObject8
type InlineObject8 struct {
	ArchiveBucket *ArchiveBucketUpdate `json:"archiveBucket,omitempty"`
}

// NewInlineObject8 instantiates a new InlineObject8 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject8() *InlineObject8 {
	this := InlineObject8{}
	return &this
}

// NewInlineObject8WithDefaults instantiates a new InlineObject8 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject8WithDefaults() *InlineObject8 {
	this := InlineObject8{}
	return &this
}

// GetArchiveBucket returns the ArchiveBucket field value if set, zero value otherwise.
func (o *InlineObject8) GetArchiveBucket() ArchiveBucketUpdate {
	if o == nil || o.ArchiveBucket == nil {
		var ret ArchiveBucketUpdate
		return ret
	}
	return *o.ArchiveBucket
}

// GetArchiveBucketOk returns a tuple with the ArchiveBucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject8) GetArchiveBucketOk() (*ArchiveBucketUpdate, bool) {
	if o == nil || o.ArchiveBucket == nil {
		return nil, false
	}
	return o.ArchiveBucket, true
}

// HasArchiveBucket returns a boolean if a field has been set.
func (o *InlineObject8) HasArchiveBucket() bool {
	if o != nil && o.ArchiveBucket != nil {
		return true
	}

	return false
}

// SetArchiveBucket gets a reference to the given ArchiveBucketUpdate and assigns it to the ArchiveBucket field.
func (o *InlineObject8) SetArchiveBucket(v ArchiveBucketUpdate) {
	o.ArchiveBucket = &v
}

func (o InlineObject8) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArchiveBucket != nil {
		toSerialize["archiveBucket"] = o.ArchiveBucket
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject8 struct {
	value *InlineObject8
	isSet bool
}

func (v NullableInlineObject8) Get() *InlineObject8 {
	return v.value
}

func (v *NullableInlineObject8) Set(val *InlineObject8) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject8) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject8) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject8(val *InlineObject8) *NullableInlineObject8 {
	return &NullableInlineObject8{value: val, isSet: true}
}

func (v NullableInlineObject8) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject8) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


