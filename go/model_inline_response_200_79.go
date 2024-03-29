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

// InlineResponse20079 struct for InlineResponse20079
type InlineResponse20079 struct {
	LoadBalancerMonitor *InlineResponse20079LoadBalancerMonitor `json:"loadBalancerMonitor,omitempty"`
}

// NewInlineResponse20079 instantiates a new InlineResponse20079 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20079() *InlineResponse20079 {
	this := InlineResponse20079{}
	return &this
}

// NewInlineResponse20079WithDefaults instantiates a new InlineResponse20079 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20079WithDefaults() *InlineResponse20079 {
	this := InlineResponse20079{}
	return &this
}

// GetLoadBalancerMonitor returns the LoadBalancerMonitor field value if set, zero value otherwise.
func (o *InlineResponse20079) GetLoadBalancerMonitor() InlineResponse20079LoadBalancerMonitor {
	if o == nil || o.LoadBalancerMonitor == nil {
		var ret InlineResponse20079LoadBalancerMonitor
		return ret
	}
	return *o.LoadBalancerMonitor
}

// GetLoadBalancerMonitorOk returns a tuple with the LoadBalancerMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetLoadBalancerMonitorOk() (*InlineResponse20079LoadBalancerMonitor, bool) {
	if o == nil || o.LoadBalancerMonitor == nil {
		return nil, false
	}
	return o.LoadBalancerMonitor, true
}

// HasLoadBalancerMonitor returns a boolean if a field has been set.
func (o *InlineResponse20079) HasLoadBalancerMonitor() bool {
	if o != nil && o.LoadBalancerMonitor != nil {
		return true
	}

	return false
}

// SetLoadBalancerMonitor gets a reference to the given InlineResponse20079LoadBalancerMonitor and assigns it to the LoadBalancerMonitor field.
func (o *InlineResponse20079) SetLoadBalancerMonitor(v InlineResponse20079LoadBalancerMonitor) {
	o.LoadBalancerMonitor = &v
}

func (o InlineResponse20079) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LoadBalancerMonitor != nil {
		toSerialize["loadBalancerMonitor"] = o.LoadBalancerMonitor
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20079 struct {
	value *InlineResponse20079
	isSet bool
}

func (v NullableInlineResponse20079) Get() *InlineResponse20079 {
	return v.value
}

func (v *NullableInlineResponse20079) Set(val *InlineResponse20079) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20079) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20079) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20079(val *InlineResponse20079) *NullableInlineResponse20079 {
	return &NullableInlineResponse20079{value: val, isSet: true}
}

func (v NullableInlineResponse20079) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20079) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


