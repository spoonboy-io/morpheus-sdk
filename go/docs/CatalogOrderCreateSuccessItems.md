# CatalogOrderCreateSuccessItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to [**AppBlueprint**](app_blueprint.md) |  | [optional] 
**Quantity** | Pointer to **int64** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCatalogOrderCreateSuccessItems

`func NewCatalogOrderCreateSuccessItems() *CatalogOrderCreateSuccessItems`

NewCatalogOrderCreateSuccessItems instantiates a new CatalogOrderCreateSuccessItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogOrderCreateSuccessItemsWithDefaults

`func NewCatalogOrderCreateSuccessItemsWithDefaults() *CatalogOrderCreateSuccessItems`

NewCatalogOrderCreateSuccessItemsWithDefaults instantiates a new CatalogOrderCreateSuccessItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CatalogOrderCreateSuccessItems) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogOrderCreateSuccessItems) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogOrderCreateSuccessItems) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CatalogOrderCreateSuccessItems) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CatalogOrderCreateSuccessItems) GetType() AppBlueprint`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogOrderCreateSuccessItems) GetTypeOk() (*AppBlueprint, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogOrderCreateSuccessItems) SetType(v AppBlueprint)`

SetType sets Type field to given value.

### HasType

`func (o *CatalogOrderCreateSuccessItems) HasType() bool`

HasType returns a boolean if a field has been set.

### GetQuantity

`func (o *CatalogOrderCreateSuccessItems) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CatalogOrderCreateSuccessItems) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CatalogOrderCreateSuccessItems) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CatalogOrderCreateSuccessItems) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetPrice

`func (o *CatalogOrderCreateSuccessItems) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CatalogOrderCreateSuccessItems) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CatalogOrderCreateSuccessItems) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CatalogOrderCreateSuccessItems) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *CatalogOrderCreateSuccessItems) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CatalogOrderCreateSuccessItems) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CatalogOrderCreateSuccessItems) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CatalogOrderCreateSuccessItems) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetUnit

`func (o *CatalogOrderCreateSuccessItems) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CatalogOrderCreateSuccessItems) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CatalogOrderCreateSuccessItems) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *CatalogOrderCreateSuccessItems) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetValid

`func (o *CatalogOrderCreateSuccessItems) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *CatalogOrderCreateSuccessItems) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *CatalogOrderCreateSuccessItems) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *CatalogOrderCreateSuccessItems) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetStatus

`func (o *CatalogOrderCreateSuccessItems) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CatalogOrderCreateSuccessItems) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CatalogOrderCreateSuccessItems) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CatalogOrderCreateSuccessItems) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *CatalogOrderCreateSuccessItems) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *CatalogOrderCreateSuccessItems) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *CatalogOrderCreateSuccessItems) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *CatalogOrderCreateSuccessItems) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *CatalogOrderCreateSuccessItems) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CatalogOrderCreateSuccessItems) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CatalogOrderCreateSuccessItems) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *CatalogOrderCreateSuccessItems) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


