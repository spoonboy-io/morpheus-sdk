# ApiTasksIdTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A unique name for the task | [optional] 
**Code** | Pointer to **string** | A unique code for the task | [optional] 
**Visibility** | Pointer to **string** | Visibility | [optional] [default to "private"]
**TaskType** | Pointer to [**ApiTasksTaskTaskType**](_api_tasks_task_taskType.md) |  | [optional] 
**Labels** | Pointer to **[]string** | An array of Category labels for filtering | [optional] 
**TaskOptions** | Pointer to **map[string]interface{}** | Map of options specific to each &#x60;task type&#x60;. eg. script | [optional] 
**ResultType** | Pointer to **NullableString** |  | [optional] 
**ExecuteTarget** | Pointer to **string** | The execution target. eg. local,remote,resource. The default value varies by task type.  | [optional] 
**Retryable** | Pointer to **bool** | If the task should be retried or not. | [optional] [default to false]
**RetryCount** | Pointer to **int32** | The number of times to retry. | [optional] 
**RetryDelaySeconds** | Pointer to **float32** | The delay, between retries. | [optional] 
**File** | Pointer to [**ApiTasksTaskFile**](_api_tasks_task_file.md) |  | [optional] 
**Credential** | Pointer to [**OneOfobjectobject**](oneOf&lt;object,object&gt;.md) | Map containing Credential ID or the default {\&quot;type\&quot;: \&quot;local\&quot;}  which means use the values set in the local task options username and password instead of associating a credential.  | [optional] 

## Methods

### NewApiTasksIdTask

`func NewApiTasksIdTask() *ApiTasksIdTask`

NewApiTasksIdTask instantiates a new ApiTasksIdTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTasksIdTaskWithDefaults

`func NewApiTasksIdTaskWithDefaults() *ApiTasksIdTask`

NewApiTasksIdTaskWithDefaults instantiates a new ApiTasksIdTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiTasksIdTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiTasksIdTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiTasksIdTask) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiTasksIdTask) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ApiTasksIdTask) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiTasksIdTask) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiTasksIdTask) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiTasksIdTask) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetVisibility

`func (o *ApiTasksIdTask) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ApiTasksIdTask) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ApiTasksIdTask) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ApiTasksIdTask) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTaskType

`func (o *ApiTasksIdTask) GetTaskType() ApiTasksTaskTaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *ApiTasksIdTask) GetTaskTypeOk() (*ApiTasksTaskTaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *ApiTasksIdTask) SetTaskType(v ApiTasksTaskTaskType)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *ApiTasksIdTask) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetLabels

`func (o *ApiTasksIdTask) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ApiTasksIdTask) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ApiTasksIdTask) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ApiTasksIdTask) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTaskOptions

`func (o *ApiTasksIdTask) GetTaskOptions() map[string]interface{}`

GetTaskOptions returns the TaskOptions field if non-nil, zero value otherwise.

### GetTaskOptionsOk

`func (o *ApiTasksIdTask) GetTaskOptionsOk() (*map[string]interface{}, bool)`

GetTaskOptionsOk returns a tuple with the TaskOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskOptions

`func (o *ApiTasksIdTask) SetTaskOptions(v map[string]interface{})`

SetTaskOptions sets TaskOptions field to given value.

### HasTaskOptions

`func (o *ApiTasksIdTask) HasTaskOptions() bool`

HasTaskOptions returns a boolean if a field has been set.

### GetResultType

`func (o *ApiTasksIdTask) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *ApiTasksIdTask) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *ApiTasksIdTask) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *ApiTasksIdTask) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### SetResultTypeNil

`func (o *ApiTasksIdTask) SetResultTypeNil(b bool)`

 SetResultTypeNil sets the value for ResultType to be an explicit nil

### UnsetResultType
`func (o *ApiTasksIdTask) UnsetResultType()`

UnsetResultType ensures that no value is present for ResultType, not even an explicit nil
### GetExecuteTarget

`func (o *ApiTasksIdTask) GetExecuteTarget() string`

GetExecuteTarget returns the ExecuteTarget field if non-nil, zero value otherwise.

### GetExecuteTargetOk

`func (o *ApiTasksIdTask) GetExecuteTargetOk() (*string, bool)`

GetExecuteTargetOk returns a tuple with the ExecuteTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteTarget

`func (o *ApiTasksIdTask) SetExecuteTarget(v string)`

SetExecuteTarget sets ExecuteTarget field to given value.

### HasExecuteTarget

`func (o *ApiTasksIdTask) HasExecuteTarget() bool`

HasExecuteTarget returns a boolean if a field has been set.

### GetRetryable

`func (o *ApiTasksIdTask) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *ApiTasksIdTask) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *ApiTasksIdTask) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *ApiTasksIdTask) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetRetryCount

`func (o *ApiTasksIdTask) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *ApiTasksIdTask) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *ApiTasksIdTask) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *ApiTasksIdTask) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetRetryDelaySeconds

`func (o *ApiTasksIdTask) GetRetryDelaySeconds() float32`

GetRetryDelaySeconds returns the RetryDelaySeconds field if non-nil, zero value otherwise.

### GetRetryDelaySecondsOk

`func (o *ApiTasksIdTask) GetRetryDelaySecondsOk() (*float32, bool)`

GetRetryDelaySecondsOk returns a tuple with the RetryDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelaySeconds

`func (o *ApiTasksIdTask) SetRetryDelaySeconds(v float32)`

SetRetryDelaySeconds sets RetryDelaySeconds field to given value.

### HasRetryDelaySeconds

`func (o *ApiTasksIdTask) HasRetryDelaySeconds() bool`

HasRetryDelaySeconds returns a boolean if a field has been set.

### GetFile

`func (o *ApiTasksIdTask) GetFile() ApiTasksTaskFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ApiTasksIdTask) GetFileOk() (*ApiTasksTaskFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ApiTasksIdTask) SetFile(v ApiTasksTaskFile)`

SetFile sets File field to given value.

### HasFile

`func (o *ApiTasksIdTask) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetCredential

`func (o *ApiTasksIdTask) GetCredential() OneOfobjectobject`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *ApiTasksIdTask) GetCredentialOk() (*OneOfobjectobject, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *ApiTasksIdTask) SetCredential(v OneOfobjectobject)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *ApiTasksIdTask) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


