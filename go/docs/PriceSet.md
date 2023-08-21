# PriceSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**PriceUnit** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**RegionCode** | Pointer to **string** |  | [optional] 
**SystemCreated** | Pointer to **bool** |  | [optional] 
**Zone** | Pointer to **NullableString** |  | [optional] 
**ZonePool** | Pointer to **NullableString** |  | [optional] 
**Account** | Pointer to **NullableString** |  | [optional] 
**Prices** | Pointer to [**[]PriceSetPrices**](PriceSetPrices.md) |  | [optional] 

## Methods

### NewPriceSet

`func NewPriceSet() *PriceSet`

NewPriceSet instantiates a new PriceSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceSetWithDefaults

`func NewPriceSetWithDefaults() *PriceSet`

NewPriceSetWithDefaults instantiates a new PriceSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PriceSet) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PriceSet) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PriceSet) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PriceSet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PriceSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PriceSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PriceSet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PriceSet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *PriceSet) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PriceSet) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PriceSet) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PriceSet) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetActive

`func (o *PriceSet) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PriceSet) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PriceSet) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PriceSet) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPriceUnit

`func (o *PriceSet) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *PriceSet) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *PriceSet) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *PriceSet) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetType

`func (o *PriceSet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PriceSet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PriceSet) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PriceSet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRegionCode

`func (o *PriceSet) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *PriceSet) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *PriceSet) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *PriceSet) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetSystemCreated

`func (o *PriceSet) GetSystemCreated() bool`

GetSystemCreated returns the SystemCreated field if non-nil, zero value otherwise.

### GetSystemCreatedOk

`func (o *PriceSet) GetSystemCreatedOk() (*bool, bool)`

GetSystemCreatedOk returns a tuple with the SystemCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemCreated

`func (o *PriceSet) SetSystemCreated(v bool)`

SetSystemCreated sets SystemCreated field to given value.

### HasSystemCreated

`func (o *PriceSet) HasSystemCreated() bool`

HasSystemCreated returns a boolean if a field has been set.

### GetZone

`func (o *PriceSet) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *PriceSet) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *PriceSet) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *PriceSet) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *PriceSet) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *PriceSet) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetZonePool

`func (o *PriceSet) GetZonePool() string`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *PriceSet) GetZonePoolOk() (*string, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *PriceSet) SetZonePool(v string)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *PriceSet) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### SetZonePoolNil

`func (o *PriceSet) SetZonePoolNil(b bool)`

 SetZonePoolNil sets the value for ZonePool to be an explicit nil

### UnsetZonePool
`func (o *PriceSet) UnsetZonePool()`

UnsetZonePool ensures that no value is present for ZonePool, not even an explicit nil
### GetAccount

`func (o *PriceSet) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PriceSet) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PriceSet) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PriceSet) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *PriceSet) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *PriceSet) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetPrices

`func (o *PriceSet) GetPrices() []PriceSetPrices`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *PriceSet) GetPricesOk() (*[]PriceSetPrices, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *PriceSet) SetPrices(v []PriceSetPrices)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *PriceSet) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


