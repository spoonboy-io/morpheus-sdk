# ApiExecuteSchedulesSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the execute schedule | 
**Description** | Pointer to **string** | A description for the execute schedule | [optional] 
**ScheduleType** | Pointer to **string** | Type of schedule | [optional] 
**ScheduleTimezone** | Pointer to **string** | Time Zone eg. America/New_York, Europe/Amsterdam, etc. | [optional] [default to "UTC"]
**Cron** | Pointer to **string** | Cron Expression. The default is daily at midnight | [optional] [default to "0 0 * * *"]
**Enabled** | Pointer to **bool** | Is enabled | [optional] [default to true]

## Methods

### NewApiExecuteSchedulesSchedule

`func NewApiExecuteSchedulesSchedule(name string, ) *ApiExecuteSchedulesSchedule`

NewApiExecuteSchedulesSchedule instantiates a new ApiExecuteSchedulesSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiExecuteSchedulesScheduleWithDefaults

`func NewApiExecuteSchedulesScheduleWithDefaults() *ApiExecuteSchedulesSchedule`

NewApiExecuteSchedulesScheduleWithDefaults instantiates a new ApiExecuteSchedulesSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiExecuteSchedulesSchedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiExecuteSchedulesSchedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiExecuteSchedulesSchedule) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ApiExecuteSchedulesSchedule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiExecuteSchedulesSchedule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiExecuteSchedulesSchedule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiExecuteSchedulesSchedule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScheduleType

`func (o *ApiExecuteSchedulesSchedule) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *ApiExecuteSchedulesSchedule) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *ApiExecuteSchedulesSchedule) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *ApiExecuteSchedulesSchedule) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetScheduleTimezone

`func (o *ApiExecuteSchedulesSchedule) GetScheduleTimezone() string`

GetScheduleTimezone returns the ScheduleTimezone field if non-nil, zero value otherwise.

### GetScheduleTimezoneOk

`func (o *ApiExecuteSchedulesSchedule) GetScheduleTimezoneOk() (*string, bool)`

GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimezone

`func (o *ApiExecuteSchedulesSchedule) SetScheduleTimezone(v string)`

SetScheduleTimezone sets ScheduleTimezone field to given value.

### HasScheduleTimezone

`func (o *ApiExecuteSchedulesSchedule) HasScheduleTimezone() bool`

HasScheduleTimezone returns a boolean if a field has been set.

### GetCron

`func (o *ApiExecuteSchedulesSchedule) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *ApiExecuteSchedulesSchedule) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *ApiExecuteSchedulesSchedule) SetCron(v string)`

SetCron sets Cron field to given value.

### HasCron

`func (o *ApiExecuteSchedulesSchedule) HasCron() bool`

HasCron returns a boolean if a field has been set.

### GetEnabled

`func (o *ApiExecuteSchedulesSchedule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiExecuteSchedulesSchedule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiExecuteSchedulesSchedule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiExecuteSchedulesSchedule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


