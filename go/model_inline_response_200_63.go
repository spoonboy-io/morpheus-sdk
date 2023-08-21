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

// InlineResponse20063 struct for InlineResponse20063
type InlineResponse20063 struct {
	Objects *[]IntegrationObject `json:"objects,omitempty"`
}

// NewInlineResponse20063 instantiates a new InlineResponse20063 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20063() *InlineResponse20063 {
	this := InlineResponse20063{}
	return &this
}

// NewInlineResponse20063WithDefaults instantiates a new InlineResponse20063 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20063WithDefaults() *InlineResponse20063 {
	this := InlineResponse20063{}
	return &this
}

// GetObjects returns the Objects field value if set, zero value otherwise.
func (o *InlineResponse20063) GetObjects() []IntegrationObject {
	if o == nil || o.Objects == nil {
		var ret []IntegrationObject
		return ret
	}
	return *o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20063) GetObjectsOk() (*[]IntegrationObject, bool) {
	if o == nil || o.Objects == nil {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *InlineResponse20063) HasObjects() bool {
	if o != nil && o.Objects != nil {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []IntegrationObject and assigns it to the Objects field.
func (o *InlineResponse20063) SetObjects(v []IntegrationObject) {
	o.Objects = &v
}

func (o InlineResponse20063) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20063 struct {
	value *InlineResponse20063
	isSet bool
}

func (v NullableInlineResponse20063) Get() *InlineResponse20063 {
	return v.value
}

func (v *NullableInlineResponse20063) Set(val *InlineResponse20063) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20063) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20063) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20063(val *InlineResponse20063) *NullableInlineResponse20063 {
	return &NullableInlineResponse20063{value: val, isSet: true}
}

func (v NullableInlineResponse20063) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20063) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

