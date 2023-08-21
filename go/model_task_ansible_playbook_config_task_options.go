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

// TaskAnsiblePlaybookConfigTaskOptions struct for TaskAnsiblePlaybookConfigTaskOptions
type TaskAnsiblePlaybookConfigTaskOptions struct {
	AnsibleOptions NullableString `json:"ansibleOptions,omitempty"`
	AnsiblePlaybook NullableString `json:"ansiblePlaybook,omitempty"`
	SshKey NullableString `json:"sshKey,omitempty"`
	Port NullableString `json:"port,omitempty"`
	LocalScriptGitRef NullableString `json:"localScriptGitRef,omitempty"`
	Password NullableString `json:"password,omitempty"`
	PasswordHash NullableString `json:"passwordHash,omitempty"`
	LocalScriptGitId NullableString `json:"localScriptGitId,omitempty"`
	AnsibleGitId *string `json:"ansibleGitId,omitempty"`
	Host NullableString `json:"host,omitempty"`
	AnsibleSkipTags NullableString `json:"ansibleSkipTags,omitempty"`
	AnsibleTags NullableString `json:"ansibleTags,omitempty"`
	Username NullableString `json:"username,omitempty"`
	AnsibleGitRef NullableString `json:"ansibleGitRef,omitempty"`
}

// NewTaskAnsiblePlaybookConfigTaskOptions instantiates a new TaskAnsiblePlaybookConfigTaskOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskAnsiblePlaybookConfigTaskOptions() *TaskAnsiblePlaybookConfigTaskOptions {
	this := TaskAnsiblePlaybookConfigTaskOptions{}
	return &this
}

// NewTaskAnsiblePlaybookConfigTaskOptionsWithDefaults instantiates a new TaskAnsiblePlaybookConfigTaskOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskAnsiblePlaybookConfigTaskOptionsWithDefaults() *TaskAnsiblePlaybookConfigTaskOptions {
	this := TaskAnsiblePlaybookConfigTaskOptions{}
	return &this
}

// GetAnsibleOptions returns the AnsibleOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetAnsibleOptions() string {
	if o == nil || o.AnsibleOptions.Get() == nil {
		var ret string
		return ret
	}
	return *o.AnsibleOptions.Get()
}

// GetAnsibleOptionsOk returns a tuple with the AnsibleOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetAnsibleOptionsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AnsibleOptions.Get(), o.AnsibleOptions.IsSet()
}

// HasAnsibleOptions returns a boolean if a field has been set.
func (o *TaskAnsiblePlaybookConfigTaskOptions) HasAnsibleOptions() bool {
	if o != nil && o.AnsibleOptions.IsSet() {
		return true
	}

	return false
}

// SetAnsibleOptions gets a reference to the given NullableString and assigns it to the AnsibleOptions field.
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetAnsibleOptions(v string) {
	o.AnsibleOptions.Set(&v)
}
// SetAnsibleOptionsNil sets the value for AnsibleOptions to be an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetAnsibleOptionsNil() {
	o.AnsibleOptions.Set(nil)
}

// UnsetAnsibleOptions ensures that no value is present for AnsibleOptions, not even an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) UnsetAnsibleOptions() {
	o.AnsibleOptions.Unset()
}

// GetAnsiblePlaybook returns the AnsiblePlaybook field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetAnsiblePlaybook() string {
	if o == nil || o.AnsiblePlaybook.Get() == nil {
		var ret string
		return ret
	}
	return *o.AnsiblePlaybook.Get()
}

// GetAnsiblePlaybookOk returns a tuple with the AnsiblePlaybook field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetAnsiblePlaybookOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AnsiblePlaybook.Get(), o.AnsiblePlaybook.IsSet()
}

// HasAnsiblePlaybook returns a boolean if a field has been set.
func (o *TaskAnsiblePlaybookConfigTaskOptions) HasAnsiblePlaybook() bool {
	if o != nil && o.AnsiblePlaybook.IsSet() {
		return true
	}

	return false
}

// SetAnsiblePlaybook gets a reference to the given NullableString and assigns it to the AnsiblePlaybook field.
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetAnsiblePlaybook(v string) {
	o.AnsiblePlaybook.Set(&v)
}
// SetAnsiblePlaybookNil sets the value for AnsiblePlaybook to be an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetAnsiblePlaybookNil() {
	o.AnsiblePlaybook.Set(nil)
}

// UnsetAnsiblePlaybook ensures that no value is present for AnsiblePlaybook, not even an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) UnsetAnsiblePlaybook() {
	o.AnsiblePlaybook.Unset()
}

// GetSshKey returns the SshKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetSshKey() string {
	if o == nil || o.SshKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.SshKey.Get()
}

// GetSshKeyOk returns a tuple with the SshKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetSshKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SshKey.Get(), o.SshKey.IsSet()
}

// HasSshKey returns a boolean if a field has been set.
func (o *TaskAnsiblePlaybookConfigTaskOptions) HasSshKey() bool {
	if o != nil && o.SshKey.IsSet() {
		return true
	}

	return false
}

// SetSshKey gets a reference to the given NullableString and assigns it to the SshKey field.
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetSshKey(v string) {
	o.SshKey.Set(&v)
}
// SetSshKeyNil sets the value for SshKey to be an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetSshKeyNil() {
	o.SshKey.Set(nil)
}

// UnsetSshKey ensures that no value is present for SshKey, not even an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) UnsetSshKey() {
	o.SshKey.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetPort() string {
	if o == nil || o.Port.Get() == nil {
		var ret string
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetPortOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *TaskAnsiblePlaybookConfigTaskOptions) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableString and assigns it to the Port field.
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetPort(v string) {
	o.Port.Set(&v)
}
// SetPortNil sets the value for Port to be an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) UnsetPort() {
	o.Port.Unset()
}

// GetLocalScriptGitRef returns the LocalScriptGitRef field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetLocalScriptGitRef() string {
	if o == nil || o.LocalScriptGitRef.Get() == nil {
		var ret string
		return ret
	}
	return *o.LocalScriptGitRef.Get()
}

// GetLocalScriptGitRefOk returns a tuple with the LocalScriptGitRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetLocalScriptGitRefOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LocalScriptGitRef.Get(), o.LocalScriptGitRef.IsSet()
}

// HasLocalScriptGitRef returns a boolean if a field has been set.
func (o *TaskAnsiblePlaybookConfigTaskOptions) HasLocalScriptGitRef() bool {
	if o != nil && o.LocalScriptGitRef.IsSet() {
		return true
	}

	return false
}

// SetLocalScriptGitRef gets a reference to the given NullableString and assigns it to the LocalScriptGitRef field.
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetLocalScriptGitRef(v string) {
	o.LocalScriptGitRef.Set(&v)
}
// SetLocalScriptGitRefNil sets the value for LocalScriptGitRef to be an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetLocalScriptGitRefNil() {
	o.LocalScriptGitRef.Set(nil)
}

// UnsetLocalScriptGitRef ensures that no value is present for LocalScriptGitRef, not even an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) UnsetLocalScriptGitRef() {
	o.LocalScriptGitRef.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *TaskAnsiblePlaybookConfigTaskOptions) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) UnsetPassword() {
	o.Password.Unset()
}

// GetPasswordHash returns the PasswordHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetPasswordHash() string {
	if o == nil || o.PasswordHash.Get() == nil {
		var ret string
		return ret
	}
	return *o.PasswordHash.Get()
}

// GetPasswordHashOk returns a tuple with the PasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetPasswordHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PasswordHash.Get(), o.PasswordHash.IsSet()
}

// HasPasswordHash returns a boolean if a field has been set.
func (o *TaskAnsiblePlaybookConfigTaskOptions) HasPasswordHash() bool {
	if o != nil && o.PasswordHash.IsSet() {
		return true
	}

	return false
}

// SetPasswordHash gets a reference to the given NullableString and assigns it to the PasswordHash field.
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetPasswordHash(v string) {
	o.PasswordHash.Set(&v)
}
// SetPasswordHashNil sets the value for PasswordHash to be an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetPasswordHashNil() {
	o.PasswordHash.Set(nil)
}

// UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) UnsetPasswordHash() {
	o.PasswordHash.Unset()
}

// GetLocalScriptGitId returns the LocalScriptGitId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetLocalScriptGitId() string {
	if o == nil || o.LocalScriptGitId.Get() == nil {
		var ret string
		return ret
	}
	return *o.LocalScriptGitId.Get()
}

// GetLocalScriptGitIdOk returns a tuple with the LocalScriptGitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetLocalScriptGitIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LocalScriptGitId.Get(), o.LocalScriptGitId.IsSet()
}

// HasLocalScriptGitId returns a boolean if a field has been set.
func (o *TaskAnsiblePlaybookConfigTaskOptions) HasLocalScriptGitId() bool {
	if o != nil && o.LocalScriptGitId.IsSet() {
		return true
	}

	return false
}

// SetLocalScriptGitId gets a reference to the given NullableString and assigns it to the LocalScriptGitId field.
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetLocalScriptGitId(v string) {
	o.LocalScriptGitId.Set(&v)
}
// SetLocalScriptGitIdNil sets the value for LocalScriptGitId to be an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetLocalScriptGitIdNil() {
	o.LocalScriptGitId.Set(nil)
}

// UnsetLocalScriptGitId ensures that no value is present for LocalScriptGitId, not even an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) UnsetLocalScriptGitId() {
	o.LocalScriptGitId.Unset()
}

// GetAnsibleGitId returns the AnsibleGitId field value if set, zero value otherwise.
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetAnsibleGitId() string {
	if o == nil || o.AnsibleGitId == nil {
		var ret string
		return ret
	}
	return *o.AnsibleGitId
}

// GetAnsibleGitIdOk returns a tuple with the AnsibleGitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetAnsibleGitIdOk() (*string, bool) {
	if o == nil || o.AnsibleGitId == nil {
		return nil, false
	}
	return o.AnsibleGitId, true
}

// HasAnsibleGitId returns a boolean if a field has been set.
func (o *TaskAnsiblePlaybookConfigTaskOptions) HasAnsibleGitId() bool {
	if o != nil && o.AnsibleGitId != nil {
		return true
	}

	return false
}

// SetAnsibleGitId gets a reference to the given string and assigns it to the AnsibleGitId field.
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetAnsibleGitId(v string) {
	o.AnsibleGitId = &v
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetHost() string {
	if o == nil || o.Host.Get() == nil {
		var ret string
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *TaskAnsiblePlaybookConfigTaskOptions) HasHost() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableString and assigns it to the Host field.
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetHost(v string) {
	o.Host.Set(&v)
}
// SetHostNil sets the value for Host to be an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) UnsetHost() {
	o.Host.Unset()
}

// GetAnsibleSkipTags returns the AnsibleSkipTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetAnsibleSkipTags() string {
	if o == nil || o.AnsibleSkipTags.Get() == nil {
		var ret string
		return ret
	}
	return *o.AnsibleSkipTags.Get()
}

// GetAnsibleSkipTagsOk returns a tuple with the AnsibleSkipTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetAnsibleSkipTagsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AnsibleSkipTags.Get(), o.AnsibleSkipTags.IsSet()
}

// HasAnsibleSkipTags returns a boolean if a field has been set.
func (o *TaskAnsiblePlaybookConfigTaskOptions) HasAnsibleSkipTags() bool {
	if o != nil && o.AnsibleSkipTags.IsSet() {
		return true
	}

	return false
}

// SetAnsibleSkipTags gets a reference to the given NullableString and assigns it to the AnsibleSkipTags field.
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetAnsibleSkipTags(v string) {
	o.AnsibleSkipTags.Set(&v)
}
// SetAnsibleSkipTagsNil sets the value for AnsibleSkipTags to be an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetAnsibleSkipTagsNil() {
	o.AnsibleSkipTags.Set(nil)
}

// UnsetAnsibleSkipTags ensures that no value is present for AnsibleSkipTags, not even an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) UnsetAnsibleSkipTags() {
	o.AnsibleSkipTags.Unset()
}

// GetAnsibleTags returns the AnsibleTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetAnsibleTags() string {
	if o == nil || o.AnsibleTags.Get() == nil {
		var ret string
		return ret
	}
	return *o.AnsibleTags.Get()
}

// GetAnsibleTagsOk returns a tuple with the AnsibleTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetAnsibleTagsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AnsibleTags.Get(), o.AnsibleTags.IsSet()
}

// HasAnsibleTags returns a boolean if a field has been set.
func (o *TaskAnsiblePlaybookConfigTaskOptions) HasAnsibleTags() bool {
	if o != nil && o.AnsibleTags.IsSet() {
		return true
	}

	return false
}

// SetAnsibleTags gets a reference to the given NullableString and assigns it to the AnsibleTags field.
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetAnsibleTags(v string) {
	o.AnsibleTags.Set(&v)
}
// SetAnsibleTagsNil sets the value for AnsibleTags to be an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetAnsibleTagsNil() {
	o.AnsibleTags.Set(nil)
}

// UnsetAnsibleTags ensures that no value is present for AnsibleTags, not even an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) UnsetAnsibleTags() {
	o.AnsibleTags.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetUsername() string {
	if o == nil || o.Username.Get() == nil {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *TaskAnsiblePlaybookConfigTaskOptions) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) UnsetUsername() {
	o.Username.Unset()
}

// GetAnsibleGitRef returns the AnsibleGitRef field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetAnsibleGitRef() string {
	if o == nil || o.AnsibleGitRef.Get() == nil {
		var ret string
		return ret
	}
	return *o.AnsibleGitRef.Get()
}

// GetAnsibleGitRefOk returns a tuple with the AnsibleGitRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskAnsiblePlaybookConfigTaskOptions) GetAnsibleGitRefOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AnsibleGitRef.Get(), o.AnsibleGitRef.IsSet()
}

// HasAnsibleGitRef returns a boolean if a field has been set.
func (o *TaskAnsiblePlaybookConfigTaskOptions) HasAnsibleGitRef() bool {
	if o != nil && o.AnsibleGitRef.IsSet() {
		return true
	}

	return false
}

// SetAnsibleGitRef gets a reference to the given NullableString and assigns it to the AnsibleGitRef field.
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetAnsibleGitRef(v string) {
	o.AnsibleGitRef.Set(&v)
}
// SetAnsibleGitRefNil sets the value for AnsibleGitRef to be an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) SetAnsibleGitRefNil() {
	o.AnsibleGitRef.Set(nil)
}

// UnsetAnsibleGitRef ensures that no value is present for AnsibleGitRef, not even an explicit nil
func (o *TaskAnsiblePlaybookConfigTaskOptions) UnsetAnsibleGitRef() {
	o.AnsibleGitRef.Unset()
}

func (o TaskAnsiblePlaybookConfigTaskOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AnsibleOptions.IsSet() {
		toSerialize["ansibleOptions"] = o.AnsibleOptions.Get()
	}
	if o.AnsiblePlaybook.IsSet() {
		toSerialize["ansiblePlaybook"] = o.AnsiblePlaybook.Get()
	}
	if o.SshKey.IsSet() {
		toSerialize["sshKey"] = o.SshKey.Get()
	}
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
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
	if o.LocalScriptGitId.IsSet() {
		toSerialize["localScriptGitId"] = o.LocalScriptGitId.Get()
	}
	if o.AnsibleGitId != nil {
		toSerialize["ansibleGitId"] = o.AnsibleGitId
	}
	if o.Host.IsSet() {
		toSerialize["host"] = o.Host.Get()
	}
	if o.AnsibleSkipTags.IsSet() {
		toSerialize["ansibleSkipTags"] = o.AnsibleSkipTags.Get()
	}
	if o.AnsibleTags.IsSet() {
		toSerialize["ansibleTags"] = o.AnsibleTags.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.AnsibleGitRef.IsSet() {
		toSerialize["ansibleGitRef"] = o.AnsibleGitRef.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTaskAnsiblePlaybookConfigTaskOptions struct {
	value *TaskAnsiblePlaybookConfigTaskOptions
	isSet bool
}

func (v NullableTaskAnsiblePlaybookConfigTaskOptions) Get() *TaskAnsiblePlaybookConfigTaskOptions {
	return v.value
}

func (v *NullableTaskAnsiblePlaybookConfigTaskOptions) Set(val *TaskAnsiblePlaybookConfigTaskOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskAnsiblePlaybookConfigTaskOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskAnsiblePlaybookConfigTaskOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskAnsiblePlaybookConfigTaskOptions(val *TaskAnsiblePlaybookConfigTaskOptions) *NullableTaskAnsiblePlaybookConfigTaskOptions {
	return &NullableTaskAnsiblePlaybookConfigTaskOptions{value: val, isSet: true}
}

func (v NullableTaskAnsiblePlaybookConfigTaskOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskAnsiblePlaybookConfigTaskOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


