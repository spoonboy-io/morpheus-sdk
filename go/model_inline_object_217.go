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

// InlineObject217 struct for InlineObject217
type InlineObject217 struct {
	Rule ApiSecurityGroupsIdRulesSgIdRule `json:"rule"`
}

// NewInlineObject217 instantiates a new InlineObject217 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject217(rule ApiSecurityGroupsIdRulesSgIdRule, ) *InlineObject217 {
	this := InlineObject217{}
	this.Rule = rule
	return &this
}

// NewInlineObject217WithDefaults instantiates a new InlineObject217 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject217WithDefaults() *InlineObject217 {
	this := InlineObject217{}
	return &this
}

// GetRule returns the Rule field value
func (o *InlineObject217) GetRule() ApiSecurityGroupsIdRulesSgIdRule {
	if o == nil  {
		var ret ApiSecurityGroupsIdRulesSgIdRule
		return ret
	}

	return o.Rule
}

// GetRuleOk returns a tuple with the Rule field value
// and a boolean to check if the value has been set.
func (o *InlineObject217) GetRuleOk() (*ApiSecurityGroupsIdRulesSgIdRule, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Rule, true
}

// SetRule sets field value
func (o *InlineObject217) SetRule(v ApiSecurityGroupsIdRulesSgIdRule) {
	o.Rule = v
}

func (o InlineObject217) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rule"] = o.Rule
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject217 struct {
	value *InlineObject217
	isSet bool
}

func (v NullableInlineObject217) Get() *InlineObject217 {
	return v.value
}

func (v *NullableInlineObject217) Set(val *InlineObject217) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject217) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject217) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject217(val *InlineObject217) *NullableInlineObject217 {
	return &NullableInlineObject217{value: val, isSet: true}
}

func (v NullableInlineObject217) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject217) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

