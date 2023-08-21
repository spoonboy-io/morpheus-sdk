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

// InlineResponse20035Actions struct for InlineResponse20035Actions
type InlineResponse20035Actions struct {
	Name *string `json:"name,omitempty"`
	Code *string `json:"code,omitempty"`
}

// NewInlineResponse20035Actions instantiates a new InlineResponse20035Actions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20035Actions() *InlineResponse20035Actions {
	this := InlineResponse20035Actions{}
	return &this
}

// NewInlineResponse20035ActionsWithDefaults instantiates a new InlineResponse20035Actions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20035ActionsWithDefaults() *InlineResponse20035Actions {
	this := InlineResponse20035Actions{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20035Actions) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Actions) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20035Actions) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20035Actions) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InlineResponse20035Actions) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Actions) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InlineResponse20035Actions) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *InlineResponse20035Actions) SetCode(v string) {
	o.Code = &v
}

func (o InlineResponse20035Actions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20035Actions struct {
	value *InlineResponse20035Actions
	isSet bool
}

func (v NullableInlineResponse20035Actions) Get() *InlineResponse20035Actions {
	return v.value
}

func (v *NullableInlineResponse20035Actions) Set(val *InlineResponse20035Actions) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20035Actions) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20035Actions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20035Actions(val *InlineResponse20035Actions) *NullableInlineResponse20035Actions {
	return &NullableInlineResponse20035Actions{value: val, isSet: true}
}

func (v NullableInlineResponse20035Actions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20035Actions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

