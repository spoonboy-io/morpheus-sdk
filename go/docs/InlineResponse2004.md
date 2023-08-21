# InlineResponse2004

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveBucket** | Pointer to [**ArchiveBucket**](archiveBucket.md) |  | [optional] 
**IsOwner** | Pointer to **bool** |  | [optional] 
**ParentDirectory** | Pointer to **NullableString** |  | [optional] 
**ArchiveFiles** | Pointer to [**[]ArchiveBucketFile**](ArchiveBucketFile.md) |  | [optional] 
**ArchiveFileCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewInlineResponse2004

`func NewInlineResponse2004() *InlineResponse2004`

NewInlineResponse2004 instantiates a new InlineResponse2004 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2004WithDefaults

`func NewInlineResponse2004WithDefaults() *InlineResponse2004`

NewInlineResponse2004WithDefaults instantiates a new InlineResponse2004 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveBucket

`func (o *InlineResponse2004) GetArchiveBucket() ArchiveBucket`

GetArchiveBucket returns the ArchiveBucket field if non-nil, zero value otherwise.

### GetArchiveBucketOk

`func (o *InlineResponse2004) GetArchiveBucketOk() (*ArchiveBucket, bool)`

GetArchiveBucketOk returns a tuple with the ArchiveBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveBucket

`func (o *InlineResponse2004) SetArchiveBucket(v ArchiveBucket)`

SetArchiveBucket sets ArchiveBucket field to given value.

### HasArchiveBucket

`func (o *InlineResponse2004) HasArchiveBucket() bool`

HasArchiveBucket returns a boolean if a field has been set.

### GetIsOwner

`func (o *InlineResponse2004) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *InlineResponse2004) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *InlineResponse2004) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.

### HasIsOwner

`func (o *InlineResponse2004) HasIsOwner() bool`

HasIsOwner returns a boolean if a field has been set.

### GetParentDirectory

`func (o *InlineResponse2004) GetParentDirectory() string`

GetParentDirectory returns the ParentDirectory field if non-nil, zero value otherwise.

### GetParentDirectoryOk

`func (o *InlineResponse2004) GetParentDirectoryOk() (*string, bool)`

GetParentDirectoryOk returns a tuple with the ParentDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDirectory

`func (o *InlineResponse2004) SetParentDirectory(v string)`

SetParentDirectory sets ParentDirectory field to given value.

### HasParentDirectory

`func (o *InlineResponse2004) HasParentDirectory() bool`

HasParentDirectory returns a boolean if a field has been set.

### SetParentDirectoryNil

`func (o *InlineResponse2004) SetParentDirectoryNil(b bool)`

 SetParentDirectoryNil sets the value for ParentDirectory to be an explicit nil

### UnsetParentDirectory
`func (o *InlineResponse2004) UnsetParentDirectory()`

UnsetParentDirectory ensures that no value is present for ParentDirectory, not even an explicit nil
### GetArchiveFiles

`func (o *InlineResponse2004) GetArchiveFiles() []ArchiveBucketFile`

GetArchiveFiles returns the ArchiveFiles field if non-nil, zero value otherwise.

### GetArchiveFilesOk

`func (o *InlineResponse2004) GetArchiveFilesOk() (*[]ArchiveBucketFile, bool)`

GetArchiveFilesOk returns a tuple with the ArchiveFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveFiles

`func (o *InlineResponse2004) SetArchiveFiles(v []ArchiveBucketFile)`

SetArchiveFiles sets ArchiveFiles field to given value.

### HasArchiveFiles

`func (o *InlineResponse2004) HasArchiveFiles() bool`

HasArchiveFiles returns a boolean if a field has been set.

### GetArchiveFileCount

`func (o *InlineResponse2004) GetArchiveFileCount() int64`

GetArchiveFileCount returns the ArchiveFileCount field if non-nil, zero value otherwise.

### GetArchiveFileCountOk

`func (o *InlineResponse2004) GetArchiveFileCountOk() (*int64, bool)`

GetArchiveFileCountOk returns a tuple with the ArchiveFileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveFileCount

`func (o *InlineResponse2004) SetArchiveFileCount(v int64)`

SetArchiveFileCount sets ArchiveFileCount field to given value.

### HasArchiveFileCount

`func (o *InlineResponse2004) HasArchiveFileCount() bool`

HasArchiveFileCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


