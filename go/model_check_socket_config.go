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

// CheckSocketConfig struct for CheckSocketConfig
type CheckSocketConfig struct {
	// Hostname or IP address of the socket server
	Host string `json:"host"`
	// TCP port
	Port string `json:"port"`
	// Connection string you might want to send to the service
	Send string `json:"send"`
	// Response from the service to match against
	ResponseMatch string `json:"responseMatch"`
	CheckUser *string `json:"checkUser,omitempty"`
	TextCheckOn *string `json:"textCheckOn,omitempty"`
	CheckPassword *string `json:"checkPassword,omitempty"`
	WebTextMatch *string `json:"webTextMatch,omitempty"`
	CheckPasswordHash *string `json:"checkPasswordHash,omitempty"`
	// Set to on to turn on tunneling
	TunnelOn *string `json:"tunnelOn,omitempty"`
	// Hostname or IP address of the proxy host
	SshHost *string `json:"sshHost,omitempty"`
	// Port for SSH on the proxy host, defaults to 22
	SshPort *int64 `json:"sshPort,omitempty"`
	// SSH user on the proxy host to login as
	SshUser *string `json:"sshUser,omitempty"`
	// Password for user, if not using key based authentication
	SshPassword *string `json:"sshPassword,omitempty"`
}

// NewCheckSocketConfig instantiates a new CheckSocketConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckSocketConfig(host string, port string, send string, responseMatch string, ) *CheckSocketConfig {
	this := CheckSocketConfig{}
	this.Host = host
	this.Port = port
	this.Send = send
	this.ResponseMatch = responseMatch
	var tunnelOn string = "off"
	this.TunnelOn = &tunnelOn
	return &this
}

// NewCheckSocketConfigWithDefaults instantiates a new CheckSocketConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckSocketConfigWithDefaults() *CheckSocketConfig {
	this := CheckSocketConfig{}
	var tunnelOn string = "off"
	this.TunnelOn = &tunnelOn
	return &this
}

// GetHost returns the Host field value
func (o *CheckSocketConfig) GetHost() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *CheckSocketConfig) GetHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *CheckSocketConfig) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value
func (o *CheckSocketConfig) GetPort() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *CheckSocketConfig) GetPortOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *CheckSocketConfig) SetPort(v string) {
	o.Port = v
}

// GetSend returns the Send field value
func (o *CheckSocketConfig) GetSend() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Send
}

// GetSendOk returns a tuple with the Send field value
// and a boolean to check if the value has been set.
func (o *CheckSocketConfig) GetSendOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Send, true
}

// SetSend sets field value
func (o *CheckSocketConfig) SetSend(v string) {
	o.Send = v
}

// GetResponseMatch returns the ResponseMatch field value
func (o *CheckSocketConfig) GetResponseMatch() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ResponseMatch
}

// GetResponseMatchOk returns a tuple with the ResponseMatch field value
// and a boolean to check if the value has been set.
func (o *CheckSocketConfig) GetResponseMatchOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResponseMatch, true
}

// SetResponseMatch sets field value
func (o *CheckSocketConfig) SetResponseMatch(v string) {
	o.ResponseMatch = v
}

// GetCheckUser returns the CheckUser field value if set, zero value otherwise.
func (o *CheckSocketConfig) GetCheckUser() string {
	if o == nil || o.CheckUser == nil {
		var ret string
		return ret
	}
	return *o.CheckUser
}

// GetCheckUserOk returns a tuple with the CheckUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSocketConfig) GetCheckUserOk() (*string, bool) {
	if o == nil || o.CheckUser == nil {
		return nil, false
	}
	return o.CheckUser, true
}

// HasCheckUser returns a boolean if a field has been set.
func (o *CheckSocketConfig) HasCheckUser() bool {
	if o != nil && o.CheckUser != nil {
		return true
	}

	return false
}

// SetCheckUser gets a reference to the given string and assigns it to the CheckUser field.
func (o *CheckSocketConfig) SetCheckUser(v string) {
	o.CheckUser = &v
}

// GetTextCheckOn returns the TextCheckOn field value if set, zero value otherwise.
func (o *CheckSocketConfig) GetTextCheckOn() string {
	if o == nil || o.TextCheckOn == nil {
		var ret string
		return ret
	}
	return *o.TextCheckOn
}

// GetTextCheckOnOk returns a tuple with the TextCheckOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSocketConfig) GetTextCheckOnOk() (*string, bool) {
	if o == nil || o.TextCheckOn == nil {
		return nil, false
	}
	return o.TextCheckOn, true
}

// HasTextCheckOn returns a boolean if a field has been set.
func (o *CheckSocketConfig) HasTextCheckOn() bool {
	if o != nil && o.TextCheckOn != nil {
		return true
	}

	return false
}

// SetTextCheckOn gets a reference to the given string and assigns it to the TextCheckOn field.
func (o *CheckSocketConfig) SetTextCheckOn(v string) {
	o.TextCheckOn = &v
}

// GetCheckPassword returns the CheckPassword field value if set, zero value otherwise.
func (o *CheckSocketConfig) GetCheckPassword() string {
	if o == nil || o.CheckPassword == nil {
		var ret string
		return ret
	}
	return *o.CheckPassword
}

// GetCheckPasswordOk returns a tuple with the CheckPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSocketConfig) GetCheckPasswordOk() (*string, bool) {
	if o == nil || o.CheckPassword == nil {
		return nil, false
	}
	return o.CheckPassword, true
}

// HasCheckPassword returns a boolean if a field has been set.
func (o *CheckSocketConfig) HasCheckPassword() bool {
	if o != nil && o.CheckPassword != nil {
		return true
	}

	return false
}

// SetCheckPassword gets a reference to the given string and assigns it to the CheckPassword field.
func (o *CheckSocketConfig) SetCheckPassword(v string) {
	o.CheckPassword = &v
}

// GetWebTextMatch returns the WebTextMatch field value if set, zero value otherwise.
func (o *CheckSocketConfig) GetWebTextMatch() string {
	if o == nil || o.WebTextMatch == nil {
		var ret string
		return ret
	}
	return *o.WebTextMatch
}

// GetWebTextMatchOk returns a tuple with the WebTextMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSocketConfig) GetWebTextMatchOk() (*string, bool) {
	if o == nil || o.WebTextMatch == nil {
		return nil, false
	}
	return o.WebTextMatch, true
}

// HasWebTextMatch returns a boolean if a field has been set.
func (o *CheckSocketConfig) HasWebTextMatch() bool {
	if o != nil && o.WebTextMatch != nil {
		return true
	}

	return false
}

// SetWebTextMatch gets a reference to the given string and assigns it to the WebTextMatch field.
func (o *CheckSocketConfig) SetWebTextMatch(v string) {
	o.WebTextMatch = &v
}

// GetCheckPasswordHash returns the CheckPasswordHash field value if set, zero value otherwise.
func (o *CheckSocketConfig) GetCheckPasswordHash() string {
	if o == nil || o.CheckPasswordHash == nil {
		var ret string
		return ret
	}
	return *o.CheckPasswordHash
}

// GetCheckPasswordHashOk returns a tuple with the CheckPasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSocketConfig) GetCheckPasswordHashOk() (*string, bool) {
	if o == nil || o.CheckPasswordHash == nil {
		return nil, false
	}
	return o.CheckPasswordHash, true
}

// HasCheckPasswordHash returns a boolean if a field has been set.
func (o *CheckSocketConfig) HasCheckPasswordHash() bool {
	if o != nil && o.CheckPasswordHash != nil {
		return true
	}

	return false
}

// SetCheckPasswordHash gets a reference to the given string and assigns it to the CheckPasswordHash field.
func (o *CheckSocketConfig) SetCheckPasswordHash(v string) {
	o.CheckPasswordHash = &v
}

// GetTunnelOn returns the TunnelOn field value if set, zero value otherwise.
func (o *CheckSocketConfig) GetTunnelOn() string {
	if o == nil || o.TunnelOn == nil {
		var ret string
		return ret
	}
	return *o.TunnelOn
}

// GetTunnelOnOk returns a tuple with the TunnelOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSocketConfig) GetTunnelOnOk() (*string, bool) {
	if o == nil || o.TunnelOn == nil {
		return nil, false
	}
	return o.TunnelOn, true
}

// HasTunnelOn returns a boolean if a field has been set.
func (o *CheckSocketConfig) HasTunnelOn() bool {
	if o != nil && o.TunnelOn != nil {
		return true
	}

	return false
}

// SetTunnelOn gets a reference to the given string and assigns it to the TunnelOn field.
func (o *CheckSocketConfig) SetTunnelOn(v string) {
	o.TunnelOn = &v
}

// GetSshHost returns the SshHost field value if set, zero value otherwise.
func (o *CheckSocketConfig) GetSshHost() string {
	if o == nil || o.SshHost == nil {
		var ret string
		return ret
	}
	return *o.SshHost
}

// GetSshHostOk returns a tuple with the SshHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSocketConfig) GetSshHostOk() (*string, bool) {
	if o == nil || o.SshHost == nil {
		return nil, false
	}
	return o.SshHost, true
}

// HasSshHost returns a boolean if a field has been set.
func (o *CheckSocketConfig) HasSshHost() bool {
	if o != nil && o.SshHost != nil {
		return true
	}

	return false
}

// SetSshHost gets a reference to the given string and assigns it to the SshHost field.
func (o *CheckSocketConfig) SetSshHost(v string) {
	o.SshHost = &v
}

// GetSshPort returns the SshPort field value if set, zero value otherwise.
func (o *CheckSocketConfig) GetSshPort() int64 {
	if o == nil || o.SshPort == nil {
		var ret int64
		return ret
	}
	return *o.SshPort
}

// GetSshPortOk returns a tuple with the SshPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSocketConfig) GetSshPortOk() (*int64, bool) {
	if o == nil || o.SshPort == nil {
		return nil, false
	}
	return o.SshPort, true
}

// HasSshPort returns a boolean if a field has been set.
func (o *CheckSocketConfig) HasSshPort() bool {
	if o != nil && o.SshPort != nil {
		return true
	}

	return false
}

// SetSshPort gets a reference to the given int64 and assigns it to the SshPort field.
func (o *CheckSocketConfig) SetSshPort(v int64) {
	o.SshPort = &v
}

// GetSshUser returns the SshUser field value if set, zero value otherwise.
func (o *CheckSocketConfig) GetSshUser() string {
	if o == nil || o.SshUser == nil {
		var ret string
		return ret
	}
	return *o.SshUser
}

// GetSshUserOk returns a tuple with the SshUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSocketConfig) GetSshUserOk() (*string, bool) {
	if o == nil || o.SshUser == nil {
		return nil, false
	}
	return o.SshUser, true
}

// HasSshUser returns a boolean if a field has been set.
func (o *CheckSocketConfig) HasSshUser() bool {
	if o != nil && o.SshUser != nil {
		return true
	}

	return false
}

// SetSshUser gets a reference to the given string and assigns it to the SshUser field.
func (o *CheckSocketConfig) SetSshUser(v string) {
	o.SshUser = &v
}

// GetSshPassword returns the SshPassword field value if set, zero value otherwise.
func (o *CheckSocketConfig) GetSshPassword() string {
	if o == nil || o.SshPassword == nil {
		var ret string
		return ret
	}
	return *o.SshPassword
}

// GetSshPasswordOk returns a tuple with the SshPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSocketConfig) GetSshPasswordOk() (*string, bool) {
	if o == nil || o.SshPassword == nil {
		return nil, false
	}
	return o.SshPassword, true
}

// HasSshPassword returns a boolean if a field has been set.
func (o *CheckSocketConfig) HasSshPassword() bool {
	if o != nil && o.SshPassword != nil {
		return true
	}

	return false
}

// SetSshPassword gets a reference to the given string and assigns it to the SshPassword field.
func (o *CheckSocketConfig) SetSshPassword(v string) {
	o.SshPassword = &v
}

func (o CheckSocketConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["host"] = o.Host
	}
	if true {
		toSerialize["port"] = o.Port
	}
	if true {
		toSerialize["send"] = o.Send
	}
	if true {
		toSerialize["responseMatch"] = o.ResponseMatch
	}
	if o.CheckUser != nil {
		toSerialize["checkUser"] = o.CheckUser
	}
	if o.TextCheckOn != nil {
		toSerialize["textCheckOn"] = o.TextCheckOn
	}
	if o.CheckPassword != nil {
		toSerialize["checkPassword"] = o.CheckPassword
	}
	if o.WebTextMatch != nil {
		toSerialize["webTextMatch"] = o.WebTextMatch
	}
	if o.CheckPasswordHash != nil {
		toSerialize["checkPasswordHash"] = o.CheckPasswordHash
	}
	if o.TunnelOn != nil {
		toSerialize["tunnelOn"] = o.TunnelOn
	}
	if o.SshHost != nil {
		toSerialize["sshHost"] = o.SshHost
	}
	if o.SshPort != nil {
		toSerialize["sshPort"] = o.SshPort
	}
	if o.SshUser != nil {
		toSerialize["sshUser"] = o.SshUser
	}
	if o.SshPassword != nil {
		toSerialize["sshPassword"] = o.SshPassword
	}
	return json.Marshal(toSerialize)
}

type NullableCheckSocketConfig struct {
	value *CheckSocketConfig
	isSet bool
}

func (v NullableCheckSocketConfig) Get() *CheckSocketConfig {
	return v.value
}

func (v *NullableCheckSocketConfig) Set(val *CheckSocketConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckSocketConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckSocketConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckSocketConfig(val *CheckSocketConfig) *NullableCheckSocketConfig {
	return &NullableCheckSocketConfig{value: val, isSet: true}
}

func (v NullableCheckSocketConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckSocketConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


