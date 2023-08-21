# WorkflowTaskSetTasks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**TaskPhase** | Pointer to **string** |  | [optional] 
**TaskOrder** | Pointer to **int64** |  | [optional] 
**Task** | Pointer to [**WorkflowTask**](workflow_task.md) |  | [optional] 

## Methods

### NewWorkflowTaskSetTasks

`func NewWorkflowTaskSetTasks() *WorkflowTaskSetTasks`

NewWorkflowTaskSetTasks instantiates a new WorkflowTaskSetTasks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskSetTasksWithDefaults

`func NewWorkflowTaskSetTasksWithDefaults() *WorkflowTaskSetTasks`

NewWorkflowTaskSetTasksWithDefaults instantiates a new WorkflowTaskSetTasks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowTaskSetTasks) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowTaskSetTasks) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowTaskSetTasks) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowTaskSetTasks) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTaskPhase

`func (o *WorkflowTaskSetTasks) GetTaskPhase() string`

GetTaskPhase returns the TaskPhase field if non-nil, zero value otherwise.

### GetTaskPhaseOk

`func (o *WorkflowTaskSetTasks) GetTaskPhaseOk() (*string, bool)`

GetTaskPhaseOk returns a tuple with the TaskPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPhase

`func (o *WorkflowTaskSetTasks) SetTaskPhase(v string)`

SetTaskPhase sets TaskPhase field to given value.

### HasTaskPhase

`func (o *WorkflowTaskSetTasks) HasTaskPhase() bool`

HasTaskPhase returns a boolean if a field has been set.

### GetTaskOrder

`func (o *WorkflowTaskSetTasks) GetTaskOrder() int64`

GetTaskOrder returns the TaskOrder field if non-nil, zero value otherwise.

### GetTaskOrderOk

`func (o *WorkflowTaskSetTasks) GetTaskOrderOk() (*int64, bool)`

GetTaskOrderOk returns a tuple with the TaskOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskOrder

`func (o *WorkflowTaskSetTasks) SetTaskOrder(v int64)`

SetTaskOrder sets TaskOrder field to given value.

### HasTaskOrder

`func (o *WorkflowTaskSetTasks) HasTaskOrder() bool`

HasTaskOrder returns a boolean if a field has been set.

### GetTask

`func (o *WorkflowTaskSetTasks) GetTask() WorkflowTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *WorkflowTaskSetTasks) GetTaskOk() (*WorkflowTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *WorkflowTaskSetTasks) SetTask(v WorkflowTask)`

SetTask sets Task field to given value.

### HasTask

`func (o *WorkflowTaskSetTasks) HasTask() bool`

HasTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


