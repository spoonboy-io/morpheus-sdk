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

// LoadBalancerInstanceUpdate struct for LoadBalancerInstanceUpdate
type LoadBalancerInstanceUpdate struct {
	// VIP Name
	VipName *string `json:"vipName,omitempty"`
	// Description
	Description *string `json:"description,omitempty"`
	// VIP Address
	VipAddress *string `json:"vipAddress,omitempty"`
	// VIP Port
	VipPort *string `json:"vipPort,omitempty"`
	// VIP Protocol
	VipProtocol *string `json:"vipProtocol,omitempty"`
	// VIP Hostname
	VipHostname *string `json:"vipHostname,omitempty"`
	Pool *int64 `json:"pool,omitempty"`
	// SSL Client Certificate ID
	SslCert *int64 `json:"sslCert,omitempty"`
	// SSL Server Certificate ID
	SslServerCert *int64 `json:"sslServerCert,omitempty"`
	// Configuration object with parameters that vary by type.
	Config *map[string]interface{} `json:"config,omitempty"`
}

// NewLoadBalancerInstanceUpdate instantiates a new LoadBalancerInstanceUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerInstanceUpdate() *LoadBalancerInstanceUpdate {
	this := LoadBalancerInstanceUpdate{}
	return &this
}

// NewLoadBalancerInstanceUpdateWithDefaults instantiates a new LoadBalancerInstanceUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerInstanceUpdateWithDefaults() *LoadBalancerInstanceUpdate {
	this := LoadBalancerInstanceUpdate{}
	return &this
}

// GetVipName returns the VipName field value if set, zero value otherwise.
func (o *LoadBalancerInstanceUpdate) GetVipName() string {
	if o == nil || o.VipName == nil {
		var ret string
		return ret
	}
	return *o.VipName
}

// GetVipNameOk returns a tuple with the VipName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerInstanceUpdate) GetVipNameOk() (*string, bool) {
	if o == nil || o.VipName == nil {
		return nil, false
	}
	return o.VipName, true
}

// HasVipName returns a boolean if a field has been set.
func (o *LoadBalancerInstanceUpdate) HasVipName() bool {
	if o != nil && o.VipName != nil {
		return true
	}

	return false
}

// SetVipName gets a reference to the given string and assigns it to the VipName field.
func (o *LoadBalancerInstanceUpdate) SetVipName(v string) {
	o.VipName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LoadBalancerInstanceUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerInstanceUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LoadBalancerInstanceUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LoadBalancerInstanceUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetVipAddress returns the VipAddress field value if set, zero value otherwise.
func (o *LoadBalancerInstanceUpdate) GetVipAddress() string {
	if o == nil || o.VipAddress == nil {
		var ret string
		return ret
	}
	return *o.VipAddress
}

// GetVipAddressOk returns a tuple with the VipAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerInstanceUpdate) GetVipAddressOk() (*string, bool) {
	if o == nil || o.VipAddress == nil {
		return nil, false
	}
	return o.VipAddress, true
}

// HasVipAddress returns a boolean if a field has been set.
func (o *LoadBalancerInstanceUpdate) HasVipAddress() bool {
	if o != nil && o.VipAddress != nil {
		return true
	}

	return false
}

// SetVipAddress gets a reference to the given string and assigns it to the VipAddress field.
func (o *LoadBalancerInstanceUpdate) SetVipAddress(v string) {
	o.VipAddress = &v
}

// GetVipPort returns the VipPort field value if set, zero value otherwise.
func (o *LoadBalancerInstanceUpdate) GetVipPort() string {
	if o == nil || o.VipPort == nil {
		var ret string
		return ret
	}
	return *o.VipPort
}

// GetVipPortOk returns a tuple with the VipPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerInstanceUpdate) GetVipPortOk() (*string, bool) {
	if o == nil || o.VipPort == nil {
		return nil, false
	}
	return o.VipPort, true
}

// HasVipPort returns a boolean if a field has been set.
func (o *LoadBalancerInstanceUpdate) HasVipPort() bool {
	if o != nil && o.VipPort != nil {
		return true
	}

	return false
}

// SetVipPort gets a reference to the given string and assigns it to the VipPort field.
func (o *LoadBalancerInstanceUpdate) SetVipPort(v string) {
	o.VipPort = &v
}

// GetVipProtocol returns the VipProtocol field value if set, zero value otherwise.
func (o *LoadBalancerInstanceUpdate) GetVipProtocol() string {
	if o == nil || o.VipProtocol == nil {
		var ret string
		return ret
	}
	return *o.VipProtocol
}

// GetVipProtocolOk returns a tuple with the VipProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerInstanceUpdate) GetVipProtocolOk() (*string, bool) {
	if o == nil || o.VipProtocol == nil {
		return nil, false
	}
	return o.VipProtocol, true
}

// HasVipProtocol returns a boolean if a field has been set.
func (o *LoadBalancerInstanceUpdate) HasVipProtocol() bool {
	if o != nil && o.VipProtocol != nil {
		return true
	}

	return false
}

// SetVipProtocol gets a reference to the given string and assigns it to the VipProtocol field.
func (o *LoadBalancerInstanceUpdate) SetVipProtocol(v string) {
	o.VipProtocol = &v
}

// GetVipHostname returns the VipHostname field value if set, zero value otherwise.
func (o *LoadBalancerInstanceUpdate) GetVipHostname() string {
	if o == nil || o.VipHostname == nil {
		var ret string
		return ret
	}
	return *o.VipHostname
}

// GetVipHostnameOk returns a tuple with the VipHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerInstanceUpdate) GetVipHostnameOk() (*string, bool) {
	if o == nil || o.VipHostname == nil {
		return nil, false
	}
	return o.VipHostname, true
}

// HasVipHostname returns a boolean if a field has been set.
func (o *LoadBalancerInstanceUpdate) HasVipHostname() bool {
	if o != nil && o.VipHostname != nil {
		return true
	}

	return false
}

// SetVipHostname gets a reference to the given string and assigns it to the VipHostname field.
func (o *LoadBalancerInstanceUpdate) SetVipHostname(v string) {
	o.VipHostname = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *LoadBalancerInstanceUpdate) GetPool() int64 {
	if o == nil || o.Pool == nil {
		var ret int64
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerInstanceUpdate) GetPoolOk() (*int64, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *LoadBalancerInstanceUpdate) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given int64 and assigns it to the Pool field.
func (o *LoadBalancerInstanceUpdate) SetPool(v int64) {
	o.Pool = &v
}

// GetSslCert returns the SslCert field value if set, zero value otherwise.
func (o *LoadBalancerInstanceUpdate) GetSslCert() int64 {
	if o == nil || o.SslCert == nil {
		var ret int64
		return ret
	}
	return *o.SslCert
}

// GetSslCertOk returns a tuple with the SslCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerInstanceUpdate) GetSslCertOk() (*int64, bool) {
	if o == nil || o.SslCert == nil {
		return nil, false
	}
	return o.SslCert, true
}

// HasSslCert returns a boolean if a field has been set.
func (o *LoadBalancerInstanceUpdate) HasSslCert() bool {
	if o != nil && o.SslCert != nil {
		return true
	}

	return false
}

// SetSslCert gets a reference to the given int64 and assigns it to the SslCert field.
func (o *LoadBalancerInstanceUpdate) SetSslCert(v int64) {
	o.SslCert = &v
}

// GetSslServerCert returns the SslServerCert field value if set, zero value otherwise.
func (o *LoadBalancerInstanceUpdate) GetSslServerCert() int64 {
	if o == nil || o.SslServerCert == nil {
		var ret int64
		return ret
	}
	return *o.SslServerCert
}

// GetSslServerCertOk returns a tuple with the SslServerCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerInstanceUpdate) GetSslServerCertOk() (*int64, bool) {
	if o == nil || o.SslServerCert == nil {
		return nil, false
	}
	return o.SslServerCert, true
}

// HasSslServerCert returns a boolean if a field has been set.
func (o *LoadBalancerInstanceUpdate) HasSslServerCert() bool {
	if o != nil && o.SslServerCert != nil {
		return true
	}

	return false
}

// SetSslServerCert gets a reference to the given int64 and assigns it to the SslServerCert field.
func (o *LoadBalancerInstanceUpdate) SetSslServerCert(v int64) {
	o.SslServerCert = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *LoadBalancerInstanceUpdate) GetConfig() map[string]interface{} {
	if o == nil || o.Config == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerInstanceUpdate) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *LoadBalancerInstanceUpdate) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *LoadBalancerInstanceUpdate) SetConfig(v map[string]interface{}) {
	o.Config = &v
}

func (o LoadBalancerInstanceUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VipName != nil {
		toSerialize["vipName"] = o.VipName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.VipAddress != nil {
		toSerialize["vipAddress"] = o.VipAddress
	}
	if o.VipPort != nil {
		toSerialize["vipPort"] = o.VipPort
	}
	if o.VipProtocol != nil {
		toSerialize["vipProtocol"] = o.VipProtocol
	}
	if o.VipHostname != nil {
		toSerialize["vipHostname"] = o.VipHostname
	}
	if o.Pool != nil {
		toSerialize["pool"] = o.Pool
	}
	if o.SslCert != nil {
		toSerialize["sslCert"] = o.SslCert
	}
	if o.SslServerCert != nil {
		toSerialize["sslServerCert"] = o.SslServerCert
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableLoadBalancerInstanceUpdate struct {
	value *LoadBalancerInstanceUpdate
	isSet bool
}

func (v NullableLoadBalancerInstanceUpdate) Get() *LoadBalancerInstanceUpdate {
	return v.value
}

func (v *NullableLoadBalancerInstanceUpdate) Set(val *LoadBalancerInstanceUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerInstanceUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerInstanceUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerInstanceUpdate(val *LoadBalancerInstanceUpdate) *NullableLoadBalancerInstanceUpdate {
	return &NullableLoadBalancerInstanceUpdate{value: val, isSet: true}
}

func (v NullableLoadBalancerInstanceUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerInstanceUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


