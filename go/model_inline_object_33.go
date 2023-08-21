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

// InlineObject33 struct for InlineObject33
type InlineObject33 struct {
	// Payload for creating a new monitoring check
	Check OneOfcheckWebcheckSqlcheckSocketcheckElasticcheckPush `json:"check"`
}

// NewInlineObject33 instantiates a new InlineObject33 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject33(check OneOfcheckWebcheckSqlcheckSocketcheckElasticcheckPush, ) *InlineObject33 {
	this := InlineObject33{}
	this.Check = check
	return &this
}

// NewInlineObject33WithDefaults instantiates a new InlineObject33 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject33WithDefaults() *InlineObject33 {
	this := InlineObject33{}
	return &this
}

// GetCheck returns the Check field value
func (o *InlineObject33) GetCheck() OneOfcheckWebcheckSqlcheckSocketcheckElasticcheckPush {
	if o == nil  {
		var ret OneOfcheckWebcheckSqlcheckSocketcheckElasticcheckPush
		return ret
	}

	return o.Check
}

// GetCheckOk returns a tuple with the Check field value
// and a boolean to check if the value has been set.
func (o *InlineObject33) GetCheckOk() (*OneOfcheckWebcheckSqlcheckSocketcheckElasticcheckPush, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Check, true
}

// SetCheck sets field value
func (o *InlineObject33) SetCheck(v OneOfcheckWebcheckSqlcheckSocketcheckElasticcheckPush) {
	o.Check = v
}

func (o InlineObject33) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["check"] = o.Check
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject33 struct {
	value *InlineObject33
	isSet bool
}

func (v NullableInlineObject33) Get() *InlineObject33 {
	return v.value
}

func (v *NullableInlineObject33) Set(val *InlineObject33) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject33) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject33) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject33(val *InlineObject33) *NullableInlineObject33 {
	return &NullableInlineObject33{value: val, isSet: true}
}

func (v NullableInlineObject33) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject33) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


