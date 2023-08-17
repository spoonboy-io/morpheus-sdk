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

// checks if the BackupSettingsUpdateDefaultStorageBucket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupSettingsUpdateDefaultStorageBucket{}

// BackupSettingsUpdateDefaultStorageBucket struct for BackupSettingsUpdateDefaultStorageBucket
type BackupSettingsUpdateDefaultStorageBucket struct {
	// ID of default storage bucket
	Id *int64 `json:"id,omitempty"`
}

// NewBackupSettingsUpdateDefaultStorageBucket instantiates a new BackupSettingsUpdateDefaultStorageBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupSettingsUpdateDefaultStorageBucket() *BackupSettingsUpdateDefaultStorageBucket {
	this := BackupSettingsUpdateDefaultStorageBucket{}
	return &this
}

// NewBackupSettingsUpdateDefaultStorageBucketWithDefaults instantiates a new BackupSettingsUpdateDefaultStorageBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupSettingsUpdateDefaultStorageBucketWithDefaults() *BackupSettingsUpdateDefaultStorageBucket {
	this := BackupSettingsUpdateDefaultStorageBucket{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BackupSettingsUpdateDefaultStorageBucket) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSettingsUpdateDefaultStorageBucket) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BackupSettingsUpdateDefaultStorageBucket) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *BackupSettingsUpdateDefaultStorageBucket) SetId(v int64) {
	o.Id = &v
}

func (o BackupSettingsUpdateDefaultStorageBucket) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupSettingsUpdateDefaultStorageBucket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableBackupSettingsUpdateDefaultStorageBucket struct {
	value *BackupSettingsUpdateDefaultStorageBucket
	isSet bool
}

func (v NullableBackupSettingsUpdateDefaultStorageBucket) Get() *BackupSettingsUpdateDefaultStorageBucket {
	return v.value
}

func (v *NullableBackupSettingsUpdateDefaultStorageBucket) Set(val *BackupSettingsUpdateDefaultStorageBucket) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupSettingsUpdateDefaultStorageBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupSettingsUpdateDefaultStorageBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupSettingsUpdateDefaultStorageBucket(val *BackupSettingsUpdateDefaultStorageBucket) *NullableBackupSettingsUpdateDefaultStorageBucket {
	return &NullableBackupSettingsUpdateDefaultStorageBucket{value: val, isSet: true}
}

func (v NullableBackupSettingsUpdateDefaultStorageBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupSettingsUpdateDefaultStorageBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

