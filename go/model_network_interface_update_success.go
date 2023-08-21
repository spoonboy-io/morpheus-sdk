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

// NetworkInterfaceUpdateSuccess struct for NetworkInterfaceUpdateSuccess
type NetworkInterfaceUpdateSuccess struct {
	NetworkInterface *NetworkInterfaceUpdateSuccessNetworkInterface `json:"networkInterface,omitempty"`
	InterfaceType *string `json:"interfaceType,omitempty"`
	NetId *int64 `json:"netId,omitempty"`
	Server *NetworkInterfaceUpdateSuccessServer `json:"server,omitempty"`
}

// NewNetworkInterfaceUpdateSuccess instantiates a new NetworkInterfaceUpdateSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkInterfaceUpdateSuccess() *NetworkInterfaceUpdateSuccess {
	this := NetworkInterfaceUpdateSuccess{}
	return &this
}

// NewNetworkInterfaceUpdateSuccessWithDefaults instantiates a new NetworkInterfaceUpdateSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkInterfaceUpdateSuccessWithDefaults() *NetworkInterfaceUpdateSuccess {
	this := NetworkInterfaceUpdateSuccess{}
	return &this
}

// GetNetworkInterface returns the NetworkInterface field value if set, zero value otherwise.
func (o *NetworkInterfaceUpdateSuccess) GetNetworkInterface() NetworkInterfaceUpdateSuccessNetworkInterface {
	if o == nil || o.NetworkInterface == nil {
		var ret NetworkInterfaceUpdateSuccessNetworkInterface
		return ret
	}
	return *o.NetworkInterface
}

// GetNetworkInterfaceOk returns a tuple with the NetworkInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceUpdateSuccess) GetNetworkInterfaceOk() (*NetworkInterfaceUpdateSuccessNetworkInterface, bool) {
	if o == nil || o.NetworkInterface == nil {
		return nil, false
	}
	return o.NetworkInterface, true
}

// HasNetworkInterface returns a boolean if a field has been set.
func (o *NetworkInterfaceUpdateSuccess) HasNetworkInterface() bool {
	if o != nil && o.NetworkInterface != nil {
		return true
	}

	return false
}

// SetNetworkInterface gets a reference to the given NetworkInterfaceUpdateSuccessNetworkInterface and assigns it to the NetworkInterface field.
func (o *NetworkInterfaceUpdateSuccess) SetNetworkInterface(v NetworkInterfaceUpdateSuccessNetworkInterface) {
	o.NetworkInterface = &v
}

// GetInterfaceType returns the InterfaceType field value if set, zero value otherwise.
func (o *NetworkInterfaceUpdateSuccess) GetInterfaceType() string {
	if o == nil || o.InterfaceType == nil {
		var ret string
		return ret
	}
	return *o.InterfaceType
}

// GetInterfaceTypeOk returns a tuple with the InterfaceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceUpdateSuccess) GetInterfaceTypeOk() (*string, bool) {
	if o == nil || o.InterfaceType == nil {
		return nil, false
	}
	return o.InterfaceType, true
}

// HasInterfaceType returns a boolean if a field has been set.
func (o *NetworkInterfaceUpdateSuccess) HasInterfaceType() bool {
	if o != nil && o.InterfaceType != nil {
		return true
	}

	return false
}

// SetInterfaceType gets a reference to the given string and assigns it to the InterfaceType field.
func (o *NetworkInterfaceUpdateSuccess) SetInterfaceType(v string) {
	o.InterfaceType = &v
}

// GetNetId returns the NetId field value if set, zero value otherwise.
func (o *NetworkInterfaceUpdateSuccess) GetNetId() int64 {
	if o == nil || o.NetId == nil {
		var ret int64
		return ret
	}
	return *o.NetId
}

// GetNetIdOk returns a tuple with the NetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceUpdateSuccess) GetNetIdOk() (*int64, bool) {
	if o == nil || o.NetId == nil {
		return nil, false
	}
	return o.NetId, true
}

// HasNetId returns a boolean if a field has been set.
func (o *NetworkInterfaceUpdateSuccess) HasNetId() bool {
	if o != nil && o.NetId != nil {
		return true
	}

	return false
}

// SetNetId gets a reference to the given int64 and assigns it to the NetId field.
func (o *NetworkInterfaceUpdateSuccess) SetNetId(v int64) {
	o.NetId = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *NetworkInterfaceUpdateSuccess) GetServer() NetworkInterfaceUpdateSuccessServer {
	if o == nil || o.Server == nil {
		var ret NetworkInterfaceUpdateSuccessServer
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceUpdateSuccess) GetServerOk() (*NetworkInterfaceUpdateSuccessServer, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *NetworkInterfaceUpdateSuccess) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given NetworkInterfaceUpdateSuccessServer and assigns it to the Server field.
func (o *NetworkInterfaceUpdateSuccess) SetServer(v NetworkInterfaceUpdateSuccessServer) {
	o.Server = &v
}

func (o NetworkInterfaceUpdateSuccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkInterface != nil {
		toSerialize["networkInterface"] = o.NetworkInterface
	}
	if o.InterfaceType != nil {
		toSerialize["interfaceType"] = o.InterfaceType
	}
	if o.NetId != nil {
		toSerialize["netId"] = o.NetId
	}
	if o.Server != nil {
		toSerialize["server"] = o.Server
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkInterfaceUpdateSuccess struct {
	value *NetworkInterfaceUpdateSuccess
	isSet bool
}

func (v NullableNetworkInterfaceUpdateSuccess) Get() *NetworkInterfaceUpdateSuccess {
	return v.value
}

func (v *NullableNetworkInterfaceUpdateSuccess) Set(val *NetworkInterfaceUpdateSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkInterfaceUpdateSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkInterfaceUpdateSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkInterfaceUpdateSuccess(val *NetworkInterfaceUpdateSuccess) *NullableNetworkInterfaceUpdateSuccess {
	return &NullableNetworkInterfaceUpdateSuccess{value: val, isSet: true}
}

func (v NullableNetworkInterfaceUpdateSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkInterfaceUpdateSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

