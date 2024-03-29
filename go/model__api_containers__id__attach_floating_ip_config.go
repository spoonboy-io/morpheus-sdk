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

// ApiContainersIdAttachFloatingIpConfig struct for ApiContainersIdAttachFloatingIpConfig
type ApiContainersIdAttachFloatingIpConfig struct {
	// The Floating IP identifier in the format: \"ip-ID\" or \"pool-ID\".  The Options API /api/options/openStack/openstackFloatingIpOptions?containerId={{containerId}} can be used to see which options are available. 
	OsExternalNetworkId string `json:"osExternalNetworkId"`
	// Bandwidth (Mbit/s) Only the following cloud types support this parameter: Huawei, OpenTelekom 
	FloatingIpBandwidth *float32 `json:"floatingIpBandwidth,omitempty"`
}

// NewApiContainersIdAttachFloatingIpConfig instantiates a new ApiContainersIdAttachFloatingIpConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiContainersIdAttachFloatingIpConfig(osExternalNetworkId string, ) *ApiContainersIdAttachFloatingIpConfig {
	this := ApiContainersIdAttachFloatingIpConfig{}
	this.OsExternalNetworkId = osExternalNetworkId
	return &this
}

// NewApiContainersIdAttachFloatingIpConfigWithDefaults instantiates a new ApiContainersIdAttachFloatingIpConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiContainersIdAttachFloatingIpConfigWithDefaults() *ApiContainersIdAttachFloatingIpConfig {
	this := ApiContainersIdAttachFloatingIpConfig{}
	return &this
}

// GetOsExternalNetworkId returns the OsExternalNetworkId field value
func (o *ApiContainersIdAttachFloatingIpConfig) GetOsExternalNetworkId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.OsExternalNetworkId
}

// GetOsExternalNetworkIdOk returns a tuple with the OsExternalNetworkId field value
// and a boolean to check if the value has been set.
func (o *ApiContainersIdAttachFloatingIpConfig) GetOsExternalNetworkIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OsExternalNetworkId, true
}

// SetOsExternalNetworkId sets field value
func (o *ApiContainersIdAttachFloatingIpConfig) SetOsExternalNetworkId(v string) {
	o.OsExternalNetworkId = v
}

// GetFloatingIpBandwidth returns the FloatingIpBandwidth field value if set, zero value otherwise.
func (o *ApiContainersIdAttachFloatingIpConfig) GetFloatingIpBandwidth() float32 {
	if o == nil || o.FloatingIpBandwidth == nil {
		var ret float32
		return ret
	}
	return *o.FloatingIpBandwidth
}

// GetFloatingIpBandwidthOk returns a tuple with the FloatingIpBandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiContainersIdAttachFloatingIpConfig) GetFloatingIpBandwidthOk() (*float32, bool) {
	if o == nil || o.FloatingIpBandwidth == nil {
		return nil, false
	}
	return o.FloatingIpBandwidth, true
}

// HasFloatingIpBandwidth returns a boolean if a field has been set.
func (o *ApiContainersIdAttachFloatingIpConfig) HasFloatingIpBandwidth() bool {
	if o != nil && o.FloatingIpBandwidth != nil {
		return true
	}

	return false
}

// SetFloatingIpBandwidth gets a reference to the given float32 and assigns it to the FloatingIpBandwidth field.
func (o *ApiContainersIdAttachFloatingIpConfig) SetFloatingIpBandwidth(v float32) {
	o.FloatingIpBandwidth = &v
}

func (o ApiContainersIdAttachFloatingIpConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["osExternalNetworkId"] = o.OsExternalNetworkId
	}
	if o.FloatingIpBandwidth != nil {
		toSerialize["floatingIpBandwidth"] = o.FloatingIpBandwidth
	}
	return json.Marshal(toSerialize)
}

type NullableApiContainersIdAttachFloatingIpConfig struct {
	value *ApiContainersIdAttachFloatingIpConfig
	isSet bool
}

func (v NullableApiContainersIdAttachFloatingIpConfig) Get() *ApiContainersIdAttachFloatingIpConfig {
	return v.value
}

func (v *NullableApiContainersIdAttachFloatingIpConfig) Set(val *ApiContainersIdAttachFloatingIpConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableApiContainersIdAttachFloatingIpConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableApiContainersIdAttachFloatingIpConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiContainersIdAttachFloatingIpConfig(val *ApiContainersIdAttachFloatingIpConfig) *NullableApiContainersIdAttachFloatingIpConfig {
	return &NullableApiContainersIdAttachFloatingIpConfig{value: val, isSet: true}
}

func (v NullableApiContainersIdAttachFloatingIpConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiContainersIdAttachFloatingIpConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


