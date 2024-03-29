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

// NetworkPoolServerCreateSolarWinds struct for NetworkPoolServerCreateSolarWinds
type NetworkPoolServerCreateSolarWinds struct {
	// Type Code (SolarWinds)
	Type string `json:"type"`
	// Name
	Name string `json:"name"`
	// Can be used to enable / disable the network pool server.
	Enabled *bool `json:"enabled,omitempty"`
	// URL
	ServiceUrl NullableString `json:"serviceUrl"`
	// Username
	ServiceUsername NullableString `json:"serviceUsername,omitempty"`
	// Password
	ServicePassword NullableString `json:"servicePassword,omitempty"`
	// Throttle Rate
	ServiceThrottleRate NullableInt64 `json:"serviceThrottleRate,omitempty"`
	// Disable SSL SNI Verification
	IgnoreSsl *bool `json:"ignoreSsl,omitempty"`
	Config *NetworkPoolServerCreateBluecatConfig `json:"config,omitempty"`
	Credential *NetworkPoolServerCreateBluecatCredential `json:"credential,omitempty"`
}

// NewNetworkPoolServerCreateSolarWinds instantiates a new NetworkPoolServerCreateSolarWinds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkPoolServerCreateSolarWinds(type_ string, name string, serviceUrl NullableString, ) *NetworkPoolServerCreateSolarWinds {
	this := NetworkPoolServerCreateSolarWinds{}
	this.Type = type_
	this.Name = name
	var enabled bool = true
	this.Enabled = &enabled
	this.ServiceUrl = serviceUrl
	var serviceThrottleRate int64 = 0
	this.ServiceThrottleRate = *NewNullableInt64(&serviceThrottleRate)
	var ignoreSsl bool = true
	this.IgnoreSsl = &ignoreSsl
	return &this
}

// NewNetworkPoolServerCreateSolarWindsWithDefaults instantiates a new NetworkPoolServerCreateSolarWinds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkPoolServerCreateSolarWindsWithDefaults() *NetworkPoolServerCreateSolarWinds {
	this := NetworkPoolServerCreateSolarWinds{}
	var enabled bool = true
	this.Enabled = &enabled
	var serviceThrottleRate int64 = 0
	this.ServiceThrottleRate = *NewNullableInt64(&serviceThrottleRate)
	var ignoreSsl bool = true
	this.IgnoreSsl = &ignoreSsl
	return &this
}

// GetType returns the Type field value
func (o *NetworkPoolServerCreateSolarWinds) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NetworkPoolServerCreateSolarWinds) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NetworkPoolServerCreateSolarWinds) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *NetworkPoolServerCreateSolarWinds) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworkPoolServerCreateSolarWinds) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworkPoolServerCreateSolarWinds) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworkPoolServerCreateSolarWinds) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServerCreateSolarWinds) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworkPoolServerCreateSolarWinds) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworkPoolServerCreateSolarWinds) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetServiceUrl returns the ServiceUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NetworkPoolServerCreateSolarWinds) GetServiceUrl() string {
	if o == nil || o.ServiceUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.ServiceUrl.Get()
}

// GetServiceUrlOk returns a tuple with the ServiceUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkPoolServerCreateSolarWinds) GetServiceUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceUrl.Get(), o.ServiceUrl.IsSet()
}

// SetServiceUrl sets field value
func (o *NetworkPoolServerCreateSolarWinds) SetServiceUrl(v string) {
	o.ServiceUrl.Set(&v)
}

// GetServiceUsername returns the ServiceUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkPoolServerCreateSolarWinds) GetServiceUsername() string {
	if o == nil || o.ServiceUsername.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceUsername.Get()
}

// GetServiceUsernameOk returns a tuple with the ServiceUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkPoolServerCreateSolarWinds) GetServiceUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceUsername.Get(), o.ServiceUsername.IsSet()
}

// HasServiceUsername returns a boolean if a field has been set.
func (o *NetworkPoolServerCreateSolarWinds) HasServiceUsername() bool {
	if o != nil && o.ServiceUsername.IsSet() {
		return true
	}

	return false
}

// SetServiceUsername gets a reference to the given NullableString and assigns it to the ServiceUsername field.
func (o *NetworkPoolServerCreateSolarWinds) SetServiceUsername(v string) {
	o.ServiceUsername.Set(&v)
}
// SetServiceUsernameNil sets the value for ServiceUsername to be an explicit nil
func (o *NetworkPoolServerCreateSolarWinds) SetServiceUsernameNil() {
	o.ServiceUsername.Set(nil)
}

// UnsetServiceUsername ensures that no value is present for ServiceUsername, not even an explicit nil
func (o *NetworkPoolServerCreateSolarWinds) UnsetServiceUsername() {
	o.ServiceUsername.Unset()
}

// GetServicePassword returns the ServicePassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkPoolServerCreateSolarWinds) GetServicePassword() string {
	if o == nil || o.ServicePassword.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServicePassword.Get()
}

// GetServicePasswordOk returns a tuple with the ServicePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkPoolServerCreateSolarWinds) GetServicePasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServicePassword.Get(), o.ServicePassword.IsSet()
}

// HasServicePassword returns a boolean if a field has been set.
func (o *NetworkPoolServerCreateSolarWinds) HasServicePassword() bool {
	if o != nil && o.ServicePassword.IsSet() {
		return true
	}

	return false
}

// SetServicePassword gets a reference to the given NullableString and assigns it to the ServicePassword field.
func (o *NetworkPoolServerCreateSolarWinds) SetServicePassword(v string) {
	o.ServicePassword.Set(&v)
}
// SetServicePasswordNil sets the value for ServicePassword to be an explicit nil
func (o *NetworkPoolServerCreateSolarWinds) SetServicePasswordNil() {
	o.ServicePassword.Set(nil)
}

// UnsetServicePassword ensures that no value is present for ServicePassword, not even an explicit nil
func (o *NetworkPoolServerCreateSolarWinds) UnsetServicePassword() {
	o.ServicePassword.Unset()
}

// GetServiceThrottleRate returns the ServiceThrottleRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkPoolServerCreateSolarWinds) GetServiceThrottleRate() int64 {
	if o == nil || o.ServiceThrottleRate.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ServiceThrottleRate.Get()
}

// GetServiceThrottleRateOk returns a tuple with the ServiceThrottleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkPoolServerCreateSolarWinds) GetServiceThrottleRateOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceThrottleRate.Get(), o.ServiceThrottleRate.IsSet()
}

// HasServiceThrottleRate returns a boolean if a field has been set.
func (o *NetworkPoolServerCreateSolarWinds) HasServiceThrottleRate() bool {
	if o != nil && o.ServiceThrottleRate.IsSet() {
		return true
	}

	return false
}

// SetServiceThrottleRate gets a reference to the given NullableInt64 and assigns it to the ServiceThrottleRate field.
func (o *NetworkPoolServerCreateSolarWinds) SetServiceThrottleRate(v int64) {
	o.ServiceThrottleRate.Set(&v)
}
// SetServiceThrottleRateNil sets the value for ServiceThrottleRate to be an explicit nil
func (o *NetworkPoolServerCreateSolarWinds) SetServiceThrottleRateNil() {
	o.ServiceThrottleRate.Set(nil)
}

// UnsetServiceThrottleRate ensures that no value is present for ServiceThrottleRate, not even an explicit nil
func (o *NetworkPoolServerCreateSolarWinds) UnsetServiceThrottleRate() {
	o.ServiceThrottleRate.Unset()
}

// GetIgnoreSsl returns the IgnoreSsl field value if set, zero value otherwise.
func (o *NetworkPoolServerCreateSolarWinds) GetIgnoreSsl() bool {
	if o == nil || o.IgnoreSsl == nil {
		var ret bool
		return ret
	}
	return *o.IgnoreSsl
}

// GetIgnoreSslOk returns a tuple with the IgnoreSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServerCreateSolarWinds) GetIgnoreSslOk() (*bool, bool) {
	if o == nil || o.IgnoreSsl == nil {
		return nil, false
	}
	return o.IgnoreSsl, true
}

// HasIgnoreSsl returns a boolean if a field has been set.
func (o *NetworkPoolServerCreateSolarWinds) HasIgnoreSsl() bool {
	if o != nil && o.IgnoreSsl != nil {
		return true
	}

	return false
}

// SetIgnoreSsl gets a reference to the given bool and assigns it to the IgnoreSsl field.
func (o *NetworkPoolServerCreateSolarWinds) SetIgnoreSsl(v bool) {
	o.IgnoreSsl = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *NetworkPoolServerCreateSolarWinds) GetConfig() NetworkPoolServerCreateBluecatConfig {
	if o == nil || o.Config == nil {
		var ret NetworkPoolServerCreateBluecatConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServerCreateSolarWinds) GetConfigOk() (*NetworkPoolServerCreateBluecatConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *NetworkPoolServerCreateSolarWinds) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given NetworkPoolServerCreateBluecatConfig and assigns it to the Config field.
func (o *NetworkPoolServerCreateSolarWinds) SetConfig(v NetworkPoolServerCreateBluecatConfig) {
	o.Config = &v
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *NetworkPoolServerCreateSolarWinds) GetCredential() NetworkPoolServerCreateBluecatCredential {
	if o == nil || o.Credential == nil {
		var ret NetworkPoolServerCreateBluecatCredential
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServerCreateSolarWinds) GetCredentialOk() (*NetworkPoolServerCreateBluecatCredential, bool) {
	if o == nil || o.Credential == nil {
		return nil, false
	}
	return o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *NetworkPoolServerCreateSolarWinds) HasCredential() bool {
	if o != nil && o.Credential != nil {
		return true
	}

	return false
}

// SetCredential gets a reference to the given NetworkPoolServerCreateBluecatCredential and assigns it to the Credential field.
func (o *NetworkPoolServerCreateSolarWinds) SetCredential(v NetworkPoolServerCreateBluecatCredential) {
	o.Credential = &v
}

func (o NetworkPoolServerCreateSolarWinds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["serviceUrl"] = o.ServiceUrl.Get()
	}
	if o.ServiceUsername.IsSet() {
		toSerialize["serviceUsername"] = o.ServiceUsername.Get()
	}
	if o.ServicePassword.IsSet() {
		toSerialize["servicePassword"] = o.ServicePassword.Get()
	}
	if o.ServiceThrottleRate.IsSet() {
		toSerialize["serviceThrottleRate"] = o.ServiceThrottleRate.Get()
	}
	if o.IgnoreSsl != nil {
		toSerialize["ignoreSsl"] = o.IgnoreSsl
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Credential != nil {
		toSerialize["credential"] = o.Credential
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkPoolServerCreateSolarWinds struct {
	value *NetworkPoolServerCreateSolarWinds
	isSet bool
}

func (v NullableNetworkPoolServerCreateSolarWinds) Get() *NetworkPoolServerCreateSolarWinds {
	return v.value
}

func (v *NullableNetworkPoolServerCreateSolarWinds) Set(val *NetworkPoolServerCreateSolarWinds) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkPoolServerCreateSolarWinds) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkPoolServerCreateSolarWinds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkPoolServerCreateSolarWinds(val *NetworkPoolServerCreateSolarWinds) *NullableNetworkPoolServerCreateSolarWinds {
	return &NullableNetworkPoolServerCreateSolarWinds{value: val, isSet: true}
}

func (v NullableNetworkPoolServerCreateSolarWinds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkPoolServerCreateSolarWinds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


