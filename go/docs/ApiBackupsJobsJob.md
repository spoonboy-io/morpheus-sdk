# ApiBackupsJobsJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the backup job | 
**Code** | Pointer to **string** | A code for the backup job | [optional] 
**ScheduleId** | Pointer to **NullableInt64** | Execute Schedule ID to use for the backup job | [optional] 
**RetentionCount** | Pointer to **int64** | Retention Count. By default the backup settings value will be used. | [optional] 

## Methods

### NewApiBackupsJobsJob

`func NewApiBackupsJobsJob(name string, ) *ApiBackupsJobsJob`

NewApiBackupsJobsJob instantiates a new ApiBackupsJobsJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiBackupsJobsJobWithDefaults

`func NewApiBackupsJobsJobWithDefaults() *ApiBackupsJobsJob`

NewApiBackupsJobsJobWithDefaults instantiates a new ApiBackupsJobsJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiBackupsJobsJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiBackupsJobsJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiBackupsJobsJob) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *ApiBackupsJobsJob) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiBackupsJobsJob) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiBackupsJobsJob) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiBackupsJobsJob) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetScheduleId

`func (o *ApiBackupsJobsJob) GetScheduleId() int64`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *ApiBackupsJobsJob) GetScheduleIdOk() (*int64, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *ApiBackupsJobsJob) SetScheduleId(v int64)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *ApiBackupsJobsJob) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *ApiBackupsJobsJob) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *ApiBackupsJobsJob) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetRetentionCount

`func (o *ApiBackupsJobsJob) GetRetentionCount() int64`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *ApiBackupsJobsJob) GetRetentionCountOk() (*int64, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *ApiBackupsJobsJob) SetRetentionCount(v int64)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *ApiBackupsJobsJob) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


