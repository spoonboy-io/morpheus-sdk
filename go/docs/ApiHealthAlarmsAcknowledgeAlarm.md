# ApiHealthAlarmsAcknowledgeAlarm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acknowledged** | **bool** | Pass &#x60;true&#x60; to ackowledge an alarm, or pass &#x60;false&#x60; to unacknowledge it. | 
**Ids** | Pointer to **[]int64** | Array of Alarm ID(s)to be updated. | [optional] [default to []]
**All** | Pointer to **bool** | Pass &#x60;true&#x60; to update all alarms instead of passing ids. This will update any active alarm that is not already acknowledged.  | [optional] [default to false]

## Methods

### NewApiHealthAlarmsAcknowledgeAlarm

`func NewApiHealthAlarmsAcknowledgeAlarm(acknowledged bool, ) *ApiHealthAlarmsAcknowledgeAlarm`

NewApiHealthAlarmsAcknowledgeAlarm instantiates a new ApiHealthAlarmsAcknowledgeAlarm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiHealthAlarmsAcknowledgeAlarmWithDefaults

`func NewApiHealthAlarmsAcknowledgeAlarmWithDefaults() *ApiHealthAlarmsAcknowledgeAlarm`

NewApiHealthAlarmsAcknowledgeAlarmWithDefaults instantiates a new ApiHealthAlarmsAcknowledgeAlarm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledged

`func (o *ApiHealthAlarmsAcknowledgeAlarm) GetAcknowledged() bool`

GetAcknowledged returns the Acknowledged field if non-nil, zero value otherwise.

### GetAcknowledgedOk

`func (o *ApiHealthAlarmsAcknowledgeAlarm) GetAcknowledgedOk() (*bool, bool)`

GetAcknowledgedOk returns a tuple with the Acknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledged

`func (o *ApiHealthAlarmsAcknowledgeAlarm) SetAcknowledged(v bool)`

SetAcknowledged sets Acknowledged field to given value.


### GetIds

`func (o *ApiHealthAlarmsAcknowledgeAlarm) GetIds() []int64`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ApiHealthAlarmsAcknowledgeAlarm) GetIdsOk() (*[]int64, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ApiHealthAlarmsAcknowledgeAlarm) SetIds(v []int64)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ApiHealthAlarmsAcknowledgeAlarm) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetAll

`func (o *ApiHealthAlarmsAcknowledgeAlarm) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *ApiHealthAlarmsAcknowledgeAlarm) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *ApiHealthAlarmsAcknowledgeAlarm) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *ApiHealthAlarmsAcknowledgeAlarm) HasAll() bool`

HasAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


