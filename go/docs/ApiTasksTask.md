# ApiTasksTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name for the task | 
**Code** | Pointer to **string** | A unique code for the task | [optional] 
**Visibility** | Pointer to **string** | private or public | [optional] [default to "private"]
**TaskType** | [**ApiTasksTaskTaskType**](_api_tasks_task_taskType.md) |  | 
**Labels** | Pointer to **[]string** | An array of Category labels for filtering | [optional] 
**TaskOptions** | Pointer to **map[string]interface{}** | Map of options specific to each &#x60;task type&#x60;. eg. script | [optional] 
**ResultType** | Pointer to **NullableString** |  | [optional] 
**ExecuteTarget** | **string** | The execution target. eg. local,remote,resource. The default value varies by task type.  | 
**Retryable** | Pointer to **bool** | If the task should be retried or not. | [optional] [default to false]
**RetryCount** | Pointer to **int32** | The number of times to retry. | [optional] 
**RetryDelaySeconds** | Pointer to **float32** | The delay, between retries. | [optional] 
**File** | Pointer to [**ApiTasksTaskFile**](_api_tasks_task_file.md) |  | [optional] 
**Credential** | Pointer to [**OneOfobjectobject**](oneOf&lt;object,object&gt;.md) | Map containing Credential ID or the default &#x60;{\&quot;type\&quot;: \&quot;local\&quot;}&#x60; which means use the values set in the local task options username and password instead of associating a credential.  | [optional] 

## Methods

### NewApiTasksTask

`func NewApiTasksTask(name string, taskType ApiTasksTaskTaskType, executeTarget string, ) *ApiTasksTask`

NewApiTasksTask instantiates a new ApiTasksTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTasksTaskWithDefaults

`func NewApiTasksTaskWithDefaults() *ApiTasksTask`

NewApiTasksTaskWithDefaults instantiates a new ApiTasksTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiTasksTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiTasksTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiTasksTask) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *ApiTasksTask) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiTasksTask) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiTasksTask) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiTasksTask) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetVisibility

`func (o *ApiTasksTask) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ApiTasksTask) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ApiTasksTask) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ApiTasksTask) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTaskType

`func (o *ApiTasksTask) GetTaskType() ApiTasksTaskTaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *ApiTasksTask) GetTaskTypeOk() (*ApiTasksTaskTaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *ApiTasksTask) SetTaskType(v ApiTasksTaskTaskType)`

SetTaskType sets TaskType field to given value.


### GetLabels

`func (o *ApiTasksTask) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ApiTasksTask) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ApiTasksTask) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ApiTasksTask) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTaskOptions

`func (o *ApiTasksTask) GetTaskOptions() map[string]interface{}`

GetTaskOptions returns the TaskOptions field if non-nil, zero value otherwise.

### GetTaskOptionsOk

`func (o *ApiTasksTask) GetTaskOptionsOk() (*map[string]interface{}, bool)`

GetTaskOptionsOk returns a tuple with the TaskOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskOptions

`func (o *ApiTasksTask) SetTaskOptions(v map[string]interface{})`

SetTaskOptions sets TaskOptions field to given value.

### HasTaskOptions

`func (o *ApiTasksTask) HasTaskOptions() bool`

HasTaskOptions returns a boolean if a field has been set.

### GetResultType

`func (o *ApiTasksTask) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *ApiTasksTask) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *ApiTasksTask) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *ApiTasksTask) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### SetResultTypeNil

`func (o *ApiTasksTask) SetResultTypeNil(b bool)`

 SetResultTypeNil sets the value for ResultType to be an explicit nil

### UnsetResultType
`func (o *ApiTasksTask) UnsetResultType()`

UnsetResultType ensures that no value is present for ResultType, not even an explicit nil
### GetExecuteTarget

`func (o *ApiTasksTask) GetExecuteTarget() string`

GetExecuteTarget returns the ExecuteTarget field if non-nil, zero value otherwise.

### GetExecuteTargetOk

`func (o *ApiTasksTask) GetExecuteTargetOk() (*string, bool)`

GetExecuteTargetOk returns a tuple with the ExecuteTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteTarget

`func (o *ApiTasksTask) SetExecuteTarget(v string)`

SetExecuteTarget sets ExecuteTarget field to given value.


### GetRetryable

`func (o *ApiTasksTask) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *ApiTasksTask) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *ApiTasksTask) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *ApiTasksTask) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetRetryCount

`func (o *ApiTasksTask) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *ApiTasksTask) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *ApiTasksTask) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *ApiTasksTask) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetRetryDelaySeconds

`func (o *ApiTasksTask) GetRetryDelaySeconds() float32`

GetRetryDelaySeconds returns the RetryDelaySeconds field if non-nil, zero value otherwise.

### GetRetryDelaySecondsOk

`func (o *ApiTasksTask) GetRetryDelaySecondsOk() (*float32, bool)`

GetRetryDelaySecondsOk returns a tuple with the RetryDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelaySeconds

`func (o *ApiTasksTask) SetRetryDelaySeconds(v float32)`

SetRetryDelaySeconds sets RetryDelaySeconds field to given value.

### HasRetryDelaySeconds

`func (o *ApiTasksTask) HasRetryDelaySeconds() bool`

HasRetryDelaySeconds returns a boolean if a field has been set.

### GetFile

`func (o *ApiTasksTask) GetFile() ApiTasksTaskFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ApiTasksTask) GetFileOk() (*ApiTasksTaskFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ApiTasksTask) SetFile(v ApiTasksTaskFile)`

SetFile sets File field to given value.

### HasFile

`func (o *ApiTasksTask) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetCredential

`func (o *ApiTasksTask) GetCredential() OneOfobjectobject`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *ApiTasksTask) GetCredentialOk() (*OneOfobjectobject, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *ApiTasksTask) SetCredential(v OneOfobjectobject)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *ApiTasksTask) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


