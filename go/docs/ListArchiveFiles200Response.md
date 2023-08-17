# ListArchiveFiles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**ArchiveBucket** | Pointer to [**ArchiveBucket**](ArchiveBucket.md) |  | [optional] 
**IsOwner** | Pointer to **bool** |  | [optional] 
**ParentDirectory** | Pointer to **NullableString** |  | [optional] 
**ArchiveFiles** | Pointer to [**[]ArchiveBucketFile**](ArchiveBucketFile.md) |  | [optional] 

## Methods

### NewListArchiveFiles200Response

`func NewListArchiveFiles200Response() *ListArchiveFiles200Response`

NewListArchiveFiles200Response instantiates a new ListArchiveFiles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListArchiveFiles200ResponseWithDefaults

`func NewListArchiveFiles200ResponseWithDefaults() *ListArchiveFiles200Response`

NewListArchiveFiles200ResponseWithDefaults instantiates a new ListArchiveFiles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListArchiveFiles200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListArchiveFiles200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListArchiveFiles200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListArchiveFiles200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetArchiveBucket

`func (o *ListArchiveFiles200Response) GetArchiveBucket() ArchiveBucket`

GetArchiveBucket returns the ArchiveBucket field if non-nil, zero value otherwise.

### GetArchiveBucketOk

`func (o *ListArchiveFiles200Response) GetArchiveBucketOk() (*ArchiveBucket, bool)`

GetArchiveBucketOk returns a tuple with the ArchiveBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveBucket

`func (o *ListArchiveFiles200Response) SetArchiveBucket(v ArchiveBucket)`

SetArchiveBucket sets ArchiveBucket field to given value.

### HasArchiveBucket

`func (o *ListArchiveFiles200Response) HasArchiveBucket() bool`

HasArchiveBucket returns a boolean if a field has been set.

### GetIsOwner

`func (o *ListArchiveFiles200Response) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *ListArchiveFiles200Response) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *ListArchiveFiles200Response) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.

### HasIsOwner

`func (o *ListArchiveFiles200Response) HasIsOwner() bool`

HasIsOwner returns a boolean if a field has been set.

### GetParentDirectory

`func (o *ListArchiveFiles200Response) GetParentDirectory() string`

GetParentDirectory returns the ParentDirectory field if non-nil, zero value otherwise.

### GetParentDirectoryOk

`func (o *ListArchiveFiles200Response) GetParentDirectoryOk() (*string, bool)`

GetParentDirectoryOk returns a tuple with the ParentDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDirectory

`func (o *ListArchiveFiles200Response) SetParentDirectory(v string)`

SetParentDirectory sets ParentDirectory field to given value.

### HasParentDirectory

`func (o *ListArchiveFiles200Response) HasParentDirectory() bool`

HasParentDirectory returns a boolean if a field has been set.

### SetParentDirectoryNil

`func (o *ListArchiveFiles200Response) SetParentDirectoryNil(b bool)`

 SetParentDirectoryNil sets the value for ParentDirectory to be an explicit nil

### UnsetParentDirectory
`func (o *ListArchiveFiles200Response) UnsetParentDirectory()`

UnsetParentDirectory ensures that no value is present for ParentDirectory, not even an explicit nil
### GetArchiveFiles

`func (o *ListArchiveFiles200Response) GetArchiveFiles() []ArchiveBucketFile`

GetArchiveFiles returns the ArchiveFiles field if non-nil, zero value otherwise.

### GetArchiveFilesOk

`func (o *ListArchiveFiles200Response) GetArchiveFilesOk() (*[]ArchiveBucketFile, bool)`

GetArchiveFilesOk returns a tuple with the ArchiveFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveFiles

`func (o *ListArchiveFiles200Response) SetArchiveFiles(v []ArchiveBucketFile)`

SetArchiveFiles sets ArchiveFiles field to given value.

### HasArchiveFiles

`func (o *ListArchiveFiles200Response) HasArchiveFiles() bool`

HasArchiveFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


