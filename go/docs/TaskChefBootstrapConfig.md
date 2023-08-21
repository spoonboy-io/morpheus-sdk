# TaskChefBootstrapConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**TaskType** | Pointer to [**TaskChefBootstrapConfigTaskType**](taskChefBootstrapConfig_taskType.md) |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**TaskOptions** | Pointer to [**TaskChefBootstrapConfigTaskOptions**](taskChefBootstrapConfig_taskOptions.md) |  | [optional] 
**File** | Pointer to [**NullableTaskAnsiblePlaybookConfigFile**](taskAnsiblePlaybookConfig_file.md) |  | [optional] 
**ResultType** | Pointer to **NullableString** |  | [optional] 
**ExecuteTarget** | Pointer to **string** |  | [optional] 
**Retryable** | Pointer to **bool** |  | [optional] 
**RetryCount** | Pointer to **int64** |  | [optional] 
**RetryDelaySeconds** | Pointer to **int64** |  | [optional] 
**AllowCustomConfig** | Pointer to **bool** |  | [optional] 
**Credential** | Pointer to [**OptionTypeListCredential**](optionTypeList_credential.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTaskChefBootstrapConfig

`func NewTaskChefBootstrapConfig() *TaskChefBootstrapConfig`

NewTaskChefBootstrapConfig instantiates a new TaskChefBootstrapConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskChefBootstrapConfigWithDefaults

`func NewTaskChefBootstrapConfigWithDefaults() *TaskChefBootstrapConfig`

NewTaskChefBootstrapConfigWithDefaults instantiates a new TaskChefBootstrapConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskChefBootstrapConfig) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskChefBootstrapConfig) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskChefBootstrapConfig) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TaskChefBootstrapConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *TaskChefBootstrapConfig) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TaskChefBootstrapConfig) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TaskChefBootstrapConfig) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *TaskChefBootstrapConfig) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *TaskChefBootstrapConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskChefBootstrapConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskChefBootstrapConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaskChefBootstrapConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *TaskChefBootstrapConfig) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TaskChefBootstrapConfig) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TaskChefBootstrapConfig) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TaskChefBootstrapConfig) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *TaskChefBootstrapConfig) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *TaskChefBootstrapConfig) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetTaskType

`func (o *TaskChefBootstrapConfig) GetTaskType() TaskChefBootstrapConfigTaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *TaskChefBootstrapConfig) GetTaskTypeOk() (*TaskChefBootstrapConfigTaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *TaskChefBootstrapConfig) SetTaskType(v TaskChefBootstrapConfigTaskType)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *TaskChefBootstrapConfig) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetLabels

`func (o *TaskChefBootstrapConfig) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *TaskChefBootstrapConfig) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *TaskChefBootstrapConfig) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *TaskChefBootstrapConfig) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetVisibility

`func (o *TaskChefBootstrapConfig) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *TaskChefBootstrapConfig) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *TaskChefBootstrapConfig) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *TaskChefBootstrapConfig) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTaskOptions

`func (o *TaskChefBootstrapConfig) GetTaskOptions() TaskChefBootstrapConfigTaskOptions`

GetTaskOptions returns the TaskOptions field if non-nil, zero value otherwise.

### GetTaskOptionsOk

`func (o *TaskChefBootstrapConfig) GetTaskOptionsOk() (*TaskChefBootstrapConfigTaskOptions, bool)`

GetTaskOptionsOk returns a tuple with the TaskOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskOptions

`func (o *TaskChefBootstrapConfig) SetTaskOptions(v TaskChefBootstrapConfigTaskOptions)`

SetTaskOptions sets TaskOptions field to given value.

### HasTaskOptions

`func (o *TaskChefBootstrapConfig) HasTaskOptions() bool`

HasTaskOptions returns a boolean if a field has been set.

### GetFile

`func (o *TaskChefBootstrapConfig) GetFile() TaskAnsiblePlaybookConfigFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *TaskChefBootstrapConfig) GetFileOk() (*TaskAnsiblePlaybookConfigFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *TaskChefBootstrapConfig) SetFile(v TaskAnsiblePlaybookConfigFile)`

SetFile sets File field to given value.

### HasFile

`func (o *TaskChefBootstrapConfig) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *TaskChefBootstrapConfig) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *TaskChefBootstrapConfig) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetResultType

`func (o *TaskChefBootstrapConfig) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *TaskChefBootstrapConfig) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *TaskChefBootstrapConfig) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *TaskChefBootstrapConfig) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### SetResultTypeNil

`func (o *TaskChefBootstrapConfig) SetResultTypeNil(b bool)`

 SetResultTypeNil sets the value for ResultType to be an explicit nil

### UnsetResultType
`func (o *TaskChefBootstrapConfig) UnsetResultType()`

UnsetResultType ensures that no value is present for ResultType, not even an explicit nil
### GetExecuteTarget

`func (o *TaskChefBootstrapConfig) GetExecuteTarget() string`

GetExecuteTarget returns the ExecuteTarget field if non-nil, zero value otherwise.

### GetExecuteTargetOk

`func (o *TaskChefBootstrapConfig) GetExecuteTargetOk() (*string, bool)`

GetExecuteTargetOk returns a tuple with the ExecuteTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteTarget

`func (o *TaskChefBootstrapConfig) SetExecuteTarget(v string)`

SetExecuteTarget sets ExecuteTarget field to given value.

### HasExecuteTarget

`func (o *TaskChefBootstrapConfig) HasExecuteTarget() bool`

HasExecuteTarget returns a boolean if a field has been set.

### GetRetryable

`func (o *TaskChefBootstrapConfig) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *TaskChefBootstrapConfig) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *TaskChefBootstrapConfig) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *TaskChefBootstrapConfig) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetRetryCount

`func (o *TaskChefBootstrapConfig) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *TaskChefBootstrapConfig) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *TaskChefBootstrapConfig) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *TaskChefBootstrapConfig) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetRetryDelaySeconds

`func (o *TaskChefBootstrapConfig) GetRetryDelaySeconds() int64`

GetRetryDelaySeconds returns the RetryDelaySeconds field if non-nil, zero value otherwise.

### GetRetryDelaySecondsOk

`func (o *TaskChefBootstrapConfig) GetRetryDelaySecondsOk() (*int64, bool)`

GetRetryDelaySecondsOk returns a tuple with the RetryDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelaySeconds

`func (o *TaskChefBootstrapConfig) SetRetryDelaySeconds(v int64)`

SetRetryDelaySeconds sets RetryDelaySeconds field to given value.

### HasRetryDelaySeconds

`func (o *TaskChefBootstrapConfig) HasRetryDelaySeconds() bool`

HasRetryDelaySeconds returns a boolean if a field has been set.

### GetAllowCustomConfig

`func (o *TaskChefBootstrapConfig) GetAllowCustomConfig() bool`

GetAllowCustomConfig returns the AllowCustomConfig field if non-nil, zero value otherwise.

### GetAllowCustomConfigOk

`func (o *TaskChefBootstrapConfig) GetAllowCustomConfigOk() (*bool, bool)`

GetAllowCustomConfigOk returns a tuple with the AllowCustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomConfig

`func (o *TaskChefBootstrapConfig) SetAllowCustomConfig(v bool)`

SetAllowCustomConfig sets AllowCustomConfig field to given value.

### HasAllowCustomConfig

`func (o *TaskChefBootstrapConfig) HasAllowCustomConfig() bool`

HasAllowCustomConfig returns a boolean if a field has been set.

### GetCredential

`func (o *TaskChefBootstrapConfig) GetCredential() OptionTypeListCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *TaskChefBootstrapConfig) GetCredentialOk() (*OptionTypeListCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *TaskChefBootstrapConfig) SetCredential(v OptionTypeListCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *TaskChefBootstrapConfig) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetDateCreated

`func (o *TaskChefBootstrapConfig) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TaskChefBootstrapConfig) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TaskChefBootstrapConfig) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TaskChefBootstrapConfig) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *TaskChefBootstrapConfig) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *TaskChefBootstrapConfig) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *TaskChefBootstrapConfig) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *TaskChefBootstrapConfig) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


