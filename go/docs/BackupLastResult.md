# BackupLastResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Last Result ID | [optional] 
**Name** | Pointer to **string** | Last Result Status | [optional] 
**DateCreated** | Pointer to **time.Time** | Last Result Date Created | [optional] 

## Methods

### NewBackupLastResult

`func NewBackupLastResult() *BackupLastResult`

NewBackupLastResult instantiates a new BackupLastResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupLastResultWithDefaults

`func NewBackupLastResultWithDefaults() *BackupLastResult`

NewBackupLastResultWithDefaults instantiates a new BackupLastResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupLastResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupLastResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupLastResult) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BackupLastResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BackupLastResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupLastResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupLastResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BackupLastResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDateCreated

`func (o *BackupLastResult) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *BackupLastResult) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *BackupLastResult) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *BackupLastResult) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


