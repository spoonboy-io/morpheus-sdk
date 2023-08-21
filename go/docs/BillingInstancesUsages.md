# BillingInstancesUsages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**InstanceName** | Pointer to **string** |  | [optional] 
**ZoneName** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** |  | [optional] 
**Volumes** | Pointer to [**[]BillingInstancesVolumes**](BillingInstancesVolumes.md) |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**MaxCpu** | Pointer to **NullableString** |  | [optional] 
**MaxCores** | Pointer to **int64** |  | [optional] 
**ServerExternalId** | Pointer to **string** |  | [optional] 
**ServerInternalId** | Pointer to **string** |  | [optional] 
**PlanName** | Pointer to **string** |  | [optional] 
**HourlyPrice** | Pointer to **float32** |  | [optional] 
**HourlyCost** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**PricesUsed** | Pointer to [**[]BillingInstancesPricesUsed**](BillingInstancesPricesUsed.md) |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**CreatedByUser** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **int64** |  | [optional] 
**SiteId** | Pointer to **int64** |  | [optional] 
**SiteName** | Pointer to **string** |  | [optional] 
**SiteUUID** | Pointer to **string** |  | [optional] 
**SiteCode** | Pointer to **NullableString** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ApplicablePrices** | Pointer to [**[]BillingInstancesApplicablePrices**](BillingInstancesApplicablePrices.md) |  | [optional] 
**ServicePlanId** | Pointer to **int64** |  | [optional] 
**ServicePlanName** | Pointer to **string** |  | [optional] 
**ResourcePoolId** | Pointer to **int64** |  | [optional] 
**ResourcePoolName** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingInstancesUsages

`func NewBillingInstancesUsages() *BillingInstancesUsages`

NewBillingInstancesUsages instantiates a new BillingInstancesUsages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInstancesUsagesWithDefaults

`func NewBillingInstancesUsagesWithDefaults() *BillingInstancesUsages`

NewBillingInstancesUsagesWithDefaults instantiates a new BillingInstancesUsages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BillingInstancesUsages) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingInstancesUsages) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingInstancesUsages) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingInstancesUsages) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInstanceName

`func (o *BillingInstancesUsages) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *BillingInstancesUsages) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *BillingInstancesUsages) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *BillingInstancesUsages) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetZoneName

`func (o *BillingInstancesUsages) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *BillingInstancesUsages) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *BillingInstancesUsages) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *BillingInstancesUsages) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### GetAccountName

`func (o *BillingInstancesUsages) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *BillingInstancesUsages) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *BillingInstancesUsages) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *BillingInstancesUsages) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetVolumes

`func (o *BillingInstancesUsages) GetVolumes() []BillingInstancesVolumes`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *BillingInstancesUsages) GetVolumesOk() (*[]BillingInstancesVolumes, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *BillingInstancesUsages) SetVolumes(v []BillingInstancesVolumes)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *BillingInstancesUsages) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetMaxMemory

`func (o *BillingInstancesUsages) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *BillingInstancesUsages) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *BillingInstancesUsages) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *BillingInstancesUsages) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxCpu

`func (o *BillingInstancesUsages) GetMaxCpu() string`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *BillingInstancesUsages) GetMaxCpuOk() (*string, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *BillingInstancesUsages) SetMaxCpu(v string)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *BillingInstancesUsages) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### SetMaxCpuNil

`func (o *BillingInstancesUsages) SetMaxCpuNil(b bool)`

 SetMaxCpuNil sets the value for MaxCpu to be an explicit nil

### UnsetMaxCpu
`func (o *BillingInstancesUsages) UnsetMaxCpu()`

UnsetMaxCpu ensures that no value is present for MaxCpu, not even an explicit nil
### GetMaxCores

`func (o *BillingInstancesUsages) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *BillingInstancesUsages) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *BillingInstancesUsages) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *BillingInstancesUsages) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetServerExternalId

`func (o *BillingInstancesUsages) GetServerExternalId() string`

GetServerExternalId returns the ServerExternalId field if non-nil, zero value otherwise.

### GetServerExternalIdOk

`func (o *BillingInstancesUsages) GetServerExternalIdOk() (*string, bool)`

GetServerExternalIdOk returns a tuple with the ServerExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerExternalId

`func (o *BillingInstancesUsages) SetServerExternalId(v string)`

SetServerExternalId sets ServerExternalId field to given value.

### HasServerExternalId

`func (o *BillingInstancesUsages) HasServerExternalId() bool`

HasServerExternalId returns a boolean if a field has been set.

### GetServerInternalId

`func (o *BillingInstancesUsages) GetServerInternalId() string`

GetServerInternalId returns the ServerInternalId field if non-nil, zero value otherwise.

### GetServerInternalIdOk

`func (o *BillingInstancesUsages) GetServerInternalIdOk() (*string, bool)`

GetServerInternalIdOk returns a tuple with the ServerInternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInternalId

`func (o *BillingInstancesUsages) SetServerInternalId(v string)`

SetServerInternalId sets ServerInternalId field to given value.

### HasServerInternalId

`func (o *BillingInstancesUsages) HasServerInternalId() bool`

HasServerInternalId returns a boolean if a field has been set.

### GetPlanName

`func (o *BillingInstancesUsages) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *BillingInstancesUsages) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *BillingInstancesUsages) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *BillingInstancesUsages) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetHourlyPrice

`func (o *BillingInstancesUsages) GetHourlyPrice() float32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *BillingInstancesUsages) GetHourlyPriceOk() (*float32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *BillingInstancesUsages) SetHourlyPrice(v float32)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *BillingInstancesUsages) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.

### GetHourlyCost

`func (o *BillingInstancesUsages) GetHourlyCost() float32`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *BillingInstancesUsages) GetHourlyCostOk() (*float32, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *BillingInstancesUsages) SetHourlyCost(v float32)`

SetHourlyCost sets HourlyCost field to given value.

### HasHourlyCost

`func (o *BillingInstancesUsages) HasHourlyCost() bool`

HasHourlyCost returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingInstancesUsages) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingInstancesUsages) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingInstancesUsages) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingInstancesUsages) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPricesUsed

`func (o *BillingInstancesUsages) GetPricesUsed() []BillingInstancesPricesUsed`

GetPricesUsed returns the PricesUsed field if non-nil, zero value otherwise.

### GetPricesUsedOk

`func (o *BillingInstancesUsages) GetPricesUsedOk() (*[]BillingInstancesPricesUsed, bool)`

GetPricesUsedOk returns a tuple with the PricesUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricesUsed

`func (o *BillingInstancesUsages) SetPricesUsed(v []BillingInstancesPricesUsed)`

SetPricesUsed sets PricesUsed field to given value.

### HasPricesUsed

`func (o *BillingInstancesUsages) HasPricesUsed() bool`

HasPricesUsed returns a boolean if a field has been set.

### GetCost

`func (o *BillingInstancesUsages) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *BillingInstancesUsages) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *BillingInstancesUsages) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *BillingInstancesUsages) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPrice

`func (o *BillingInstancesUsages) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingInstancesUsages) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingInstancesUsages) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingInstancesUsages) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCreatedByUser

`func (o *BillingInstancesUsages) GetCreatedByUser() string`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *BillingInstancesUsages) GetCreatedByUserOk() (*string, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *BillingInstancesUsages) SetCreatedByUser(v string)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *BillingInstancesUsages) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *BillingInstancesUsages) GetCreatedByUserId() int64`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *BillingInstancesUsages) GetCreatedByUserIdOk() (*int64, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *BillingInstancesUsages) SetCreatedByUserId(v int64)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *BillingInstancesUsages) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetSiteId

`func (o *BillingInstancesUsages) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *BillingInstancesUsages) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *BillingInstancesUsages) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *BillingInstancesUsages) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSiteName

`func (o *BillingInstancesUsages) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *BillingInstancesUsages) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *BillingInstancesUsages) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *BillingInstancesUsages) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSiteUUID

`func (o *BillingInstancesUsages) GetSiteUUID() string`

GetSiteUUID returns the SiteUUID field if non-nil, zero value otherwise.

### GetSiteUUIDOk

`func (o *BillingInstancesUsages) GetSiteUUIDOk() (*string, bool)`

GetSiteUUIDOk returns a tuple with the SiteUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteUUID

`func (o *BillingInstancesUsages) SetSiteUUID(v string)`

SetSiteUUID sets SiteUUID field to given value.

### HasSiteUUID

`func (o *BillingInstancesUsages) HasSiteUUID() bool`

HasSiteUUID returns a boolean if a field has been set.

### GetSiteCode

`func (o *BillingInstancesUsages) GetSiteCode() string`

GetSiteCode returns the SiteCode field if non-nil, zero value otherwise.

### GetSiteCodeOk

`func (o *BillingInstancesUsages) GetSiteCodeOk() (*string, bool)`

GetSiteCodeOk returns a tuple with the SiteCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteCode

`func (o *BillingInstancesUsages) SetSiteCode(v string)`

SetSiteCode sets SiteCode field to given value.

### HasSiteCode

`func (o *BillingInstancesUsages) HasSiteCode() bool`

HasSiteCode returns a boolean if a field has been set.

### SetSiteCodeNil

`func (o *BillingInstancesUsages) SetSiteCodeNil(b bool)`

 SetSiteCodeNil sets the value for SiteCode to be an explicit nil

### UnsetSiteCode
`func (o *BillingInstancesUsages) UnsetSiteCode()`

UnsetSiteCode ensures that no value is present for SiteCode, not even an explicit nil
### GetStartDate

`func (o *BillingInstancesUsages) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BillingInstancesUsages) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BillingInstancesUsages) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BillingInstancesUsages) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *BillingInstancesUsages) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BillingInstancesUsages) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BillingInstancesUsages) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BillingInstancesUsages) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStatus

`func (o *BillingInstancesUsages) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingInstancesUsages) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingInstancesUsages) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BillingInstancesUsages) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *BillingInstancesUsages) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BillingInstancesUsages) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BillingInstancesUsages) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *BillingInstancesUsages) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *BillingInstancesUsages) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *BillingInstancesUsages) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetApplicablePrices

`func (o *BillingInstancesUsages) GetApplicablePrices() []BillingInstancesApplicablePrices`

GetApplicablePrices returns the ApplicablePrices field if non-nil, zero value otherwise.

### GetApplicablePricesOk

`func (o *BillingInstancesUsages) GetApplicablePricesOk() (*[]BillingInstancesApplicablePrices, bool)`

GetApplicablePricesOk returns a tuple with the ApplicablePrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicablePrices

`func (o *BillingInstancesUsages) SetApplicablePrices(v []BillingInstancesApplicablePrices)`

SetApplicablePrices sets ApplicablePrices field to given value.

### HasApplicablePrices

`func (o *BillingInstancesUsages) HasApplicablePrices() bool`

HasApplicablePrices returns a boolean if a field has been set.

### GetServicePlanId

`func (o *BillingInstancesUsages) GetServicePlanId() int64`

GetServicePlanId returns the ServicePlanId field if non-nil, zero value otherwise.

### GetServicePlanIdOk

`func (o *BillingInstancesUsages) GetServicePlanIdOk() (*int64, bool)`

GetServicePlanIdOk returns a tuple with the ServicePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanId

`func (o *BillingInstancesUsages) SetServicePlanId(v int64)`

SetServicePlanId sets ServicePlanId field to given value.

### HasServicePlanId

`func (o *BillingInstancesUsages) HasServicePlanId() bool`

HasServicePlanId returns a boolean if a field has been set.

### GetServicePlanName

`func (o *BillingInstancesUsages) GetServicePlanName() string`

GetServicePlanName returns the ServicePlanName field if non-nil, zero value otherwise.

### GetServicePlanNameOk

`func (o *BillingInstancesUsages) GetServicePlanNameOk() (*string, bool)`

GetServicePlanNameOk returns a tuple with the ServicePlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanName

`func (o *BillingInstancesUsages) SetServicePlanName(v string)`

SetServicePlanName sets ServicePlanName field to given value.

### HasServicePlanName

`func (o *BillingInstancesUsages) HasServicePlanName() bool`

HasServicePlanName returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *BillingInstancesUsages) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *BillingInstancesUsages) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *BillingInstancesUsages) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *BillingInstancesUsages) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetResourcePoolName

`func (o *BillingInstancesUsages) GetResourcePoolName() string`

GetResourcePoolName returns the ResourcePoolName field if non-nil, zero value otherwise.

### GetResourcePoolNameOk

`func (o *BillingInstancesUsages) GetResourcePoolNameOk() (*string, bool)`

GetResourcePoolNameOk returns a tuple with the ResourcePoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolName

`func (o *BillingInstancesUsages) SetResourcePoolName(v string)`

SetResourcePoolName sets ResourcePoolName field to given value.

### HasResourcePoolName

`func (o *BillingInstancesUsages) HasResourcePoolName() bool`

HasResourcePoolName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

