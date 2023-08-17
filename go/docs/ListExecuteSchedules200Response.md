# ListExecuteSchedules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Schedules** | Pointer to [**[]ExecuteSchedule**](ExecuteSchedule.md) |  | [optional] 

## Methods

### NewListExecuteSchedules200Response

`func NewListExecuteSchedules200Response() *ListExecuteSchedules200Response`

NewListExecuteSchedules200Response instantiates a new ListExecuteSchedules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListExecuteSchedules200ResponseWithDefaults

`func NewListExecuteSchedules200ResponseWithDefaults() *ListExecuteSchedules200Response`

NewListExecuteSchedules200ResponseWithDefaults instantiates a new ListExecuteSchedules200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListExecuteSchedules200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListExecuteSchedules200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListExecuteSchedules200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListExecuteSchedules200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSchedules

`func (o *ListExecuteSchedules200Response) GetSchedules() []ExecuteSchedule`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *ListExecuteSchedules200Response) GetSchedulesOk() (*[]ExecuteSchedule, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *ListExecuteSchedules200Response) SetSchedules(v []ExecuteSchedule)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *ListExecuteSchedules200Response) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


