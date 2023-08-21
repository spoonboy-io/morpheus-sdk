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

// InstanceConnectionInfo struct for InstanceConnectionInfo
type InstanceConnectionInfo struct {
	Ip *string `json:"ip,omitempty"`
	Port NullableInt32 `json:"port,omitempty"`
}

// NewInstanceConnectionInfo instantiates a new InstanceConnectionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceConnectionInfo() *InstanceConnectionInfo {
	this := InstanceConnectionInfo{}
	return &this
}

// NewInstanceConnectionInfoWithDefaults instantiates a new InstanceConnectionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceConnectionInfoWithDefaults() *InstanceConnectionInfo {
	this := InstanceConnectionInfo{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *InstanceConnectionInfo) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceConnectionInfo) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *InstanceConnectionInfo) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *InstanceConnectionInfo) SetIp(v string) {
	o.Ip = &v
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceConnectionInfo) GetPort() int32 {
	if o == nil || o.Port.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceConnectionInfo) GetPortOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *InstanceConnectionInfo) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt32 and assigns it to the Port field.
func (o *InstanceConnectionInfo) SetPort(v int32) {
	o.Port.Set(&v)
}
// SetPortNil sets the value for Port to be an explicit nil
func (o *InstanceConnectionInfo) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *InstanceConnectionInfo) UnsetPort() {
	o.Port.Unset()
}

func (o InstanceConnectionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInstanceConnectionInfo struct {
	value *InstanceConnectionInfo
	isSet bool
}

func (v NullableInstanceConnectionInfo) Get() *InstanceConnectionInfo {
	return v.value
}

func (v *NullableInstanceConnectionInfo) Set(val *InstanceConnectionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceConnectionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceConnectionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceConnectionInfo(val *InstanceConnectionInfo) *NullableInstanceConnectionInfo {
	return &NullableInstanceConnectionInfo{value: val, isSet: true}
}

func (v NullableInstanceConnectionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceConnectionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

