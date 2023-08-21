# BackupRestore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Backup Result ID | [optional] 
**BackupResultId** | Pointer to **int64** |  | [optional] 
**BackupId** | Pointer to **int64** |  | [optional] 
**Backup** | Pointer to [**NullableBackupRestoreBackup**](backupRestore_backup.md) |  | [optional] 
**ContainerId** | Pointer to **NullableInt64** |  | [optional] 
**Container** | Pointer to [**NullableBackupRestoreContainer**](backupRestore_container.md) |  | [optional] 
**Instance** | Pointer to [**NullableBackupInstance**](backup_instance.md) |  | [optional] 
**RestoreToNew** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**StartDate** | Pointer to **NullableTime** |  | [optional] 
**EndDate** | Pointer to **NullableTime** |  | [optional] 
**DurationMillis** | Pointer to **NullableInt64** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalStatusRef** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **time.Time** | Date Created | [optional] 
**LastUpdated** | Pointer to **time.Time** | Last Updated | [optional] 

## Methods

### NewBackupRestore

`func NewBackupRestore() *BackupRestore`

NewBackupRestore instantiates a new BackupRestore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreWithDefaults

`func NewBackupRestoreWithDefaults() *BackupRestore`

NewBackupRestoreWithDefaults instantiates a new BackupRestore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupRestore) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupRestore) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupRestore) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BackupRestore) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBackupResultId

`func (o *BackupRestore) GetBackupResultId() int64`

GetBackupResultId returns the BackupResultId field if non-nil, zero value otherwise.

### GetBackupResultIdOk

`func (o *BackupRestore) GetBackupResultIdOk() (*int64, bool)`

GetBackupResultIdOk returns a tuple with the BackupResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupResultId

`func (o *BackupRestore) SetBackupResultId(v int64)`

SetBackupResultId sets BackupResultId field to given value.

### HasBackupResultId

`func (o *BackupRestore) HasBackupResultId() bool`

HasBackupResultId returns a boolean if a field has been set.

### GetBackupId

`func (o *BackupRestore) GetBackupId() int64`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *BackupRestore) GetBackupIdOk() (*int64, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *BackupRestore) SetBackupId(v int64)`

SetBackupId sets BackupId field to given value.

### HasBackupId

`func (o *BackupRestore) HasBackupId() bool`

HasBackupId returns a boolean if a field has been set.

### GetBackup

`func (o *BackupRestore) GetBackup() BackupRestoreBackup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *BackupRestore) GetBackupOk() (*BackupRestoreBackup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *BackupRestore) SetBackup(v BackupRestoreBackup)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *BackupRestore) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### SetBackupNil

`func (o *BackupRestore) SetBackupNil(b bool)`

 SetBackupNil sets the value for Backup to be an explicit nil

### UnsetBackup
`func (o *BackupRestore) UnsetBackup()`

UnsetBackup ensures that no value is present for Backup, not even an explicit nil
### GetContainerId

`func (o *BackupRestore) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *BackupRestore) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *BackupRestore) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *BackupRestore) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### SetContainerIdNil

`func (o *BackupRestore) SetContainerIdNil(b bool)`

 SetContainerIdNil sets the value for ContainerId to be an explicit nil

### UnsetContainerId
`func (o *BackupRestore) UnsetContainerId()`

UnsetContainerId ensures that no value is present for ContainerId, not even an explicit nil
### GetContainer

`func (o *BackupRestore) GetContainer() BackupRestoreContainer`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *BackupRestore) GetContainerOk() (*BackupRestoreContainer, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *BackupRestore) SetContainer(v BackupRestoreContainer)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *BackupRestore) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *BackupRestore) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *BackupRestore) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil
### GetInstance

`func (o *BackupRestore) GetInstance() BackupInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *BackupRestore) GetInstanceOk() (*BackupInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *BackupRestore) SetInstance(v BackupInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *BackupRestore) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### SetInstanceNil

`func (o *BackupRestore) SetInstanceNil(b bool)`

 SetInstanceNil sets the value for Instance to be an explicit nil

### UnsetInstance
`func (o *BackupRestore) UnsetInstance()`

UnsetInstance ensures that no value is present for Instance, not even an explicit nil
### GetRestoreToNew

`func (o *BackupRestore) GetRestoreToNew() bool`

GetRestoreToNew returns the RestoreToNew field if non-nil, zero value otherwise.

### GetRestoreToNewOk

`func (o *BackupRestore) GetRestoreToNewOk() (*bool, bool)`

GetRestoreToNewOk returns a tuple with the RestoreToNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreToNew

`func (o *BackupRestore) SetRestoreToNew(v bool)`

SetRestoreToNew sets RestoreToNew field to given value.

### HasRestoreToNew

`func (o *BackupRestore) HasRestoreToNew() bool`

HasRestoreToNew returns a boolean if a field has been set.

### GetStatus

`func (o *BackupRestore) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BackupRestore) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BackupRestore) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BackupRestore) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BackupRestore) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BackupRestore) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetErrorMessage

`func (o *BackupRestore) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BackupRestore) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BackupRestore) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *BackupRestore) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *BackupRestore) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *BackupRestore) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetStartDate

`func (o *BackupRestore) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BackupRestore) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BackupRestore) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BackupRestore) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *BackupRestore) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *BackupRestore) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *BackupRestore) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BackupRestore) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BackupRestore) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BackupRestore) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *BackupRestore) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *BackupRestore) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetDurationMillis

`func (o *BackupRestore) GetDurationMillis() int64`

GetDurationMillis returns the DurationMillis field if non-nil, zero value otherwise.

### GetDurationMillisOk

`func (o *BackupRestore) GetDurationMillisOk() (*int64, bool)`

GetDurationMillisOk returns a tuple with the DurationMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMillis

`func (o *BackupRestore) SetDurationMillis(v int64)`

SetDurationMillis sets DurationMillis field to given value.

### HasDurationMillis

`func (o *BackupRestore) HasDurationMillis() bool`

HasDurationMillis returns a boolean if a field has been set.

### SetDurationMillisNil

`func (o *BackupRestore) SetDurationMillisNil(b bool)`

 SetDurationMillisNil sets the value for DurationMillis to be an explicit nil

### UnsetDurationMillis
`func (o *BackupRestore) UnsetDurationMillis()`

UnsetDurationMillis ensures that no value is present for DurationMillis, not even an explicit nil
### GetExternalId

`func (o *BackupRestore) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *BackupRestore) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *BackupRestore) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *BackupRestore) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *BackupRestore) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *BackupRestore) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetExternalStatusRef

`func (o *BackupRestore) GetExternalStatusRef() string`

GetExternalStatusRef returns the ExternalStatusRef field if non-nil, zero value otherwise.

### GetExternalStatusRefOk

`func (o *BackupRestore) GetExternalStatusRefOk() (*string, bool)`

GetExternalStatusRefOk returns a tuple with the ExternalStatusRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStatusRef

`func (o *BackupRestore) SetExternalStatusRef(v string)`

SetExternalStatusRef sets ExternalStatusRef field to given value.

### HasExternalStatusRef

`func (o *BackupRestore) HasExternalStatusRef() bool`

HasExternalStatusRef returns a boolean if a field has been set.

### SetExternalStatusRefNil

`func (o *BackupRestore) SetExternalStatusRefNil(b bool)`

 SetExternalStatusRefNil sets the value for ExternalStatusRef to be an explicit nil

### UnsetExternalStatusRef
`func (o *BackupRestore) UnsetExternalStatusRef()`

UnsetExternalStatusRef ensures that no value is present for ExternalStatusRef, not even an explicit nil
### GetDateCreated

`func (o *BackupRestore) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *BackupRestore) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *BackupRestore) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *BackupRestore) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *BackupRestore) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *BackupRestore) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *BackupRestore) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *BackupRestore) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


