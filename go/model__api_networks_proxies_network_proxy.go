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

// ApiNetworksProxiesNetworkProxy struct for ApiNetworksProxiesNetworkProxy
type ApiNetworksProxiesNetworkProxy struct {
	// Name
	Name *string `json:"name,omitempty"`
	// Proxy Host
	ProxyHost *string `json:"proxyHost,omitempty"`
	// Proxy Port
	ProxyPort *string `json:"proxyPort,omitempty"`
	// Proxy Username
	ProxyUser *string `json:"proxyUser,omitempty"`
	// Proxy Password
	ProxyPassword *string `json:"proxyPassword,omitempty"`
	// Proxy Domain
	ProxyDomain *string `json:"proxyDomain,omitempty"`
	// Proxy Workstation
	ProxyWorkstation *string `json:"proxyWorkstation,omitempty"`
	// Visibility
	Visibility *string `json:"visibility,omitempty"`
	Account *ApiNetworksProxiesNetworkProxyAccount `json:"account,omitempty"`
}

// NewApiNetworksProxiesNetworkProxy instantiates a new ApiNetworksProxiesNetworkProxy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiNetworksProxiesNetworkProxy() *ApiNetworksProxiesNetworkProxy {
	this := ApiNetworksProxiesNetworkProxy{}
	var visibility string = "private"
	this.Visibility = &visibility
	return &this
}

// NewApiNetworksProxiesNetworkProxyWithDefaults instantiates a new ApiNetworksProxiesNetworkProxy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiNetworksProxiesNetworkProxyWithDefaults() *ApiNetworksProxiesNetworkProxy {
	this := ApiNetworksProxiesNetworkProxy{}
	var visibility string = "private"
	this.Visibility = &visibility
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiNetworksProxiesNetworkProxy) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNetworksProxiesNetworkProxy) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiNetworksProxiesNetworkProxy) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiNetworksProxiesNetworkProxy) SetName(v string) {
	o.Name = &v
}

// GetProxyHost returns the ProxyHost field value if set, zero value otherwise.
func (o *ApiNetworksProxiesNetworkProxy) GetProxyHost() string {
	if o == nil || o.ProxyHost == nil {
		var ret string
		return ret
	}
	return *o.ProxyHost
}

// GetProxyHostOk returns a tuple with the ProxyHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNetworksProxiesNetworkProxy) GetProxyHostOk() (*string, bool) {
	if o == nil || o.ProxyHost == nil {
		return nil, false
	}
	return o.ProxyHost, true
}

// HasProxyHost returns a boolean if a field has been set.
func (o *ApiNetworksProxiesNetworkProxy) HasProxyHost() bool {
	if o != nil && o.ProxyHost != nil {
		return true
	}

	return false
}

// SetProxyHost gets a reference to the given string and assigns it to the ProxyHost field.
func (o *ApiNetworksProxiesNetworkProxy) SetProxyHost(v string) {
	o.ProxyHost = &v
}

// GetProxyPort returns the ProxyPort field value if set, zero value otherwise.
func (o *ApiNetworksProxiesNetworkProxy) GetProxyPort() string {
	if o == nil || o.ProxyPort == nil {
		var ret string
		return ret
	}
	return *o.ProxyPort
}

// GetProxyPortOk returns a tuple with the ProxyPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNetworksProxiesNetworkProxy) GetProxyPortOk() (*string, bool) {
	if o == nil || o.ProxyPort == nil {
		return nil, false
	}
	return o.ProxyPort, true
}

// HasProxyPort returns a boolean if a field has been set.
func (o *ApiNetworksProxiesNetworkProxy) HasProxyPort() bool {
	if o != nil && o.ProxyPort != nil {
		return true
	}

	return false
}

// SetProxyPort gets a reference to the given string and assigns it to the ProxyPort field.
func (o *ApiNetworksProxiesNetworkProxy) SetProxyPort(v string) {
	o.ProxyPort = &v
}

// GetProxyUser returns the ProxyUser field value if set, zero value otherwise.
func (o *ApiNetworksProxiesNetworkProxy) GetProxyUser() string {
	if o == nil || o.ProxyUser == nil {
		var ret string
		return ret
	}
	return *o.ProxyUser
}

// GetProxyUserOk returns a tuple with the ProxyUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNetworksProxiesNetworkProxy) GetProxyUserOk() (*string, bool) {
	if o == nil || o.ProxyUser == nil {
		return nil, false
	}
	return o.ProxyUser, true
}

// HasProxyUser returns a boolean if a field has been set.
func (o *ApiNetworksProxiesNetworkProxy) HasProxyUser() bool {
	if o != nil && o.ProxyUser != nil {
		return true
	}

	return false
}

// SetProxyUser gets a reference to the given string and assigns it to the ProxyUser field.
func (o *ApiNetworksProxiesNetworkProxy) SetProxyUser(v string) {
	o.ProxyUser = &v
}

// GetProxyPassword returns the ProxyPassword field value if set, zero value otherwise.
func (o *ApiNetworksProxiesNetworkProxy) GetProxyPassword() string {
	if o == nil || o.ProxyPassword == nil {
		var ret string
		return ret
	}
	return *o.ProxyPassword
}

// GetProxyPasswordOk returns a tuple with the ProxyPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNetworksProxiesNetworkProxy) GetProxyPasswordOk() (*string, bool) {
	if o == nil || o.ProxyPassword == nil {
		return nil, false
	}
	return o.ProxyPassword, true
}

// HasProxyPassword returns a boolean if a field has been set.
func (o *ApiNetworksProxiesNetworkProxy) HasProxyPassword() bool {
	if o != nil && o.ProxyPassword != nil {
		return true
	}

	return false
}

// SetProxyPassword gets a reference to the given string and assigns it to the ProxyPassword field.
func (o *ApiNetworksProxiesNetworkProxy) SetProxyPassword(v string) {
	o.ProxyPassword = &v
}

// GetProxyDomain returns the ProxyDomain field value if set, zero value otherwise.
func (o *ApiNetworksProxiesNetworkProxy) GetProxyDomain() string {
	if o == nil || o.ProxyDomain == nil {
		var ret string
		return ret
	}
	return *o.ProxyDomain
}

// GetProxyDomainOk returns a tuple with the ProxyDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNetworksProxiesNetworkProxy) GetProxyDomainOk() (*string, bool) {
	if o == nil || o.ProxyDomain == nil {
		return nil, false
	}
	return o.ProxyDomain, true
}

// HasProxyDomain returns a boolean if a field has been set.
func (o *ApiNetworksProxiesNetworkProxy) HasProxyDomain() bool {
	if o != nil && o.ProxyDomain != nil {
		return true
	}

	return false
}

// SetProxyDomain gets a reference to the given string and assigns it to the ProxyDomain field.
func (o *ApiNetworksProxiesNetworkProxy) SetProxyDomain(v string) {
	o.ProxyDomain = &v
}

// GetProxyWorkstation returns the ProxyWorkstation field value if set, zero value otherwise.
func (o *ApiNetworksProxiesNetworkProxy) GetProxyWorkstation() string {
	if o == nil || o.ProxyWorkstation == nil {
		var ret string
		return ret
	}
	return *o.ProxyWorkstation
}

// GetProxyWorkstationOk returns a tuple with the ProxyWorkstation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNetworksProxiesNetworkProxy) GetProxyWorkstationOk() (*string, bool) {
	if o == nil || o.ProxyWorkstation == nil {
		return nil, false
	}
	return o.ProxyWorkstation, true
}

// HasProxyWorkstation returns a boolean if a field has been set.
func (o *ApiNetworksProxiesNetworkProxy) HasProxyWorkstation() bool {
	if o != nil && o.ProxyWorkstation != nil {
		return true
	}

	return false
}

// SetProxyWorkstation gets a reference to the given string and assigns it to the ProxyWorkstation field.
func (o *ApiNetworksProxiesNetworkProxy) SetProxyWorkstation(v string) {
	o.ProxyWorkstation = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *ApiNetworksProxiesNetworkProxy) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNetworksProxiesNetworkProxy) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *ApiNetworksProxiesNetworkProxy) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *ApiNetworksProxiesNetworkProxy) SetVisibility(v string) {
	o.Visibility = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ApiNetworksProxiesNetworkProxy) GetAccount() ApiNetworksProxiesNetworkProxyAccount {
	if o == nil || o.Account == nil {
		var ret ApiNetworksProxiesNetworkProxyAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNetworksProxiesNetworkProxy) GetAccountOk() (*ApiNetworksProxiesNetworkProxyAccount, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ApiNetworksProxiesNetworkProxy) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given ApiNetworksProxiesNetworkProxyAccount and assigns it to the Account field.
func (o *ApiNetworksProxiesNetworkProxy) SetAccount(v ApiNetworksProxiesNetworkProxyAccount) {
	o.Account = &v
}

func (o ApiNetworksProxiesNetworkProxy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ProxyHost != nil {
		toSerialize["proxyHost"] = o.ProxyHost
	}
	if o.ProxyPort != nil {
		toSerialize["proxyPort"] = o.ProxyPort
	}
	if o.ProxyUser != nil {
		toSerialize["proxyUser"] = o.ProxyUser
	}
	if o.ProxyPassword != nil {
		toSerialize["proxyPassword"] = o.ProxyPassword
	}
	if o.ProxyDomain != nil {
		toSerialize["proxyDomain"] = o.ProxyDomain
	}
	if o.ProxyWorkstation != nil {
		toSerialize["proxyWorkstation"] = o.ProxyWorkstation
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	return json.Marshal(toSerialize)
}

type NullableApiNetworksProxiesNetworkProxy struct {
	value *ApiNetworksProxiesNetworkProxy
	isSet bool
}

func (v NullableApiNetworksProxiesNetworkProxy) Get() *ApiNetworksProxiesNetworkProxy {
	return v.value
}

func (v *NullableApiNetworksProxiesNetworkProxy) Set(val *ApiNetworksProxiesNetworkProxy) {
	v.value = val
	v.isSet = true
}

func (v NullableApiNetworksProxiesNetworkProxy) IsSet() bool {
	return v.isSet
}

func (v *NullableApiNetworksProxiesNetworkProxy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiNetworksProxiesNetworkProxy(val *ApiNetworksProxiesNetworkProxy) *NullableApiNetworksProxiesNetworkProxy {
	return &NullableApiNetworksProxiesNetworkProxy{value: val, isSet: true}
}

func (v NullableApiNetworksProxiesNetworkProxy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiNetworksProxiesNetworkProxy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

