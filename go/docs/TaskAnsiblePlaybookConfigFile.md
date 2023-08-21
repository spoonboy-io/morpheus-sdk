# TaskAnsiblePlaybookConfigFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** |  | [optional] 
**SourceType** | Pointer to **NullableString** |  | [optional] 
**ContentRef** | Pointer to **NullableString** |  | [optional] 
**ContentPath** | Pointer to **NullableString** |  | [optional] 
**Repository** | Pointer to [**NullableInlineResponse20082LoadBalancerInstanceSslCert**](inline_response_200_82_loadBalancerInstance_sslCert.md) |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTaskAnsiblePlaybookConfigFile

`func NewTaskAnsiblePlaybookConfigFile() *TaskAnsiblePlaybookConfigFile`

NewTaskAnsiblePlaybookConfigFile instantiates a new TaskAnsiblePlaybookConfigFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskAnsiblePlaybookConfigFileWithDefaults

`func NewTaskAnsiblePlaybookConfigFileWithDefaults() *TaskAnsiblePlaybookConfigFile`

NewTaskAnsiblePlaybookConfigFileWithDefaults instantiates a new TaskAnsiblePlaybookConfigFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskAnsiblePlaybookConfigFile) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskAnsiblePlaybookConfigFile) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskAnsiblePlaybookConfigFile) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TaskAnsiblePlaybookConfigFile) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TaskAnsiblePlaybookConfigFile) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TaskAnsiblePlaybookConfigFile) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSourceType

`func (o *TaskAnsiblePlaybookConfigFile) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *TaskAnsiblePlaybookConfigFile) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *TaskAnsiblePlaybookConfigFile) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *TaskAnsiblePlaybookConfigFile) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceTypeNil

`func (o *TaskAnsiblePlaybookConfigFile) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *TaskAnsiblePlaybookConfigFile) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetContentRef

`func (o *TaskAnsiblePlaybookConfigFile) GetContentRef() string`

GetContentRef returns the ContentRef field if non-nil, zero value otherwise.

### GetContentRefOk

`func (o *TaskAnsiblePlaybookConfigFile) GetContentRefOk() (*string, bool)`

GetContentRefOk returns a tuple with the ContentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentRef

`func (o *TaskAnsiblePlaybookConfigFile) SetContentRef(v string)`

SetContentRef sets ContentRef field to given value.

### HasContentRef

`func (o *TaskAnsiblePlaybookConfigFile) HasContentRef() bool`

HasContentRef returns a boolean if a field has been set.

### SetContentRefNil

`func (o *TaskAnsiblePlaybookConfigFile) SetContentRefNil(b bool)`

 SetContentRefNil sets the value for ContentRef to be an explicit nil

### UnsetContentRef
`func (o *TaskAnsiblePlaybookConfigFile) UnsetContentRef()`

UnsetContentRef ensures that no value is present for ContentRef, not even an explicit nil
### GetContentPath

`func (o *TaskAnsiblePlaybookConfigFile) GetContentPath() string`

GetContentPath returns the ContentPath field if non-nil, zero value otherwise.

### GetContentPathOk

`func (o *TaskAnsiblePlaybookConfigFile) GetContentPathOk() (*string, bool)`

GetContentPathOk returns a tuple with the ContentPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentPath

`func (o *TaskAnsiblePlaybookConfigFile) SetContentPath(v string)`

SetContentPath sets ContentPath field to given value.

### HasContentPath

`func (o *TaskAnsiblePlaybookConfigFile) HasContentPath() bool`

HasContentPath returns a boolean if a field has been set.

### SetContentPathNil

`func (o *TaskAnsiblePlaybookConfigFile) SetContentPathNil(b bool)`

 SetContentPathNil sets the value for ContentPath to be an explicit nil

### UnsetContentPath
`func (o *TaskAnsiblePlaybookConfigFile) UnsetContentPath()`

UnsetContentPath ensures that no value is present for ContentPath, not even an explicit nil
### GetRepository

`func (o *TaskAnsiblePlaybookConfigFile) GetRepository() InlineResponse20082LoadBalancerInstanceSslCert`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *TaskAnsiblePlaybookConfigFile) GetRepositoryOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *TaskAnsiblePlaybookConfigFile) SetRepository(v InlineResponse20082LoadBalancerInstanceSslCert)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *TaskAnsiblePlaybookConfigFile) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *TaskAnsiblePlaybookConfigFile) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *TaskAnsiblePlaybookConfigFile) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetContent

`func (o *TaskAnsiblePlaybookConfigFile) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *TaskAnsiblePlaybookConfigFile) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *TaskAnsiblePlaybookConfigFile) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *TaskAnsiblePlaybookConfigFile) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *TaskAnsiblePlaybookConfigFile) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *TaskAnsiblePlaybookConfigFile) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


