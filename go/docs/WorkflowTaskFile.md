# WorkflowTaskFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**SourceType** | Pointer to **string** |  | [optional] 
**ContentRef** | Pointer to **NullableString** |  | [optional] 
**ContentPath** | Pointer to **NullableString** |  | [optional] 
**Repository** | Pointer to [**NullableInlineResponse20082LoadBalancerInstanceSslCert**](inline_response_200_82_loadBalancerInstance_sslCert.md) |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWorkflowTaskFile

`func NewWorkflowTaskFile() *WorkflowTaskFile`

NewWorkflowTaskFile instantiates a new WorkflowTaskFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskFileWithDefaults

`func NewWorkflowTaskFileWithDefaults() *WorkflowTaskFile`

NewWorkflowTaskFileWithDefaults instantiates a new WorkflowTaskFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowTaskFile) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowTaskFile) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowTaskFile) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowTaskFile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSourceType

`func (o *WorkflowTaskFile) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *WorkflowTaskFile) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *WorkflowTaskFile) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *WorkflowTaskFile) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetContentRef

`func (o *WorkflowTaskFile) GetContentRef() string`

GetContentRef returns the ContentRef field if non-nil, zero value otherwise.

### GetContentRefOk

`func (o *WorkflowTaskFile) GetContentRefOk() (*string, bool)`

GetContentRefOk returns a tuple with the ContentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentRef

`func (o *WorkflowTaskFile) SetContentRef(v string)`

SetContentRef sets ContentRef field to given value.

### HasContentRef

`func (o *WorkflowTaskFile) HasContentRef() bool`

HasContentRef returns a boolean if a field has been set.

### SetContentRefNil

`func (o *WorkflowTaskFile) SetContentRefNil(b bool)`

 SetContentRefNil sets the value for ContentRef to be an explicit nil

### UnsetContentRef
`func (o *WorkflowTaskFile) UnsetContentRef()`

UnsetContentRef ensures that no value is present for ContentRef, not even an explicit nil
### GetContentPath

`func (o *WorkflowTaskFile) GetContentPath() string`

GetContentPath returns the ContentPath field if non-nil, zero value otherwise.

### GetContentPathOk

`func (o *WorkflowTaskFile) GetContentPathOk() (*string, bool)`

GetContentPathOk returns a tuple with the ContentPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentPath

`func (o *WorkflowTaskFile) SetContentPath(v string)`

SetContentPath sets ContentPath field to given value.

### HasContentPath

`func (o *WorkflowTaskFile) HasContentPath() bool`

HasContentPath returns a boolean if a field has been set.

### SetContentPathNil

`func (o *WorkflowTaskFile) SetContentPathNil(b bool)`

 SetContentPathNil sets the value for ContentPath to be an explicit nil

### UnsetContentPath
`func (o *WorkflowTaskFile) UnsetContentPath()`

UnsetContentPath ensures that no value is present for ContentPath, not even an explicit nil
### GetRepository

`func (o *WorkflowTaskFile) GetRepository() InlineResponse20082LoadBalancerInstanceSslCert`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *WorkflowTaskFile) GetRepositoryOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *WorkflowTaskFile) SetRepository(v InlineResponse20082LoadBalancerInstanceSslCert)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *WorkflowTaskFile) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *WorkflowTaskFile) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *WorkflowTaskFile) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetContent

`func (o *WorkflowTaskFile) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WorkflowTaskFile) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WorkflowTaskFile) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *WorkflowTaskFile) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *WorkflowTaskFile) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *WorkflowTaskFile) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


