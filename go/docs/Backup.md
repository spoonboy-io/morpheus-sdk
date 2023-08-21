# Backup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Backup ID | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**LocationType** | Pointer to **string** | Source Type (instance, server, storage) | [optional] 
**Instance** | Pointer to [**NullableBackupInstance**](backup_instance.md) |  | [optional] 
**ContainerId** | Pointer to **NullableInt64** |  | [optional] 
**Job** | Pointer to [**BackupJob**](backup_job.md) |  | [optional] 
**Schedule** | Pointer to [**NullableBackupSchedule**](backup_schedule.md) |  | [optional] 
**RetentionCount** | Pointer to **NullableInt64** |  | [optional] 
**BackupType** | Pointer to [**BackupBackupType**](backup_backupType.md) |  | [optional] 
**StorageProvider** | Pointer to [**NullableBackupStorageProvider**](backup_storageProvider.md) |  | [optional] 
**BackupProvider** | Pointer to [**NullableBackupBackupProvider**](backup_backupProvider.md) |  | [optional] 
**BackupRespository** | Pointer to [**NullableBackupBackupRespository**](backup_backupRespository.md) |  | [optional] 
**CronExpression** | Pointer to **NullableString** | Cron Expression | [optional] 
**NextFire** | Pointer to **NullableTime** | Next Fire | [optional] 
**LastStatus** | Pointer to **NullableString** | Last Status | [optional] 
**LastResult** | Pointer to [**NullableBackupLastResult**](backup_lastResult.md) |  | [optional] 
**Stats** | Pointer to [**BackupStats**](backup_stats.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Enabled | [optional] 
**DateCreated** | Pointer to **time.Time** | Date Created | [optional] 
**LastUpdated** | Pointer to **time.Time** | Last Updated | [optional] 

## Methods

### NewBackup

`func NewBackup() *Backup`

NewBackup instantiates a new Backup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupWithDefaults

`func NewBackupWithDefaults() *Backup`

NewBackupWithDefaults instantiates a new Backup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Backup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Backup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Backup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Backup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Backup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Backup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Backup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Backup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocationType

`func (o *Backup) GetLocationType() string`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *Backup) GetLocationTypeOk() (*string, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *Backup) SetLocationType(v string)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *Backup) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### GetInstance

`func (o *Backup) GetInstance() BackupInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *Backup) GetInstanceOk() (*BackupInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *Backup) SetInstance(v BackupInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *Backup) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### SetInstanceNil

`func (o *Backup) SetInstanceNil(b bool)`

 SetInstanceNil sets the value for Instance to be an explicit nil

### UnsetInstance
`func (o *Backup) UnsetInstance()`

UnsetInstance ensures that no value is present for Instance, not even an explicit nil
### GetContainerId

`func (o *Backup) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *Backup) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *Backup) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *Backup) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### SetContainerIdNil

`func (o *Backup) SetContainerIdNil(b bool)`

 SetContainerIdNil sets the value for ContainerId to be an explicit nil

### UnsetContainerId
`func (o *Backup) UnsetContainerId()`

UnsetContainerId ensures that no value is present for ContainerId, not even an explicit nil
### GetJob

`func (o *Backup) GetJob() BackupJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *Backup) GetJobOk() (*BackupJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *Backup) SetJob(v BackupJob)`

SetJob sets Job field to given value.

### HasJob

`func (o *Backup) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetSchedule

`func (o *Backup) GetSchedule() BackupSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *Backup) GetScheduleOk() (*BackupSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *Backup) SetSchedule(v BackupSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *Backup) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *Backup) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *Backup) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetRetentionCount

`func (o *Backup) GetRetentionCount() int64`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *Backup) GetRetentionCountOk() (*int64, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *Backup) SetRetentionCount(v int64)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *Backup) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.

### SetRetentionCountNil

`func (o *Backup) SetRetentionCountNil(b bool)`

 SetRetentionCountNil sets the value for RetentionCount to be an explicit nil

### UnsetRetentionCount
`func (o *Backup) UnsetRetentionCount()`

UnsetRetentionCount ensures that no value is present for RetentionCount, not even an explicit nil
### GetBackupType

`func (o *Backup) GetBackupType() BackupBackupType`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *Backup) GetBackupTypeOk() (*BackupBackupType, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *Backup) SetBackupType(v BackupBackupType)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *Backup) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### GetStorageProvider

`func (o *Backup) GetStorageProvider() BackupStorageProvider`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *Backup) GetStorageProviderOk() (*BackupStorageProvider, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *Backup) SetStorageProvider(v BackupStorageProvider)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *Backup) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### SetStorageProviderNil

`func (o *Backup) SetStorageProviderNil(b bool)`

 SetStorageProviderNil sets the value for StorageProvider to be an explicit nil

### UnsetStorageProvider
`func (o *Backup) UnsetStorageProvider()`

UnsetStorageProvider ensures that no value is present for StorageProvider, not even an explicit nil
### GetBackupProvider

`func (o *Backup) GetBackupProvider() BackupBackupProvider`

GetBackupProvider returns the BackupProvider field if non-nil, zero value otherwise.

### GetBackupProviderOk

`func (o *Backup) GetBackupProviderOk() (*BackupBackupProvider, bool)`

GetBackupProviderOk returns a tuple with the BackupProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupProvider

`func (o *Backup) SetBackupProvider(v BackupBackupProvider)`

SetBackupProvider sets BackupProvider field to given value.

### HasBackupProvider

`func (o *Backup) HasBackupProvider() bool`

HasBackupProvider returns a boolean if a field has been set.

### SetBackupProviderNil

`func (o *Backup) SetBackupProviderNil(b bool)`

 SetBackupProviderNil sets the value for BackupProvider to be an explicit nil

### UnsetBackupProvider
`func (o *Backup) UnsetBackupProvider()`

UnsetBackupProvider ensures that no value is present for BackupProvider, not even an explicit nil
### GetBackupRespository

`func (o *Backup) GetBackupRespository() BackupBackupRespository`

GetBackupRespository returns the BackupRespository field if non-nil, zero value otherwise.

### GetBackupRespositoryOk

`func (o *Backup) GetBackupRespositoryOk() (*BackupBackupRespository, bool)`

GetBackupRespositoryOk returns a tuple with the BackupRespository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRespository

`func (o *Backup) SetBackupRespository(v BackupBackupRespository)`

SetBackupRespository sets BackupRespository field to given value.

### HasBackupRespository

`func (o *Backup) HasBackupRespository() bool`

HasBackupRespository returns a boolean if a field has been set.

### SetBackupRespositoryNil

`func (o *Backup) SetBackupRespositoryNil(b bool)`

 SetBackupRespositoryNil sets the value for BackupRespository to be an explicit nil

### UnsetBackupRespository
`func (o *Backup) UnsetBackupRespository()`

UnsetBackupRespository ensures that no value is present for BackupRespository, not even an explicit nil
### GetCronExpression

`func (o *Backup) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *Backup) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *Backup) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *Backup) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### SetCronExpressionNil

`func (o *Backup) SetCronExpressionNil(b bool)`

 SetCronExpressionNil sets the value for CronExpression to be an explicit nil

### UnsetCronExpression
`func (o *Backup) UnsetCronExpression()`

UnsetCronExpression ensures that no value is present for CronExpression, not even an explicit nil
### GetNextFire

`func (o *Backup) GetNextFire() time.Time`

GetNextFire returns the NextFire field if non-nil, zero value otherwise.

### GetNextFireOk

`func (o *Backup) GetNextFireOk() (*time.Time, bool)`

GetNextFireOk returns a tuple with the NextFire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextFire

`func (o *Backup) SetNextFire(v time.Time)`

SetNextFire sets NextFire field to given value.

### HasNextFire

`func (o *Backup) HasNextFire() bool`

HasNextFire returns a boolean if a field has been set.

### SetNextFireNil

`func (o *Backup) SetNextFireNil(b bool)`

 SetNextFireNil sets the value for NextFire to be an explicit nil

### UnsetNextFire
`func (o *Backup) UnsetNextFire()`

UnsetNextFire ensures that no value is present for NextFire, not even an explicit nil
### GetLastStatus

`func (o *Backup) GetLastStatus() string`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *Backup) GetLastStatusOk() (*string, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *Backup) SetLastStatus(v string)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *Backup) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.

### SetLastStatusNil

`func (o *Backup) SetLastStatusNil(b bool)`

 SetLastStatusNil sets the value for LastStatus to be an explicit nil

### UnsetLastStatus
`func (o *Backup) UnsetLastStatus()`

UnsetLastStatus ensures that no value is present for LastStatus, not even an explicit nil
### GetLastResult

`func (o *Backup) GetLastResult() BackupLastResult`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *Backup) GetLastResultOk() (*BackupLastResult, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *Backup) SetLastResult(v BackupLastResult)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *Backup) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### SetLastResultNil

`func (o *Backup) SetLastResultNil(b bool)`

 SetLastResultNil sets the value for LastResult to be an explicit nil

### UnsetLastResult
`func (o *Backup) UnsetLastResult()`

UnsetLastResult ensures that no value is present for LastResult, not even an explicit nil
### GetStats

`func (o *Backup) GetStats() BackupStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Backup) GetStatsOk() (*BackupStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Backup) SetStats(v BackupStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *Backup) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetEnabled

`func (o *Backup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Backup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Backup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Backup) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDateCreated

`func (o *Backup) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Backup) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Backup) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Backup) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Backup) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Backup) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Backup) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Backup) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


