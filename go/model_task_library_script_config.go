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
	"time"
)

// TaskLibraryScriptConfig struct for TaskLibraryScriptConfig
type TaskLibraryScriptConfig struct {
	Id *int64 `json:"id,omitempty"`
	AccountId *int64 `json:"accountId,omitempty"`
	Name *string `json:"name,omitempty"`
	Code NullableString `json:"code,omitempty"`
	TaskType *TaskLibraryScriptConfigTaskType `json:"taskType,omitempty"`
	Labels *[]string `json:"labels,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
	TaskOptions *TaskLibraryScriptConfigTaskOptions `json:"taskOptions,omitempty"`
	File NullableTaskAnsiblePlaybookConfigFile `json:"file,omitempty"`
	ResultType NullableString `json:"resultType,omitempty"`
	ExecuteTarget *string `json:"executeTarget,omitempty"`
	Retryable *bool `json:"retryable,omitempty"`
	RetryCount *int64 `json:"retryCount,omitempty"`
	RetryDelaySeconds *int64 `json:"retryDelaySeconds,omitempty"`
	AllowCustomConfig *bool `json:"allowCustomConfig,omitempty"`
	Credential *OptionTypeListCredential `json:"credential,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
}

// NewTaskLibraryScriptConfig instantiates a new TaskLibraryScriptConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskLibraryScriptConfig() *TaskLibraryScriptConfig {
	this := TaskLibraryScriptConfig{}
	return &this
}

// NewTaskLibraryScriptConfigWithDefaults instantiates a new TaskLibraryScriptConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskLibraryScriptConfigWithDefaults() *TaskLibraryScriptConfig {
	this := TaskLibraryScriptConfig{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TaskLibraryScriptConfig) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskLibraryScriptConfig) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfig) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *TaskLibraryScriptConfig) SetId(v int64) {
	o.Id = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *TaskLibraryScriptConfig) GetAccountId() int64 {
	if o == nil || o.AccountId == nil {
		var ret int64
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskLibraryScriptConfig) GetAccountIdOk() (*int64, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfig) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given int64 and assigns it to the AccountId field.
func (o *TaskLibraryScriptConfig) SetAccountId(v int64) {
	o.AccountId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TaskLibraryScriptConfig) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskLibraryScriptConfig) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfig) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TaskLibraryScriptConfig) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskLibraryScriptConfig) GetCode() string {
	if o == nil || o.Code.Get() == nil {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskLibraryScriptConfig) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfig) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *TaskLibraryScriptConfig) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *TaskLibraryScriptConfig) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *TaskLibraryScriptConfig) UnsetCode() {
	o.Code.Unset()
}

// GetTaskType returns the TaskType field value if set, zero value otherwise.
func (o *TaskLibraryScriptConfig) GetTaskType() TaskLibraryScriptConfigTaskType {
	if o == nil || o.TaskType == nil {
		var ret TaskLibraryScriptConfigTaskType
		return ret
	}
	return *o.TaskType
}

// GetTaskTypeOk returns a tuple with the TaskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskLibraryScriptConfig) GetTaskTypeOk() (*TaskLibraryScriptConfigTaskType, bool) {
	if o == nil || o.TaskType == nil {
		return nil, false
	}
	return o.TaskType, true
}

// HasTaskType returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfig) HasTaskType() bool {
	if o != nil && o.TaskType != nil {
		return true
	}

	return false
}

// SetTaskType gets a reference to the given TaskLibraryScriptConfigTaskType and assigns it to the TaskType field.
func (o *TaskLibraryScriptConfig) SetTaskType(v TaskLibraryScriptConfigTaskType) {
	o.TaskType = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *TaskLibraryScriptConfig) GetLabels() []string {
	if o == nil || o.Labels == nil {
		var ret []string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskLibraryScriptConfig) GetLabelsOk() (*[]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfig) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *TaskLibraryScriptConfig) SetLabels(v []string) {
	o.Labels = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *TaskLibraryScriptConfig) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskLibraryScriptConfig) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfig) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *TaskLibraryScriptConfig) SetVisibility(v string) {
	o.Visibility = &v
}

// GetTaskOptions returns the TaskOptions field value if set, zero value otherwise.
func (o *TaskLibraryScriptConfig) GetTaskOptions() TaskLibraryScriptConfigTaskOptions {
	if o == nil || o.TaskOptions == nil {
		var ret TaskLibraryScriptConfigTaskOptions
		return ret
	}
	return *o.TaskOptions
}

// GetTaskOptionsOk returns a tuple with the TaskOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskLibraryScriptConfig) GetTaskOptionsOk() (*TaskLibraryScriptConfigTaskOptions, bool) {
	if o == nil || o.TaskOptions == nil {
		return nil, false
	}
	return o.TaskOptions, true
}

// HasTaskOptions returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfig) HasTaskOptions() bool {
	if o != nil && o.TaskOptions != nil {
		return true
	}

	return false
}

// SetTaskOptions gets a reference to the given TaskLibraryScriptConfigTaskOptions and assigns it to the TaskOptions field.
func (o *TaskLibraryScriptConfig) SetTaskOptions(v TaskLibraryScriptConfigTaskOptions) {
	o.TaskOptions = &v
}

// GetFile returns the File field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskLibraryScriptConfig) GetFile() TaskAnsiblePlaybookConfigFile {
	if o == nil || o.File.Get() == nil {
		var ret TaskAnsiblePlaybookConfigFile
		return ret
	}
	return *o.File.Get()
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskLibraryScriptConfig) GetFileOk() (*TaskAnsiblePlaybookConfigFile, bool) {
	if o == nil  {
		return nil, false
	}
	return o.File.Get(), o.File.IsSet()
}

// HasFile returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfig) HasFile() bool {
	if o != nil && o.File.IsSet() {
		return true
	}

	return false
}

// SetFile gets a reference to the given NullableTaskAnsiblePlaybookConfigFile and assigns it to the File field.
func (o *TaskLibraryScriptConfig) SetFile(v TaskAnsiblePlaybookConfigFile) {
	o.File.Set(&v)
}
// SetFileNil sets the value for File to be an explicit nil
func (o *TaskLibraryScriptConfig) SetFileNil() {
	o.File.Set(nil)
}

// UnsetFile ensures that no value is present for File, not even an explicit nil
func (o *TaskLibraryScriptConfig) UnsetFile() {
	o.File.Unset()
}

// GetResultType returns the ResultType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskLibraryScriptConfig) GetResultType() string {
	if o == nil || o.ResultType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResultType.Get()
}

// GetResultTypeOk returns a tuple with the ResultType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskLibraryScriptConfig) GetResultTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResultType.Get(), o.ResultType.IsSet()
}

// HasResultType returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfig) HasResultType() bool {
	if o != nil && o.ResultType.IsSet() {
		return true
	}

	return false
}

// SetResultType gets a reference to the given NullableString and assigns it to the ResultType field.
func (o *TaskLibraryScriptConfig) SetResultType(v string) {
	o.ResultType.Set(&v)
}
// SetResultTypeNil sets the value for ResultType to be an explicit nil
func (o *TaskLibraryScriptConfig) SetResultTypeNil() {
	o.ResultType.Set(nil)
}

// UnsetResultType ensures that no value is present for ResultType, not even an explicit nil
func (o *TaskLibraryScriptConfig) UnsetResultType() {
	o.ResultType.Unset()
}

// GetExecuteTarget returns the ExecuteTarget field value if set, zero value otherwise.
func (o *TaskLibraryScriptConfig) GetExecuteTarget() string {
	if o == nil || o.ExecuteTarget == nil {
		var ret string
		return ret
	}
	return *o.ExecuteTarget
}

// GetExecuteTargetOk returns a tuple with the ExecuteTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskLibraryScriptConfig) GetExecuteTargetOk() (*string, bool) {
	if o == nil || o.ExecuteTarget == nil {
		return nil, false
	}
	return o.ExecuteTarget, true
}

// HasExecuteTarget returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfig) HasExecuteTarget() bool {
	if o != nil && o.ExecuteTarget != nil {
		return true
	}

	return false
}

// SetExecuteTarget gets a reference to the given string and assigns it to the ExecuteTarget field.
func (o *TaskLibraryScriptConfig) SetExecuteTarget(v string) {
	o.ExecuteTarget = &v
}

// GetRetryable returns the Retryable field value if set, zero value otherwise.
func (o *TaskLibraryScriptConfig) GetRetryable() bool {
	if o == nil || o.Retryable == nil {
		var ret bool
		return ret
	}
	return *o.Retryable
}

// GetRetryableOk returns a tuple with the Retryable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskLibraryScriptConfig) GetRetryableOk() (*bool, bool) {
	if o == nil || o.Retryable == nil {
		return nil, false
	}
	return o.Retryable, true
}

// HasRetryable returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfig) HasRetryable() bool {
	if o != nil && o.Retryable != nil {
		return true
	}

	return false
}

// SetRetryable gets a reference to the given bool and assigns it to the Retryable field.
func (o *TaskLibraryScriptConfig) SetRetryable(v bool) {
	o.Retryable = &v
}

// GetRetryCount returns the RetryCount field value if set, zero value otherwise.
func (o *TaskLibraryScriptConfig) GetRetryCount() int64 {
	if o == nil || o.RetryCount == nil {
		var ret int64
		return ret
	}
	return *o.RetryCount
}

// GetRetryCountOk returns a tuple with the RetryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskLibraryScriptConfig) GetRetryCountOk() (*int64, bool) {
	if o == nil || o.RetryCount == nil {
		return nil, false
	}
	return o.RetryCount, true
}

// HasRetryCount returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfig) HasRetryCount() bool {
	if o != nil && o.RetryCount != nil {
		return true
	}

	return false
}

// SetRetryCount gets a reference to the given int64 and assigns it to the RetryCount field.
func (o *TaskLibraryScriptConfig) SetRetryCount(v int64) {
	o.RetryCount = &v
}

// GetRetryDelaySeconds returns the RetryDelaySeconds field value if set, zero value otherwise.
func (o *TaskLibraryScriptConfig) GetRetryDelaySeconds() int64 {
	if o == nil || o.RetryDelaySeconds == nil {
		var ret int64
		return ret
	}
	return *o.RetryDelaySeconds
}

// GetRetryDelaySecondsOk returns a tuple with the RetryDelaySeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskLibraryScriptConfig) GetRetryDelaySecondsOk() (*int64, bool) {
	if o == nil || o.RetryDelaySeconds == nil {
		return nil, false
	}
	return o.RetryDelaySeconds, true
}

// HasRetryDelaySeconds returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfig) HasRetryDelaySeconds() bool {
	if o != nil && o.RetryDelaySeconds != nil {
		return true
	}

	return false
}

// SetRetryDelaySeconds gets a reference to the given int64 and assigns it to the RetryDelaySeconds field.
func (o *TaskLibraryScriptConfig) SetRetryDelaySeconds(v int64) {
	o.RetryDelaySeconds = &v
}

// GetAllowCustomConfig returns the AllowCustomConfig field value if set, zero value otherwise.
func (o *TaskLibraryScriptConfig) GetAllowCustomConfig() bool {
	if o == nil || o.AllowCustomConfig == nil {
		var ret bool
		return ret
	}
	return *o.AllowCustomConfig
}

// GetAllowCustomConfigOk returns a tuple with the AllowCustomConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskLibraryScriptConfig) GetAllowCustomConfigOk() (*bool, bool) {
	if o == nil || o.AllowCustomConfig == nil {
		return nil, false
	}
	return o.AllowCustomConfig, true
}

// HasAllowCustomConfig returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfig) HasAllowCustomConfig() bool {
	if o != nil && o.AllowCustomConfig != nil {
		return true
	}

	return false
}

// SetAllowCustomConfig gets a reference to the given bool and assigns it to the AllowCustomConfig field.
func (o *TaskLibraryScriptConfig) SetAllowCustomConfig(v bool) {
	o.AllowCustomConfig = &v
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *TaskLibraryScriptConfig) GetCredential() OptionTypeListCredential {
	if o == nil || o.Credential == nil {
		var ret OptionTypeListCredential
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskLibraryScriptConfig) GetCredentialOk() (*OptionTypeListCredential, bool) {
	if o == nil || o.Credential == nil {
		return nil, false
	}
	return o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfig) HasCredential() bool {
	if o != nil && o.Credential != nil {
		return true
	}

	return false
}

// SetCredential gets a reference to the given OptionTypeListCredential and assigns it to the Credential field.
func (o *TaskLibraryScriptConfig) SetCredential(v OptionTypeListCredential) {
	o.Credential = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *TaskLibraryScriptConfig) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskLibraryScriptConfig) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfig) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *TaskLibraryScriptConfig) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *TaskLibraryScriptConfig) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskLibraryScriptConfig) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *TaskLibraryScriptConfig) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *TaskLibraryScriptConfig) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o TaskLibraryScriptConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AccountId != nil {
		toSerialize["accountId"] = o.AccountId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.TaskType != nil {
		toSerialize["taskType"] = o.TaskType
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.TaskOptions != nil {
		toSerialize["taskOptions"] = o.TaskOptions
	}
	if o.File.IsSet() {
		toSerialize["file"] = o.File.Get()
	}
	if o.ResultType.IsSet() {
		toSerialize["resultType"] = o.ResultType.Get()
	}
	if o.ExecuteTarget != nil {
		toSerialize["executeTarget"] = o.ExecuteTarget
	}
	if o.Retryable != nil {
		toSerialize["retryable"] = o.Retryable
	}
	if o.RetryCount != nil {
		toSerialize["retryCount"] = o.RetryCount
	}
	if o.RetryDelaySeconds != nil {
		toSerialize["retryDelaySeconds"] = o.RetryDelaySeconds
	}
	if o.AllowCustomConfig != nil {
		toSerialize["allowCustomConfig"] = o.AllowCustomConfig
	}
	if o.Credential != nil {
		toSerialize["credential"] = o.Credential
	}
	if o.DateCreated != nil {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	return json.Marshal(toSerialize)
}

type NullableTaskLibraryScriptConfig struct {
	value *TaskLibraryScriptConfig
	isSet bool
}

func (v NullableTaskLibraryScriptConfig) Get() *TaskLibraryScriptConfig {
	return v.value
}

func (v *NullableTaskLibraryScriptConfig) Set(val *TaskLibraryScriptConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskLibraryScriptConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskLibraryScriptConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskLibraryScriptConfig(val *TaskLibraryScriptConfig) *NullableTaskLibraryScriptConfig {
	return &NullableTaskLibraryScriptConfig{value: val, isSet: true}
}

func (v NullableTaskLibraryScriptConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskLibraryScriptConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


