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
	"time"
)

// checks if the ArchiveFileLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArchiveFileLinks{}

// ArchiveFileLinks struct for ArchiveFileLinks
type ArchiveFileLinks struct {
	Id *int64 `json:"id,omitempty"`
	SecretAccessKey *string `json:"secretAccessKey,omitempty"`
	ArchiveFile *ArchiveFileLinksArchiveFile `json:"archiveFile,omitempty"`
	CreatedBy *ArchiveBucketFileCreatedBy `json:"createdBy,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	LastAccessDate NullableTime `json:"lastAccessDate,omitempty"`
	ExpirationDate NullableTime `json:"expirationDate,omitempty"`
	DownloadCount *int64 `json:"downloadCount,omitempty"`
}

// NewArchiveFileLinks instantiates a new ArchiveFileLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArchiveFileLinks() *ArchiveFileLinks {
	this := ArchiveFileLinks{}
	return &this
}

// NewArchiveFileLinksWithDefaults instantiates a new ArchiveFileLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArchiveFileLinksWithDefaults() *ArchiveFileLinks {
	this := ArchiveFileLinks{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ArchiveFileLinks) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveFileLinks) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ArchiveFileLinks) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ArchiveFileLinks) SetId(v int64) {
	o.Id = &v
}

// GetSecretAccessKey returns the SecretAccessKey field value if set, zero value otherwise.
func (o *ArchiveFileLinks) GetSecretAccessKey() string {
	if o == nil || IsNil(o.SecretAccessKey) {
		var ret string
		return ret
	}
	return *o.SecretAccessKey
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveFileLinks) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SecretAccessKey) {
		return nil, false
	}
	return o.SecretAccessKey, true
}

// HasSecretAccessKey returns a boolean if a field has been set.
func (o *ArchiveFileLinks) HasSecretAccessKey() bool {
	if o != nil && !IsNil(o.SecretAccessKey) {
		return true
	}

	return false
}

// SetSecretAccessKey gets a reference to the given string and assigns it to the SecretAccessKey field.
func (o *ArchiveFileLinks) SetSecretAccessKey(v string) {
	o.SecretAccessKey = &v
}

// GetArchiveFile returns the ArchiveFile field value if set, zero value otherwise.
func (o *ArchiveFileLinks) GetArchiveFile() ArchiveFileLinksArchiveFile {
	if o == nil || IsNil(o.ArchiveFile) {
		var ret ArchiveFileLinksArchiveFile
		return ret
	}
	return *o.ArchiveFile
}

// GetArchiveFileOk returns a tuple with the ArchiveFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveFileLinks) GetArchiveFileOk() (*ArchiveFileLinksArchiveFile, bool) {
	if o == nil || IsNil(o.ArchiveFile) {
		return nil, false
	}
	return o.ArchiveFile, true
}

// HasArchiveFile returns a boolean if a field has been set.
func (o *ArchiveFileLinks) HasArchiveFile() bool {
	if o != nil && !IsNil(o.ArchiveFile) {
		return true
	}

	return false
}

// SetArchiveFile gets a reference to the given ArchiveFileLinksArchiveFile and assigns it to the ArchiveFile field.
func (o *ArchiveFileLinks) SetArchiveFile(v ArchiveFileLinksArchiveFile) {
	o.ArchiveFile = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ArchiveFileLinks) GetCreatedBy() ArchiveBucketFileCreatedBy {
	if o == nil || IsNil(o.CreatedBy) {
		var ret ArchiveBucketFileCreatedBy
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveFileLinks) GetCreatedByOk() (*ArchiveBucketFileCreatedBy, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ArchiveFileLinks) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given ArchiveBucketFileCreatedBy and assigns it to the CreatedBy field.
func (o *ArchiveFileLinks) SetCreatedBy(v ArchiveBucketFileCreatedBy) {
	o.CreatedBy = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *ArchiveFileLinks) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveFileLinks) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *ArchiveFileLinks) HasDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *ArchiveFileLinks) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ArchiveFileLinks) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveFileLinks) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ArchiveFileLinks) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ArchiveFileLinks) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLastAccessDate returns the LastAccessDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArchiveFileLinks) GetLastAccessDate() time.Time {
	if o == nil || IsNil(o.LastAccessDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastAccessDate.Get()
}

// GetLastAccessDateOk returns a tuple with the LastAccessDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchiveFileLinks) GetLastAccessDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastAccessDate.Get(), o.LastAccessDate.IsSet()
}

// HasLastAccessDate returns a boolean if a field has been set.
func (o *ArchiveFileLinks) HasLastAccessDate() bool {
	if o != nil && o.LastAccessDate.IsSet() {
		return true
	}

	return false
}

// SetLastAccessDate gets a reference to the given NullableTime and assigns it to the LastAccessDate field.
func (o *ArchiveFileLinks) SetLastAccessDate(v time.Time) {
	o.LastAccessDate.Set(&v)
}
// SetLastAccessDateNil sets the value for LastAccessDate to be an explicit nil
func (o *ArchiveFileLinks) SetLastAccessDateNil() {
	o.LastAccessDate.Set(nil)
}

// UnsetLastAccessDate ensures that no value is present for LastAccessDate, not even an explicit nil
func (o *ArchiveFileLinks) UnsetLastAccessDate() {
	o.LastAccessDate.Unset()
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArchiveFileLinks) GetExpirationDate() time.Time {
	if o == nil || IsNil(o.ExpirationDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate.Get()
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchiveFileLinks) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpirationDate.Get(), o.ExpirationDate.IsSet()
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *ArchiveFileLinks) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate.IsSet() {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given NullableTime and assigns it to the ExpirationDate field.
func (o *ArchiveFileLinks) SetExpirationDate(v time.Time) {
	o.ExpirationDate.Set(&v)
}
// SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil
func (o *ArchiveFileLinks) SetExpirationDateNil() {
	o.ExpirationDate.Set(nil)
}

// UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
func (o *ArchiveFileLinks) UnsetExpirationDate() {
	o.ExpirationDate.Unset()
}

// GetDownloadCount returns the DownloadCount field value if set, zero value otherwise.
func (o *ArchiveFileLinks) GetDownloadCount() int64 {
	if o == nil || IsNil(o.DownloadCount) {
		var ret int64
		return ret
	}
	return *o.DownloadCount
}

// GetDownloadCountOk returns a tuple with the DownloadCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveFileLinks) GetDownloadCountOk() (*int64, bool) {
	if o == nil || IsNil(o.DownloadCount) {
		return nil, false
	}
	return o.DownloadCount, true
}

// HasDownloadCount returns a boolean if a field has been set.
func (o *ArchiveFileLinks) HasDownloadCount() bool {
	if o != nil && !IsNil(o.DownloadCount) {
		return true
	}

	return false
}

// SetDownloadCount gets a reference to the given int64 and assigns it to the DownloadCount field.
func (o *ArchiveFileLinks) SetDownloadCount(v int64) {
	o.DownloadCount = &v
}

func (o ArchiveFileLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArchiveFileLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SecretAccessKey) {
		toSerialize["secretAccessKey"] = o.SecretAccessKey
	}
	if !IsNil(o.ArchiveFile) {
		toSerialize["archiveFile"] = o.ArchiveFile
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.DateCreated) {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.LastAccessDate.IsSet() {
		toSerialize["lastAccessDate"] = o.LastAccessDate.Get()
	}
	if o.ExpirationDate.IsSet() {
		toSerialize["expirationDate"] = o.ExpirationDate.Get()
	}
	if !IsNil(o.DownloadCount) {
		toSerialize["downloadCount"] = o.DownloadCount
	}
	return toSerialize, nil
}

type NullableArchiveFileLinks struct {
	value *ArchiveFileLinks
	isSet bool
}

func (v NullableArchiveFileLinks) Get() *ArchiveFileLinks {
	return v.value
}

func (v *NullableArchiveFileLinks) Set(val *ArchiveFileLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableArchiveFileLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableArchiveFileLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArchiveFileLinks(val *ArchiveFileLinks) *NullableArchiveFileLinks {
	return &NullableArchiveFileLinks{value: val, isSet: true}
}

func (v NullableArchiveFileLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArchiveFileLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

