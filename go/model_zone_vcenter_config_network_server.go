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

// ZoneVcenterConfigNetworkServer struct for ZoneVcenterConfigNetworkServer
type ZoneVcenterConfigNetworkServer struct {
	Id *string `json:"id,omitempty"`
}

// NewZoneVcenterConfigNetworkServer instantiates a new ZoneVcenterConfigNetworkServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneVcenterConfigNetworkServer() *ZoneVcenterConfigNetworkServer {
	this := ZoneVcenterConfigNetworkServer{}
	return &this
}

// NewZoneVcenterConfigNetworkServerWithDefaults instantiates a new ZoneVcenterConfigNetworkServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneVcenterConfigNetworkServerWithDefaults() *ZoneVcenterConfigNetworkServer {
	this := ZoneVcenterConfigNetworkServer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ZoneVcenterConfigNetworkServer) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneVcenterConfigNetworkServer) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ZoneVcenterConfigNetworkServer) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ZoneVcenterConfigNetworkServer) SetId(v string) {
	o.Id = &v
}

func (o ZoneVcenterConfigNetworkServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableZoneVcenterConfigNetworkServer struct {
	value *ZoneVcenterConfigNetworkServer
	isSet bool
}

func (v NullableZoneVcenterConfigNetworkServer) Get() *ZoneVcenterConfigNetworkServer {
	return v.value
}

func (v *NullableZoneVcenterConfigNetworkServer) Set(val *ZoneVcenterConfigNetworkServer) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneVcenterConfigNetworkServer) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneVcenterConfigNetworkServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneVcenterConfigNetworkServer(val *ZoneVcenterConfigNetworkServer) *NullableZoneVcenterConfigNetworkServer {
	return &NullableZoneVcenterConfigNetworkServer{value: val, isSet: true}
}

func (v NullableZoneVcenterConfigNetworkServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneVcenterConfigNetworkServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


