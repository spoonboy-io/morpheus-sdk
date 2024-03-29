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

// InlineObject229 struct for InlineObject229
type InlineObject229 struct {
	ServicePlan ApiServicePlansIdServicePlan `json:"servicePlan"`
}

// NewInlineObject229 instantiates a new InlineObject229 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject229(servicePlan ApiServicePlansIdServicePlan, ) *InlineObject229 {
	this := InlineObject229{}
	this.ServicePlan = servicePlan
	return &this
}

// NewInlineObject229WithDefaults instantiates a new InlineObject229 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject229WithDefaults() *InlineObject229 {
	this := InlineObject229{}
	return &this
}

// GetServicePlan returns the ServicePlan field value
func (o *InlineObject229) GetServicePlan() ApiServicePlansIdServicePlan {
	if o == nil  {
		var ret ApiServicePlansIdServicePlan
		return ret
	}

	return o.ServicePlan
}

// GetServicePlanOk returns a tuple with the ServicePlan field value
// and a boolean to check if the value has been set.
func (o *InlineObject229) GetServicePlanOk() (*ApiServicePlansIdServicePlan, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServicePlan, true
}

// SetServicePlan sets field value
func (o *InlineObject229) SetServicePlan(v ApiServicePlansIdServicePlan) {
	o.ServicePlan = v
}

func (o InlineObject229) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["servicePlan"] = o.ServicePlan
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject229 struct {
	value *InlineObject229
	isSet bool
}

func (v NullableInlineObject229) Get() *InlineObject229 {
	return v.value
}

func (v *NullableInlineObject229) Set(val *InlineObject229) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject229) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject229) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject229(val *InlineObject229) *NullableInlineObject229 {
	return &NullableInlineObject229{value: val, isSet: true}
}

func (v NullableInlineObject229) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject229) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


