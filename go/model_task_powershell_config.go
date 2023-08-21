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

// TaskPowershellConfig struct for TaskPowershellConfig
type TaskPowershellConfig struct {
	Id *int64 `json:"id,omitempty"`
	AccountId *int64 `json:"accountId,omitempty"`
	Name *string `json:"name,omitempty"`
	Code NullableString `json:"code,omitempty"`
	TaskType *TaskPowershellConfigTaskType `json:"taskType,omitempty"`
	Labels *[]string `json:"labels,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
	TaskOptions *TaskPowershellConfigTaskOptions `json:"taskOptions,omitempty"`
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

// NewTaskPowershellConfig instantiates a new TaskPowershellConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskPowershellConfig() *TaskPowershellConfig {
	this := TaskPowershellConfig{}
	return &this
}

// NewTaskPowershellConfigWithDefaults instantiates a new TaskPowershellConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskPowershellConfigWithDefaults() *TaskPowershellConfig {
	this := TaskPowershellConfig{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TaskPowershellConfig) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPowershellConfig) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TaskPowershellConfig) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *TaskPowershellConfig) SetId(v int64) {
	o.Id = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *TaskPowershellConfig) GetAccountId() int64 {
	if o == nil || o.AccountId == nil {
		var ret int64
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPowershellConfig) GetAccountIdOk() (*int64, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *TaskPowershellConfig) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given int64 and assigns it to the AccountId field.
func (o *TaskPowershellConfig) SetAccountId(v int64) {
	o.AccountId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TaskPowershellConfig) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPowershellConfig) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TaskPowershellConfig) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TaskPowershellConfig) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskPowershellConfig) GetCode() string {
	if o == nil || o.Code.Get() == nil {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskPowershellConfig) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *TaskPowershellConfig) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *TaskPowershellConfig) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *TaskPowershellConfig) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *TaskPowershellConfig) UnsetCode() {
	o.Code.Unset()
}

// GetTaskType returns the TaskType field value if set, zero value otherwise.
func (o *TaskPowershellConfig) GetTaskType() TaskPowershellConfigTaskType {
	if o == nil || o.TaskType == nil {
		var ret TaskPowershellConfigTaskType
		return ret
	}
	return *o.TaskType
}

// GetTaskTypeOk returns a tuple with the TaskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPowershellConfig) GetTaskTypeOk() (*TaskPowershellConfigTaskType, bool) {
	if o == nil || o.TaskType == nil {
		return nil, false
	}
	return o.TaskType, true
}

// HasTaskType returns a boolean if a field has been set.
func (o *TaskPowershellConfig) HasTaskType() bool {
	if o != nil && o.TaskType != nil {
		return true
	}

	return false
}

// SetTaskType gets a reference to the given TaskPowershellConfigTaskType and assigns it to the TaskType field.
func (o *TaskPowershellConfig) SetTaskType(v TaskPowershellConfigTaskType) {
	o.TaskType = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *TaskPowershellConfig) GetLabels() []string {
	if o == nil || o.Labels == nil {
		var ret []string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPowershellConfig) GetLabelsOk() (*[]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *TaskPowershellConfig) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *TaskPowershellConfig) SetLabels(v []string) {
	o.Labels = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *TaskPowershellConfig) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPowershellConfig) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *TaskPowershellConfig) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *TaskPowershellConfig) SetVisibility(v string) {
	o.Visibility = &v
}

// GetTaskOptions returns the TaskOptions field value if set, zero value otherwise.
func (o *TaskPowershellConfig) GetTaskOptions() TaskPowershellConfigTaskOptions {
	if o == nil || o.TaskOptions == nil {
		var ret TaskPowershellConfigTaskOptions
		return ret
	}
	return *o.TaskOptions
}

// GetTaskOptionsOk returns a tuple with the TaskOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPowershellConfig) GetTaskOptionsOk() (*TaskPowershellConfigTaskOptions, bool) {
	if o == nil || o.TaskOptions == nil {
		return nil, false
	}
	return o.TaskOptions, true
}

// HasTaskOptions returns a boolean if a field has been set.
func (o *TaskPowershellConfig) HasTaskOptions() bool {
	if o != nil && o.TaskOptions != nil {
		return true
	}

	return false
}

// SetTaskOptions gets a reference to the given TaskPowershellConfigTaskOptions and assigns it to the TaskOptions field.
func (o *TaskPowershellConfig) SetTaskOptions(v TaskPowershellConfigTaskOptions) {
	o.TaskOptions = &v
}

// GetFile returns the File field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskPowershellConfig) GetFile() TaskAnsiblePlaybookConfigFile {
	if o == nil || o.File.Get() == nil {
		var ret TaskAnsiblePlaybookConfigFile
		return ret
	}
	return *o.File.Get()
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskPowershellConfig) GetFileOk() (*TaskAnsiblePlaybookConfigFile, bool) {
	if o == nil  {
		return nil, false
	}
	return o.File.Get(), o.File.IsSet()
}

// HasFile returns a boolean if a field has been set.
func (o *TaskPowershellConfig) HasFile() bool {
	if o != nil && o.File.IsSet() {
		return true
	}

	return false
}

// SetFile gets a reference to the given NullableTaskAnsiblePlaybookConfigFile and assigns it to the File field.
func (o *TaskPowershellConfig) SetFile(v TaskAnsiblePlaybookConfigFile) {
	o.File.Set(&v)
}
// SetFileNil sets the value for File to be an explicit nil
func (o *TaskPowershellConfig) SetFileNil() {
	o.File.Set(nil)
}

// UnsetFile ensures that no value is present for File, not even an explicit nil
func (o *TaskPowershellConfig) UnsetFile() {
	o.File.Unset()
}

// GetResultType returns the ResultType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskPowershellConfig) GetResultType() string {
	if o == nil || o.ResultType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResultType.Get()
}

// GetResultTypeOk returns a tuple with the ResultType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskPowershellConfig) GetResultTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResultType.Get(), o.ResultType.IsSet()
}

// HasResultType returns a boolean if a field has been set.
func (o *TaskPowershellConfig) HasResultType() bool {
	if o != nil && o.ResultType.IsSet() {
		return true
	}

	return false
}

// SetResultType gets a reference to the given NullableString and assigns it to the ResultType field.
func (o *TaskPowershellConfig) SetResultType(v string) {
	o.ResultType.Set(&v)
}
// SetResultTypeNil sets the value for ResultType to be an explicit nil
func (o *TaskPowershellConfig) SetResultTypeNil() {
	o.ResultType.Set(nil)
}

// UnsetResultType ensures that no value is present for ResultType, not even an explicit nil
func (o *TaskPowershellConfig) UnsetResultType() {
	o.ResultType.Unset()
}

// GetExecuteTarget returns the ExecuteTarget field value if set, zero value otherwise.
func (o *TaskPowershellConfig) GetExecuteTarget() string {
	if o == nil || o.ExecuteTarget == nil {
		var ret string
		return ret
	}
	return *o.ExecuteTarget
}

// GetExecuteTargetOk returns a tuple with the ExecuteTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPowershellConfig) GetExecuteTargetOk() (*string, bool) {
	if o == nil || o.ExecuteTarget == nil {
		return nil, false
	}
	return o.ExecuteTarget, true
}

// HasExecuteTarget returns a boolean if a field has been set.
func (o *TaskPowershellConfig) HasExecuteTarget() bool {
	if o != nil && o.ExecuteTarget != nil {
		return true
	}

	return false
}

// SetExecuteTarget gets a reference to the given string and assigns it to the ExecuteTarget field.
func (o *TaskPowershellConfig) SetExecuteTarget(v string) {
	o.ExecuteTarget = &v
}

// GetRetryable returns the Retryable field value if set, zero value otherwise.
func (o *TaskPowershellConfig) GetRetryable() bool {
	if o == nil || o.Retryable == nil {
		var ret bool
		return ret
	}
	return *o.Retryable
}

// GetRetryableOk returns a tuple with the Retryable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPowershellConfig) GetRetryableOk() (*bool, bool) {
	if o == nil || o.Retryable == nil {
		return nil, false
	}
	return o.Retryable, true
}

// HasRetryable returns a boolean if a field has been set.
func (o *TaskPowershellConfig) HasRetryable() bool {
	if o != nil && o.Retryable != nil {
		return true
	}

	return false
}

// SetRetryable gets a reference to the given bool and assigns it to the Retryable field.
func (o *TaskPowershellConfig) SetRetryable(v bool) {
	o.Retryable = &v
}

// GetRetryCount returns the RetryCount field value if set, zero value otherwise.
func (o *TaskPowershellConfig) GetRetryCount() int64 {
	if o == nil || o.RetryCount == nil {
		var ret int64
		return ret
	}
	return *o.RetryCount
}

// GetRetryCountOk returns a tuple with the RetryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPowershellConfig) GetRetryCountOk() (*int64, bool) {
	if o == nil || o.RetryCount == nil {
		return nil, false
	}
	return o.RetryCount, true
}

// HasRetryCount returns a boolean if a field has been set.
func (o *TaskPowershellConfig) HasRetryCount() bool {
	if o != nil && o.RetryCount != nil {
		return true
	}

	return false
}

// SetRetryCount gets a reference to the given int64 and assigns it to the RetryCount field.
func (o *TaskPowershellConfig) SetRetryCount(v int64) {
	o.RetryCount = &v
}

// GetRetryDelaySeconds returns the RetryDelaySeconds field value if set, zero value otherwise.
func (o *TaskPowershellConfig) GetRetryDelaySeconds() int64 {
	if o == nil || o.RetryDelaySeconds == nil {
		var ret int64
		return ret
	}
	return *o.RetryDelaySeconds
}

// GetRetryDelaySecondsOk returns a tuple with the RetryDelaySeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPowershellConfig) GetRetryDelaySecondsOk() (*int64, bool) {
	if o == nil || o.RetryDelaySeconds == nil {
		return nil, false
	}
	return o.RetryDelaySeconds, true
}

// HasRetryDelaySeconds returns a boolean if a field has been set.
func (o *TaskPowershellConfig) HasRetryDelaySeconds() bool {
	if o != nil && o.RetryDelaySeconds != nil {
		return true
	}

	return false
}

// SetRetryDelaySeconds gets a reference to the given int64 and assigns it to the RetryDelaySeconds field.
func (o *TaskPowershellConfig) SetRetryDelaySeconds(v int64) {
	o.RetryDelaySeconds = &v
}

// GetAllowCustomConfig returns the AllowCustomConfig field value if set, zero value otherwise.
func (o *TaskPowershellConfig) GetAllowCustomConfig() bool {
	if o == nil || o.AllowCustomConfig == nil {
		var ret bool
		return ret
	}
	return *o.AllowCustomConfig
}

// GetAllowCustomConfigOk returns a tuple with the AllowCustomConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPowershellConfig) GetAllowCustomConfigOk() (*bool, bool) {
	if o == nil || o.AllowCustomConfig == nil {
		return nil, false
	}
	return o.AllowCustomConfig, true
}

// HasAllowCustomConfig returns a boolean if a field has been set.
func (o *TaskPowershellConfig) HasAllowCustomConfig() bool {
	if o != nil && o.AllowCustomConfig != nil {
		return true
	}

	return false
}

// SetAllowCustomConfig gets a reference to the given bool and assigns it to the AllowCustomConfig field.
func (o *TaskPowershellConfig) SetAllowCustomConfig(v bool) {
	o.AllowCustomConfig = &v
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *TaskPowershellConfig) GetCredential() OptionTypeListCredential {
	if o == nil || o.Credential == nil {
		var ret OptionTypeListCredential
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPowershellConfig) GetCredentialOk() (*OptionTypeListCredential, bool) {
	if o == nil || o.Credential == nil {
		return nil, false
	}
	return o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *TaskPowershellConfig) HasCredential() bool {
	if o != nil && o.Credential != nil {
		return true
	}

	return false
}

// SetCredential gets a reference to the given OptionTypeListCredential and assigns it to the Credential field.
func (o *TaskPowershellConfig) SetCredential(v OptionTypeListCredential) {
	o.Credential = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *TaskPowershellConfig) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPowershellConfig) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *TaskPowershellConfig) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *TaskPowershellConfig) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *TaskPowershellConfig) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPowershellConfig) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *TaskPowershellConfig) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *TaskPowershellConfig) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o TaskPowershellConfig) MarshalJSON() ([]byte, error) {
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

type NullableTaskPowershellConfig struct {
	value *TaskPowershellConfig
	isSet bool
}

func (v NullableTaskPowershellConfig) Get() *TaskPowershellConfig {
	return v.value
}

func (v *NullableTaskPowershellConfig) Set(val *TaskPowershellConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskPowershellConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskPowershellConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskPowershellConfig(val *TaskPowershellConfig) *NullableTaskPowershellConfig {
	return &NullableTaskPowershellConfig{value: val, isSet: true}
}

func (v NullableTaskPowershellConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskPowershellConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


