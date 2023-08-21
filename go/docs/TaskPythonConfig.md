# TaskPythonConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**TaskType** | Pointer to [**TaskPythonConfigTaskType**](taskPythonConfig_taskType.md) |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**TaskOptions** | Pointer to [**TaskPythonConfigTaskOptions**](taskPythonConfig_taskOptions.md) |  | [optional] 
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

### NewTaskPythonConfig

`func NewTaskPythonConfig() *TaskPythonConfig`

NewTaskPythonConfig instantiates a new TaskPythonConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskPythonConfigWithDefaults

`func NewTaskPythonConfigWithDefaults() *TaskPythonConfig`

NewTaskPythonConfigWithDefaults instantiates a new TaskPythonConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskPythonConfig) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskPythonConfig) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskPythonConfig) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TaskPythonConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *TaskPythonConfig) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TaskPythonConfig) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TaskPythonConfig) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *TaskPythonConfig) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *TaskPythonConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskPythonConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskPythonConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaskPythonConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *TaskPythonConfig) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TaskPythonConfig) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TaskPythonConfig) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TaskPythonConfig) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *TaskPythonConfig) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *TaskPythonConfig) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetTaskType

`func (o *TaskPythonConfig) GetTaskType() TaskPythonConfigTaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *TaskPythonConfig) GetTaskTypeOk() (*TaskPythonConfigTaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *TaskPythonConfig) SetTaskType(v TaskPythonConfigTaskType)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *TaskPythonConfig) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetLabels

`func (o *TaskPythonConfig) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *TaskPythonConfig) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *TaskPythonConfig) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *TaskPythonConfig) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetVisibility

`func (o *TaskPythonConfig) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *TaskPythonConfig) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *TaskPythonConfig) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *TaskPythonConfig) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTaskOptions

`func (o *TaskPythonConfig) GetTaskOptions() TaskPythonConfigTaskOptions`

GetTaskOptions returns the TaskOptions field if non-nil, zero value otherwise.

### GetTaskOptionsOk

`func (o *TaskPythonConfig) GetTaskOptionsOk() (*TaskPythonConfigTaskOptions, bool)`

GetTaskOptionsOk returns a tuple with the TaskOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskOptions

`func (o *TaskPythonConfig) SetTaskOptions(v TaskPythonConfigTaskOptions)`

SetTaskOptions sets TaskOptions field to given value.

### HasTaskOptions

`func (o *TaskPythonConfig) HasTaskOptions() bool`

HasTaskOptions returns a boolean if a field has been set.

### GetFile

`func (o *TaskPythonConfig) GetFile() TaskAnsiblePlaybookConfigFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *TaskPythonConfig) GetFileOk() (*TaskAnsiblePlaybookConfigFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *TaskPythonConfig) SetFile(v TaskAnsiblePlaybookConfigFile)`

SetFile sets File field to given value.

### HasFile

`func (o *TaskPythonConfig) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *TaskPythonConfig) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *TaskPythonConfig) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetResultType

`func (o *TaskPythonConfig) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *TaskPythonConfig) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *TaskPythonConfig) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *TaskPythonConfig) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### SetResultTypeNil

`func (o *TaskPythonConfig) SetResultTypeNil(b bool)`

 SetResultTypeNil sets the value for ResultType to be an explicit nil

### UnsetResultType
`func (o *TaskPythonConfig) UnsetResultType()`

UnsetResultType ensures that no value is present for ResultType, not even an explicit nil
### GetExecuteTarget

`func (o *TaskPythonConfig) GetExecuteTarget() string`

GetExecuteTarget returns the ExecuteTarget field if non-nil, zero value otherwise.

### GetExecuteTargetOk

`func (o *TaskPythonConfig) GetExecuteTargetOk() (*string, bool)`

GetExecuteTargetOk returns a tuple with the ExecuteTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteTarget

`func (o *TaskPythonConfig) SetExecuteTarget(v string)`

SetExecuteTarget sets ExecuteTarget field to given value.

### HasExecuteTarget

`func (o *TaskPythonConfig) HasExecuteTarget() bool`

HasExecuteTarget returns a boolean if a field has been set.

### GetRetryable

`func (o *TaskPythonConfig) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *TaskPythonConfig) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *TaskPythonConfig) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *TaskPythonConfig) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetRetryCount

`func (o *TaskPythonConfig) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *TaskPythonConfig) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *TaskPythonConfig) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *TaskPythonConfig) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetRetryDelaySeconds

`func (o *TaskPythonConfig) GetRetryDelaySeconds() int64`

GetRetryDelaySeconds returns the RetryDelaySeconds field if non-nil, zero value otherwise.

### GetRetryDelaySecondsOk

`func (o *TaskPythonConfig) GetRetryDelaySecondsOk() (*int64, bool)`

GetRetryDelaySecondsOk returns a tuple with the RetryDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelaySeconds

`func (o *TaskPythonConfig) SetRetryDelaySeconds(v int64)`

SetRetryDelaySeconds sets RetryDelaySeconds field to given value.

### HasRetryDelaySeconds

`func (o *TaskPythonConfig) HasRetryDelaySeconds() bool`

HasRetryDelaySeconds returns a boolean if a field has been set.

### GetAllowCustomConfig

`func (o *TaskPythonConfig) GetAllowCustomConfig() bool`

GetAllowCustomConfig returns the AllowCustomConfig field if non-nil, zero value otherwise.

### GetAllowCustomConfigOk

`func (o *TaskPythonConfig) GetAllowCustomConfigOk() (*bool, bool)`

GetAllowCustomConfigOk returns a tuple with the AllowCustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomConfig

`func (o *TaskPythonConfig) SetAllowCustomConfig(v bool)`

SetAllowCustomConfig sets AllowCustomConfig field to given value.

### HasAllowCustomConfig

`func (o *TaskPythonConfig) HasAllowCustomConfig() bool`

HasAllowCustomConfig returns a boolean if a field has been set.

### GetCredential

`func (o *TaskPythonConfig) GetCredential() OptionTypeListCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *TaskPythonConfig) GetCredentialOk() (*OptionTypeListCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *TaskPythonConfig) SetCredential(v OptionTypeListCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *TaskPythonConfig) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetDateCreated

`func (o *TaskPythonConfig) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TaskPythonConfig) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TaskPythonConfig) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TaskPythonConfig) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *TaskPythonConfig) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *TaskPythonConfig) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *TaskPythonConfig) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *TaskPythonConfig) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


