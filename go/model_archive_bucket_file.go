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
	"time"
)

// ArchiveBucketFile struct for ArchiveBucketFile
type ArchiveBucketFile struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	FilePath *string `json:"filePath,omitempty"`
	ArchiveBucket *ArchiveBucketFileArchiveBucket `json:"archiveBucket,omitempty"`
	CreatedBy *ArchiveBucketFileCreatedBy `json:"createdBy,omitempty"`
	IsDirectory *bool `json:"isDirectory,omitempty"`
	Status *string `json:"status,omitempty"`
	RawSize *int64 `json:"rawSize,omitempty"`
	ContentType NullableString `json:"contentType,omitempty"`
	DownloadCount *int64 `json:"downloadCount,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
}

// NewArchiveBucketFile instantiates a new ArchiveBucketFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArchiveBucketFile() *ArchiveBucketFile {
	this := ArchiveBucketFile{}
	return &this
}

// NewArchiveBucketFileWithDefaults instantiates a new ArchiveBucketFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArchiveBucketFileWithDefaults() *ArchiveBucketFile {
	this := ArchiveBucketFile{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ArchiveBucketFile) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveBucketFile) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ArchiveBucketFile) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ArchiveBucketFile) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ArchiveBucketFile) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveBucketFile) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ArchiveBucketFile) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ArchiveBucketFile) SetName(v string) {
	o.Name = &v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *ArchiveBucketFile) GetFilePath() string {
	if o == nil || o.FilePath == nil {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveBucketFile) GetFilePathOk() (*string, bool) {
	if o == nil || o.FilePath == nil {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *ArchiveBucketFile) HasFilePath() bool {
	if o != nil && o.FilePath != nil {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *ArchiveBucketFile) SetFilePath(v string) {
	o.FilePath = &v
}

// GetArchiveBucket returns the ArchiveBucket field value if set, zero value otherwise.
func (o *ArchiveBucketFile) GetArchiveBucket() ArchiveBucketFileArchiveBucket {
	if o == nil || o.ArchiveBucket == nil {
		var ret ArchiveBucketFileArchiveBucket
		return ret
	}
	return *o.ArchiveBucket
}

// GetArchiveBucketOk returns a tuple with the ArchiveBucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveBucketFile) GetArchiveBucketOk() (*ArchiveBucketFileArchiveBucket, bool) {
	if o == nil || o.ArchiveBucket == nil {
		return nil, false
	}
	return o.ArchiveBucket, true
}

// HasArchiveBucket returns a boolean if a field has been set.
func (o *ArchiveBucketFile) HasArchiveBucket() bool {
	if o != nil && o.ArchiveBucket != nil {
		return true
	}

	return false
}

// SetArchiveBucket gets a reference to the given ArchiveBucketFileArchiveBucket and assigns it to the ArchiveBucket field.
func (o *ArchiveBucketFile) SetArchiveBucket(v ArchiveBucketFileArchiveBucket) {
	o.ArchiveBucket = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ArchiveBucketFile) GetCreatedBy() ArchiveBucketFileCreatedBy {
	if o == nil || o.CreatedBy == nil {
		var ret ArchiveBucketFileCreatedBy
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveBucketFile) GetCreatedByOk() (*ArchiveBucketFileCreatedBy, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ArchiveBucketFile) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given ArchiveBucketFileCreatedBy and assigns it to the CreatedBy field.
func (o *ArchiveBucketFile) SetCreatedBy(v ArchiveBucketFileCreatedBy) {
	o.CreatedBy = &v
}

// GetIsDirectory returns the IsDirectory field value if set, zero value otherwise.
func (o *ArchiveBucketFile) GetIsDirectory() bool {
	if o == nil || o.IsDirectory == nil {
		var ret bool
		return ret
	}
	return *o.IsDirectory
}

// GetIsDirectoryOk returns a tuple with the IsDirectory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveBucketFile) GetIsDirectoryOk() (*bool, bool) {
	if o == nil || o.IsDirectory == nil {
		return nil, false
	}
	return o.IsDirectory, true
}

// HasIsDirectory returns a boolean if a field has been set.
func (o *ArchiveBucketFile) HasIsDirectory() bool {
	if o != nil && o.IsDirectory != nil {
		return true
	}

	return false
}

// SetIsDirectory gets a reference to the given bool and assigns it to the IsDirectory field.
func (o *ArchiveBucketFile) SetIsDirectory(v bool) {
	o.IsDirectory = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ArchiveBucketFile) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveBucketFile) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ArchiveBucketFile) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ArchiveBucketFile) SetStatus(v string) {
	o.Status = &v
}

// GetRawSize returns the RawSize field value if set, zero value otherwise.
func (o *ArchiveBucketFile) GetRawSize() int64 {
	if o == nil || o.RawSize == nil {
		var ret int64
		return ret
	}
	return *o.RawSize
}

// GetRawSizeOk returns a tuple with the RawSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveBucketFile) GetRawSizeOk() (*int64, bool) {
	if o == nil || o.RawSize == nil {
		return nil, false
	}
	return o.RawSize, true
}

// HasRawSize returns a boolean if a field has been set.
func (o *ArchiveBucketFile) HasRawSize() bool {
	if o != nil && o.RawSize != nil {
		return true
	}

	return false
}

// SetRawSize gets a reference to the given int64 and assigns it to the RawSize field.
func (o *ArchiveBucketFile) SetRawSize(v int64) {
	o.RawSize = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArchiveBucketFile) GetContentType() string {
	if o == nil || o.ContentType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchiveBucketFile) GetContentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *ArchiveBucketFile) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *ArchiveBucketFile) SetContentType(v string) {
	o.ContentType.Set(&v)
}
// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *ArchiveBucketFile) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *ArchiveBucketFile) UnsetContentType() {
	o.ContentType.Unset()
}

// GetDownloadCount returns the DownloadCount field value if set, zero value otherwise.
func (o *ArchiveBucketFile) GetDownloadCount() int64 {
	if o == nil || o.DownloadCount == nil {
		var ret int64
		return ret
	}
	return *o.DownloadCount
}

// GetDownloadCountOk returns a tuple with the DownloadCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveBucketFile) GetDownloadCountOk() (*int64, bool) {
	if o == nil || o.DownloadCount == nil {
		return nil, false
	}
	return o.DownloadCount, true
}

// HasDownloadCount returns a boolean if a field has been set.
func (o *ArchiveBucketFile) HasDownloadCount() bool {
	if o != nil && o.DownloadCount != nil {
		return true
	}

	return false
}

// SetDownloadCount gets a reference to the given int64 and assigns it to the DownloadCount field.
func (o *ArchiveBucketFile) SetDownloadCount(v int64) {
	o.DownloadCount = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *ArchiveBucketFile) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveBucketFile) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *ArchiveBucketFile) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *ArchiveBucketFile) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ArchiveBucketFile) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveBucketFile) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ArchiveBucketFile) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ArchiveBucketFile) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o ArchiveBucketFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FilePath != nil {
		toSerialize["filePath"] = o.FilePath
	}
	if o.ArchiveBucket != nil {
		toSerialize["archiveBucket"] = o.ArchiveBucket
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.IsDirectory != nil {
		toSerialize["isDirectory"] = o.IsDirectory
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.RawSize != nil {
		toSerialize["rawSize"] = o.RawSize
	}
	if o.ContentType.IsSet() {
		toSerialize["contentType"] = o.ContentType.Get()
	}
	if o.DownloadCount != nil {
		toSerialize["downloadCount"] = o.DownloadCount
	}
	if o.DateCreated != nil {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	return json.Marshal(toSerialize)
}

type NullableArchiveBucketFile struct {
	value *ArchiveBucketFile
	isSet bool
}

func (v NullableArchiveBucketFile) Get() *ArchiveBucketFile {
	return v.value
}

func (v *NullableArchiveBucketFile) Set(val *ArchiveBucketFile) {
	v.value = val
	v.isSet = true
}

func (v NullableArchiveBucketFile) IsSet() bool {
	return v.isSet
}

func (v *NullableArchiveBucketFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArchiveBucketFile(val *ArchiveBucketFile) *NullableArchiveBucketFile {
	return &NullableArchiveBucketFile{value: val, isSet: true}
}

func (v NullableArchiveBucketFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArchiveBucketFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


