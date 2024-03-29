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

// InlineObject228 struct for InlineObject228
type InlineObject228 struct {
	ServicePlan ApiServicePlansServicePlan `json:"servicePlan"`
}

// NewInlineObject228 instantiates a new InlineObject228 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject228(servicePlan ApiServicePlansServicePlan, ) *InlineObject228 {
	this := InlineObject228{}
	this.ServicePlan = servicePlan
	return &this
}

// NewInlineObject228WithDefaults instantiates a new InlineObject228 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject228WithDefaults() *InlineObject228 {
	this := InlineObject228{}
	return &this
}

// GetServicePlan returns the ServicePlan field value
func (o *InlineObject228) GetServicePlan() ApiServicePlansServicePlan {
	if o == nil  {
		var ret ApiServicePlansServicePlan
		return ret
	}

	return o.ServicePlan
}

// GetServicePlanOk returns a tuple with the ServicePlan field value
// and a boolean to check if the value has been set.
func (o *InlineObject228) GetServicePlanOk() (*ApiServicePlansServicePlan, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServicePlan, true
}

// SetServicePlan sets field value
func (o *InlineObject228) SetServicePlan(v ApiServicePlansServicePlan) {
	o.ServicePlan = v
}

func (o InlineObject228) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["servicePlan"] = o.ServicePlan
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject228 struct {
	value *InlineObject228
	isSet bool
}

func (v NullableInlineObject228) Get() *InlineObject228 {
	return v.value
}

func (v *NullableInlineObject228) Set(val *InlineObject228) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject228) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject228) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject228(val *InlineObject228) *NullableInlineObject228 {
	return &NullableInlineObject228{value: val, isSet: true}
}

func (v NullableInlineObject228) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject228) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


