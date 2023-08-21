# BillingSnapshots

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **float32** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**Snapshots** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 

## Methods

### NewBillingSnapshots

`func NewBillingSnapshots() *BillingSnapshots`

NewBillingSnapshots instantiates a new BillingSnapshots object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingSnapshotsWithDefaults

`func NewBillingSnapshotsWithDefaults() *BillingSnapshots`

NewBillingSnapshotsWithDefaults instantiates a new BillingSnapshots object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *BillingSnapshots) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingSnapshots) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingSnapshots) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingSnapshots) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCost

`func (o *BillingSnapshots) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *BillingSnapshots) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *BillingSnapshots) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *BillingSnapshots) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetSnapshots

`func (o *BillingSnapshots) GetSnapshots() []map[string]interface{}`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *BillingSnapshots) GetSnapshotsOk() (*[]map[string]interface{}, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *BillingSnapshots) SetSnapshots(v []map[string]interface{})`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *BillingSnapshots) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### GetCount

`func (o *BillingSnapshots) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BillingSnapshots) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BillingSnapshots) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *BillingSnapshots) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


