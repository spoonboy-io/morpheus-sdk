# BackupRestoreContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Container ID | [optional] 
**Name** | Pointer to **string** | Container Name | [optional] 

## Methods

### NewBackupRestoreContainer

`func NewBackupRestoreContainer() *BackupRestoreContainer`

NewBackupRestoreContainer instantiates a new BackupRestoreContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreContainerWithDefaults

`func NewBackupRestoreContainerWithDefaults() *BackupRestoreContainer`

NewBackupRestoreContainerWithDefaults instantiates a new BackupRestoreContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupRestoreContainer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupRestoreContainer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupRestoreContainer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BackupRestoreContainer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BackupRestoreContainer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupRestoreContainer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupRestoreContainer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BackupRestoreContainer) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


