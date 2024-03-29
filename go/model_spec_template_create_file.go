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

// SpecTemplateCreateFile File, object specifying file type and content
type SpecTemplateCreateFile struct {
	// File Source i.e. local, repository, url.
	SourceType string `json:"sourceType"`
	// File content, the template text. Only required when sourceType is `local`.
	Content *string `json:"content,omitempty"`
	// Content Path, the repo file location or url. Required when sourceType is repository or url.
	ContentPath *string `json:"contentPath,omitempty"`
	// Content Ref, the branch/tag. Only used when sourceType is repo.
	ContentRef *string `json:"contentRef,omitempty"`
	Repository *SpecTemplateCreateFileRepository `json:"repository,omitempty"`
}

// NewSpecTemplateCreateFile instantiates a new SpecTemplateCreateFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpecTemplateCreateFile(sourceType string, ) *SpecTemplateCreateFile {
	this := SpecTemplateCreateFile{}
	this.SourceType = sourceType
	return &this
}

// NewSpecTemplateCreateFileWithDefaults instantiates a new SpecTemplateCreateFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpecTemplateCreateFileWithDefaults() *SpecTemplateCreateFile {
	this := SpecTemplateCreateFile{}
	var sourceType string = "local"
	this.SourceType = sourceType
	return &this
}

// GetSourceType returns the SourceType field value
func (o *SpecTemplateCreateFile) GetSourceType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *SpecTemplateCreateFile) GetSourceTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *SpecTemplateCreateFile) SetSourceType(v string) {
	o.SourceType = v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *SpecTemplateCreateFile) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecTemplateCreateFile) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *SpecTemplateCreateFile) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *SpecTemplateCreateFile) SetContent(v string) {
	o.Content = &v
}

// GetContentPath returns the ContentPath field value if set, zero value otherwise.
func (o *SpecTemplateCreateFile) GetContentPath() string {
	if o == nil || o.ContentPath == nil {
		var ret string
		return ret
	}
	return *o.ContentPath
}

// GetContentPathOk returns a tuple with the ContentPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecTemplateCreateFile) GetContentPathOk() (*string, bool) {
	if o == nil || o.ContentPath == nil {
		return nil, false
	}
	return o.ContentPath, true
}

// HasContentPath returns a boolean if a field has been set.
func (o *SpecTemplateCreateFile) HasContentPath() bool {
	if o != nil && o.ContentPath != nil {
		return true
	}

	return false
}

// SetContentPath gets a reference to the given string and assigns it to the ContentPath field.
func (o *SpecTemplateCreateFile) SetContentPath(v string) {
	o.ContentPath = &v
}

// GetContentRef returns the ContentRef field value if set, zero value otherwise.
func (o *SpecTemplateCreateFile) GetContentRef() string {
	if o == nil || o.ContentRef == nil {
		var ret string
		return ret
	}
	return *o.ContentRef
}

// GetContentRefOk returns a tuple with the ContentRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecTemplateCreateFile) GetContentRefOk() (*string, bool) {
	if o == nil || o.ContentRef == nil {
		return nil, false
	}
	return o.ContentRef, true
}

// HasContentRef returns a boolean if a field has been set.
func (o *SpecTemplateCreateFile) HasContentRef() bool {
	if o != nil && o.ContentRef != nil {
		return true
	}

	return false
}

// SetContentRef gets a reference to the given string and assigns it to the ContentRef field.
func (o *SpecTemplateCreateFile) SetContentRef(v string) {
	o.ContentRef = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *SpecTemplateCreateFile) GetRepository() SpecTemplateCreateFileRepository {
	if o == nil || o.Repository == nil {
		var ret SpecTemplateCreateFileRepository
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecTemplateCreateFile) GetRepositoryOk() (*SpecTemplateCreateFileRepository, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *SpecTemplateCreateFile) HasRepository() bool {
	if o != nil && o.Repository != nil {
		return true
	}

	return false
}

// SetRepository gets a reference to the given SpecTemplateCreateFileRepository and assigns it to the Repository field.
func (o *SpecTemplateCreateFile) SetRepository(v SpecTemplateCreateFileRepository) {
	o.Repository = &v
}

func (o SpecTemplateCreateFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceType"] = o.SourceType
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.ContentPath != nil {
		toSerialize["contentPath"] = o.ContentPath
	}
	if o.ContentRef != nil {
		toSerialize["contentRef"] = o.ContentRef
	}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	return json.Marshal(toSerialize)
}

type NullableSpecTemplateCreateFile struct {
	value *SpecTemplateCreateFile
	isSet bool
}

func (v NullableSpecTemplateCreateFile) Get() *SpecTemplateCreateFile {
	return v.value
}

func (v *NullableSpecTemplateCreateFile) Set(val *SpecTemplateCreateFile) {
	v.value = val
	v.isSet = true
}

func (v NullableSpecTemplateCreateFile) IsSet() bool {
	return v.isSet
}

func (v *NullableSpecTemplateCreateFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpecTemplateCreateFile(val *SpecTemplateCreateFile) *NullableSpecTemplateCreateFile {
	return &NullableSpecTemplateCreateFile{value: val, isSet: true}
}

func (v NullableSpecTemplateCreateFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpecTemplateCreateFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


