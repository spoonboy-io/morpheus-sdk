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

// InlineObject197 struct for InlineObject197
type InlineObject197 struct {
	PreseedScript *PreseedScriptsCreate `json:"preseedScript,omitempty"`
}

// NewInlineObject197 instantiates a new InlineObject197 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject197() *InlineObject197 {
	this := InlineObject197{}
	return &this
}

// NewInlineObject197WithDefaults instantiates a new InlineObject197 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject197WithDefaults() *InlineObject197 {
	this := InlineObject197{}
	return &this
}

// GetPreseedScript returns the PreseedScript field value if set, zero value otherwise.
func (o *InlineObject197) GetPreseedScript() PreseedScriptsCreate {
	if o == nil || o.PreseedScript == nil {
		var ret PreseedScriptsCreate
		return ret
	}
	return *o.PreseedScript
}

// GetPreseedScriptOk returns a tuple with the PreseedScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject197) GetPreseedScriptOk() (*PreseedScriptsCreate, bool) {
	if o == nil || o.PreseedScript == nil {
		return nil, false
	}
	return o.PreseedScript, true
}

// HasPreseedScript returns a boolean if a field has been set.
func (o *InlineObject197) HasPreseedScript() bool {
	if o != nil && o.PreseedScript != nil {
		return true
	}

	return false
}

// SetPreseedScript gets a reference to the given PreseedScriptsCreate and assigns it to the PreseedScript field.
func (o *InlineObject197) SetPreseedScript(v PreseedScriptsCreate) {
	o.PreseedScript = &v
}

func (o InlineObject197) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PreseedScript != nil {
		toSerialize["preseedScript"] = o.PreseedScript
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject197 struct {
	value *InlineObject197
	isSet bool
}

func (v NullableInlineObject197) Get() *InlineObject197 {
	return v.value
}

func (v *NullableInlineObject197) Set(val *InlineObject197) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject197) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject197) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject197(val *InlineObject197) *NullableInlineObject197 {
	return &NullableInlineObject197{value: val, isSet: true}
}

func (v NullableInlineObject197) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject197) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


