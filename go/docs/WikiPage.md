# WikiPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UrlName** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **NullableString** |  | [optional] 
**RefType** | Pointer to **NullableString** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to [**InlineResponse200107NetworkPoolCreatedBy**](inline_response_200_107_networkPool_createdBy.md) |  | [optional] 
**UpdatedBy** | Pointer to [**InlineResponse200107NetworkPoolCreatedBy**](inline_response_200_107_networkPool_createdBy.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewWikiPage

`func NewWikiPage() *WikiPage`

NewWikiPage instantiates a new WikiPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWikiPageWithDefaults

`func NewWikiPageWithDefaults() *WikiPage`

NewWikiPageWithDefaults instantiates a new WikiPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WikiPage) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WikiPage) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WikiPage) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *WikiPage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WikiPage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WikiPage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WikiPage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WikiPage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrlName

`func (o *WikiPage) GetUrlName() string`

GetUrlName returns the UrlName field if non-nil, zero value otherwise.

### GetUrlNameOk

`func (o *WikiPage) GetUrlNameOk() (*string, bool)`

GetUrlNameOk returns a tuple with the UrlName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlName

`func (o *WikiPage) SetUrlName(v string)`

SetUrlName sets UrlName field to given value.

### HasUrlName

`func (o *WikiPage) HasUrlName() bool`

HasUrlName returns a boolean if a field has been set.

### GetCategory

`func (o *WikiPage) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *WikiPage) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *WikiPage) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *WikiPage) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetRefId

`func (o *WikiPage) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *WikiPage) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *WikiPage) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *WikiPage) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *WikiPage) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *WikiPage) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetRefType

`func (o *WikiPage) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *WikiPage) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *WikiPage) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *WikiPage) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *WikiPage) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *WikiPage) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetFormat

`func (o *WikiPage) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *WikiPage) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *WikiPage) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *WikiPage) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetContent

`func (o *WikiPage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WikiPage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WikiPage) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *WikiPage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedBy

`func (o *WikiPage) GetCreatedBy() InlineResponse200107NetworkPoolCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WikiPage) GetCreatedByOk() (*InlineResponse200107NetworkPoolCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WikiPage) SetCreatedBy(v InlineResponse200107NetworkPoolCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WikiPage) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *WikiPage) GetUpdatedBy() InlineResponse200107NetworkPoolCreatedBy`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *WikiPage) GetUpdatedByOk() (*InlineResponse200107NetworkPoolCreatedBy, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *WikiPage) SetUpdatedBy(v InlineResponse200107NetworkPoolCreatedBy)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *WikiPage) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *WikiPage) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *WikiPage) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *WikiPage) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *WikiPage) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *WikiPage) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *WikiPage) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *WikiPage) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *WikiPage) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


