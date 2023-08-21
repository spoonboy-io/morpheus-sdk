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

// NetworkRouterFirewallRuleCreate For a full list of available firewall rule options, see ruleOptionTypes in the specific Network Router Type detail
type NetworkRouterFirewallRuleCreate struct {
	// Firewall rule name
	Name string `json:"name"`
	// Can be used to enable / disable the rule (true, false). Default is on
	Enabled *bool `json:"enabled,omitempty"`
	// Priority for rule
	Priority *int64 `json:"priority,omitempty"`
}

// NewNetworkRouterFirewallRuleCreate instantiates a new NetworkRouterFirewallRuleCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkRouterFirewallRuleCreate(name string, ) *NetworkRouterFirewallRuleCreate {
	this := NetworkRouterFirewallRuleCreate{}
	this.Name = name
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// NewNetworkRouterFirewallRuleCreateWithDefaults instantiates a new NetworkRouterFirewallRuleCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkRouterFirewallRuleCreateWithDefaults() *NetworkRouterFirewallRuleCreate {
	this := NetworkRouterFirewallRuleCreate{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetName returns the Name field value
func (o *NetworkRouterFirewallRuleCreate) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworkRouterFirewallRuleCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworkRouterFirewallRuleCreate) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworkRouterFirewallRuleCreate) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRouterFirewallRuleCreate) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworkRouterFirewallRuleCreate) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworkRouterFirewallRuleCreate) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *NetworkRouterFirewallRuleCreate) GetPriority() int64 {
	if o == nil || o.Priority == nil {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRouterFirewallRuleCreate) GetPriorityOk() (*int64, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *NetworkRouterFirewallRuleCreate) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *NetworkRouterFirewallRuleCreate) SetPriority(v int64) {
	o.Priority = &v
}

func (o NetworkRouterFirewallRuleCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkRouterFirewallRuleCreate struct {
	value *NetworkRouterFirewallRuleCreate
	isSet bool
}

func (v NullableNetworkRouterFirewallRuleCreate) Get() *NetworkRouterFirewallRuleCreate {
	return v.value
}

func (v *NullableNetworkRouterFirewallRuleCreate) Set(val *NetworkRouterFirewallRuleCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkRouterFirewallRuleCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkRouterFirewallRuleCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkRouterFirewallRuleCreate(val *NetworkRouterFirewallRuleCreate) *NullableNetworkRouterFirewallRuleCreate {
	return &NullableNetworkRouterFirewallRuleCreate{value: val, isSet: true}
}

func (v NullableNetworkRouterFirewallRuleCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkRouterFirewallRuleCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


