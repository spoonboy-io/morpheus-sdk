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

// checks if the ListArchiveFiles200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListArchiveFiles200Response{}

// ListArchiveFiles200Response struct for ListArchiveFiles200Response
type ListArchiveFiles200Response struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	ArchiveBucket *ArchiveBucket `json:"archiveBucket,omitempty"`
	IsOwner *bool `json:"isOwner,omitempty"`
	ParentDirectory NullableString `json:"parentDirectory,omitempty"`
	ArchiveFiles []ArchiveBucketFile `json:"archiveFiles,omitempty"`
}

// NewListArchiveFiles200Response instantiates a new ListArchiveFiles200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListArchiveFiles200Response() *ListArchiveFiles200Response {
	this := ListArchiveFiles200Response{}
	return &this
}

// NewListArchiveFiles200ResponseWithDefaults instantiates a new ListArchiveFiles200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListArchiveFiles200ResponseWithDefaults() *ListArchiveFiles200Response {
	this := ListArchiveFiles200Response{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListArchiveFiles200Response) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListArchiveFiles200Response) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListArchiveFiles200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ListArchiveFiles200Response) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetArchiveBucket returns the ArchiveBucket field value if set, zero value otherwise.
func (o *ListArchiveFiles200Response) GetArchiveBucket() ArchiveBucket {
	if o == nil || IsNil(o.ArchiveBucket) {
		var ret ArchiveBucket
		return ret
	}
	return *o.ArchiveBucket
}

// GetArchiveBucketOk returns a tuple with the ArchiveBucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListArchiveFiles200Response) GetArchiveBucketOk() (*ArchiveBucket, bool) {
	if o == nil || IsNil(o.ArchiveBucket) {
		return nil, false
	}
	return o.ArchiveBucket, true
}

// HasArchiveBucket returns a boolean if a field has been set.
func (o *ListArchiveFiles200Response) HasArchiveBucket() bool {
	if o != nil && !IsNil(o.ArchiveBucket) {
		return true
	}

	return false
}

// SetArchiveBucket gets a reference to the given ArchiveBucket and assigns it to the ArchiveBucket field.
func (o *ListArchiveFiles200Response) SetArchiveBucket(v ArchiveBucket) {
	o.ArchiveBucket = &v
}

// GetIsOwner returns the IsOwner field value if set, zero value otherwise.
func (o *ListArchiveFiles200Response) GetIsOwner() bool {
	if o == nil || IsNil(o.IsOwner) {
		var ret bool
		return ret
	}
	return *o.IsOwner
}

// GetIsOwnerOk returns a tuple with the IsOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListArchiveFiles200Response) GetIsOwnerOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOwner) {
		return nil, false
	}
	return o.IsOwner, true
}

// HasIsOwner returns a boolean if a field has been set.
func (o *ListArchiveFiles200Response) HasIsOwner() bool {
	if o != nil && !IsNil(o.IsOwner) {
		return true
	}

	return false
}

// SetIsOwner gets a reference to the given bool and assigns it to the IsOwner field.
func (o *ListArchiveFiles200Response) SetIsOwner(v bool) {
	o.IsOwner = &v
}

// GetParentDirectory returns the ParentDirectory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListArchiveFiles200Response) GetParentDirectory() string {
	if o == nil || IsNil(o.ParentDirectory.Get()) {
		var ret string
		return ret
	}
	return *o.ParentDirectory.Get()
}

// GetParentDirectoryOk returns a tuple with the ParentDirectory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListArchiveFiles200Response) GetParentDirectoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentDirectory.Get(), o.ParentDirectory.IsSet()
}

// HasParentDirectory returns a boolean if a field has been set.
func (o *ListArchiveFiles200Response) HasParentDirectory() bool {
	if o != nil && o.ParentDirectory.IsSet() {
		return true
	}

	return false
}

// SetParentDirectory gets a reference to the given NullableString and assigns it to the ParentDirectory field.
func (o *ListArchiveFiles200Response) SetParentDirectory(v string) {
	o.ParentDirectory.Set(&v)
}
// SetParentDirectoryNil sets the value for ParentDirectory to be an explicit nil
func (o *ListArchiveFiles200Response) SetParentDirectoryNil() {
	o.ParentDirectory.Set(nil)
}

// UnsetParentDirectory ensures that no value is present for ParentDirectory, not even an explicit nil
func (o *ListArchiveFiles200Response) UnsetParentDirectory() {
	o.ParentDirectory.Unset()
}

// GetArchiveFiles returns the ArchiveFiles field value if set, zero value otherwise.
func (o *ListArchiveFiles200Response) GetArchiveFiles() []ArchiveBucketFile {
	if o == nil || IsNil(o.ArchiveFiles) {
		var ret []ArchiveBucketFile
		return ret
	}
	return o.ArchiveFiles
}

// GetArchiveFilesOk returns a tuple with the ArchiveFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListArchiveFiles200Response) GetArchiveFilesOk() ([]ArchiveBucketFile, bool) {
	if o == nil || IsNil(o.ArchiveFiles) {
		return nil, false
	}
	return o.ArchiveFiles, true
}

// HasArchiveFiles returns a boolean if a field has been set.
func (o *ListArchiveFiles200Response) HasArchiveFiles() bool {
	if o != nil && !IsNil(o.ArchiveFiles) {
		return true
	}

	return false
}

// SetArchiveFiles gets a reference to the given []ArchiveBucketFile and assigns it to the ArchiveFiles field.
func (o *ListArchiveFiles200Response) SetArchiveFiles(v []ArchiveBucketFile) {
	o.ArchiveFiles = v
}

func (o ListArchiveFiles200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListArchiveFiles200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.ArchiveBucket) {
		toSerialize["archiveBucket"] = o.ArchiveBucket
	}
	if !IsNil(o.IsOwner) {
		toSerialize["isOwner"] = o.IsOwner
	}
	if o.ParentDirectory.IsSet() {
		toSerialize["parentDirectory"] = o.ParentDirectory.Get()
	}
	if !IsNil(o.ArchiveFiles) {
		toSerialize["archiveFiles"] = o.ArchiveFiles
	}
	return toSerialize, nil
}

type NullableListArchiveFiles200Response struct {
	value *ListArchiveFiles200Response
	isSet bool
}

func (v NullableListArchiveFiles200Response) Get() *ListArchiveFiles200Response {
	return v.value
}

func (v *NullableListArchiveFiles200Response) Set(val *ListArchiveFiles200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListArchiveFiles200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListArchiveFiles200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListArchiveFiles200Response(val *ListArchiveFiles200Response) *NullableListArchiveFiles200Response {
	return &NullableListArchiveFiles200Response{value: val, isSet: true}
}

func (v NullableListArchiveFiles200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListArchiveFiles200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


