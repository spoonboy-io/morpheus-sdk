# ApiTasksTaskFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | **string** | File Source i.e. &#x60;local&#x60;, &#x60;repository&#x60;, &#x60;url&#x60;. Default is &#x60;local&#x60;. | [default to "local"]
**Content** | Pointer to **string** | File content, the script text. Only required when sourceType is &#x60;local&#x60;. | [optional] 
**ContentPath** | Pointer to **string** | Content Path, the repo file location or url. Required when sourceType is &#x60;repository&#x60; or &#x60;url&#x60;. | [optional] 
**ContentRef** | Pointer to **string** | Content Ref, the branch/tag. Only used when sourceType is &#x60;repository&#x60;. | [optional] 
**Repository** | Pointer to [**ApiTasksTaskFileRepository**](_api_tasks_task_file_repository.md) |  | [optional] 

## Methods

### NewApiTasksTaskFile

`func NewApiTasksTaskFile(sourceType string, ) *ApiTasksTaskFile`

NewApiTasksTaskFile instantiates a new ApiTasksTaskFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTasksTaskFileWithDefaults

`func NewApiTasksTaskFileWithDefaults() *ApiTasksTaskFile`

NewApiTasksTaskFileWithDefaults instantiates a new ApiTasksTaskFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *ApiTasksTaskFile) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ApiTasksTaskFile) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ApiTasksTaskFile) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetContent

`func (o *ApiTasksTaskFile) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ApiTasksTaskFile) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ApiTasksTaskFile) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ApiTasksTaskFile) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentPath

`func (o *ApiTasksTaskFile) GetContentPath() string`

GetContentPath returns the ContentPath field if non-nil, zero value otherwise.

### GetContentPathOk

`func (o *ApiTasksTaskFile) GetContentPathOk() (*string, bool)`

GetContentPathOk returns a tuple with the ContentPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentPath

`func (o *ApiTasksTaskFile) SetContentPath(v string)`

SetContentPath sets ContentPath field to given value.

### HasContentPath

`func (o *ApiTasksTaskFile) HasContentPath() bool`

HasContentPath returns a boolean if a field has been set.

### GetContentRef

`func (o *ApiTasksTaskFile) GetContentRef() string`

GetContentRef returns the ContentRef field if non-nil, zero value otherwise.

### GetContentRefOk

`func (o *ApiTasksTaskFile) GetContentRefOk() (*string, bool)`

GetContentRefOk returns a tuple with the ContentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentRef

`func (o *ApiTasksTaskFile) SetContentRef(v string)`

SetContentRef sets ContentRef field to given value.

### HasContentRef

`func (o *ApiTasksTaskFile) HasContentRef() bool`

HasContentRef returns a boolean if a field has been set.

### GetRepository

`func (o *ApiTasksTaskFile) GetRepository() ApiTasksTaskFileRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ApiTasksTaskFile) GetRepositoryOk() (*ApiTasksTaskFileRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ApiTasksTaskFile) SetRepository(v ApiTasksTaskFileRepository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ApiTasksTaskFile) HasRepository() bool`

HasRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


