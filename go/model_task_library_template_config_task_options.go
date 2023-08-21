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

// TaskLibraryTemplateConfigTaskOptions struct for TaskLibraryTemplateConfigTaskOptions
type TaskLibraryTemplateConfigTaskOptions struct {
	Username NullableString `json:"username,omitempty"`
	LocalScriptGitId NullableString `json:"localScriptGitId,omitempty"`
	SshKey NullableString `json:"sshKey,omitempty"`
	Port NullableString `json:"port,omitempty"`
	Password NullableString `json:"password,omitempty"`
	PasswordHash NullableString `json:"passwordHash,omitempty"`
	ContainerTemplateId NullableString `json:"containerTemplateId,omitempty"`
	ContainerTemplate NullableString `json:"containerTemplate,omitempty"`
	LocalScriptGitRef NullableString `json:"localScriptGitRef,omitempty"`
	Host NullableString `json:"host,omitempty"`
}

// NewTaskLibraryTemplateConfigTaskOptions instantiates a new TaskLibraryTemplateConfigTaskOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskLibraryTemplateConfigTaskOptions() *TaskLibraryTemplateConfigTaskOptions {
	this := TaskLibraryTemplateConfigTaskOptions{}
	return &this
}

// NewTaskLibraryTemplateConfigTaskOptionsWithDefaults instantiates a new TaskLibraryTemplateConfigTaskOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskLibraryTemplateConfigTaskOptionsWithDefaults() *TaskLibraryTemplateConfigTaskOptions {
	this := TaskLibraryTemplateConfigTaskOptions{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskLibraryTemplateConfigTaskOptions) GetUsername() string {
	if o == nil || o.Username.Get() == nil {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskLibraryTemplateConfigTaskOptions) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *TaskLibraryTemplateConfigTaskOptions) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *TaskLibraryTemplateConfigTaskOptions) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *TaskLibraryTemplateConfigTaskOptions) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *TaskLibraryTemplateConfigTaskOptions) UnsetUsername() {
	o.Username.Unset()
}

// GetLocalScriptGitId returns the LocalScriptGitId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskLibraryTemplateConfigTaskOptions) GetLocalScriptGitId() string {
	if o == nil || o.LocalScriptGitId.Get() == nil {
		var ret string
		return ret
	}
	return *o.LocalScriptGitId.Get()
}

// GetLocalScriptGitIdOk returns a tuple with the LocalScriptGitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskLibraryTemplateConfigTaskOptions) GetLocalScriptGitIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LocalScriptGitId.Get(), o.LocalScriptGitId.IsSet()
}

// HasLocalScriptGitId returns a boolean if a field has been set.
func (o *TaskLibraryTemplateConfigTaskOptions) HasLocalScriptGitId() bool {
	if o != nil && o.LocalScriptGitId.IsSet() {
		return true
	}

	return false
}

// SetLocalScriptGitId gets a reference to the given NullableString and assigns it to the LocalScriptGitId field.
func (o *TaskLibraryTemplateConfigTaskOptions) SetLocalScriptGitId(v string) {
	o.LocalScriptGitId.Set(&v)
}
// SetLocalScriptGitIdNil sets the value for LocalScriptGitId to be an explicit nil
func (o *TaskLibraryTemplateConfigTaskOptions) SetLocalScriptGitIdNil() {
	o.LocalScriptGitId.Set(nil)
}

// UnsetLocalScriptGitId ensures that no value is present for LocalScriptGitId, not even an explicit nil
func (o *TaskLibraryTemplateConfigTaskOptions) UnsetLocalScriptGitId() {
	o.LocalScriptGitId.Unset()
}

// GetSshKey returns the SshKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskLibraryTemplateConfigTaskOptions) GetSshKey() string {
	if o == nil || o.SshKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.SshKey.Get()
}

// GetSshKeyOk returns a tuple with the SshKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskLibraryTemplateConfigTaskOptions) GetSshKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SshKey.Get(), o.SshKey.IsSet()
}

// HasSshKey returns a boolean if a field has been set.
func (o *TaskLibraryTemplateConfigTaskOptions) HasSshKey() bool {
	if o != nil && o.SshKey.IsSet() {
		return true
	}

	return false
}

// SetSshKey gets a reference to the given NullableString and assigns it to the SshKey field.
func (o *TaskLibraryTemplateConfigTaskOptions) SetSshKey(v string) {
	o.SshKey.Set(&v)
}
// SetSshKeyNil sets the value for SshKey to be an explicit nil
func (o *TaskLibraryTemplateConfigTaskOptions) SetSshKeyNil() {
	o.SshKey.Set(nil)
}

// UnsetSshKey ensures that no value is present for SshKey, not even an explicit nil
func (o *TaskLibraryTemplateConfigTaskOptions) UnsetSshKey() {
	o.SshKey.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskLibraryTemplateConfigTaskOptions) GetPort() string {
	if o == nil || o.Port.Get() == nil {
		var ret string
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskLibraryTemplateConfigTaskOptions) GetPortOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *TaskLibraryTemplateConfigTaskOptions) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableString and assigns it to the Port field.
func (o *TaskLibraryTemplateConfigTaskOptions) SetPort(v string) {
	o.Port.Set(&v)
}
// SetPortNil sets the value for Port to be an explicit nil
func (o *TaskLibraryTemplateConfigTaskOptions) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *TaskLibraryTemplateConfigTaskOptions) UnsetPort() {
	o.Port.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskLibraryTemplateConfigTaskOptions) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskLibraryTemplateConfigTaskOptions) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *TaskLibraryTemplateConfigTaskOptions) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *TaskLibraryTemplateConfigTaskOptions) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *TaskLibraryTemplateConfigTaskOptions) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *TaskLibraryTemplateConfigTaskOptions) UnsetPassword() {
	o.Password.Unset()
}

// GetPasswordHash returns the PasswordHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskLibraryTemplateConfigTaskOptions) GetPasswordHash() string {
	if o == nil || o.PasswordHash.Get() == nil {
		var ret string
		return ret
	}
	return *o.PasswordHash.Get()
}

// GetPasswordHashOk returns a tuple with the PasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskLibraryTemplateConfigTaskOptions) GetPasswordHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PasswordHash.Get(), o.PasswordHash.IsSet()
}

// HasPasswordHash returns a boolean if a field has been set.
func (o *TaskLibraryTemplateConfigTaskOptions) HasPasswordHash() bool {
	if o != nil && o.PasswordHash.IsSet() {
		return true
	}

	return false
}

// SetPasswordHash gets a reference to the given NullableString and assigns it to the PasswordHash field.
func (o *TaskLibraryTemplateConfigTaskOptions) SetPasswordHash(v string) {
	o.PasswordHash.Set(&v)
}
// SetPasswordHashNil sets the value for PasswordHash to be an explicit nil
func (o *TaskLibraryTemplateConfigTaskOptions) SetPasswordHashNil() {
	o.PasswordHash.Set(nil)
}

// UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
func (o *TaskLibraryTemplateConfigTaskOptions) UnsetPasswordHash() {
	o.PasswordHash.Unset()
}

// GetContainerTemplateId returns the ContainerTemplateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskLibraryTemplateConfigTaskOptions) GetContainerTemplateId() string {
	if o == nil || o.ContainerTemplateId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContainerTemplateId.Get()
}

// GetContainerTemplateIdOk returns a tuple with the ContainerTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskLibraryTemplateConfigTaskOptions) GetContainerTemplateIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContainerTemplateId.Get(), o.ContainerTemplateId.IsSet()
}

// HasContainerTemplateId returns a boolean if a field has been set.
func (o *TaskLibraryTemplateConfigTaskOptions) HasContainerTemplateId() bool {
	if o != nil && o.ContainerTemplateId.IsSet() {
		return true
	}

	return false
}

// SetContainerTemplateId gets a reference to the given NullableString and assigns it to the ContainerTemplateId field.
func (o *TaskLibraryTemplateConfigTaskOptions) SetContainerTemplateId(v string) {
	o.ContainerTemplateId.Set(&v)
}
// SetContainerTemplateIdNil sets the value for ContainerTemplateId to be an explicit nil
func (o *TaskLibraryTemplateConfigTaskOptions) SetContainerTemplateIdNil() {
	o.ContainerTemplateId.Set(nil)
}

// UnsetContainerTemplateId ensures that no value is present for ContainerTemplateId, not even an explicit nil
func (o *TaskLibraryTemplateConfigTaskOptions) UnsetContainerTemplateId() {
	o.ContainerTemplateId.Unset()
}

// GetContainerTemplate returns the ContainerTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskLibraryTemplateConfigTaskOptions) GetContainerTemplate() string {
	if o == nil || o.ContainerTemplate.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContainerTemplate.Get()
}

// GetContainerTemplateOk returns a tuple with the ContainerTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskLibraryTemplateConfigTaskOptions) GetContainerTemplateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContainerTemplate.Get(), o.ContainerTemplate.IsSet()
}

// HasContainerTemplate returns a boolean if a field has been set.
func (o *TaskLibraryTemplateConfigTaskOptions) HasContainerTemplate() bool {
	if o != nil && o.ContainerTemplate.IsSet() {
		return true
	}

	return false
}

// SetContainerTemplate gets a reference to the given NullableString and assigns it to the ContainerTemplate field.
func (o *TaskLibraryTemplateConfigTaskOptions) SetContainerTemplate(v string) {
	o.ContainerTemplate.Set(&v)
}
// SetContainerTemplateNil sets the value for ContainerTemplate to be an explicit nil
func (o *TaskLibraryTemplateConfigTaskOptions) SetContainerTemplateNil() {
	o.ContainerTemplate.Set(nil)
}

// UnsetContainerTemplate ensures that no value is present for ContainerTemplate, not even an explicit nil
func (o *TaskLibraryTemplateConfigTaskOptions) UnsetContainerTemplate() {
	o.ContainerTemplate.Unset()
}

// GetLocalScriptGitRef returns the LocalScriptGitRef field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskLibraryTemplateConfigTaskOptions) GetLocalScriptGitRef() string {
	if o == nil || o.LocalScriptGitRef.Get() == nil {
		var ret string
		return ret
	}
	return *o.LocalScriptGitRef.Get()
}

// GetLocalScriptGitRefOk returns a tuple with the LocalScriptGitRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskLibraryTemplateConfigTaskOptions) GetLocalScriptGitRefOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LocalScriptGitRef.Get(), o.LocalScriptGitRef.IsSet()
}

// HasLocalScriptGitRef returns a boolean if a field has been set.
func (o *TaskLibraryTemplateConfigTaskOptions) HasLocalScriptGitRef() bool {
	if o != nil && o.LocalScriptGitRef.IsSet() {
		return true
	}

	return false
}

// SetLocalScriptGitRef gets a reference to the given NullableString and assigns it to the LocalScriptGitRef field.
func (o *TaskLibraryTemplateConfigTaskOptions) SetLocalScriptGitRef(v string) {
	o.LocalScriptGitRef.Set(&v)
}
// SetLocalScriptGitRefNil sets the value for LocalScriptGitRef to be an explicit nil
func (o *TaskLibraryTemplateConfigTaskOptions) SetLocalScriptGitRefNil() {
	o.LocalScriptGitRef.Set(nil)
}

// UnsetLocalScriptGitRef ensures that no value is present for LocalScriptGitRef, not even an explicit nil
func (o *TaskLibraryTemplateConfigTaskOptions) UnsetLocalScriptGitRef() {
	o.LocalScriptGitRef.Unset()
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskLibraryTemplateConfigTaskOptions) GetHost() string {
	if o == nil || o.Host.Get() == nil {
		var ret string
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskLibraryTemplateConfigTaskOptions) GetHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *TaskLibraryTemplateConfigTaskOptions) HasHost() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableString and assigns it to the Host field.
func (o *TaskLibraryTemplateConfigTaskOptions) SetHost(v string) {
	o.Host.Set(&v)
}
// SetHostNil sets the value for Host to be an explicit nil
func (o *TaskLibraryTemplateConfigTaskOptions) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *TaskLibraryTemplateConfigTaskOptions) UnsetHost() {
	o.Host.Unset()
}

func (o TaskLibraryTemplateConfigTaskOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.LocalScriptGitId.IsSet() {
		toSerialize["localScriptGitId"] = o.LocalScriptGitId.Get()
	}
	if o.SshKey.IsSet() {
		toSerialize["sshKey"] = o.SshKey.Get()
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
	if o.ContainerTemplateId.IsSet() {
		toSerialize["containerTemplateId"] = o.ContainerTemplateId.Get()
	}
	if o.ContainerTemplate.IsSet() {
		toSerialize["containerTemplate"] = o.ContainerTemplate.Get()
	}
	if o.LocalScriptGitRef.IsSet() {
		toSerialize["localScriptGitRef"] = o.LocalScriptGitRef.Get()
	}
	if o.Host.IsSet() {
		toSerialize["host"] = o.Host.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTaskLibraryTemplateConfigTaskOptions struct {
	value *TaskLibraryTemplateConfigTaskOptions
	isSet bool
}

func (v NullableTaskLibraryTemplateConfigTaskOptions) Get() *TaskLibraryTemplateConfigTaskOptions {
	return v.value
}

func (v *NullableTaskLibraryTemplateConfigTaskOptions) Set(val *TaskLibraryTemplateConfigTaskOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskLibraryTemplateConfigTaskOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskLibraryTemplateConfigTaskOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskLibraryTemplateConfigTaskOptions(val *TaskLibraryTemplateConfigTaskOptions) *NullableTaskLibraryTemplateConfigTaskOptions {
	return &NullableTaskLibraryTemplateConfigTaskOptions{value: val, isSet: true}
}

func (v NullableTaskLibraryTemplateConfigTaskOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskLibraryTemplateConfigTaskOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

