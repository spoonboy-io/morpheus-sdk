# InlineObject225

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskSet** | Pointer to [**ApiServersIdWorkflowTaskSet**](_api_servers__id__workflow_taskSet.md) |  | [optional] 
**TaskPhase** | Pointer to **string** | Task Phase to run for Provisioning workflows. The default is &#x60;provision&#x60;. | [optional] [default to "provision"]

## Methods

### NewInlineObject225

`func NewInlineObject225() *InlineObject225`

NewInlineObject225 instantiates a new InlineObject225 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject225WithDefaults

`func NewInlineObject225WithDefaults() *InlineObject225`

NewInlineObject225WithDefaults instantiates a new InlineObject225 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskSet

`func (o *InlineObject225) GetTaskSet() ApiServersIdWorkflowTaskSet`

GetTaskSet returns the TaskSet field if non-nil, zero value otherwise.

### GetTaskSetOk

`func (o *InlineObject225) GetTaskSetOk() (*ApiServersIdWorkflowTaskSet, bool)`

GetTaskSetOk returns a tuple with the TaskSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSet

`func (o *InlineObject225) SetTaskSet(v ApiServersIdWorkflowTaskSet)`

SetTaskSet sets TaskSet field to given value.

### HasTaskSet

`func (o *InlineObject225) HasTaskSet() bool`

HasTaskSet returns a boolean if a field has been set.

### GetTaskPhase

`func (o *InlineObject225) GetTaskPhase() string`

GetTaskPhase returns the TaskPhase field if non-nil, zero value otherwise.

### GetTaskPhaseOk

`func (o *InlineObject225) GetTaskPhaseOk() (*string, bool)`

GetTaskPhaseOk returns a tuple with the TaskPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPhase

`func (o *InlineObject225) SetTaskPhase(v string)`

SetTaskPhase sets TaskPhase field to given value.

### HasTaskPhase

`func (o *InlineObject225) HasTaskPhase() bool`

HasTaskPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


