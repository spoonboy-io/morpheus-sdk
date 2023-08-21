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

// InlineObject35 struct for InlineObject35
type InlineObject35 struct {
	// Payload for updating a monitoring check
	Check OneOfcheckWebcheckSqlcheckSocketcheckElasticcheckPush `json:"check"`
}

// NewInlineObject35 instantiates a new InlineObject35 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject35(check OneOfcheckWebcheckSqlcheckSocketcheckElasticcheckPush, ) *InlineObject35 {
	this := InlineObject35{}
	this.Check = check
	return &this
}

// NewInlineObject35WithDefaults instantiates a new InlineObject35 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject35WithDefaults() *InlineObject35 {
	this := InlineObject35{}
	return &this
}

// GetCheck returns the Check field value
func (o *InlineObject35) GetCheck() OneOfcheckWebcheckSqlcheckSocketcheckElasticcheckPush {
	if o == nil  {
		var ret OneOfcheckWebcheckSqlcheckSocketcheckElasticcheckPush
		return ret
	}

	return o.Check
}

// GetCheckOk returns a tuple with the Check field value
// and a boolean to check if the value has been set.
func (o *InlineObject35) GetCheckOk() (*OneOfcheckWebcheckSqlcheckSocketcheckElasticcheckPush, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Check, true
}

// SetCheck sets field value
func (o *InlineObject35) SetCheck(v OneOfcheckWebcheckSqlcheckSocketcheckElasticcheckPush) {
	o.Check = v
}

func (o InlineObject35) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["check"] = o.Check
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject35 struct {
	value *InlineObject35
	isSet bool
}

func (v NullableInlineObject35) Get() *InlineObject35 {
	return v.value
}

func (v *NullableInlineObject35) Set(val *InlineObject35) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject35) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject35) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject35(val *InlineObject35) *NullableInlineObject35 {
	return &NullableInlineObject35{value: val, isSet: true}
}

func (v NullableInlineObject35) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject35) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


