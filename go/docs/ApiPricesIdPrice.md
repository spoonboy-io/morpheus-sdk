# ApiPricesIdPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Price name | [optional] 
**Code** | Pointer to **string** | Price code, must be unique | [optional] 
**Account** | Pointer to [**ApiPricesPriceAccount**](_api_prices_price_account.md) |  | [optional] 
**PriceType** | Pointer to **string** | Restricts query to only load only prices with specified priceType. * &#x60;fixed&#x60; - Everything * &#x60;compute&#x60; - Memory + CPU * &#x60;memory&#x60; - Memory * &#x60;cores&#x60; - Cores * &#x60;storage&#x60; - Storage * &#x60;datastore&#x60; - Datastore * &#x60;platform&#x60; - Platform * &#x60;software&#x60; - Software * &#x60;load_balancer&#x60; - Load Balancer * &#x60;load_balancer_virtual_server&#x60; - Load Balancer Virtual Server  | [optional] 
**PriceUnit** | Pointer to **string** | The unit of pricing | [optional] 
**IncurCharges** | Pointer to **string** | Indicates when to incur charge | [optional] 
**Currency** | Pointer to **string** | ISO Currency code | [optional] 
**Cost** | Pointer to **float32** | Cost | [optional] 
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

### NewApiPricesIdPrice

`func NewApiPricesIdPrice() *ApiPricesIdPrice`

NewApiPricesIdPrice instantiates a new ApiPricesIdPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPricesIdPriceWithDefaults

`func NewApiPricesIdPriceWithDefaults() *ApiPricesIdPrice`

NewApiPricesIdPriceWithDefaults instantiates a new ApiPricesIdPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiPricesIdPrice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiPricesIdPrice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiPricesIdPrice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiPricesIdPrice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ApiPricesIdPrice) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiPricesIdPrice) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiPricesIdPrice) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiPricesIdPrice) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetAccount

`func (o *ApiPricesIdPrice) GetAccount() ApiPricesPriceAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApiPricesIdPrice) GetAccountOk() (*ApiPricesPriceAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApiPricesIdPrice) SetAccount(v ApiPricesPriceAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApiPricesIdPrice) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetPriceType

`func (o *ApiPricesIdPrice) GetPriceType() string`

GetPriceType returns the PriceType field if non-nil, zero value otherwise.

### GetPriceTypeOk

`func (o *ApiPricesIdPrice) GetPriceTypeOk() (*string, bool)`

GetPriceTypeOk returns a tuple with the PriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceType

`func (o *ApiPricesIdPrice) SetPriceType(v string)`

SetPriceType sets PriceType field to given value.

### HasPriceType

`func (o *ApiPricesIdPrice) HasPriceType() bool`

HasPriceType returns a boolean if a field has been set.

### GetPriceUnit

`func (o *ApiPricesIdPrice) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *ApiPricesIdPrice) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *ApiPricesIdPrice) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *ApiPricesIdPrice) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetIncurCharges

`func (o *ApiPricesIdPrice) GetIncurCharges() string`

GetIncurCharges returns the IncurCharges field if non-nil, zero value otherwise.

### GetIncurChargesOk

`func (o *ApiPricesIdPrice) GetIncurChargesOk() (*string, bool)`

GetIncurChargesOk returns a tuple with the IncurCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncurCharges

`func (o *ApiPricesIdPrice) SetIncurCharges(v string)`

SetIncurCharges sets IncurCharges field to given value.

### HasIncurCharges

`func (o *ApiPricesIdPrice) HasIncurCharges() bool`

HasIncurCharges returns a boolean if a field has been set.

### GetCurrency

`func (o *ApiPricesIdPrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ApiPricesIdPrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ApiPricesIdPrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ApiPricesIdPrice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCost

`func (o *ApiPricesIdPrice) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ApiPricesIdPrice) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ApiPricesIdPrice) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ApiPricesIdPrice) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetMarkupType

`func (o *ApiPricesIdPrice) GetMarkupType() string`

GetMarkupType returns the MarkupType field if non-nil, zero value otherwise.

### GetMarkupTypeOk

`func (o *ApiPricesIdPrice) GetMarkupTypeOk() (*string, bool)`

GetMarkupTypeOk returns a tuple with the MarkupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupType

`func (o *ApiPricesIdPrice) SetMarkupType(v string)`

SetMarkupType sets MarkupType field to given value.

### HasMarkupType

`func (o *ApiPricesIdPrice) HasMarkupType() bool`

HasMarkupType returns a boolean if a field has been set.

### GetMarkup

`func (o *ApiPricesIdPrice) GetMarkup() float32`

GetMarkup returns the Markup field if non-nil, zero value otherwise.

### GetMarkupOk

`func (o *ApiPricesIdPrice) GetMarkupOk() (*float32, bool)`

GetMarkupOk returns a tuple with the Markup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkup

`func (o *ApiPricesIdPrice) SetMarkup(v float32)`

SetMarkup sets Markup field to given value.

### HasMarkup

`func (o *ApiPricesIdPrice) HasMarkup() bool`

HasMarkup returns a boolean if a field has been set.

### GetMarkupPercent

`func (o *ApiPricesIdPrice) GetMarkupPercent() float32`

GetMarkupPercent returns the MarkupPercent field if non-nil, zero value otherwise.

### GetMarkupPercentOk

`func (o *ApiPricesIdPrice) GetMarkupPercentOk() (*float32, bool)`

GetMarkupPercentOk returns a tuple with the MarkupPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupPercent

`func (o *ApiPricesIdPrice) SetMarkupPercent(v float32)`

SetMarkupPercent sets MarkupPercent field to given value.

### HasMarkupPercent

`func (o *ApiPricesIdPrice) HasMarkupPercent() bool`

HasMarkupPercent returns a boolean if a field has been set.

### GetCustomPrice

`func (o *ApiPricesIdPrice) GetCustomPrice() float32`

GetCustomPrice returns the CustomPrice field if non-nil, zero value otherwise.

### GetCustomPriceOk

`func (o *ApiPricesIdPrice) GetCustomPriceOk() (*float32, bool)`

GetCustomPriceOk returns a tuple with the CustomPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPrice

`func (o *ApiPricesIdPrice) SetCustomPrice(v float32)`

SetCustomPrice sets CustomPrice field to given value.

### HasCustomPrice

`func (o *ApiPricesIdPrice) HasCustomPrice() bool`

HasCustomPrice returns a boolean if a field has been set.

### GetPlatform

`func (o *ApiPricesIdPrice) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ApiPricesIdPrice) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ApiPricesIdPrice) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ApiPricesIdPrice) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSoftware

`func (o *ApiPricesIdPrice) GetSoftware() string`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *ApiPricesIdPrice) GetSoftwareOk() (*string, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *ApiPricesIdPrice) SetSoftware(v string)`

SetSoftware sets Software field to given value.

### HasSoftware

`func (o *ApiPricesIdPrice) HasSoftware() bool`

HasSoftware returns a boolean if a field has been set.

### GetVolumeType

`func (o *ApiPricesIdPrice) GetVolumeType() ApiPricesPriceVolumeType`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *ApiPricesIdPrice) GetVolumeTypeOk() (*ApiPricesPriceVolumeType, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *ApiPricesIdPrice) SetVolumeType(v ApiPricesPriceVolumeType)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *ApiPricesIdPrice) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetDatastore

`func (o *ApiPricesIdPrice) GetDatastore() ApiPricesPriceDatastore`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *ApiPricesIdPrice) GetDatastoreOk() (*ApiPricesPriceDatastore, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *ApiPricesIdPrice) SetDatastore(v ApiPricesPriceDatastore)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *ApiPricesIdPrice) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetCrossCloudApply

`func (o *ApiPricesIdPrice) GetCrossCloudApply() bool`

GetCrossCloudApply returns the CrossCloudApply field if non-nil, zero value otherwise.

### GetCrossCloudApplyOk

`func (o *ApiPricesIdPrice) GetCrossCloudApplyOk() (*bool, bool)`

GetCrossCloudApplyOk returns a tuple with the CrossCloudApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossCloudApply

`func (o *ApiPricesIdPrice) SetCrossCloudApply(v bool)`

SetCrossCloudApply sets CrossCloudApply field to given value.

### HasCrossCloudApply

`func (o *ApiPricesIdPrice) HasCrossCloudApply() bool`

HasCrossCloudApply returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


