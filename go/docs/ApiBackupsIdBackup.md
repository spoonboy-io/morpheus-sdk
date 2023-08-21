# ApiBackupsIdBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the backup | [optional] 
**JobId** | Pointer to **int64** | The Backup Job ID to assign the backup to. This determines when the backup is run. | [optional] 
**Enabled** | Pointer to **bool** | Can be used to enable or disable the backup | [optional] 

## Methods

### NewApiBackupsIdBackup

`func NewApiBackupsIdBackup() *ApiBackupsIdBackup`

NewApiBackupsIdBackup instantiates a new ApiBackupsIdBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiBackupsIdBackupWithDefaults

`func NewApiBackupsIdBackupWithDefaults() *ApiBackupsIdBackup`

NewApiBackupsIdBackupWithDefaults instantiates a new ApiBackupsIdBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiBackupsIdBackup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiBackupsIdBackup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiBackupsIdBackup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiBackupsIdBackup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetJobId

`func (o *ApiBackupsIdBackup) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ApiBackupsIdBackup) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ApiBackupsIdBackup) SetJobId(v int64)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *ApiBackupsIdBackup) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetEnabled

`func (o *ApiBackupsIdBackup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiBackupsIdBackup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiBackupsIdBackup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiBackupsIdBackup) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


