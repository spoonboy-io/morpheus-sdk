# ApiPriceSetsIdPriceSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Price set name | [optional] 
**Code** | Pointer to **string** | Price set code. Must be unique. | [optional] 
**RegionCode** | Pointer to **string** | Price set region code | [optional] 
**Zone** | Pointer to [**ApiPriceSetsPriceSetZone**](_api_price_sets_priceSet_zone.md) |  | [optional] 
**ZonePool** | Pointer to [**ApiPriceSetsPriceSetZonePool**](_api_price_sets_priceSet_zonePool.md) |  | [optional] 
**PriceUnit** | Pointer to **string** | Price Unit | [optional] 
**Type** | Pointer to **string** | Price set type | [optional] 
**Prices** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewApiPriceSetsIdPriceSet

`func NewApiPriceSetsIdPriceSet() *ApiPriceSetsIdPriceSet`

NewApiPriceSetsIdPriceSet instantiates a new ApiPriceSetsIdPriceSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPriceSetsIdPriceSetWithDefaults

`func NewApiPriceSetsIdPriceSetWithDefaults() *ApiPriceSetsIdPriceSet`

NewApiPriceSetsIdPriceSetWithDefaults instantiates a new ApiPriceSetsIdPriceSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiPriceSetsIdPriceSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiPriceSetsIdPriceSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiPriceSetsIdPriceSet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiPriceSetsIdPriceSet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ApiPriceSetsIdPriceSet) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiPriceSetsIdPriceSet) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiPriceSetsIdPriceSet) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiPriceSetsIdPriceSet) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetRegionCode

`func (o *ApiPriceSetsIdPriceSet) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *ApiPriceSetsIdPriceSet) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *ApiPriceSetsIdPriceSet) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *ApiPriceSetsIdPriceSet) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetZone

`func (o *ApiPriceSetsIdPriceSet) GetZone() ApiPriceSetsPriceSetZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ApiPriceSetsIdPriceSet) GetZoneOk() (*ApiPriceSetsPriceSetZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ApiPriceSetsIdPriceSet) SetZone(v ApiPriceSetsPriceSetZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ApiPriceSetsIdPriceSet) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZonePool

`func (o *ApiPriceSetsIdPriceSet) GetZonePool() ApiPriceSetsPriceSetZonePool`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *ApiPriceSetsIdPriceSet) GetZonePoolOk() (*ApiPriceSetsPriceSetZonePool, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *ApiPriceSetsIdPriceSet) SetZonePool(v ApiPriceSetsPriceSetZonePool)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *ApiPriceSetsIdPriceSet) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### GetPriceUnit

`func (o *ApiPriceSetsIdPriceSet) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *ApiPriceSetsIdPriceSet) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *ApiPriceSetsIdPriceSet) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *ApiPriceSetsIdPriceSet) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetType

`func (o *ApiPriceSetsIdPriceSet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiPriceSetsIdPriceSet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiPriceSetsIdPriceSet) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiPriceSetsIdPriceSet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrices

`func (o *ApiPriceSetsIdPriceSet) GetPrices() []int64`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *ApiPriceSetsIdPriceSet) GetPricesOk() (*[]int64, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *ApiPriceSetsIdPriceSet) SetPrices(v []int64)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *ApiPriceSetsIdPriceSet) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


