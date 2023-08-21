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

// TaskShellConfigTaskOptions struct for TaskShellConfigTaskOptions
type TaskShellConfigTaskOptions struct {
	LocalScriptGitRef NullableString `json:"localScriptGitRef,omitempty"`
	Port NullableString `json:"port,omitempty"`
	Password NullableString `json:"password,omitempty"`
	PasswordHash NullableString `json:"passwordHash,omitempty"`
	Username NullableString `json:"username,omitempty"`
	Host NullableString `json:"host,omitempty"`
	SshKey NullableString `json:"sshKey,omitempty"`
	LocalScriptGitId NullableString `json:"localScriptGitId,omitempty"`
}

// NewTaskShellConfigTaskOptions instantiates a new TaskShellConfigTaskOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskShellConfigTaskOptions() *TaskShellConfigTaskOptions {
	this := TaskShellConfigTaskOptions{}
	return &this
}

// NewTaskShellConfigTaskOptionsWithDefaults instantiates a new TaskShellConfigTaskOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskShellConfigTaskOptionsWithDefaults() *TaskShellConfigTaskOptions {
	this := TaskShellConfigTaskOptions{}
	return &this
}

// GetLocalScriptGitRef returns the LocalScriptGitRef field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskShellConfigTaskOptions) GetLocalScriptGitRef() string {
	if o == nil || o.LocalScriptGitRef.Get() == nil {
		var ret string
		return ret
	}
	return *o.LocalScriptGitRef.Get()
}

// GetLocalScriptGitRefOk returns a tuple with the LocalScriptGitRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskShellConfigTaskOptions) GetLocalScriptGitRefOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LocalScriptGitRef.Get(), o.LocalScriptGitRef.IsSet()
}

// HasLocalScriptGitRef returns a boolean if a field has been set.
func (o *TaskShellConfigTaskOptions) HasLocalScriptGitRef() bool {
	if o != nil && o.LocalScriptGitRef.IsSet() {
		return true
	}

	return false
}

// SetLocalScriptGitRef gets a reference to the given NullableString and assigns it to the LocalScriptGitRef field.
func (o *TaskShellConfigTaskOptions) SetLocalScriptGitRef(v string) {
	o.LocalScriptGitRef.Set(&v)
}
// SetLocalScriptGitRefNil sets the value for LocalScriptGitRef to be an explicit nil
func (o *TaskShellConfigTaskOptions) SetLocalScriptGitRefNil() {
	o.LocalScriptGitRef.Set(nil)
}

// UnsetLocalScriptGitRef ensures that no value is present for LocalScriptGitRef, not even an explicit nil
func (o *TaskShellConfigTaskOptions) UnsetLocalScriptGitRef() {
	o.LocalScriptGitRef.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskShellConfigTaskOptions) GetPort() string {
	if o == nil || o.Port.Get() == nil {
		var ret string
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskShellConfigTaskOptions) GetPortOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *TaskShellConfigTaskOptions) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableString and assigns it to the Port field.
func (o *TaskShellConfigTaskOptions) SetPort(v string) {
	o.Port.Set(&v)
}
// SetPortNil sets the value for Port to be an explicit nil
func (o *TaskShellConfigTaskOptions) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *TaskShellConfigTaskOptions) UnsetPort() {
	o.Port.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskShellConfigTaskOptions) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskShellConfigTaskOptions) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *TaskShellConfigTaskOptions) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *TaskShellConfigTaskOptions) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *TaskShellConfigTaskOptions) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *TaskShellConfigTaskOptions) UnsetPassword() {
	o.Password.Unset()
}

// GetPasswordHash returns the PasswordHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskShellConfigTaskOptions) GetPasswordHash() string {
	if o == nil || o.PasswordHash.Get() == nil {
		var ret string
		return ret
	}
	return *o.PasswordHash.Get()
}

// GetPasswordHashOk returns a tuple with the PasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskShellConfigTaskOptions) GetPasswordHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PasswordHash.Get(), o.PasswordHash.IsSet()
}

// HasPasswordHash returns a boolean if a field has been set.
func (o *TaskShellConfigTaskOptions) HasPasswordHash() bool {
	if o != nil && o.PasswordHash.IsSet() {
		return true
	}

	return false
}

// SetPasswordHash gets a reference to the given NullableString and assigns it to the PasswordHash field.
func (o *TaskShellConfigTaskOptions) SetPasswordHash(v string) {
	o.PasswordHash.Set(&v)
}
// SetPasswordHashNil sets the value for PasswordHash to be an explicit nil
func (o *TaskShellConfigTaskOptions) SetPasswordHashNil() {
	o.PasswordHash.Set(nil)
}

// UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
func (o *TaskShellConfigTaskOptions) UnsetPasswordHash() {
	o.PasswordHash.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskShellConfigTaskOptions) GetUsername() string {
	if o == nil || o.Username.Get() == nil {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskShellConfigTaskOptions) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *TaskShellConfigTaskOptions) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *TaskShellConfigTaskOptions) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *TaskShellConfigTaskOptions) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *TaskShellConfigTaskOptions) UnsetUsername() {
	o.Username.Unset()
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskShellConfigTaskOptions) GetHost() string {
	if o == nil || o.Host.Get() == nil {
		var ret string
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskShellConfigTaskOptions) GetHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *TaskShellConfigTaskOptions) HasHost() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableString and assigns it to the Host field.
func (o *TaskShellConfigTaskOptions) SetHost(v string) {
	o.Host.Set(&v)
}
// SetHostNil sets the value for Host to be an explicit nil
func (o *TaskShellConfigTaskOptions) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *TaskShellConfigTaskOptions) UnsetHost() {
	o.Host.Unset()
}

// GetSshKey returns the SshKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskShellConfigTaskOptions) GetSshKey() string {
	if o == nil || o.SshKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.SshKey.Get()
}

// GetSshKeyOk returns a tuple with the SshKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskShellConfigTaskOptions) GetSshKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SshKey.Get(), o.SshKey.IsSet()
}

// HasSshKey returns a boolean if a field has been set.
func (o *TaskShellConfigTaskOptions) HasSshKey() bool {
	if o != nil && o.SshKey.IsSet() {
		return true
	}

	return false
}

// SetSshKey gets a reference to the given NullableString and assigns it to the SshKey field.
func (o *TaskShellConfigTaskOptions) SetSshKey(v string) {
	o.SshKey.Set(&v)
}
// SetSshKeyNil sets the value for SshKey to be an explicit nil
func (o *TaskShellConfigTaskOptions) SetSshKeyNil() {
	o.SshKey.Set(nil)
}

// UnsetSshKey ensures that no value is present for SshKey, not even an explicit nil
func (o *TaskShellConfigTaskOptions) UnsetSshKey() {
	o.SshKey.Unset()
}

// GetLocalScriptGitId returns the LocalScriptGitId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskShellConfigTaskOptions) GetLocalScriptGitId() string {
	if o == nil || o.LocalScriptGitId.Get() == nil {
		var ret string
		return ret
	}
	return *o.LocalScriptGitId.Get()
}

// GetLocalScriptGitIdOk returns a tuple with the LocalScriptGitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskShellConfigTaskOptions) GetLocalScriptGitIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LocalScriptGitId.Get(), o.LocalScriptGitId.IsSet()
}

// HasLocalScriptGitId returns a boolean if a field has been set.
func (o *TaskShellConfigTaskOptions) HasLocalScriptGitId() bool {
	if o != nil && o.LocalScriptGitId.IsSet() {
		return true
	}

	return false
}

// SetLocalScriptGitId gets a reference to the given NullableString and assigns it to the LocalScriptGitId field.
func (o *TaskShellConfigTaskOptions) SetLocalScriptGitId(v string) {
	o.LocalScriptGitId.Set(&v)
}
// SetLocalScriptGitIdNil sets the value for LocalScriptGitId to be an explicit nil
func (o *TaskShellConfigTaskOptions) SetLocalScriptGitIdNil() {
	o.LocalScriptGitId.Set(nil)
}

// UnsetLocalScriptGitId ensures that no value is present for LocalScriptGitId, not even an explicit nil
func (o *TaskShellConfigTaskOptions) UnsetLocalScriptGitId() {
	o.LocalScriptGitId.Unset()
}

func (o TaskShellConfigTaskOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LocalScriptGitRef.IsSet() {
		toSerialize["localScriptGitRef"] = o.LocalScriptGitRef.Get()
	}
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.PasswordHash.IsSet() {
		toSerialize["passwordHash"] = o.PasswordHash.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Host.IsSet() {
		toSerialize["host"] = o.Host.Get()
	}
	if o.SshKey.IsSet() {
		toSerialize["sshKey"] = o.SshKey.Get()
	}
	if o.LocalScriptGitId.IsSet() {
		toSerialize["localScriptGitId"] = o.LocalScriptGitId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTaskShellConfigTaskOptions struct {
	value *TaskShellConfigTaskOptions
	isSet bool
}

func (v NullableTaskShellConfigTaskOptions) Get() *TaskShellConfigTaskOptions {
	return v.value
}

func (v *NullableTaskShellConfigTaskOptions) Set(val *TaskShellConfigTaskOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskShellConfigTaskOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskShellConfigTaskOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskShellConfigTaskOptions(val *TaskShellConfigTaskOptions) *NullableTaskShellConfigTaskOptions {
	return &NullableTaskShellConfigTaskOptions{value: val, isSet: true}
}

func (v NullableTaskShellConfigTaskOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskShellConfigTaskOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


