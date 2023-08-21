# BillingServersServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefType** | Pointer to **string** |  | [optional] 
**RefUUID** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**NumUnits** | Pointer to **float32** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Usages** | Pointer to [**[]BillingServersUsages**](BillingServersUsages.md) |  | [optional] 
**NumUsages** | Pointer to **int64** |  | [optional] 
**TotalUsages** | Pointer to **int64** |  | [optional] 
**HasMoreUsages** | Pointer to **bool** |  | [optional] 
**FoundPricing** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ServerUUID** | Pointer to **string** |  | [optional] 
**ServerUniqueId** | Pointer to **NullableString** |  | [optional] 
**ServerExternalId** | Pointer to **NullableString** |  | [optional] 
**ServerInternalId** | Pointer to **NullableString** |  | [optional] 
**ResourcePoolId** | Pointer to **NullableString** |  | [optional] 
**ResourcePoolName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBillingServersServers

`func NewBillingServersServers() *BillingServersServers`

NewBillingServersServers instantiates a new BillingServersServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingServersServersWithDefaults

`func NewBillingServersServersWithDefaults() *BillingServersServers`

NewBillingServersServersWithDefaults instantiates a new BillingServersServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefType

`func (o *BillingServersServers) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *BillingServersServers) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *BillingServersServers) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *BillingServersServers) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefUUID

`func (o *BillingServersServers) GetRefUUID() string`

GetRefUUID returns the RefUUID field if non-nil, zero value otherwise.

### GetRefUUIDOk

`func (o *BillingServersServers) GetRefUUIDOk() (*string, bool)`

GetRefUUIDOk returns a tuple with the RefUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUUID

`func (o *BillingServersServers) SetRefUUID(v string)`

SetRefUUID sets RefUUID field to given value.

### HasRefUUID

`func (o *BillingServersServers) HasRefUUID() bool`

HasRefUUID returns a boolean if a field has been set.

### GetRefId

`func (o *BillingServersServers) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *BillingServersServers) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *BillingServersServers) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *BillingServersServers) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetStartDate

`func (o *BillingServersServers) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BillingServersServers) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BillingServersServers) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BillingServersServers) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *BillingServersServers) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BillingServersServers) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BillingServersServers) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BillingServersServers) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetCost

`func (o *BillingServersServers) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *BillingServersServers) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *BillingServersServers) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *BillingServersServers) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPrice

`func (o *BillingServersServers) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingServersServers) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingServersServers) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingServersServers) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetNumUnits

`func (o *BillingServersServers) GetNumUnits() float32`

GetNumUnits returns the NumUnits field if non-nil, zero value otherwise.

### GetNumUnitsOk

`func (o *BillingServersServers) GetNumUnitsOk() (*float32, bool)`

GetNumUnitsOk returns a tuple with the NumUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUnits

`func (o *BillingServersServers) SetNumUnits(v float32)`

SetNumUnits sets NumUnits field to given value.

### HasNumUnits

`func (o *BillingServersServers) HasNumUnits() bool`

HasNumUnits returns a boolean if a field has been set.

### GetUnit

`func (o *BillingServersServers) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *BillingServersServers) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *BillingServersServers) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *BillingServersServers) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingServersServers) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingServersServers) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingServersServers) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingServersServers) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetUsages

`func (o *BillingServersServers) GetUsages() []BillingServersUsages`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *BillingServersServers) GetUsagesOk() (*[]BillingServersUsages, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *BillingServersServers) SetUsages(v []BillingServersUsages)`

SetUsages sets Usages field to given value.

### HasUsages

`func (o *BillingServersServers) HasUsages() bool`

HasUsages returns a boolean if a field has been set.

### GetNumUsages

`func (o *BillingServersServers) GetNumUsages() int64`

GetNumUsages returns the NumUsages field if non-nil, zero value otherwise.

### GetNumUsagesOk

`func (o *BillingServersServers) GetNumUsagesOk() (*int64, bool)`

GetNumUsagesOk returns a tuple with the NumUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUsages

`func (o *BillingServersServers) SetNumUsages(v int64)`

SetNumUsages sets NumUsages field to given value.

### HasNumUsages

`func (o *BillingServersServers) HasNumUsages() bool`

HasNumUsages returns a boolean if a field has been set.

### GetTotalUsages

`func (o *BillingServersServers) GetTotalUsages() int64`

GetTotalUsages returns the TotalUsages field if non-nil, zero value otherwise.

### GetTotalUsagesOk

`func (o *BillingServersServers) GetTotalUsagesOk() (*int64, bool)`

GetTotalUsagesOk returns a tuple with the TotalUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsages

`func (o *BillingServersServers) SetTotalUsages(v int64)`

SetTotalUsages sets TotalUsages field to given value.

### HasTotalUsages

`func (o *BillingServersServers) HasTotalUsages() bool`

HasTotalUsages returns a boolean if a field has been set.

### GetHasMoreUsages

`func (o *BillingServersServers) GetHasMoreUsages() bool`

GetHasMoreUsages returns the HasMoreUsages field if non-nil, zero value otherwise.

### GetHasMoreUsagesOk

`func (o *BillingServersServers) GetHasMoreUsagesOk() (*bool, bool)`

GetHasMoreUsagesOk returns a tuple with the HasMoreUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMoreUsages

`func (o *BillingServersServers) SetHasMoreUsages(v bool)`

SetHasMoreUsages sets HasMoreUsages field to given value.

### HasHasMoreUsages

`func (o *BillingServersServers) HasHasMoreUsages() bool`

HasHasMoreUsages returns a boolean if a field has been set.

### GetFoundPricing

`func (o *BillingServersServers) GetFoundPricing() bool`

GetFoundPricing returns the FoundPricing field if non-nil, zero value otherwise.

### GetFoundPricingOk

`func (o *BillingServersServers) GetFoundPricingOk() (*bool, bool)`

GetFoundPricingOk returns a tuple with the FoundPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundPricing

`func (o *BillingServersServers) SetFoundPricing(v bool)`

SetFoundPricing sets FoundPricing field to given value.

### HasFoundPricing

`func (o *BillingServersServers) HasFoundPricing() bool`

HasFoundPricing returns a boolean if a field has been set.

### GetName

`func (o *BillingServersServers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingServersServers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingServersServers) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingServersServers) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServerUUID

`func (o *BillingServersServers) GetServerUUID() string`

GetServerUUID returns the ServerUUID field if non-nil, zero value otherwise.

### GetServerUUIDOk

`func (o *BillingServersServers) GetServerUUIDOk() (*string, bool)`

GetServerUUIDOk returns a tuple with the ServerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUUID

`func (o *BillingServersServers) SetServerUUID(v string)`

SetServerUUID sets ServerUUID field to given value.

### HasServerUUID

`func (o *BillingServersServers) HasServerUUID() bool`

HasServerUUID returns a boolean if a field has been set.

### GetServerUniqueId

`func (o *BillingServersServers) GetServerUniqueId() string`

GetServerUniqueId returns the ServerUniqueId field if non-nil, zero value otherwise.

### GetServerUniqueIdOk

`func (o *BillingServersServers) GetServerUniqueIdOk() (*string, bool)`

GetServerUniqueIdOk returns a tuple with the ServerUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUniqueId

`func (o *BillingServersServers) SetServerUniqueId(v string)`

SetServerUniqueId sets ServerUniqueId field to given value.

### HasServerUniqueId

`func (o *BillingServersServers) HasServerUniqueId() bool`

HasServerUniqueId returns a boolean if a field has been set.

### SetServerUniqueIdNil

`func (o *BillingServersServers) SetServerUniqueIdNil(b bool)`

 SetServerUniqueIdNil sets the value for ServerUniqueId to be an explicit nil

### UnsetServerUniqueId
`func (o *BillingServersServers) UnsetServerUniqueId()`

UnsetServerUniqueId ensures that no value is present for ServerUniqueId, not even an explicit nil
### GetServerExternalId

`func (o *BillingServersServers) GetServerExternalId() string`

GetServerExternalId returns the ServerExternalId field if non-nil, zero value otherwise.

### GetServerExternalIdOk

`func (o *BillingServersServers) GetServerExternalIdOk() (*string, bool)`

GetServerExternalIdOk returns a tuple with the ServerExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerExternalId

`func (o *BillingServersServers) SetServerExternalId(v string)`

SetServerExternalId sets ServerExternalId field to given value.

### HasServerExternalId

`func (o *BillingServersServers) HasServerExternalId() bool`

HasServerExternalId returns a boolean if a field has been set.

### SetServerExternalIdNil

`func (o *BillingServersServers) SetServerExternalIdNil(b bool)`

 SetServerExternalIdNil sets the value for ServerExternalId to be an explicit nil

### UnsetServerExternalId
`func (o *BillingServersServers) UnsetServerExternalId()`

UnsetServerExternalId ensures that no value is present for ServerExternalId, not even an explicit nil
### GetServerInternalId

`func (o *BillingServersServers) GetServerInternalId() string`

GetServerInternalId returns the ServerInternalId field if non-nil, zero value otherwise.

### GetServerInternalIdOk

`func (o *BillingServersServers) GetServerInternalIdOk() (*string, bool)`

GetServerInternalIdOk returns a tuple with the ServerInternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInternalId

`func (o *BillingServersServers) SetServerInternalId(v string)`

SetServerInternalId sets ServerInternalId field to given value.

### HasServerInternalId

`func (o *BillingServersServers) HasServerInternalId() bool`

HasServerInternalId returns a boolean if a field has been set.

### SetServerInternalIdNil

`func (o *BillingServersServers) SetServerInternalIdNil(b bool)`

 SetServerInternalIdNil sets the value for ServerInternalId to be an explicit nil

### UnsetServerInternalId
`func (o *BillingServersServers) UnsetServerInternalId()`

UnsetServerInternalId ensures that no value is present for ServerInternalId, not even an explicit nil
### GetResourcePoolId

`func (o *BillingServersServers) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *BillingServersServers) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *BillingServersServers) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *BillingServersServers) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### SetResourcePoolIdNil

`func (o *BillingServersServers) SetResourcePoolIdNil(b bool)`

 SetResourcePoolIdNil sets the value for ResourcePoolId to be an explicit nil

### UnsetResourcePoolId
`func (o *BillingServersServers) UnsetResourcePoolId()`

UnsetResourcePoolId ensures that no value is present for ResourcePoolId, not even an explicit nil
### GetResourcePoolName

`func (o *BillingServersServers) GetResourcePoolName() string`

GetResourcePoolName returns the ResourcePoolName field if non-nil, zero value otherwise.

### GetResourcePoolNameOk

`func (o *BillingServersServers) GetResourcePoolNameOk() (*string, bool)`

GetResourcePoolNameOk returns a tuple with the ResourcePoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolName

`func (o *BillingServersServers) SetResourcePoolName(v string)`

SetResourcePoolName sets ResourcePoolName field to given value.

### HasResourcePoolName

`func (o *BillingServersServers) HasResourcePoolName() bool`

HasResourcePoolName returns a boolean if a field has been set.

### SetResourcePoolNameNil

`func (o *BillingServersServers) SetResourcePoolNameNil(b bool)`

 SetResourcePoolNameNil sets the value for ResourcePoolName to be an explicit nil

### UnsetResourcePoolName
`func (o *BillingServersServers) UnsetResourcePoolName()`

UnsetResourcePoolName ensures that no value is present for ResourcePoolName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


