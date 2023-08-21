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

// CheckSshConfig struct for CheckSshConfig
type CheckSshConfig struct {
	SshPort *int64 `json:"sshPort,omitempty"`
	CheckUser *string `json:"checkUser,omitempty"`
	TunnelOn *string `json:"tunnelOn,omitempty"`
	TextCheckOn *string `json:"textCheckOn,omitempty"`
	CheckPassword *string `json:"checkPassword,omitempty"`
	SshHost *string `json:"sshHost,omitempty"`
	SshUser *string `json:"sshUser,omitempty"`
	WebTextMatch *string `json:"webTextMatch,omitempty"`
	CheckPasswordHash *string `json:"checkPasswordHash,omitempty"`
}

// NewCheckSshConfig instantiates a new CheckSshConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckSshConfig() *CheckSshConfig {
	this := CheckSshConfig{}
	return &this
}

// NewCheckSshConfigWithDefaults instantiates a new CheckSshConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckSshConfigWithDefaults() *CheckSshConfig {
	this := CheckSshConfig{}
	return &this
}

// GetSshPort returns the SshPort field value if set, zero value otherwise.
func (o *CheckSshConfig) GetSshPort() int64 {
	if o == nil || o.SshPort == nil {
		var ret int64
		return ret
	}
	return *o.SshPort
}

// GetSshPortOk returns a tuple with the SshPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSshConfig) GetSshPortOk() (*int64, bool) {
	if o == nil || o.SshPort == nil {
		return nil, false
	}
	return o.SshPort, true
}

// HasSshPort returns a boolean if a field has been set.
func (o *CheckSshConfig) HasSshPort() bool {
	if o != nil && o.SshPort != nil {
		return true
	}

	return false
}

// SetSshPort gets a reference to the given int64 and assigns it to the SshPort field.
func (o *CheckSshConfig) SetSshPort(v int64) {
	o.SshPort = &v
}

// GetCheckUser returns the CheckUser field value if set, zero value otherwise.
func (o *CheckSshConfig) GetCheckUser() string {
	if o == nil || o.CheckUser == nil {
		var ret string
		return ret
	}
	return *o.CheckUser
}

// GetCheckUserOk returns a tuple with the CheckUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSshConfig) GetCheckUserOk() (*string, bool) {
	if o == nil || o.CheckUser == nil {
		return nil, false
	}
	return o.CheckUser, true
}

// HasCheckUser returns a boolean if a field has been set.
func (o *CheckSshConfig) HasCheckUser() bool {
	if o != nil && o.CheckUser != nil {
		return true
	}

	return false
}

// SetCheckUser gets a reference to the given string and assigns it to the CheckUser field.
func (o *CheckSshConfig) SetCheckUser(v string) {
	o.CheckUser = &v
}

// GetTunnelOn returns the TunnelOn field value if set, zero value otherwise.
func (o *CheckSshConfig) GetTunnelOn() string {
	if o == nil || o.TunnelOn == nil {
		var ret string
		return ret
	}
	return *o.TunnelOn
}

// GetTunnelOnOk returns a tuple with the TunnelOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSshConfig) GetTunnelOnOk() (*string, bool) {
	if o == nil || o.TunnelOn == nil {
		return nil, false
	}
	return o.TunnelOn, true
}

// HasTunnelOn returns a boolean if a field has been set.
func (o *CheckSshConfig) HasTunnelOn() bool {
	if o != nil && o.TunnelOn != nil {
		return true
	}

	return false
}

// SetTunnelOn gets a reference to the given string and assigns it to the TunnelOn field.
func (o *CheckSshConfig) SetTunnelOn(v string) {
	o.TunnelOn = &v
}

// GetTextCheckOn returns the TextCheckOn field value if set, zero value otherwise.
func (o *CheckSshConfig) GetTextCheckOn() string {
	if o == nil || o.TextCheckOn == nil {
		var ret string
		return ret
	}
	return *o.TextCheckOn
}

// GetTextCheckOnOk returns a tuple with the TextCheckOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSshConfig) GetTextCheckOnOk() (*string, bool) {
	if o == nil || o.TextCheckOn == nil {
		return nil, false
	}
	return o.TextCheckOn, true
}

// HasTextCheckOn returns a boolean if a field has been set.
func (o *CheckSshConfig) HasTextCheckOn() bool {
	if o != nil && o.TextCheckOn != nil {
		return true
	}

	return false
}

// SetTextCheckOn gets a reference to the given string and assigns it to the TextCheckOn field.
func (o *CheckSshConfig) SetTextCheckOn(v string) {
	o.TextCheckOn = &v
}

// GetCheckPassword returns the CheckPassword field value if set, zero value otherwise.
func (o *CheckSshConfig) GetCheckPassword() string {
	if o == nil || o.CheckPassword == nil {
		var ret string
		return ret
	}
	return *o.CheckPassword
}

// GetCheckPasswordOk returns a tuple with the CheckPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSshConfig) GetCheckPasswordOk() (*string, bool) {
	if o == nil || o.CheckPassword == nil {
		return nil, false
	}
	return o.CheckPassword, true
}

// HasCheckPassword returns a boolean if a field has been set.
func (o *CheckSshConfig) HasCheckPassword() bool {
	if o != nil && o.CheckPassword != nil {
		return true
	}

	return false
}

// SetCheckPassword gets a reference to the given string and assigns it to the CheckPassword field.
func (o *CheckSshConfig) SetCheckPassword(v string) {
	o.CheckPassword = &v
}

// GetSshHost returns the SshHost field value if set, zero value otherwise.
func (o *CheckSshConfig) GetSshHost() string {
	if o == nil || o.SshHost == nil {
		var ret string
		return ret
	}
	return *o.SshHost
}

// GetSshHostOk returns a tuple with the SshHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSshConfig) GetSshHostOk() (*string, bool) {
	if o == nil || o.SshHost == nil {
		return nil, false
	}
	return o.SshHost, true
}

// HasSshHost returns a boolean if a field has been set.
func (o *CheckSshConfig) HasSshHost() bool {
	if o != nil && o.SshHost != nil {
		return true
	}

	return false
}

// SetSshHost gets a reference to the given string and assigns it to the SshHost field.
func (o *CheckSshConfig) SetSshHost(v string) {
	o.SshHost = &v
}

// GetSshUser returns the SshUser field value if set, zero value otherwise.
func (o *CheckSshConfig) GetSshUser() string {
	if o == nil || o.SshUser == nil {
		var ret string
		return ret
	}
	return *o.SshUser
}

// GetSshUserOk returns a tuple with the SshUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSshConfig) GetSshUserOk() (*string, bool) {
	if o == nil || o.SshUser == nil {
		return nil, false
	}
	return o.SshUser, true
}

// HasSshUser returns a boolean if a field has been set.
func (o *CheckSshConfig) HasSshUser() bool {
	if o != nil && o.SshUser != nil {
		return true
	}

	return false
}

// SetSshUser gets a reference to the given string and assigns it to the SshUser field.
func (o *CheckSshConfig) SetSshUser(v string) {
	o.SshUser = &v
}

// GetWebTextMatch returns the WebTextMatch field value if set, zero value otherwise.
func (o *CheckSshConfig) GetWebTextMatch() string {
	if o == nil || o.WebTextMatch == nil {
		var ret string
		return ret
	}
	return *o.WebTextMatch
}

// GetWebTextMatchOk returns a tuple with the WebTextMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSshConfig) GetWebTextMatchOk() (*string, bool) {
	if o == nil || o.WebTextMatch == nil {
		return nil, false
	}
	return o.WebTextMatch, true
}

// HasWebTextMatch returns a boolean if a field has been set.
func (o *CheckSshConfig) HasWebTextMatch() bool {
	if o != nil && o.WebTextMatch != nil {
		return true
	}

	return false
}

// SetWebTextMatch gets a reference to the given string and assigns it to the WebTextMatch field.
func (o *CheckSshConfig) SetWebTextMatch(v string) {
	o.WebTextMatch = &v
}

// GetCheckPasswordHash returns the CheckPasswordHash field value if set, zero value otherwise.
func (o *CheckSshConfig) GetCheckPasswordHash() string {
	if o == nil || o.CheckPasswordHash == nil {
		var ret string
		return ret
	}
	return *o.CheckPasswordHash
}

// GetCheckPasswordHashOk returns a tuple with the CheckPasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSshConfig) GetCheckPasswordHashOk() (*string, bool) {
	if o == nil || o.CheckPasswordHash == nil {
		return nil, false
	}
	return o.CheckPasswordHash, true
}

// HasCheckPasswordHash returns a boolean if a field has been set.
func (o *CheckSshConfig) HasCheckPasswordHash() bool {
	if o != nil && o.CheckPasswordHash != nil {
		return true
	}

	return false
}

// SetCheckPasswordHash gets a reference to the given string and assigns it to the CheckPasswordHash field.
func (o *CheckSshConfig) SetCheckPasswordHash(v string) {
	o.CheckPasswordHash = &v
}

func (o CheckSshConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SshPort != nil {
		toSerialize["sshPort"] = o.SshPort
	}
	if o.CheckUser != nil {
		toSerialize["checkUser"] = o.CheckUser
	}
	if o.TunnelOn != nil {
		toSerialize["tunnelOn"] = o.TunnelOn
	}
	if o.TextCheckOn != nil {
		toSerialize["textCheckOn"] = o.TextCheckOn
	}
	if o.CheckPassword != nil {
		toSerialize["checkPassword"] = o.CheckPassword
	}
	if o.SshHost != nil {
		toSerialize["sshHost"] = o.SshHost
	}
	if o.SshUser != nil {
		toSerialize["sshUser"] = o.SshUser
	}
	if o.WebTextMatch != nil {
		toSerialize["webTextMatch"] = o.WebTextMatch
	}
	if o.CheckPasswordHash != nil {
		toSerialize["checkPasswordHash"] = o.CheckPasswordHash
	}
	return json.Marshal(toSerialize)
}

type NullableCheckSshConfig struct {
	value *CheckSshConfig
	isSet bool
}

func (v NullableCheckSshConfig) Get() *CheckSshConfig {
	return v.value
}

func (v *NullableCheckSshConfig) Set(val *CheckSshConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckSshConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckSshConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckSshConfig(val *CheckSshConfig) *NullableCheckSshConfig {
	return &NullableCheckSshConfig{value: val, isSet: true}
}

func (v NullableCheckSshConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckSshConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

