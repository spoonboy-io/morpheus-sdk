# ListCatalogItemTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**CatalogItemTypes** | Pointer to [**[]CatalogItemType**](CatalogItemType.md) |  | [optional] 

## Methods

### NewListCatalogItemTypes200Response

`func NewListCatalogItemTypes200Response() *ListCatalogItemTypes200Response`

NewListCatalogItemTypes200Response instantiates a new ListCatalogItemTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCatalogItemTypes200ResponseWithDefaults

`func NewListCatalogItemTypes200ResponseWithDefaults() *ListCatalogItemTypes200Response`

NewListCatalogItemTypes200ResponseWithDefaults instantiates a new ListCatalogItemTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListCatalogItemTypes200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListCatalogItemTypes200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListCatalogItemTypes200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListCatalogItemTypes200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetCatalogItemTypes

`func (o *ListCatalogItemTypes200Response) GetCatalogItemTypes() []CatalogItemType`

GetCatalogItemTypes returns the CatalogItemTypes field if non-nil, zero value otherwise.

### GetCatalogItemTypesOk

`func (o *ListCatalogItemTypes200Response) GetCatalogItemTypesOk() (*[]CatalogItemType, bool)`

GetCatalogItemTypesOk returns a tuple with the CatalogItemTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItemTypes

`func (o *ListCatalogItemTypes200Response) SetCatalogItemTypes(v []CatalogItemType)`

SetCatalogItemTypes sets CatalogItemTypes field to given value.

### HasCatalogItemTypes

`func (o *ListCatalogItemTypes200Response) HasCatalogItemTypes() bool`

HasCatalogItemTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


