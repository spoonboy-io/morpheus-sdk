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

// TaskLibraryScriptConfigTaskOptions struct for TaskLibraryScriptConfigTaskOptions
type TaskLibraryScriptConfigTaskOptions struct {
	Host NullableString `json:"host,omitempty"`
	Username NullableString `json:"username,omitempty"`
	Password NullableString `json:"password,omitempty"`
	PasswordHash NullableString `json:"passwordHash,omitempty"`
	LocalScriptGitRef NullableString `json:"localScriptGitRef,omitempty"`
	LocalScriptGitId NullableString `json:"localScriptGitId,omitempty"`
	ContainerScriptId NullableString `json:"containerScriptId,omitempty"`
	SshKey NullableString `json:"sshKey,omitempty"`
	ContainerScript NullableString `json:"containerScript,omitempty"`
	Port NullableString `json:"port,omitempty"`
}

// NewTaskLibraryScriptConfigTaskOptions instantiates a new TaskLibraryScriptConfigTaskOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskLibraryScriptConfigTaskOptions() *TaskLibraryScriptConfigTaskOptions {
	this := TaskLibraryScriptConfigTaskOptions{}
	return &this
}

// NewTaskLibraryScriptConfigTaskOptionsWithDefaults instantiates a new TaskLibraryScriptConfigTaskOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskLibraryScriptConfigTaskOptionsWithDefaults() *TaskLibraryScriptConfigTaskOptions {
	this := TaskLibraryScriptConfigTaskOptions{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskLibraryScriptConfigTaskOptions) GetHost() string {
	if o == nil || o.Host.Get() == nil {
		var ret string
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskLibraryScriptConfigTaskOptions) GetHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfigTaskOptions) HasHost() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableString and assigns it to the Host field.
func (o *TaskLibraryScriptConfigTaskOptions) SetHost(v string) {
	o.Host.Set(&v)
}
// SetHostNil sets the value for Host to be an explicit nil
func (o *TaskLibraryScriptConfigTaskOptions) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *TaskLibraryScriptConfigTaskOptions) UnsetHost() {
	o.Host.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskLibraryScriptConfigTaskOptions) GetUsername() string {
	if o == nil || o.Username.Get() == nil {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskLibraryScriptConfigTaskOptions) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfigTaskOptions) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *TaskLibraryScriptConfigTaskOptions) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *TaskLibraryScriptConfigTaskOptions) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *TaskLibraryScriptConfigTaskOptions) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskLibraryScriptConfigTaskOptions) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskLibraryScriptConfigTaskOptions) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfigTaskOptions) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *TaskLibraryScriptConfigTaskOptions) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *TaskLibraryScriptConfigTaskOptions) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *TaskLibraryScriptConfigTaskOptions) UnsetPassword() {
	o.Password.Unset()
}

// GetPasswordHash returns the PasswordHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskLibraryScriptConfigTaskOptions) GetPasswordHash() string {
	if o == nil || o.PasswordHash.Get() == nil {
		var ret string
		return ret
	}
	return *o.PasswordHash.Get()
}

// GetPasswordHashOk returns a tuple with the PasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskLibraryScriptConfigTaskOptions) GetPasswordHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PasswordHash.Get(), o.PasswordHash.IsSet()
}

// HasPasswordHash returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfigTaskOptions) HasPasswordHash() bool {
	if o != nil && o.PasswordHash.IsSet() {
		return true
	}

	return false
}

// SetPasswordHash gets a reference to the given NullableString and assigns it to the PasswordHash field.
func (o *TaskLibraryScriptConfigTaskOptions) SetPasswordHash(v string) {
	o.PasswordHash.Set(&v)
}
// SetPasswordHashNil sets the value for PasswordHash to be an explicit nil
func (o *TaskLibraryScriptConfigTaskOptions) SetPasswordHashNil() {
	o.PasswordHash.Set(nil)
}

// UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
func (o *TaskLibraryScriptConfigTaskOptions) UnsetPasswordHash() {
	o.PasswordHash.Unset()
}

// GetLocalScriptGitRef returns the LocalScriptGitRef field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskLibraryScriptConfigTaskOptions) GetLocalScriptGitRef() string {
	if o == nil || o.LocalScriptGitRef.Get() == nil {
		var ret string
		return ret
	}
	return *o.LocalScriptGitRef.Get()
}

// GetLocalScriptGitRefOk returns a tuple with the LocalScriptGitRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskLibraryScriptConfigTaskOptions) GetLocalScriptGitRefOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LocalScriptGitRef.Get(), o.LocalScriptGitRef.IsSet()
}

// HasLocalScriptGitRef returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfigTaskOptions) HasLocalScriptGitRef() bool {
	if o != nil && o.LocalScriptGitRef.IsSet() {
		return true
	}

	return false
}

// SetLocalScriptGitRef gets a reference to the given NullableString and assigns it to the LocalScriptGitRef field.
func (o *TaskLibraryScriptConfigTaskOptions) SetLocalScriptGitRef(v string) {
	o.LocalScriptGitRef.Set(&v)
}
// SetLocalScriptGitRefNil sets the value for LocalScriptGitRef to be an explicit nil
func (o *TaskLibraryScriptConfigTaskOptions) SetLocalScriptGitRefNil() {
	o.LocalScriptGitRef.Set(nil)
}

// UnsetLocalScriptGitRef ensures that no value is present for LocalScriptGitRef, not even an explicit nil
func (o *TaskLibraryScriptConfigTaskOptions) UnsetLocalScriptGitRef() {
	o.LocalScriptGitRef.Unset()
}

// GetLocalScriptGitId returns the LocalScriptGitId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskLibraryScriptConfigTaskOptions) GetLocalScriptGitId() string {
	if o == nil || o.LocalScriptGitId.Get() == nil {
		var ret string
		return ret
	}
	return *o.LocalScriptGitId.Get()
}

// GetLocalScriptGitIdOk returns a tuple with the LocalScriptGitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskLibraryScriptConfigTaskOptions) GetLocalScriptGitIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LocalScriptGitId.Get(), o.LocalScriptGitId.IsSet()
}

// HasLocalScriptGitId returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfigTaskOptions) HasLocalScriptGitId() bool {
	if o != nil && o.LocalScriptGitId.IsSet() {
		return true
	}

	return false
}

// SetLocalScriptGitId gets a reference to the given NullableString and assigns it to the LocalScriptGitId field.
func (o *TaskLibraryScriptConfigTaskOptions) SetLocalScriptGitId(v string) {
	o.LocalScriptGitId.Set(&v)
}
// SetLocalScriptGitIdNil sets the value for LocalScriptGitId to be an explicit nil
func (o *TaskLibraryScriptConfigTaskOptions) SetLocalScriptGitIdNil() {
	o.LocalScriptGitId.Set(nil)
}

// UnsetLocalScriptGitId ensures that no value is present for LocalScriptGitId, not even an explicit nil
func (o *TaskLibraryScriptConfigTaskOptions) UnsetLocalScriptGitId() {
	o.LocalScriptGitId.Unset()
}

// GetContainerScriptId returns the ContainerScriptId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskLibraryScriptConfigTaskOptions) GetContainerScriptId() string {
	if o == nil || o.ContainerScriptId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContainerScriptId.Get()
}

// GetContainerScriptIdOk returns a tuple with the ContainerScriptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskLibraryScriptConfigTaskOptions) GetContainerScriptIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContainerScriptId.Get(), o.ContainerScriptId.IsSet()
}

// HasContainerScriptId returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfigTaskOptions) HasContainerScriptId() bool {
	if o != nil && o.ContainerScriptId.IsSet() {
		return true
	}

	return false
}

// SetContainerScriptId gets a reference to the given NullableString and assigns it to the ContainerScriptId field.
func (o *TaskLibraryScriptConfigTaskOptions) SetContainerScriptId(v string) {
	o.ContainerScriptId.Set(&v)
}
// SetContainerScriptIdNil sets the value for ContainerScriptId to be an explicit nil
func (o *TaskLibraryScriptConfigTaskOptions) SetContainerScriptIdNil() {
	o.ContainerScriptId.Set(nil)
}

// UnsetContainerScriptId ensures that no value is present for ContainerScriptId, not even an explicit nil
func (o *TaskLibraryScriptConfigTaskOptions) UnsetContainerScriptId() {
	o.ContainerScriptId.Unset()
}

// GetSshKey returns the SshKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskLibraryScriptConfigTaskOptions) GetSshKey() string {
	if o == nil || o.SshKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.SshKey.Get()
}

// GetSshKeyOk returns a tuple with the SshKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskLibraryScriptConfigTaskOptions) GetSshKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SshKey.Get(), o.SshKey.IsSet()
}

// HasSshKey returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfigTaskOptions) HasSshKey() bool {
	if o != nil && o.SshKey.IsSet() {
		return true
	}

	return false
}

// SetSshKey gets a reference to the given NullableString and assigns it to the SshKey field.
func (o *TaskLibraryScriptConfigTaskOptions) SetSshKey(v string) {
	o.SshKey.Set(&v)
}
// SetSshKeyNil sets the value for SshKey to be an explicit nil
func (o *TaskLibraryScriptConfigTaskOptions) SetSshKeyNil() {
	o.SshKey.Set(nil)
}

// UnsetSshKey ensures that no value is present for SshKey, not even an explicit nil
func (o *TaskLibraryScriptConfigTaskOptions) UnsetSshKey() {
	o.SshKey.Unset()
}

// GetContainerScript returns the ContainerScript field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskLibraryScriptConfigTaskOptions) GetContainerScript() string {
	if o == nil || o.ContainerScript.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContainerScript.Get()
}

// GetContainerScriptOk returns a tuple with the ContainerScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskLibraryScriptConfigTaskOptions) GetContainerScriptOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContainerScript.Get(), o.ContainerScript.IsSet()
}

// HasContainerScript returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfigTaskOptions) HasContainerScript() bool {
	if o != nil && o.ContainerScript.IsSet() {
		return true
	}

	return false
}

// SetContainerScript gets a reference to the given NullableString and assigns it to the ContainerScript field.
func (o *TaskLibraryScriptConfigTaskOptions) SetContainerScript(v string) {
	o.ContainerScript.Set(&v)
}
// SetContainerScriptNil sets the value for ContainerScript to be an explicit nil
func (o *TaskLibraryScriptConfigTaskOptions) SetContainerScriptNil() {
	o.ContainerScript.Set(nil)
}

// UnsetContainerScript ensures that no value is present for ContainerScript, not even an explicit nil
func (o *TaskLibraryScriptConfigTaskOptions) UnsetContainerScript() {
	o.ContainerScript.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskLibraryScriptConfigTaskOptions) GetPort() string {
	if o == nil || o.Port.Get() == nil {
		var ret string
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskLibraryScriptConfigTaskOptions) GetPortOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfigTaskOptions) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableString and assigns it to the Port field.
func (o *TaskLibraryScriptConfigTaskOptions) SetPort(v string) {
	o.Port.Set(&v)
}
// SetPortNil sets the value for Port to be an explicit nil
func (o *TaskLibraryScriptConfigTaskOptions) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *TaskLibraryScriptConfigTaskOptions) UnsetPort() {
	o.Port.Unset()
}

func (o TaskLibraryScriptConfigTaskOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Host.IsSet() {
		toSerialize["host"] = o.Host.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.PasswordHash.IsSet() {
		toSerialize["passwordHash"] = o.PasswordHash.Get()
	}
	if o.LocalScriptGitRef.IsSet() {
		toSerialize["localScriptGitRef"] = o.LocalScriptGitRef.Get()
	}
	if o.LocalScriptGitId.IsSet() {
		toSerialize["localScriptGitId"] = o.LocalScriptGitId.Get()
	}
	if o.ContainerScriptId.IsSet() {
		toSerialize["containerScriptId"] = o.ContainerScriptId.Get()
	}
	if o.SshKey.IsSet() {
		toSerialize["sshKey"] = o.SshKey.Get()
	}
	if o.ContainerScript.IsSet() {
		toSerialize["containerScript"] = o.ContainerScript.Get()
	}
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTaskLibraryScriptConfigTaskOptions struct {
	value *TaskLibraryScriptConfigTaskOptions
	isSet bool
}

func (v NullableTaskLibraryScriptConfigTaskOptions) Get() *TaskLibraryScriptConfigTaskOptions {
	return v.value
}

func (v *NullableTaskLibraryScriptConfigTaskOptions) Set(val *TaskLibraryScriptConfigTaskOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskLibraryScriptConfigTaskOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskLibraryScriptConfigTaskOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskLibraryScriptConfigTaskOptions(val *TaskLibraryScriptConfigTaskOptions) *NullableTaskLibraryScriptConfigTaskOptions {
	return &NullableTaskLibraryScriptConfigTaskOptions{value: val, isSet: true}
}

func (v NullableTaskLibraryScriptConfigTaskOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskLibraryScriptConfigTaskOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


