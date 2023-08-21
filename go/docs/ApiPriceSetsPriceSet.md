# ApiPriceSetsPriceSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Price set name | 
**Code** | **string** | Price set code. Must be unique. | 
**RegionCode** | Pointer to **string** | Price set region code | [optional] 
**Zone** | Pointer to [**ApiPriceSetsPriceSetZone**](_api_price_sets_priceSet_zone.md) |  | [optional] 
**ZonePool** | Pointer to [**ApiPriceSetsPriceSetZonePool**](_api_price_sets_priceSet_zonePool.md) |  | [optional] 
**PriceUnit** | **string** | Price Unit | 
**Type** | **string** | Price set type | 
**Prices** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewApiPriceSetsPriceSet

`func NewApiPriceSetsPriceSet(name string, code string, priceUnit string, type_ string, ) *ApiPriceSetsPriceSet`

NewApiPriceSetsPriceSet instantiates a new ApiPriceSetsPriceSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPriceSetsPriceSetWithDefaults

`func NewApiPriceSetsPriceSetWithDefaults() *ApiPriceSetsPriceSet`

NewApiPriceSetsPriceSetWithDefaults instantiates a new ApiPriceSetsPriceSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiPriceSetsPriceSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiPriceSetsPriceSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiPriceSetsPriceSet) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *ApiPriceSetsPriceSet) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiPriceSetsPriceSet) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiPriceSetsPriceSet) SetCode(v string)`

SetCode sets Code field to given value.


### GetRegionCode

`func (o *ApiPriceSetsPriceSet) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *ApiPriceSetsPriceSet) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *ApiPriceSetsPriceSet) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *ApiPriceSetsPriceSet) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetZone

`func (o *ApiPriceSetsPriceSet) GetZone() ApiPriceSetsPriceSetZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ApiPriceSetsPriceSet) GetZoneOk() (*ApiPriceSetsPriceSetZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ApiPriceSetsPriceSet) SetZone(v ApiPriceSetsPriceSetZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ApiPriceSetsPriceSet) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZonePool

`func (o *ApiPriceSetsPriceSet) GetZonePool() ApiPriceSetsPriceSetZonePool`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *ApiPriceSetsPriceSet) GetZonePoolOk() (*ApiPriceSetsPriceSetZonePool, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *ApiPriceSetsPriceSet) SetZonePool(v ApiPriceSetsPriceSetZonePool)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *ApiPriceSetsPriceSet) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### GetPriceUnit

`func (o *ApiPriceSetsPriceSet) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *ApiPriceSetsPriceSet) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *ApiPriceSetsPriceSet) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.


### GetType

`func (o *ApiPriceSetsPriceSet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiPriceSetsPriceSet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiPriceSetsPriceSet) SetType(v string)`

SetType sets Type field to given value.


### GetPrices

`func (o *ApiPriceSetsPriceSet) GetPrices() []int64`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *ApiPriceSetsPriceSet) GetPricesOk() (*[]int64, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *ApiPriceSetsPriceSet) SetPrices(v []int64)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *ApiPriceSetsPriceSet) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


