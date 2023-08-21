# BillingServersUsages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ZoneName** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** |  | [optional] 
**Volumes** | Pointer to [**[]BillingServersVolumes**](BillingServersVolumes.md) |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**MaxCpu** | Pointer to **NullableString** |  | [optional] 
**MaxCores** | Pointer to **int64** |  | [optional] 
**ServerExternalId** | Pointer to **NullableString** |  | [optional] 
**ServerInternalId** | Pointer to **NullableString** |  | [optional] 
**PlanName** | Pointer to **string** |  | [optional] 
**HourlyPrice** | Pointer to **float32** |  | [optional] 
**HourlyCost** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**PricesUsed** | Pointer to [**[]BillingServersPricesUsed**](BillingServersPricesUsed.md) |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**CreatedByUser** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **int64** |  | [optional] 
**SiteId** | Pointer to **NullableString** |  | [optional] 
**SiteName** | Pointer to **NullableString** |  | [optional] 
**SiteUUID** | Pointer to **NullableString** |  | [optional] 
**SiteCode** | Pointer to **NullableString** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ApplicablePrices** | Pointer to [**[]BillingServersApplicablePrices**](BillingServersApplicablePrices.md) |  | [optional] 
**ServicePlanId** | Pointer to **int64** |  | [optional] 
**ServicePlanName** | Pointer to **string** |  | [optional] 
**ResourcePoolId** | Pointer to **NullableString** |  | [optional] 
**ResourcePoolName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBillingServersUsages

`func NewBillingServersUsages() *BillingServersUsages`

NewBillingServersUsages instantiates a new BillingServersUsages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingServersUsagesWithDefaults

`func NewBillingServersUsagesWithDefaults() *BillingServersUsages`

NewBillingServersUsagesWithDefaults instantiates a new BillingServersUsages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BillingServersUsages) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingServersUsages) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingServersUsages) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingServersUsages) HasName() bool`

HasName returns a boolean if a field has been set.

### GetZoneName

`func (o *BillingServersUsages) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *BillingServersUsages) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *BillingServersUsages) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *BillingServersUsages) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### GetAccountName

`func (o *BillingServersUsages) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *BillingServersUsages) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *BillingServersUsages) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *BillingServersUsages) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetVolumes

`func (o *BillingServersUsages) GetVolumes() []BillingServersVolumes`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *BillingServersUsages) GetVolumesOk() (*[]BillingServersVolumes, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *BillingServersUsages) SetVolumes(v []BillingServersVolumes)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *BillingServersUsages) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetMaxMemory

`func (o *BillingServersUsages) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *BillingServersUsages) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *BillingServersUsages) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *BillingServersUsages) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxCpu

`func (o *BillingServersUsages) GetMaxCpu() string`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *BillingServersUsages) GetMaxCpuOk() (*string, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *BillingServersUsages) SetMaxCpu(v string)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *BillingServersUsages) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### SetMaxCpuNil

`func (o *BillingServersUsages) SetMaxCpuNil(b bool)`

 SetMaxCpuNil sets the value for MaxCpu to be an explicit nil

### UnsetMaxCpu
`func (o *BillingServersUsages) UnsetMaxCpu()`

UnsetMaxCpu ensures that no value is present for MaxCpu, not even an explicit nil
### GetMaxCores

`func (o *BillingServersUsages) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *BillingServersUsages) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *BillingServersUsages) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *BillingServersUsages) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetServerExternalId

`func (o *BillingServersUsages) GetServerExternalId() string`

GetServerExternalId returns the ServerExternalId field if non-nil, zero value otherwise.

### GetServerExternalIdOk

`func (o *BillingServersUsages) GetServerExternalIdOk() (*string, bool)`

GetServerExternalIdOk returns a tuple with the ServerExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerExternalId

`func (o *BillingServersUsages) SetServerExternalId(v string)`

SetServerExternalId sets ServerExternalId field to given value.

### HasServerExternalId

`func (o *BillingServersUsages) HasServerExternalId() bool`

HasServerExternalId returns a boolean if a field has been set.

### SetServerExternalIdNil

`func (o *BillingServersUsages) SetServerExternalIdNil(b bool)`

 SetServerExternalIdNil sets the value for ServerExternalId to be an explicit nil

### UnsetServerExternalId
`func (o *BillingServersUsages) UnsetServerExternalId()`

UnsetServerExternalId ensures that no value is present for ServerExternalId, not even an explicit nil
### GetServerInternalId

`func (o *BillingServersUsages) GetServerInternalId() string`

GetServerInternalId returns the ServerInternalId field if non-nil, zero value otherwise.

### GetServerInternalIdOk

`func (o *BillingServersUsages) GetServerInternalIdOk() (*string, bool)`

GetServerInternalIdOk returns a tuple with the ServerInternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInternalId

`func (o *BillingServersUsages) SetServerInternalId(v string)`

SetServerInternalId sets ServerInternalId field to given value.

### HasServerInternalId

`func (o *BillingServersUsages) HasServerInternalId() bool`

HasServerInternalId returns a boolean if a field has been set.

### SetServerInternalIdNil

`func (o *BillingServersUsages) SetServerInternalIdNil(b bool)`

 SetServerInternalIdNil sets the value for ServerInternalId to be an explicit nil

### UnsetServerInternalId
`func (o *BillingServersUsages) UnsetServerInternalId()`

UnsetServerInternalId ensures that no value is present for ServerInternalId, not even an explicit nil
### GetPlanName

`func (o *BillingServersUsages) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *BillingServersUsages) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *BillingServersUsages) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *BillingServersUsages) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetHourlyPrice

`func (o *BillingServersUsages) GetHourlyPrice() float32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *BillingServersUsages) GetHourlyPriceOk() (*float32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *BillingServersUsages) SetHourlyPrice(v float32)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *BillingServersUsages) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.

### GetHourlyCost

`func (o *BillingServersUsages) GetHourlyCost() float32`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *BillingServersUsages) GetHourlyCostOk() (*float32, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *BillingServersUsages) SetHourlyCost(v float32)`

SetHourlyCost sets HourlyCost field to given value.

### HasHourlyCost

`func (o *BillingServersUsages) HasHourlyCost() bool`

HasHourlyCost returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingServersUsages) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingServersUsages) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingServersUsages) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingServersUsages) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPricesUsed

`func (o *BillingServersUsages) GetPricesUsed() []BillingServersPricesUsed`

GetPricesUsed returns the PricesUsed field if non-nil, zero value otherwise.

### GetPricesUsedOk

`func (o *BillingServersUsages) GetPricesUsedOk() (*[]BillingServersPricesUsed, bool)`

GetPricesUsedOk returns a tuple with the PricesUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricesUsed

`func (o *BillingServersUsages) SetPricesUsed(v []BillingServersPricesUsed)`

SetPricesUsed sets PricesUsed field to given value.

### HasPricesUsed

`func (o *BillingServersUsages) HasPricesUsed() bool`

HasPricesUsed returns a boolean if a field has been set.

### GetCost

`func (o *BillingServersUsages) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *BillingServersUsages) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *BillingServersUsages) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *BillingServersUsages) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPrice

`func (o *BillingServersUsages) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingServersUsages) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingServersUsages) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingServersUsages) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCreatedByUser

`func (o *BillingServersUsages) GetCreatedByUser() string`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *BillingServersUsages) GetCreatedByUserOk() (*string, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *BillingServersUsages) SetCreatedByUser(v string)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *BillingServersUsages) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *BillingServersUsages) GetCreatedByUserId() int64`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *BillingServersUsages) GetCreatedByUserIdOk() (*int64, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *BillingServersUsages) SetCreatedByUserId(v int64)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *BillingServersUsages) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetSiteId

`func (o *BillingServersUsages) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *BillingServersUsages) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *BillingServersUsages) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *BillingServersUsages) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteIdNil

`func (o *BillingServersUsages) SetSiteIdNil(b bool)`

 SetSiteIdNil sets the value for SiteId to be an explicit nil

### UnsetSiteId
`func (o *BillingServersUsages) UnsetSiteId()`

UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
### GetSiteName

`func (o *BillingServersUsages) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *BillingServersUsages) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *BillingServersUsages) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *BillingServersUsages) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### SetSiteNameNil

`func (o *BillingServersUsages) SetSiteNameNil(b bool)`

 SetSiteNameNil sets the value for SiteName to be an explicit nil

### UnsetSiteName
`func (o *BillingServersUsages) UnsetSiteName()`

UnsetSiteName ensures that no value is present for SiteName, not even an explicit nil
### GetSiteUUID

`func (o *BillingServersUsages) GetSiteUUID() string`

GetSiteUUID returns the SiteUUID field if non-nil, zero value otherwise.

### GetSiteUUIDOk

`func (o *BillingServersUsages) GetSiteUUIDOk() (*string, bool)`

GetSiteUUIDOk returns a tuple with the SiteUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteUUID

`func (o *BillingServersUsages) SetSiteUUID(v string)`

SetSiteUUID sets SiteUUID field to given value.

### HasSiteUUID

`func (o *BillingServersUsages) HasSiteUUID() bool`

HasSiteUUID returns a boolean if a field has been set.

### SetSiteUUIDNil

`func (o *BillingServersUsages) SetSiteUUIDNil(b bool)`

 SetSiteUUIDNil sets the value for SiteUUID to be an explicit nil

### UnsetSiteUUID
`func (o *BillingServersUsages) UnsetSiteUUID()`

UnsetSiteUUID ensures that no value is present for SiteUUID, not even an explicit nil
### GetSiteCode

`func (o *BillingServersUsages) GetSiteCode() string`

GetSiteCode returns the SiteCode field if non-nil, zero value otherwise.

### GetSiteCodeOk

`func (o *BillingServersUsages) GetSiteCodeOk() (*string, bool)`

GetSiteCodeOk returns a tuple with the SiteCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteCode

`func (o *BillingServersUsages) SetSiteCode(v string)`

SetSiteCode sets SiteCode field to given value.

### HasSiteCode

`func (o *BillingServersUsages) HasSiteCode() bool`

HasSiteCode returns a boolean if a field has been set.

### SetSiteCodeNil

`func (o *BillingServersUsages) SetSiteCodeNil(b bool)`

 SetSiteCodeNil sets the value for SiteCode to be an explicit nil

### UnsetSiteCode
`func (o *BillingServersUsages) UnsetSiteCode()`

UnsetSiteCode ensures that no value is present for SiteCode, not even an explicit nil
### GetStartDate

`func (o *BillingServersUsages) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BillingServersUsages) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BillingServersUsages) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BillingServersUsages) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *BillingServersUsages) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BillingServersUsages) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BillingServersUsages) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BillingServersUsages) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStatus

`func (o *BillingServersUsages) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingServersUsages) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingServersUsages) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BillingServersUsages) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *BillingServersUsages) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BillingServersUsages) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BillingServersUsages) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *BillingServersUsages) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetApplicablePrices

`func (o *BillingServersUsages) GetApplicablePrices() []BillingServersApplicablePrices`

GetApplicablePrices returns the ApplicablePrices field if non-nil, zero value otherwise.

### GetApplicablePricesOk

`func (o *BillingServersUsages) GetApplicablePricesOk() (*[]BillingServersApplicablePrices, bool)`

GetApplicablePricesOk returns a tuple with the ApplicablePrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicablePrices

`func (o *BillingServersUsages) SetApplicablePrices(v []BillingServersApplicablePrices)`

SetApplicablePrices sets ApplicablePrices field to given value.

### HasApplicablePrices

`func (o *BillingServersUsages) HasApplicablePrices() bool`

HasApplicablePrices returns a boolean if a field has been set.

### GetServicePlanId

`func (o *BillingServersUsages) GetServicePlanId() int64`

GetServicePlanId returns the ServicePlanId field if non-nil, zero value otherwise.

### GetServicePlanIdOk

`func (o *BillingServersUsages) GetServicePlanIdOk() (*int64, bool)`

GetServicePlanIdOk returns a tuple with the ServicePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanId

`func (o *BillingServersUsages) SetServicePlanId(v int64)`

SetServicePlanId sets ServicePlanId field to given value.

### HasServicePlanId

`func (o *BillingServersUsages) HasServicePlanId() bool`

HasServicePlanId returns a boolean if a field has been set.

### GetServicePlanName

`func (o *BillingServersUsages) GetServicePlanName() string`

GetServicePlanName returns the ServicePlanName field if non-nil, zero value otherwise.

### GetServicePlanNameOk

`func (o *BillingServersUsages) GetServicePlanNameOk() (*string, bool)`

GetServicePlanNameOk returns a tuple with the ServicePlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanName

`func (o *BillingServersUsages) SetServicePlanName(v string)`

SetServicePlanName sets ServicePlanName field to given value.

### HasServicePlanName

`func (o *BillingServersUsages) HasServicePlanName() bool`

HasServicePlanName returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *BillingServersUsages) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *BillingServersUsages) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *BillingServersUsages) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *BillingServersUsages) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### SetResourcePoolIdNil

`func (o *BillingServersUsages) SetResourcePoolIdNil(b bool)`

 SetResourcePoolIdNil sets the value for ResourcePoolId to be an explicit nil

### UnsetResourcePoolId
`func (o *BillingServersUsages) UnsetResourcePoolId()`

UnsetResourcePoolId ensures that no value is present for ResourcePoolId, not even an explicit nil
### GetResourcePoolName

`func (o *BillingServersUsages) GetResourcePoolName() string`

GetResourcePoolName returns the ResourcePoolName field if non-nil, zero value otherwise.

### GetResourcePoolNameOk

`func (o *BillingServersUsages) GetResourcePoolNameOk() (*string, bool)`

GetResourcePoolNameOk returns a tuple with the ResourcePoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolName

`func (o *BillingServersUsages) SetResourcePoolName(v string)`

SetResourcePoolName sets ResourcePoolName field to given value.

### HasResourcePoolName

`func (o *BillingServersUsages) HasResourcePoolName() bool`

HasResourcePoolName returns a boolean if a field has been set.

### SetResourcePoolNameNil

`func (o *BillingServersUsages) SetResourcePoolNameNil(b bool)`

 SetResourcePoolNameNil sets the value for ResourcePoolName to be an explicit nil

### UnsetResourcePoolName
`func (o *BillingServersUsages) UnsetResourcePoolName()`

UnsetResourcePoolName ensures that no value is present for ResourcePoolName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


