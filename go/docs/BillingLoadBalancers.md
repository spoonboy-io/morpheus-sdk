# BillingLoadBalancers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **float32** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**LoadBalancers** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 

## Methods

### NewBillingLoadBalancers

`func NewBillingLoadBalancers() *BillingLoadBalancers`

NewBillingLoadBalancers instantiates a new BillingLoadBalancers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingLoadBalancersWithDefaults

`func NewBillingLoadBalancersWithDefaults() *BillingLoadBalancers`

NewBillingLoadBalancersWithDefaults instantiates a new BillingLoadBalancers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *BillingLoadBalancers) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingLoadBalancers) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingLoadBalancers) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingLoadBalancers) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCost

`func (o *BillingLoadBalancers) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *BillingLoadBalancers) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *BillingLoadBalancers) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *BillingLoadBalancers) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetLoadBalancers

`func (o *BillingLoadBalancers) GetLoadBalancers() []map[string]interface{}`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *BillingLoadBalancers) GetLoadBalancersOk() (*[]map[string]interface{}, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *BillingLoadBalancers) SetLoadBalancers(v []map[string]interface{})`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *BillingLoadBalancers) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.

### GetCount

`func (o *BillingLoadBalancers) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BillingLoadBalancers) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BillingLoadBalancers) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *BillingLoadBalancers) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


