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

// TaskJavaConfigTaskOptions struct for TaskJavaConfigTaskOptions
type TaskJavaConfigTaskOptions struct {
	Username NullableString `json:"username,omitempty"`
	Port NullableString `json:"port,omitempty"`
	JsScript NullableString `json:"jsScript,omitempty"`
	Host NullableString `json:"host,omitempty"`
	LocalScriptGitRef NullableString `json:"localScriptGitRef,omitempty"`
	Password NullableString `json:"password,omitempty"`
	PasswordHash NullableString `json:"passwordHash,omitempty"`
	SshKey NullableString `json:"sshKey,omitempty"`
	LocalScriptGitId NullableString `json:"localScriptGitId,omitempty"`
}

// NewTaskJavaConfigTaskOptions instantiates a new TaskJavaConfigTaskOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskJavaConfigTaskOptions() *TaskJavaConfigTaskOptions {
	this := TaskJavaConfigTaskOptions{}
	return &this
}

// NewTaskJavaConfigTaskOptionsWithDefaults instantiates a new TaskJavaConfigTaskOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskJavaConfigTaskOptionsWithDefaults() *TaskJavaConfigTaskOptions {
	this := TaskJavaConfigTaskOptions{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskJavaConfigTaskOptions) GetUsername() string {
	if o == nil || o.Username.Get() == nil {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskJavaConfigTaskOptions) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *TaskJavaConfigTaskOptions) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *TaskJavaConfigTaskOptions) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *TaskJavaConfigTaskOptions) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *TaskJavaConfigTaskOptions) UnsetUsername() {
	o.Username.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskJavaConfigTaskOptions) GetPort() string {
	if o == nil || o.Port.Get() == nil {
		var ret string
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskJavaConfigTaskOptions) GetPortOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *TaskJavaConfigTaskOptions) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableString and assigns it to the Port field.
func (o *TaskJavaConfigTaskOptions) SetPort(v string) {
	o.Port.Set(&v)
}
// SetPortNil sets the value for Port to be an explicit nil
func (o *TaskJavaConfigTaskOptions) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *TaskJavaConfigTaskOptions) UnsetPort() {
	o.Port.Unset()
}

// GetJsScript returns the JsScript field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskJavaConfigTaskOptions) GetJsScript() string {
	if o == nil || o.JsScript.Get() == nil {
		var ret string
		return ret
	}
	return *o.JsScript.Get()
}

// GetJsScriptOk returns a tuple with the JsScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskJavaConfigTaskOptions) GetJsScriptOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.JsScript.Get(), o.JsScript.IsSet()
}

// HasJsScript returns a boolean if a field has been set.
func (o *TaskJavaConfigTaskOptions) HasJsScript() bool {
	if o != nil && o.JsScript.IsSet() {
		return true
	}

	return false
}

// SetJsScript gets a reference to the given NullableString and assigns it to the JsScript field.
func (o *TaskJavaConfigTaskOptions) SetJsScript(v string) {
	o.JsScript.Set(&v)
}
// SetJsScriptNil sets the value for JsScript to be an explicit nil
func (o *TaskJavaConfigTaskOptions) SetJsScriptNil() {
	o.JsScript.Set(nil)
}

// UnsetJsScript ensures that no value is present for JsScript, not even an explicit nil
func (o *TaskJavaConfigTaskOptions) UnsetJsScript() {
	o.JsScript.Unset()
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskJavaConfigTaskOptions) GetHost() string {
	if o == nil || o.Host.Get() == nil {
		var ret string
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskJavaConfigTaskOptions) GetHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *TaskJavaConfigTaskOptions) HasHost() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableString and assigns it to the Host field.
func (o *TaskJavaConfigTaskOptions) SetHost(v string) {
	o.Host.Set(&v)
}
// SetHostNil sets the value for Host to be an explicit nil
func (o *TaskJavaConfigTaskOptions) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *TaskJavaConfigTaskOptions) UnsetHost() {
	o.Host.Unset()
}

// GetLocalScriptGitRef returns the LocalScriptGitRef field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskJavaConfigTaskOptions) GetLocalScriptGitRef() string {
	if o == nil || o.LocalScriptGitRef.Get() == nil {
		var ret string
		return ret
	}
	return *o.LocalScriptGitRef.Get()
}

// GetLocalScriptGitRefOk returns a tuple with the LocalScriptGitRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskJavaConfigTaskOptions) GetLocalScriptGitRefOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LocalScriptGitRef.Get(), o.LocalScriptGitRef.IsSet()
}

// HasLocalScriptGitRef returns a boolean if a field has been set.
func (o *TaskJavaConfigTaskOptions) HasLocalScriptGitRef() bool {
	if o != nil && o.LocalScriptGitRef.IsSet() {
		return true
	}

	return false
}

// SetLocalScriptGitRef gets a reference to the given NullableString and assigns it to the LocalScriptGitRef field.
func (o *TaskJavaConfigTaskOptions) SetLocalScriptGitRef(v string) {
	o.LocalScriptGitRef.Set(&v)
}
// SetLocalScriptGitRefNil sets the value for LocalScriptGitRef to be an explicit nil
func (o *TaskJavaConfigTaskOptions) SetLocalScriptGitRefNil() {
	o.LocalScriptGitRef.Set(nil)
}

// UnsetLocalScriptGitRef ensures that no value is present for LocalScriptGitRef, not even an explicit nil
func (o *TaskJavaConfigTaskOptions) UnsetLocalScriptGitRef() {
	o.LocalScriptGitRef.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskJavaConfigTaskOptions) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskJavaConfigTaskOptions) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *TaskJavaConfigTaskOptions) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *TaskJavaConfigTaskOptions) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *TaskJavaConfigTaskOptions) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *TaskJavaConfigTaskOptions) UnsetPassword() {
	o.Password.Unset()
}

// GetPasswordHash returns the PasswordHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskJavaConfigTaskOptions) GetPasswordHash() string {
	if o == nil || o.PasswordHash.Get() == nil {
		var ret string
		return ret
	}
	return *o.PasswordHash.Get()
}

// GetPasswordHashOk returns a tuple with the PasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskJavaConfigTaskOptions) GetPasswordHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PasswordHash.Get(), o.PasswordHash.IsSet()
}

// HasPasswordHash returns a boolean if a field has been set.
func (o *TaskJavaConfigTaskOptions) HasPasswordHash() bool {
	if o != nil && o.PasswordHash.IsSet() {
		return true
	}

	return false
}

// SetPasswordHash gets a reference to the given NullableString and assigns it to the PasswordHash field.
func (o *TaskJavaConfigTaskOptions) SetPasswordHash(v string) {
	o.PasswordHash.Set(&v)
}
// SetPasswordHashNil sets the value for PasswordHash to be an explicit nil
func (o *TaskJavaConfigTaskOptions) SetPasswordHashNil() {
	o.PasswordHash.Set(nil)
}

// UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
func (o *TaskJavaConfigTaskOptions) UnsetPasswordHash() {
	o.PasswordHash.Unset()
}

// GetSshKey returns the SshKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskJavaConfigTaskOptions) GetSshKey() string {
	if o == nil || o.SshKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.SshKey.Get()
}

// GetSshKeyOk returns a tuple with the SshKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskJavaConfigTaskOptions) GetSshKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SshKey.Get(), o.SshKey.IsSet()
}

// HasSshKey returns a boolean if a field has been set.
func (o *TaskJavaConfigTaskOptions) HasSshKey() bool {
	if o != nil && o.SshKey.IsSet() {
		return true
	}

	return false
}

// SetSshKey gets a reference to the given NullableString and assigns it to the SshKey field.
func (o *TaskJavaConfigTaskOptions) SetSshKey(v string) {
	o.SshKey.Set(&v)
}
// SetSshKeyNil sets the value for SshKey to be an explicit nil
func (o *TaskJavaConfigTaskOptions) SetSshKeyNil() {
	o.SshKey.Set(nil)
}

// UnsetSshKey ensures that no value is present for SshKey, not even an explicit nil
func (o *TaskJavaConfigTaskOptions) UnsetSshKey() {
	o.SshKey.Unset()
}

// GetLocalScriptGitId returns the LocalScriptGitId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskJavaConfigTaskOptions) GetLocalScriptGitId() string {
	if o == nil || o.LocalScriptGitId.Get() == nil {
		var ret string
		return ret
	}
	return *o.LocalScriptGitId.Get()
}

// GetLocalScriptGitIdOk returns a tuple with the LocalScriptGitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskJavaConfigTaskOptions) GetLocalScriptGitIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LocalScriptGitId.Get(), o.LocalScriptGitId.IsSet()
}

// HasLocalScriptGitId returns a boolean if a field has been set.
func (o *TaskJavaConfigTaskOptions) HasLocalScriptGitId() bool {
	if o != nil && o.LocalScriptGitId.IsSet() {
		return true
	}

	return false
}

// SetLocalScriptGitId gets a reference to the given NullableString and assigns it to the LocalScriptGitId field.
func (o *TaskJavaConfigTaskOptions) SetLocalScriptGitId(v string) {
	o.LocalScriptGitId.Set(&v)
}
// SetLocalScriptGitIdNil sets the value for LocalScriptGitId to be an explicit nil
func (o *TaskJavaConfigTaskOptions) SetLocalScriptGitIdNil() {
	o.LocalScriptGitId.Set(nil)
}

// UnsetLocalScriptGitId ensures that no value is present for LocalScriptGitId, not even an explicit nil
func (o *TaskJavaConfigTaskOptions) UnsetLocalScriptGitId() {
	o.LocalScriptGitId.Unset()
}

func (o TaskJavaConfigTaskOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	if o.JsScript.IsSet() {
		toSerialize["jsScript"] = o.JsScript.Get()
	}
	if o.Host.IsSet() {
		toSerialize["host"] = o.Host.Get()
	}
	if o.LocalScriptGitRef.IsSet() {
		toSerialize["localScriptGitRef"] = o.LocalScriptGitRef.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.PasswordHash.IsSet() {
		toSerialize["passwordHash"] = o.PasswordHash.Get()
	}
	if o.SshKey.IsSet() {
		toSerialize["sshKey"] = o.SshKey.Get()
	}
	if o.LocalScriptGitId.IsSet() {
		toSerialize["localScriptGitId"] = o.LocalScriptGitId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTaskJavaConfigTaskOptions struct {
	value *TaskJavaConfigTaskOptions
	isSet bool
}

func (v NullableTaskJavaConfigTaskOptions) Get() *TaskJavaConfigTaskOptions {
	return v.value
}

func (v *NullableTaskJavaConfigTaskOptions) Set(val *TaskJavaConfigTaskOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskJavaConfigTaskOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskJavaConfigTaskOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskJavaConfigTaskOptions(val *TaskJavaConfigTaskOptions) *NullableTaskJavaConfigTaskOptions {
	return &NullableTaskJavaConfigTaskOptions{value: val, isSet: true}
}

func (v NullableTaskJavaConfigTaskOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskJavaConfigTaskOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


