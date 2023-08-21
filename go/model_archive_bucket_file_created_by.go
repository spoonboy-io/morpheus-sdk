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

// ArchiveBucketFileCreatedBy struct for ArchiveBucketFileCreatedBy
type ArchiveBucketFileCreatedBy struct {
	Username *string `json:"username,omitempty"`
}

// NewArchiveBucketFileCreatedBy instantiates a new ArchiveBucketFileCreatedBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArchiveBucketFileCreatedBy() *ArchiveBucketFileCreatedBy {
	this := ArchiveBucketFileCreatedBy{}
	return &this
}

// NewArchiveBucketFileCreatedByWithDefaults instantiates a new ArchiveBucketFileCreatedBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArchiveBucketFileCreatedByWithDefaults() *ArchiveBucketFileCreatedBy {
	this := ArchiveBucketFileCreatedBy{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ArchiveBucketFileCreatedBy) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveBucketFileCreatedBy) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ArchiveBucketFileCreatedBy) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ArchiveBucketFileCreatedBy) SetUsername(v string) {
	o.Username = &v
}

func (o ArchiveBucketFileCreatedBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableArchiveBucketFileCreatedBy struct {
	value *ArchiveBucketFileCreatedBy
	isSet bool
}

func (v NullableArchiveBucketFileCreatedBy) Get() *ArchiveBucketFileCreatedBy {
	return v.value
}

func (v *NullableArchiveBucketFileCreatedBy) Set(val *ArchiveBucketFileCreatedBy) {
	v.value = val
	v.isSet = true
}

func (v NullableArchiveBucketFileCreatedBy) IsSet() bool {
	return v.isSet
}

func (v *NullableArchiveBucketFileCreatedBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArchiveBucketFileCreatedBy(val *ArchiveBucketFileCreatedBy) *NullableArchiveBucketFileCreatedBy {
	return &NullableArchiveBucketFileCreatedBy{value: val, isSet: true}
}

func (v NullableArchiveBucketFileCreatedBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArchiveBucketFileCreatedBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


