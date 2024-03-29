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

// InlineObject154 struct for InlineObject154
type InlineObject154 struct {
	RuleGroup *NetworkFirewallRuleGroupCreate `json:"ruleGroup,omitempty"`
}

// NewInlineObject154 instantiates a new InlineObject154 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject154() *InlineObject154 {
	this := InlineObject154{}
	return &this
}

// NewInlineObject154WithDefaults instantiates a new InlineObject154 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject154WithDefaults() *InlineObject154 {
	this := InlineObject154{}
	return &this
}

// GetRuleGroup returns the RuleGroup field value if set, zero value otherwise.
func (o *InlineObject154) GetRuleGroup() NetworkFirewallRuleGroupCreate {
	if o == nil || o.RuleGroup == nil {
		var ret NetworkFirewallRuleGroupCreate
		return ret
	}
	return *o.RuleGroup
}

// GetRuleGroupOk returns a tuple with the RuleGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject154) GetRuleGroupOk() (*NetworkFirewallRuleGroupCreate, bool) {
	if o == nil || o.RuleGroup == nil {
		return nil, false
	}
	return o.RuleGroup, true
}

// HasRuleGroup returns a boolean if a field has been set.
func (o *InlineObject154) HasRuleGroup() bool {
	if o != nil && o.RuleGroup != nil {
		return true
	}

	return false
}

// SetRuleGroup gets a reference to the given NetworkFirewallRuleGroupCreate and assigns it to the RuleGroup field.
func (o *InlineObject154) SetRuleGroup(v NetworkFirewallRuleGroupCreate) {
	o.RuleGroup = &v
}

func (o InlineObject154) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RuleGroup != nil {
		toSerialize["ruleGroup"] = o.RuleGroup
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject154 struct {
	value *InlineObject154
	isSet bool
}

func (v NullableInlineObject154) Get() *InlineObject154 {
	return v.value
}

func (v *NullableInlineObject154) Set(val *InlineObject154) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject154) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject154) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject154(val *InlineObject154) *NullableInlineObject154 {
	return &NullableInlineObject154{value: val, isSet: true}
}

func (v NullableInlineObject154) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject154) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


