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

// InlineObject31 struct for InlineObject31
type InlineObject31 struct {
	MonitorApp ApiMonitoringAppsIdMonitorApp `json:"monitorApp"`
}

// NewInlineObject31 instantiates a new InlineObject31 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject31(monitorApp ApiMonitoringAppsIdMonitorApp, ) *InlineObject31 {
	this := InlineObject31{}
	this.MonitorApp = monitorApp
	return &this
}

// NewInlineObject31WithDefaults instantiates a new InlineObject31 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject31WithDefaults() *InlineObject31 {
	this := InlineObject31{}
	return &this
}

// GetMonitorApp returns the MonitorApp field value
func (o *InlineObject31) GetMonitorApp() ApiMonitoringAppsIdMonitorApp {
	if o == nil  {
		var ret ApiMonitoringAppsIdMonitorApp
		return ret
	}

	return o.MonitorApp
}

// GetMonitorAppOk returns a tuple with the MonitorApp field value
// and a boolean to check if the value has been set.
func (o *InlineObject31) GetMonitorAppOk() (*ApiMonitoringAppsIdMonitorApp, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MonitorApp, true
}

// SetMonitorApp sets field value
func (o *InlineObject31) SetMonitorApp(v ApiMonitoringAppsIdMonitorApp) {
	o.MonitorApp = v
}

func (o InlineObject31) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["monitorApp"] = o.MonitorApp
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject31 struct {
	value *InlineObject31
	isSet bool
}

func (v NullableInlineObject31) Get() *InlineObject31 {
	return v.value
}

func (v *NullableInlineObject31) Set(val *InlineObject31) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject31) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject31) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject31(val *InlineObject31) *NullableInlineObject31 {
	return &NullableInlineObject31{value: val, isSet: true}
}

func (v NullableInlineObject31) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject31) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


