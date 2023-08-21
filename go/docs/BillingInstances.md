# BillingInstances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **float32** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**Instances** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 

## Methods

### NewBillingInstances

`func NewBillingInstances() *BillingInstances`

NewBillingInstances instantiates a new BillingInstances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInstancesWithDefaults

`func NewBillingInstancesWithDefaults() *BillingInstances`

NewBillingInstancesWithDefaults instantiates a new BillingInstances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *BillingInstances) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingInstances) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingInstances) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingInstances) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCost

`func (o *BillingInstances) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *BillingInstances) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *BillingInstances) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *BillingInstances) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetInstances

`func (o *BillingInstances) GetInstances() []map[string]interface{}`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *BillingInstances) GetInstancesOk() (*[]map[string]interface{}, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *BillingInstances) SetInstances(v []map[string]interface{})`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *BillingInstances) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetCount

`func (o *BillingInstances) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BillingInstances) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BillingInstances) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *BillingInstances) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


