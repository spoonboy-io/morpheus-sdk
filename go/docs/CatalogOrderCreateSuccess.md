# CatalogOrderCreateSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Items** | Pointer to [**[]CatalogOrderCreateSuccessItems**](CatalogOrderCreateSuccessItems.md) |  | [optional] 
**Stats** | Pointer to [**CatalogCartStats**](catalogCart_stats.md) |  | [optional] 

## Methods

### NewCatalogOrderCreateSuccess

`func NewCatalogOrderCreateSuccess() *CatalogOrderCreateSuccess`

NewCatalogOrderCreateSuccess instantiates a new CatalogOrderCreateSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogOrderCreateSuccessWithDefaults

`func NewCatalogOrderCreateSuccessWithDefaults() *CatalogOrderCreateSuccess`

NewCatalogOrderCreateSuccessWithDefaults instantiates a new CatalogOrderCreateSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CatalogOrderCreateSuccess) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogOrderCreateSuccess) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogOrderCreateSuccess) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CatalogOrderCreateSuccess) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CatalogOrderCreateSuccess) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogOrderCreateSuccess) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogOrderCreateSuccess) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogOrderCreateSuccess) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CatalogOrderCreateSuccess) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CatalogOrderCreateSuccess) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetItems

`func (o *CatalogOrderCreateSuccess) GetItems() []CatalogOrderCreateSuccessItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CatalogOrderCreateSuccess) GetItemsOk() (*[]CatalogOrderCreateSuccessItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CatalogOrderCreateSuccess) SetItems(v []CatalogOrderCreateSuccessItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *CatalogOrderCreateSuccess) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetStats

`func (o *CatalogOrderCreateSuccess) GetStats() CatalogCartStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *CatalogOrderCreateSuccess) GetStatsOk() (*CatalogCartStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *CatalogOrderCreateSuccess) SetStats(v CatalogCartStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *CatalogOrderCreateSuccess) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

