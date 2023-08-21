# BillingServerPrices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**PricePerUnit** | Pointer to **float32** |  | [optional] 
**CostPerUnit** | Pointer to **float32** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Quantity** | Pointer to **int64** |  | [optional] 
**DatastoreId** | Pointer to **NullableString** |  | [optional] 
**VolumeType** | Pointer to **string** |  | [optional] 
**Datastore** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingServerPrices

`func NewBillingServerPrices() *BillingServerPrices`

NewBillingServerPrices instantiates a new BillingServerPrices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingServerPricesWithDefaults

`func NewBillingServerPricesWithDefaults() *BillingServerPrices`

NewBillingServerPricesWithDefaults instantiates a new BillingServerPrices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BillingServerPrices) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillingServerPrices) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillingServerPrices) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BillingServerPrices) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPricePerUnit

`func (o *BillingServerPrices) GetPricePerUnit() float32`

GetPricePerUnit returns the PricePerUnit field if non-nil, zero value otherwise.

### GetPricePerUnitOk

`func (o *BillingServerPrices) GetPricePerUnitOk() (*float32, bool)`

GetPricePerUnitOk returns a tuple with the PricePerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerUnit

`func (o *BillingServerPrices) SetPricePerUnit(v float32)`

SetPricePerUnit sets PricePerUnit field to given value.

### HasPricePerUnit

`func (o *BillingServerPrices) HasPricePerUnit() bool`

HasPricePerUnit returns a boolean if a field has been set.

### GetCostPerUnit

`func (o *BillingServerPrices) GetCostPerUnit() float32`

GetCostPerUnit returns the CostPerUnit field if non-nil, zero value otherwise.

### GetCostPerUnitOk

`func (o *BillingServerPrices) GetCostPerUnitOk() (*float32, bool)`

GetCostPerUnitOk returns a tuple with the CostPerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerUnit

`func (o *BillingServerPrices) SetCostPerUnit(v float32)`

SetCostPerUnit sets CostPerUnit field to given value.

### HasCostPerUnit

`func (o *BillingServerPrices) HasCostPerUnit() bool`

HasCostPerUnit returns a boolean if a field has been set.

### GetCost

`func (o *BillingServerPrices) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *BillingServerPrices) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *BillingServerPrices) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *BillingServerPrices) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPrice

`func (o *BillingServerPrices) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingServerPrices) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingServerPrices) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingServerPrices) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *BillingServerPrices) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BillingServerPrices) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BillingServerPrices) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *BillingServerPrices) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetDatastoreId

`func (o *BillingServerPrices) GetDatastoreId() string`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *BillingServerPrices) GetDatastoreIdOk() (*string, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *BillingServerPrices) SetDatastoreId(v string)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *BillingServerPrices) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### SetDatastoreIdNil

`func (o *BillingServerPrices) SetDatastoreIdNil(b bool)`

 SetDatastoreIdNil sets the value for DatastoreId to be an explicit nil

### UnsetDatastoreId
`func (o *BillingServerPrices) UnsetDatastoreId()`

UnsetDatastoreId ensures that no value is present for DatastoreId, not even an explicit nil
### GetVolumeType

`func (o *BillingServerPrices) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *BillingServerPrices) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *BillingServerPrices) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *BillingServerPrices) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetDatastore

`func (o *BillingServerPrices) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *BillingServerPrices) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *BillingServerPrices) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *BillingServerPrices) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


