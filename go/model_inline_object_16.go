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

// InlineObject16 struct for InlineObject16
type InlineObject16 struct {
	Backup AnyOfbackupTypeInstancebackupTypeServerbackupTypeProvider `json:"backup"`
}

// NewInlineObject16 instantiates a new InlineObject16 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject16(backup AnyOfbackupTypeInstancebackupTypeServerbackupTypeProvider, ) *InlineObject16 {
	this := InlineObject16{}
	this.Backup = backup
	return &this
}

// NewInlineObject16WithDefaults instantiates a new InlineObject16 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject16WithDefaults() *InlineObject16 {
	this := InlineObject16{}
	return &this
}

// GetBackup returns the Backup field value
func (o *InlineObject16) GetBackup() AnyOfbackupTypeInstancebackupTypeServerbackupTypeProvider {
	if o == nil  {
		var ret AnyOfbackupTypeInstancebackupTypeServerbackupTypeProvider
		return ret
	}

	return o.Backup
}

// GetBackupOk returns a tuple with the Backup field value
// and a boolean to check if the value has been set.
func (o *InlineObject16) GetBackupOk() (*AnyOfbackupTypeInstancebackupTypeServerbackupTypeProvider, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Backup, true
}

// SetBackup sets field value
func (o *InlineObject16) SetBackup(v AnyOfbackupTypeInstancebackupTypeServerbackupTypeProvider) {
	o.Backup = v
}

func (o InlineObject16) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["backup"] = o.Backup
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject16 struct {
	value *InlineObject16
	isSet bool
}

func (v NullableInlineObject16) Get() *InlineObject16 {
	return v.value
}

func (v *NullableInlineObject16) Set(val *InlineObject16) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject16) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject16) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject16(val *InlineObject16) *NullableInlineObject16 {
	return &NullableInlineObject16{value: val, isSet: true}
}

func (v NullableInlineObject16) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject16) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

