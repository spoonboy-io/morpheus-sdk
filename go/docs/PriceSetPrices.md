# PriceSetPrices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**PriceType** | Pointer to **string** |  | [optional] 
**PriceUnit** | Pointer to **string** |  | [optional] 
**AdditionalPriceUnit** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **NullableFloat32** |  | [optional] 
**CustomPrice** | Pointer to **NullableFloat32** |  | [optional] 
**MarkupType** | Pointer to **NullableString** |  | [optional] 
**Markup** | Pointer to **int64** |  | [optional] 
**MarkupPercent** | Pointer to **NullableFloat32** |  | [optional] 
**Cost** | Pointer to **NullableFloat32** |  | [optional] 
**Currency** | Pointer to **NullableString** |  | [optional] 
**IncurCharges** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **NullableString** |  | [optional] 
**Software** | Pointer to **NullableString** |  | [optional] 
**VolumeType** | Pointer to [**NullablePriceSetVolumeType**](priceSet_volumeType.md) |  | [optional] 
**Datastore** | Pointer to **NullableString** |  | [optional] 
**CrossCloudApply** | Pointer to **NullableBool** |  | [optional] 
**Account** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPriceSetPrices

`func NewPriceSetPrices() *PriceSetPrices`

NewPriceSetPrices instantiates a new PriceSetPrices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceSetPricesWithDefaults

`func NewPriceSetPricesWithDefaults() *PriceSetPrices`

NewPriceSetPricesWithDefaults instantiates a new PriceSetPrices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PriceSetPrices) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PriceSetPrices) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PriceSetPrices) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PriceSetPrices) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PriceSetPrices) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PriceSetPrices) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PriceSetPrices) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PriceSetPrices) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *PriceSetPrices) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PriceSetPrices) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PriceSetPrices) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PriceSetPrices) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetActive

`func (o *PriceSetPrices) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PriceSetPrices) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PriceSetPrices) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PriceSetPrices) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPriceType

`func (o *PriceSetPrices) GetPriceType() string`

GetPriceType returns the PriceType field if non-nil, zero value otherwise.

### GetPriceTypeOk

`func (o *PriceSetPrices) GetPriceTypeOk() (*string, bool)`

GetPriceTypeOk returns a tuple with the PriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceType

`func (o *PriceSetPrices) SetPriceType(v string)`

SetPriceType sets PriceType field to given value.

### HasPriceType

`func (o *PriceSetPrices) HasPriceType() bool`

HasPriceType returns a boolean if a field has been set.

### GetPriceUnit

`func (o *PriceSetPrices) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *PriceSetPrices) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *PriceSetPrices) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *PriceSetPrices) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetAdditionalPriceUnit

`func (o *PriceSetPrices) GetAdditionalPriceUnit() string`

GetAdditionalPriceUnit returns the AdditionalPriceUnit field if non-nil, zero value otherwise.

### GetAdditionalPriceUnitOk

`func (o *PriceSetPrices) GetAdditionalPriceUnitOk() (*string, bool)`

GetAdditionalPriceUnitOk returns a tuple with the AdditionalPriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPriceUnit

`func (o *PriceSetPrices) SetAdditionalPriceUnit(v string)`

SetAdditionalPriceUnit sets AdditionalPriceUnit field to given value.

### HasAdditionalPriceUnit

`func (o *PriceSetPrices) HasAdditionalPriceUnit() bool`

HasAdditionalPriceUnit returns a boolean if a field has been set.

### SetAdditionalPriceUnitNil

`func (o *PriceSetPrices) SetAdditionalPriceUnitNil(b bool)`

 SetAdditionalPriceUnitNil sets the value for AdditionalPriceUnit to be an explicit nil

### UnsetAdditionalPriceUnit
`func (o *PriceSetPrices) UnsetAdditionalPriceUnit()`

UnsetAdditionalPriceUnit ensures that no value is present for AdditionalPriceUnit, not even an explicit nil
### GetPrice

`func (o *PriceSetPrices) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PriceSetPrices) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PriceSetPrices) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PriceSetPrices) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *PriceSetPrices) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *PriceSetPrices) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetCustomPrice

`func (o *PriceSetPrices) GetCustomPrice() float32`

GetCustomPrice returns the CustomPrice field if non-nil, zero value otherwise.

### GetCustomPriceOk

`func (o *PriceSetPrices) GetCustomPriceOk() (*float32, bool)`

GetCustomPriceOk returns a tuple with the CustomPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPrice

`func (o *PriceSetPrices) SetCustomPrice(v float32)`

SetCustomPrice sets CustomPrice field to given value.

### HasCustomPrice

`func (o *PriceSetPrices) HasCustomPrice() bool`

HasCustomPrice returns a boolean if a field has been set.

### SetCustomPriceNil

`func (o *PriceSetPrices) SetCustomPriceNil(b bool)`

 SetCustomPriceNil sets the value for CustomPrice to be an explicit nil

### UnsetCustomPrice
`func (o *PriceSetPrices) UnsetCustomPrice()`

UnsetCustomPrice ensures that no value is present for CustomPrice, not even an explicit nil
### GetMarkupType

`func (o *PriceSetPrices) GetMarkupType() string`

GetMarkupType returns the MarkupType field if non-nil, zero value otherwise.

### GetMarkupTypeOk

`func (o *PriceSetPrices) GetMarkupTypeOk() (*string, bool)`

GetMarkupTypeOk returns a tuple with the MarkupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupType

`func (o *PriceSetPrices) SetMarkupType(v string)`

SetMarkupType sets MarkupType field to given value.

### HasMarkupType

`func (o *PriceSetPrices) HasMarkupType() bool`

HasMarkupType returns a boolean if a field has been set.

### SetMarkupTypeNil

`func (o *PriceSetPrices) SetMarkupTypeNil(b bool)`

 SetMarkupTypeNil sets the value for MarkupType to be an explicit nil

### UnsetMarkupType
`func (o *PriceSetPrices) UnsetMarkupType()`

UnsetMarkupType ensures that no value is present for MarkupType, not even an explicit nil
### GetMarkup

`func (o *PriceSetPrices) GetMarkup() int64`

GetMarkup returns the Markup field if non-nil, zero value otherwise.

### GetMarkupOk

`func (o *PriceSetPrices) GetMarkupOk() (*int64, bool)`

GetMarkupOk returns a tuple with the Markup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkup

`func (o *PriceSetPrices) SetMarkup(v int64)`

SetMarkup sets Markup field to given value.

### HasMarkup

`func (o *PriceSetPrices) HasMarkup() bool`

HasMarkup returns a boolean if a field has been set.

### GetMarkupPercent

`func (o *PriceSetPrices) GetMarkupPercent() float32`

GetMarkupPercent returns the MarkupPercent field if non-nil, zero value otherwise.

### GetMarkupPercentOk

`func (o *PriceSetPrices) GetMarkupPercentOk() (*float32, bool)`

GetMarkupPercentOk returns a tuple with the MarkupPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupPercent

`func (o *PriceSetPrices) SetMarkupPercent(v float32)`

SetMarkupPercent sets MarkupPercent field to given value.

### HasMarkupPercent

`func (o *PriceSetPrices) HasMarkupPercent() bool`

HasMarkupPercent returns a boolean if a field has been set.

### SetMarkupPercentNil

`func (o *PriceSetPrices) SetMarkupPercentNil(b bool)`

 SetMarkupPercentNil sets the value for MarkupPercent to be an explicit nil

### UnsetMarkupPercent
`func (o *PriceSetPrices) UnsetMarkupPercent()`

UnsetMarkupPercent ensures that no value is present for MarkupPercent, not even an explicit nil
### GetCost

`func (o *PriceSetPrices) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *PriceSetPrices) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *PriceSetPrices) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *PriceSetPrices) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *PriceSetPrices) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *PriceSetPrices) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCurrency

`func (o *PriceSetPrices) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PriceSetPrices) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PriceSetPrices) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PriceSetPrices) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *PriceSetPrices) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *PriceSetPrices) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetIncurCharges

`func (o *PriceSetPrices) GetIncurCharges() string`

GetIncurCharges returns the IncurCharges field if non-nil, zero value otherwise.

### GetIncurChargesOk

`func (o *PriceSetPrices) GetIncurChargesOk() (*string, bool)`

GetIncurChargesOk returns a tuple with the IncurCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncurCharges

`func (o *PriceSetPrices) SetIncurCharges(v string)`

SetIncurCharges sets IncurCharges field to given value.

### HasIncurCharges

`func (o *PriceSetPrices) HasIncurCharges() bool`

HasIncurCharges returns a boolean if a field has been set.

### GetPlatform

`func (o *PriceSetPrices) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PriceSetPrices) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PriceSetPrices) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PriceSetPrices) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *PriceSetPrices) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *PriceSetPrices) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetSoftware

`func (o *PriceSetPrices) GetSoftware() string`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *PriceSetPrices) GetSoftwareOk() (*string, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *PriceSetPrices) SetSoftware(v string)`

SetSoftware sets Software field to given value.

### HasSoftware

`func (o *PriceSetPrices) HasSoftware() bool`

HasSoftware returns a boolean if a field has been set.

### SetSoftwareNil

`func (o *PriceSetPrices) SetSoftwareNil(b bool)`

 SetSoftwareNil sets the value for Software to be an explicit nil

### UnsetSoftware
`func (o *PriceSetPrices) UnsetSoftware()`

UnsetSoftware ensures that no value is present for Software, not even an explicit nil
### GetVolumeType

`func (o *PriceSetPrices) GetVolumeType() PriceSetVolumeType`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *PriceSetPrices) GetVolumeTypeOk() (*PriceSetVolumeType, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *PriceSetPrices) SetVolumeType(v PriceSetVolumeType)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *PriceSetPrices) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### SetVolumeTypeNil

`func (o *PriceSetPrices) SetVolumeTypeNil(b bool)`

 SetVolumeTypeNil sets the value for VolumeType to be an explicit nil

### UnsetVolumeType
`func (o *PriceSetPrices) UnsetVolumeType()`

UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil
### GetDatastore

`func (o *PriceSetPrices) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *PriceSetPrices) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *PriceSetPrices) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *PriceSetPrices) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### SetDatastoreNil

`func (o *PriceSetPrices) SetDatastoreNil(b bool)`

 SetDatastoreNil sets the value for Datastore to be an explicit nil

### UnsetDatastore
`func (o *PriceSetPrices) UnsetDatastore()`

UnsetDatastore ensures that no value is present for Datastore, not even an explicit nil
### GetCrossCloudApply

`func (o *PriceSetPrices) GetCrossCloudApply() bool`

GetCrossCloudApply returns the CrossCloudApply field if non-nil, zero value otherwise.

### GetCrossCloudApplyOk

`func (o *PriceSetPrices) GetCrossCloudApplyOk() (*bool, bool)`

GetCrossCloudApplyOk returns a tuple with the CrossCloudApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossCloudApply

`func (o *PriceSetPrices) SetCrossCloudApply(v bool)`

SetCrossCloudApply sets CrossCloudApply field to given value.

### HasCrossCloudApply

`func (o *PriceSetPrices) HasCrossCloudApply() bool`

HasCrossCloudApply returns a boolean if a field has been set.

### SetCrossCloudApplyNil

`func (o *PriceSetPrices) SetCrossCloudApplyNil(b bool)`

 SetCrossCloudApplyNil sets the value for CrossCloudApply to be an explicit nil

### UnsetCrossCloudApply
`func (o *PriceSetPrices) UnsetCrossCloudApply()`

UnsetCrossCloudApply ensures that no value is present for CrossCloudApply, not even an explicit nil
### GetAccount

`func (o *PriceSetPrices) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PriceSetPrices) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PriceSetPrices) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PriceSetPrices) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *PriceSetPrices) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *PriceSetPrices) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


