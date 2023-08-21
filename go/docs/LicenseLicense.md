# LicenseLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductTier** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**MaxInstances** | Pointer to **int64** |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**HardLimit** | Pointer to **bool** |  | [optional] 
**FreeTrial** | Pointer to **bool** |  | [optional] 
**MultiTenant** | Pointer to **bool** |  | [optional] 
**Whitelabel** | Pointer to **bool** |  | [optional] 
**ReportStatus** | Pointer to **bool** |  | [optional] 
**SupportLevel** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**AmazonProductCodes** | Pointer to **NullableString** |  | [optional] 
**Features** | Pointer to [**LicenseLicenseFeatures**](license_license_features.md) |  | [optional] 
**ZoneTypes** | Pointer to **NullableString** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewLicenseLicense

`func NewLicenseLicense() *LicenseLicense`

NewLicenseLicense instantiates a new LicenseLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseLicenseWithDefaults

`func NewLicenseLicenseWithDefaults() *LicenseLicense`

NewLicenseLicenseWithDefaults instantiates a new LicenseLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductTier

`func (o *LicenseLicense) GetProductTier() string`

GetProductTier returns the ProductTier field if non-nil, zero value otherwise.

### GetProductTierOk

`func (o *LicenseLicense) GetProductTierOk() (*string, bool)`

GetProductTierOk returns a tuple with the ProductTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTier

`func (o *LicenseLicense) SetProductTier(v string)`

SetProductTier sets ProductTier field to given value.

### HasProductTier

`func (o *LicenseLicense) HasProductTier() bool`

HasProductTier returns a boolean if a field has been set.

### GetStartDate

`func (o *LicenseLicense) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *LicenseLicense) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *LicenseLicense) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *LicenseLicense) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *LicenseLicense) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *LicenseLicense) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *LicenseLicense) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *LicenseLicense) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetMaxInstances

`func (o *LicenseLicense) GetMaxInstances() int64`

GetMaxInstances returns the MaxInstances field if non-nil, zero value otherwise.

### GetMaxInstancesOk

`func (o *LicenseLicense) GetMaxInstancesOk() (*int64, bool)`

GetMaxInstancesOk returns a tuple with the MaxInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInstances

`func (o *LicenseLicense) SetMaxInstances(v int64)`

SetMaxInstances sets MaxInstances field to given value.

### HasMaxInstances

`func (o *LicenseLicense) HasMaxInstances() bool`

HasMaxInstances returns a boolean if a field has been set.

### GetMaxMemory

`func (o *LicenseLicense) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *LicenseLicense) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *LicenseLicense) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *LicenseLicense) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *LicenseLicense) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *LicenseLicense) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *LicenseLicense) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *LicenseLicense) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetHardLimit

`func (o *LicenseLicense) GetHardLimit() bool`

GetHardLimit returns the HardLimit field if non-nil, zero value otherwise.

### GetHardLimitOk

`func (o *LicenseLicense) GetHardLimitOk() (*bool, bool)`

GetHardLimitOk returns a tuple with the HardLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardLimit

`func (o *LicenseLicense) SetHardLimit(v bool)`

SetHardLimit sets HardLimit field to given value.

### HasHardLimit

`func (o *LicenseLicense) HasHardLimit() bool`

HasHardLimit returns a boolean if a field has been set.

### GetFreeTrial

`func (o *LicenseLicense) GetFreeTrial() bool`

GetFreeTrial returns the FreeTrial field if non-nil, zero value otherwise.

### GetFreeTrialOk

`func (o *LicenseLicense) GetFreeTrialOk() (*bool, bool)`

GetFreeTrialOk returns a tuple with the FreeTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeTrial

`func (o *LicenseLicense) SetFreeTrial(v bool)`

SetFreeTrial sets FreeTrial field to given value.

### HasFreeTrial

`func (o *LicenseLicense) HasFreeTrial() bool`

HasFreeTrial returns a boolean if a field has been set.

### GetMultiTenant

`func (o *LicenseLicense) GetMultiTenant() bool`

GetMultiTenant returns the MultiTenant field if non-nil, zero value otherwise.

### GetMultiTenantOk

`func (o *LicenseLicense) GetMultiTenantOk() (*bool, bool)`

GetMultiTenantOk returns a tuple with the MultiTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiTenant

`func (o *LicenseLicense) SetMultiTenant(v bool)`

SetMultiTenant sets MultiTenant field to given value.

### HasMultiTenant

`func (o *LicenseLicense) HasMultiTenant() bool`

HasMultiTenant returns a boolean if a field has been set.

### GetWhitelabel

`func (o *LicenseLicense) GetWhitelabel() bool`

GetWhitelabel returns the Whitelabel field if non-nil, zero value otherwise.

### GetWhitelabelOk

`func (o *LicenseLicense) GetWhitelabelOk() (*bool, bool)`

GetWhitelabelOk returns a tuple with the Whitelabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelabel

`func (o *LicenseLicense) SetWhitelabel(v bool)`

SetWhitelabel sets Whitelabel field to given value.

### HasWhitelabel

`func (o *LicenseLicense) HasWhitelabel() bool`

HasWhitelabel returns a boolean if a field has been set.

### GetReportStatus

`func (o *LicenseLicense) GetReportStatus() bool`

GetReportStatus returns the ReportStatus field if non-nil, zero value otherwise.

### GetReportStatusOk

`func (o *LicenseLicense) GetReportStatusOk() (*bool, bool)`

GetReportStatusOk returns a tuple with the ReportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportStatus

`func (o *LicenseLicense) SetReportStatus(v bool)`

SetReportStatus sets ReportStatus field to given value.

### HasReportStatus

`func (o *LicenseLicense) HasReportStatus() bool`

HasReportStatus returns a boolean if a field has been set.

### GetSupportLevel

`func (o *LicenseLicense) GetSupportLevel() string`

GetSupportLevel returns the SupportLevel field if non-nil, zero value otherwise.

### GetSupportLevelOk

`func (o *LicenseLicense) GetSupportLevelOk() (*string, bool)`

GetSupportLevelOk returns a tuple with the SupportLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLevel

`func (o *LicenseLicense) SetSupportLevel(v string)`

SetSupportLevel sets SupportLevel field to given value.

### HasSupportLevel

`func (o *LicenseLicense) HasSupportLevel() bool`

HasSupportLevel returns a boolean if a field has been set.

### GetAccountName

`func (o *LicenseLicense) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *LicenseLicense) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *LicenseLicense) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *LicenseLicense) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetConfig

`func (o *LicenseLicense) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *LicenseLicense) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *LicenseLicense) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *LicenseLicense) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *LicenseLicense) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *LicenseLicense) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetAmazonProductCodes

`func (o *LicenseLicense) GetAmazonProductCodes() string`

GetAmazonProductCodes returns the AmazonProductCodes field if non-nil, zero value otherwise.

### GetAmazonProductCodesOk

`func (o *LicenseLicense) GetAmazonProductCodesOk() (*string, bool)`

GetAmazonProductCodesOk returns a tuple with the AmazonProductCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonProductCodes

`func (o *LicenseLicense) SetAmazonProductCodes(v string)`

SetAmazonProductCodes sets AmazonProductCodes field to given value.

### HasAmazonProductCodes

`func (o *LicenseLicense) HasAmazonProductCodes() bool`

HasAmazonProductCodes returns a boolean if a field has been set.

### SetAmazonProductCodesNil

`func (o *LicenseLicense) SetAmazonProductCodesNil(b bool)`

 SetAmazonProductCodesNil sets the value for AmazonProductCodes to be an explicit nil

### UnsetAmazonProductCodes
`func (o *LicenseLicense) UnsetAmazonProductCodes()`

UnsetAmazonProductCodes ensures that no value is present for AmazonProductCodes, not even an explicit nil
### GetFeatures

`func (o *LicenseLicense) GetFeatures() LicenseLicenseFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *LicenseLicense) GetFeaturesOk() (*LicenseLicenseFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *LicenseLicense) SetFeatures(v LicenseLicenseFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *LicenseLicense) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetZoneTypes

`func (o *LicenseLicense) GetZoneTypes() string`

GetZoneTypes returns the ZoneTypes field if non-nil, zero value otherwise.

### GetZoneTypesOk

`func (o *LicenseLicense) GetZoneTypesOk() (*string, bool)`

GetZoneTypesOk returns a tuple with the ZoneTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneTypes

`func (o *LicenseLicense) SetZoneTypes(v string)`

SetZoneTypes sets ZoneTypes field to given value.

### HasZoneTypes

`func (o *LicenseLicense) HasZoneTypes() bool`

HasZoneTypes returns a boolean if a field has been set.

### SetZoneTypesNil

`func (o *LicenseLicense) SetZoneTypesNil(b bool)`

 SetZoneTypesNil sets the value for ZoneTypes to be an explicit nil

### UnsetZoneTypes
`func (o *LicenseLicense) UnsetZoneTypes()`

UnsetZoneTypes ensures that no value is present for ZoneTypes, not even an explicit nil
### GetLastUpdated

`func (o *LicenseLicense) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *LicenseLicense) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *LicenseLicense) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *LicenseLicense) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDateCreated

`func (o *LicenseLicense) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *LicenseLicense) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *LicenseLicense) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *LicenseLicense) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


