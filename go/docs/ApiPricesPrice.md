# ApiPricesPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Price name | 
**Code** | **string** | Price code, must be unique | 
**Account** | Pointer to [**ApiPricesPriceAccount**](_api_prices_price_account.md) |  | [optional] 
**PriceType** | **string** | Restricts query to only load only prices with specified priceType. * &#x60;fixed&#x60; - Everything * &#x60;compute&#x60; - Memory + CPU * &#x60;memory&#x60; - Memory * &#x60;cores&#x60; - Cores * &#x60;storage&#x60; - Storage * &#x60;datastore&#x60; - Datastore * &#x60;platform&#x60; - Platform * &#x60;software&#x60; - Software * &#x60;load_balancer&#x60; - Load Balancer * &#x60;load_balancer_virtual_server&#x60; - Load Balancer Virtual Server  | 
**PriceUnit** | **string** | The unit of pricing | 
**IncurCharges** | **string** | Indicates when to incur charge | 
**Currency** | **string** | ISO Currency code | 
**Cost** | **float32** | Cost | 
**MarkupType** | Pointer to **string** | Price adjustment type | [optional] 
**Markup** | Pointer to **float32** | Amount for &#x60;fixed&#x60; price adjustment type | [optional] 
**MarkupPercent** | Pointer to **float32** | Percent for &#x60;percent&#x60; price adjustment type | [optional] 
**CustomPrice** | Pointer to **float32** | Custom price for &#x60;custom&#x60; price adjustment type | [optional] 
**Platform** | Pointer to **string** | Platform.  Required for &#x60;platform&#x60; price type | [optional] 
**Software** | Pointer to **string** | Software.  Required for software price type | [optional] 
**VolumeType** | Pointer to [**ApiPricesPriceVolumeType**](_api_prices_price_volumeType.md) |  | [optional] 
**Datastore** | Pointer to [**ApiPricesPriceDatastore**](_api_prices_price_datastore.md) |  | [optional] 
**CrossCloudApply** | Pointer to **bool** | Apply price across clouds, optional true/false flag for datastore price type | [optional] 

## Methods

### NewApiPricesPrice

`func NewApiPricesPrice(name string, code string, priceType string, priceUnit string, incurCharges string, currency string, cost float32, ) *ApiPricesPrice`

NewApiPricesPrice instantiates a new ApiPricesPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPricesPriceWithDefaults

`func NewApiPricesPriceWithDefaults() *ApiPricesPrice`

NewApiPricesPriceWithDefaults instantiates a new ApiPricesPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiPricesPrice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiPricesPrice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiPricesPrice) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *ApiPricesPrice) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiPricesPrice) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiPricesPrice) SetCode(v string)`

SetCode sets Code field to given value.


### GetAccount

`func (o *ApiPricesPrice) GetAccount() ApiPricesPriceAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApiPricesPrice) GetAccountOk() (*ApiPricesPriceAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApiPricesPrice) SetAccount(v ApiPricesPriceAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApiPricesPrice) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetPriceType

`func (o *ApiPricesPrice) GetPriceType() string`

GetPriceType returns the PriceType field if non-nil, zero value otherwise.

### GetPriceTypeOk

`func (o *ApiPricesPrice) GetPriceTypeOk() (*string, bool)`

GetPriceTypeOk returns a tuple with the PriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceType

`func (o *ApiPricesPrice) SetPriceType(v string)`

SetPriceType sets PriceType field to given value.


### GetPriceUnit

`func (o *ApiPricesPrice) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *ApiPricesPrice) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *ApiPricesPrice) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.


### GetIncurCharges

`func (o *ApiPricesPrice) GetIncurCharges() string`

GetIncurCharges returns the IncurCharges field if non-nil, zero value otherwise.

### GetIncurChargesOk

`func (o *ApiPricesPrice) GetIncurChargesOk() (*string, bool)`

GetIncurChargesOk returns a tuple with the IncurCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncurCharges

`func (o *ApiPricesPrice) SetIncurCharges(v string)`

SetIncurCharges sets IncurCharges field to given value.


### GetCurrency

`func (o *ApiPricesPrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ApiPricesPrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ApiPricesPrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCost

`func (o *ApiPricesPrice) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ApiPricesPrice) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ApiPricesPrice) SetCost(v float32)`

SetCost sets Cost field to given value.


### GetMarkupType

`func (o *ApiPricesPrice) GetMarkupType() string`

GetMarkupType returns the MarkupType field if non-nil, zero value otherwise.

### GetMarkupTypeOk

`func (o *ApiPricesPrice) GetMarkupTypeOk() (*string, bool)`

GetMarkupTypeOk returns a tuple with the MarkupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupType

`func (o *ApiPricesPrice) SetMarkupType(v string)`

SetMarkupType sets MarkupType field to given value.

### HasMarkupType

`func (o *ApiPricesPrice) HasMarkupType() bool`

HasMarkupType returns a boolean if a field has been set.

### GetMarkup

`func (o *ApiPricesPrice) GetMarkup() float32`

GetMarkup returns the Markup field if non-nil, zero value otherwise.

### GetMarkupOk

`func (o *ApiPricesPrice) GetMarkupOk() (*float32, bool)`

GetMarkupOk returns a tuple with the Markup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkup

`func (o *ApiPricesPrice) SetMarkup(v float32)`

SetMarkup sets Markup field to given value.

### HasMarkup

`func (o *ApiPricesPrice) HasMarkup() bool`

HasMarkup returns a boolean if a field has been set.

### GetMarkupPercent

`func (o *ApiPricesPrice) GetMarkupPercent() float32`

GetMarkupPercent returns the MarkupPercent field if non-nil, zero value otherwise.

### GetMarkupPercentOk

`func (o *ApiPricesPrice) GetMarkupPercentOk() (*float32, bool)`

GetMarkupPercentOk returns a tuple with the MarkupPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupPercent

`func (o *ApiPricesPrice) SetMarkupPercent(v float32)`

SetMarkupPercent sets MarkupPercent field to given value.

### HasMarkupPercent

`func (o *ApiPricesPrice) HasMarkupPercent() bool`

HasMarkupPercent returns a boolean if a field has been set.

### GetCustomPrice

`func (o *ApiPricesPrice) GetCustomPrice() float32`

GetCustomPrice returns the CustomPrice field if non-nil, zero value otherwise.

### GetCustomPriceOk

`func (o *ApiPricesPrice) GetCustomPriceOk() (*float32, bool)`

GetCustomPriceOk returns a tuple with the CustomPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPrice

`func (o *ApiPricesPrice) SetCustomPrice(v float32)`

SetCustomPrice sets CustomPrice field to given value.

### HasCustomPrice

`func (o *ApiPricesPrice) HasCustomPrice() bool`

HasCustomPrice returns a boolean if a field has been set.

### GetPlatform

`func (o *ApiPricesPrice) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ApiPricesPrice) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ApiPricesPrice) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ApiPricesPrice) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSoftware

`func (o *ApiPricesPrice) GetSoftware() string`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *ApiPricesPrice) GetSoftwareOk() (*string, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *ApiPricesPrice) SetSoftware(v string)`

SetSoftware sets Software field to given value.

### HasSoftware

`func (o *ApiPricesPrice) HasSoftware() bool`

HasSoftware returns a boolean if a field has been set.

### GetVolumeType

`func (o *ApiPricesPrice) GetVolumeType() ApiPricesPriceVolumeType`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *ApiPricesPrice) GetVolumeTypeOk() (*ApiPricesPriceVolumeType, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *ApiPricesPrice) SetVolumeType(v ApiPricesPriceVolumeType)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *ApiPricesPrice) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetDatastore

`func (o *ApiPricesPrice) GetDatastore() ApiPricesPriceDatastore`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *ApiPricesPrice) GetDatastoreOk() (*ApiPricesPriceDatastore, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *ApiPricesPrice) SetDatastore(v ApiPricesPriceDatastore)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *ApiPricesPrice) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetCrossCloudApply

`func (o *ApiPricesPrice) GetCrossCloudApply() bool`

GetCrossCloudApply returns the CrossCloudApply field if non-nil, zero value otherwise.

### GetCrossCloudApplyOk

`func (o *ApiPricesPrice) GetCrossCloudApplyOk() (*bool, bool)`

GetCrossCloudApplyOk returns a tuple with the CrossCloudApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossCloudApply

`func (o *ApiPricesPrice) SetCrossCloudApply(v bool)`

SetCrossCloudApply sets CrossCloudApply field to given value.

### HasCrossCloudApply

`func (o *ApiPricesPrice) HasCrossCloudApply() bool`

HasCrossCloudApply returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


