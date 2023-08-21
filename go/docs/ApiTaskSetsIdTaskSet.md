# ApiTaskSetsIdTaskSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A unique name for the workflow | [optional] 
**Description** | Pointer to **string** | A description of the workflow | [optional] 
**Labels** | Pointer to **[]string** | An array of Category labels for filtering | [optional] 
**Type** | Pointer to **string** | Workflow type | [optional] [default to "provision"]
**OptionTypes** | Pointer to **[]int64** | List of option type IDs for use with operational workflow configuration. | [optional] 
**Tasks** | Pointer to [**ApiTaskSetsTaskSetTasks**](_api_task_sets_taskSet_tasks.md) |  | [optional] 

## Methods

### NewApiTaskSetsIdTaskSet

`func NewApiTaskSetsIdTaskSet() *ApiTaskSetsIdTaskSet`

NewApiTaskSetsIdTaskSet instantiates a new ApiTaskSetsIdTaskSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTaskSetsIdTaskSetWithDefaults

`func NewApiTaskSetsIdTaskSetWithDefaults() *ApiTaskSetsIdTaskSet`

NewApiTaskSetsIdTaskSetWithDefaults instantiates a new ApiTaskSetsIdTaskSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiTaskSetsIdTaskSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiTaskSetsIdTaskSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiTaskSetsIdTaskSet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiTaskSetsIdTaskSet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApiTaskSetsIdTaskSet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiTaskSetsIdTaskSet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiTaskSetsIdTaskSet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiTaskSetsIdTaskSet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *ApiTaskSetsIdTaskSet) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ApiTaskSetsIdTaskSet) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ApiTaskSetsIdTaskSet) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ApiTaskSetsIdTaskSet) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *ApiTaskSetsIdTaskSet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiTaskSetsIdTaskSet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiTaskSetsIdTaskSet) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiTaskSetsIdTaskSet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ApiTaskSetsIdTaskSet) GetOptionTypes() []int64`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ApiTaskSetsIdTaskSet) GetOptionTypesOk() (*[]int64, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ApiTaskSetsIdTaskSet) SetOptionTypes(v []int64)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ApiTaskSetsIdTaskSet) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetTasks

`func (o *ApiTaskSetsIdTaskSet) GetTasks() ApiTaskSetsTaskSetTasks`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ApiTaskSetsIdTaskSet) GetTasksOk() (*ApiTaskSetsTaskSetTasks, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ApiTaskSetsIdTaskSet) SetTasks(v ApiTaskSetsTaskSetTasks)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ApiTaskSetsIdTaskSet) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


