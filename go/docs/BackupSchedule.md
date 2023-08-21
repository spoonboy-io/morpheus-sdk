# BackupSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Schedule ID | [optional] 
**Name** | Pointer to **string** | Schedule Name | [optional] 
**Cron** | Pointer to **string** | Schedule Cron Expression | [optional] 

## Methods

### NewBackupSchedule

`func NewBackupSchedule() *BackupSchedule`

NewBackupSchedule instantiates a new BackupSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupScheduleWithDefaults

`func NewBackupScheduleWithDefaults() *BackupSchedule`

NewBackupScheduleWithDefaults instantiates a new BackupSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupSchedule) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupSchedule) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupSchedule) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BackupSchedule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BackupSchedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupSchedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupSchedule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BackupSchedule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCron

`func (o *BackupSchedule) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *BackupSchedule) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *BackupSchedule) SetCron(v string)`

SetCron sets Cron field to given value.

### HasCron

`func (o *BackupSchedule) HasCron() bool`

HasCron returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


