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

// InlineResponse20023 struct for InlineResponse20023
type InlineResponse20023 struct {
	Folder *ZoneFolder `json:"folder,omitempty"`
}

// NewInlineResponse20023 instantiates a new InlineResponse20023 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20023() *InlineResponse20023 {
	this := InlineResponse20023{}
	return &this
}

// NewInlineResponse20023WithDefaults instantiates a new InlineResponse20023 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20023WithDefaults() *InlineResponse20023 {
	this := InlineResponse20023{}
	return &this
}

// GetFolder returns the Folder field value if set, zero value otherwise.
func (o *InlineResponse20023) GetFolder() ZoneFolder {
	if o == nil || o.Folder == nil {
		var ret ZoneFolder
		return ret
	}
	return *o.Folder
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20023) GetFolderOk() (*ZoneFolder, bool) {
	if o == nil || o.Folder == nil {
		return nil, false
	}
	return o.Folder, true
}

// HasFolder returns a boolean if a field has been set.
func (o *InlineResponse20023) HasFolder() bool {
	if o != nil && o.Folder != nil {
		return true
	}

	return false
}

// SetFolder gets a reference to the given ZoneFolder and assigns it to the Folder field.
func (o *InlineResponse20023) SetFolder(v ZoneFolder) {
	o.Folder = &v
}

func (o InlineResponse20023) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Folder != nil {
		toSerialize["folder"] = o.Folder
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20023 struct {
	value *InlineResponse20023
	isSet bool
}

func (v NullableInlineResponse20023) Get() *InlineResponse20023 {
	return v.value
}

func (v *NullableInlineResponse20023) Set(val *InlineResponse20023) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20023) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20023) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20023(val *InlineResponse20023) *NullableInlineResponse20023 {
	return &NullableInlineResponse20023{value: val, isSet: true}
}

func (v NullableInlineResponse20023) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20023) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

