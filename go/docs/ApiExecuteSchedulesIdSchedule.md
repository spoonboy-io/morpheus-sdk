# ApiExecuteSchedulesIdSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the execute schedule | [optional] 
**Description** | Pointer to **string** | A description for the execute schedule | [optional] 
**ScheduleType** | Pointer to **string** | Type of schedule | [optional] 
**ScheduleTimezone** | Pointer to **string** | Time Zone eg. America/New_York, Europe/Amsterdam, etc. | [optional] [default to "UTC"]
**Cron** | Pointer to **string** | Cron Expression. The default is daily at midnight | [optional] [default to "0 0 * * *"]
**Enabled** | Pointer to **bool** | Is enabled | [optional] [default to true]

## Methods

### NewApiExecuteSchedulesIdSchedule

`func NewApiExecuteSchedulesIdSchedule() *ApiExecuteSchedulesIdSchedule`

NewApiExecuteSchedulesIdSchedule instantiates a new ApiExecuteSchedulesIdSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiExecuteSchedulesIdScheduleWithDefaults

`func NewApiExecuteSchedulesIdScheduleWithDefaults() *ApiExecuteSchedulesIdSchedule`

NewApiExecuteSchedulesIdScheduleWithDefaults instantiates a new ApiExecuteSchedulesIdSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiExecuteSchedulesIdSchedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiExecuteSchedulesIdSchedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiExecuteSchedulesIdSchedule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiExecuteSchedulesIdSchedule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApiExecuteSchedulesIdSchedule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiExecuteSchedulesIdSchedule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiExecuteSchedulesIdSchedule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiExecuteSchedulesIdSchedule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScheduleType

`func (o *ApiExecuteSchedulesIdSchedule) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *ApiExecuteSchedulesIdSchedule) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *ApiExecuteSchedulesIdSchedule) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *ApiExecuteSchedulesIdSchedule) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetScheduleTimezone

`func (o *ApiExecuteSchedulesIdSchedule) GetScheduleTimezone() string`

GetScheduleTimezone returns the ScheduleTimezone field if non-nil, zero value otherwise.

### GetScheduleTimezoneOk

`func (o *ApiExecuteSchedulesIdSchedule) GetScheduleTimezoneOk() (*string, bool)`

GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimezone

`func (o *ApiExecuteSchedulesIdSchedule) SetScheduleTimezone(v string)`

SetScheduleTimezone sets ScheduleTimezone field to given value.

### HasScheduleTimezone

`func (o *ApiExecuteSchedulesIdSchedule) HasScheduleTimezone() bool`

HasScheduleTimezone returns a boolean if a field has been set.

### GetCron

`func (o *ApiExecuteSchedulesIdSchedule) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *ApiExecuteSchedulesIdSchedule) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *ApiExecuteSchedulesIdSchedule) SetCron(v string)`

SetCron sets Cron field to given value.

### HasCron

`func (o *ApiExecuteSchedulesIdSchedule) HasCron() bool`

HasCron returns a boolean if a field has been set.

### GetEnabled

`func (o *ApiExecuteSchedulesIdSchedule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiExecuteSchedulesIdSchedule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiExecuteSchedulesIdSchedule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiExecuteSchedulesIdSchedule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


