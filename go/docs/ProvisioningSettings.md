# ProvisioningSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowZoneSelection** | Pointer to **bool** |  | [optional] 
**AllowServerSelection** | Pointer to **bool** |  | [optional] 
**RequireEnvironments** | Pointer to **bool** |  | [optional] 
**ShowPricing** | Pointer to **bool** |  | [optional] 
**HideDatastoreStats** | Pointer to **bool** |  | [optional] 
**CrossTenantNamingPolicies** | Pointer to **bool** |  | [optional] 
**ReuseSequence** | Pointer to **bool** |  | [optional] 
**CloudInitUsername** | Pointer to **string** |  | [optional] 
**CloudInitPassword** | Pointer to **string** |  | [optional] 
**CloudInitKeyPair** | Pointer to [**InlineResponse20040AppDeployInstance**](inline_response_200_40_appDeploy_instance.md) |  | [optional] 
**WindowsPassword** | Pointer to **NullableString** |  | [optional] 
**PxeRootPassword** | Pointer to **NullableString** |  | [optional] 
**DefaultTemplateType** | Pointer to [**InlineResponse20094Network**](inline_response_200_94_network.md) |  | [optional] 
**DeployStorageProvider** | Pointer to [**InlineResponse20040AppDeployInstance**](inline_response_200_40_appDeploy_instance.md) |  | [optional] 

## Methods

### NewProvisioningSettings

`func NewProvisioningSettings() *ProvisioningSettings`

NewProvisioningSettings instantiates a new ProvisioningSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningSettingsWithDefaults

`func NewProvisioningSettingsWithDefaults() *ProvisioningSettings`

NewProvisioningSettingsWithDefaults instantiates a new ProvisioningSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowZoneSelection

`func (o *ProvisioningSettings) GetAllowZoneSelection() bool`

GetAllowZoneSelection returns the AllowZoneSelection field if non-nil, zero value otherwise.

### GetAllowZoneSelectionOk

`func (o *ProvisioningSettings) GetAllowZoneSelectionOk() (*bool, bool)`

GetAllowZoneSelectionOk returns a tuple with the AllowZoneSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowZoneSelection

`func (o *ProvisioningSettings) SetAllowZoneSelection(v bool)`

SetAllowZoneSelection sets AllowZoneSelection field to given value.

### HasAllowZoneSelection

`func (o *ProvisioningSettings) HasAllowZoneSelection() bool`

HasAllowZoneSelection returns a boolean if a field has been set.

### GetAllowServerSelection

`func (o *ProvisioningSettings) GetAllowServerSelection() bool`

GetAllowServerSelection returns the AllowServerSelection field if non-nil, zero value otherwise.

### GetAllowServerSelectionOk

`func (o *ProvisioningSettings) GetAllowServerSelectionOk() (*bool, bool)`

GetAllowServerSelectionOk returns a tuple with the AllowServerSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowServerSelection

`func (o *ProvisioningSettings) SetAllowServerSelection(v bool)`

SetAllowServerSelection sets AllowServerSelection field to given value.

### HasAllowServerSelection

`func (o *ProvisioningSettings) HasAllowServerSelection() bool`

HasAllowServerSelection returns a boolean if a field has been set.

### GetRequireEnvironments

`func (o *ProvisioningSettings) GetRequireEnvironments() bool`

GetRequireEnvironments returns the RequireEnvironments field if non-nil, zero value otherwise.

### GetRequireEnvironmentsOk

`func (o *ProvisioningSettings) GetRequireEnvironmentsOk() (*bool, bool)`

GetRequireEnvironmentsOk returns a tuple with the RequireEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireEnvironments

`func (o *ProvisioningSettings) SetRequireEnvironments(v bool)`

SetRequireEnvironments sets RequireEnvironments field to given value.

### HasRequireEnvironments

`func (o *ProvisioningSettings) HasRequireEnvironments() bool`

HasRequireEnvironments returns a boolean if a field has been set.

### GetShowPricing

`func (o *ProvisioningSettings) GetShowPricing() bool`

GetShowPricing returns the ShowPricing field if non-nil, zero value otherwise.

### GetShowPricingOk

`func (o *ProvisioningSettings) GetShowPricingOk() (*bool, bool)`

GetShowPricingOk returns a tuple with the ShowPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPricing

`func (o *ProvisioningSettings) SetShowPricing(v bool)`

SetShowPricing sets ShowPricing field to given value.

### HasShowPricing

`func (o *ProvisioningSettings) HasShowPricing() bool`

HasShowPricing returns a boolean if a field has been set.

### GetHideDatastoreStats

`func (o *ProvisioningSettings) GetHideDatastoreStats() bool`

GetHideDatastoreStats returns the HideDatastoreStats field if non-nil, zero value otherwise.

### GetHideDatastoreStatsOk

`func (o *ProvisioningSettings) GetHideDatastoreStatsOk() (*bool, bool)`

GetHideDatastoreStatsOk returns a tuple with the HideDatastoreStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideDatastoreStats

`func (o *ProvisioningSettings) SetHideDatastoreStats(v bool)`

SetHideDatastoreStats sets HideDatastoreStats field to given value.

### HasHideDatastoreStats

`func (o *ProvisioningSettings) HasHideDatastoreStats() bool`

HasHideDatastoreStats returns a boolean if a field has been set.

### GetCrossTenantNamingPolicies

`func (o *ProvisioningSettings) GetCrossTenantNamingPolicies() bool`

GetCrossTenantNamingPolicies returns the CrossTenantNamingPolicies field if non-nil, zero value otherwise.

### GetCrossTenantNamingPoliciesOk

`func (o *ProvisioningSettings) GetCrossTenantNamingPoliciesOk() (*bool, bool)`

GetCrossTenantNamingPoliciesOk returns a tuple with the CrossTenantNamingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossTenantNamingPolicies

`func (o *ProvisioningSettings) SetCrossTenantNamingPolicies(v bool)`

SetCrossTenantNamingPolicies sets CrossTenantNamingPolicies field to given value.

### HasCrossTenantNamingPolicies

`func (o *ProvisioningSettings) HasCrossTenantNamingPolicies() bool`

HasCrossTenantNamingPolicies returns a boolean if a field has been set.

### GetReuseSequence

`func (o *ProvisioningSettings) GetReuseSequence() bool`

GetReuseSequence returns the ReuseSequence field if non-nil, zero value otherwise.

### GetReuseSequenceOk

`func (o *ProvisioningSettings) GetReuseSequenceOk() (*bool, bool)`

GetReuseSequenceOk returns a tuple with the ReuseSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReuseSequence

`func (o *ProvisioningSettings) SetReuseSequence(v bool)`

SetReuseSequence sets ReuseSequence field to given value.

### HasReuseSequence

`func (o *ProvisioningSettings) HasReuseSequence() bool`

HasReuseSequence returns a boolean if a field has been set.

### GetCloudInitUsername

`func (o *ProvisioningSettings) GetCloudInitUsername() string`

GetCloudInitUsername returns the CloudInitUsername field if non-nil, zero value otherwise.

### GetCloudInitUsernameOk

`func (o *ProvisioningSettings) GetCloudInitUsernameOk() (*string, bool)`

GetCloudInitUsernameOk returns a tuple with the CloudInitUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitUsername

`func (o *ProvisioningSettings) SetCloudInitUsername(v string)`

SetCloudInitUsername sets CloudInitUsername field to given value.

### HasCloudInitUsername

`func (o *ProvisioningSettings) HasCloudInitUsername() bool`

HasCloudInitUsername returns a boolean if a field has been set.

### GetCloudInitPassword

`func (o *ProvisioningSettings) GetCloudInitPassword() string`

GetCloudInitPassword returns the CloudInitPassword field if non-nil, zero value otherwise.

### GetCloudInitPasswordOk

`func (o *ProvisioningSettings) GetCloudInitPasswordOk() (*string, bool)`

GetCloudInitPasswordOk returns a tuple with the CloudInitPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitPassword

`func (o *ProvisioningSettings) SetCloudInitPassword(v string)`

SetCloudInitPassword sets CloudInitPassword field to given value.

### HasCloudInitPassword

`func (o *ProvisioningSettings) HasCloudInitPassword() bool`

HasCloudInitPassword returns a boolean if a field has been set.

### GetCloudInitKeyPair

`func (o *ProvisioningSettings) GetCloudInitKeyPair() InlineResponse20040AppDeployInstance`

GetCloudInitKeyPair returns the CloudInitKeyPair field if non-nil, zero value otherwise.

### GetCloudInitKeyPairOk

`func (o *ProvisioningSettings) GetCloudInitKeyPairOk() (*InlineResponse20040AppDeployInstance, bool)`

GetCloudInitKeyPairOk returns a tuple with the CloudInitKeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitKeyPair

`func (o *ProvisioningSettings) SetCloudInitKeyPair(v InlineResponse20040AppDeployInstance)`

SetCloudInitKeyPair sets CloudInitKeyPair field to given value.

### HasCloudInitKeyPair

`func (o *ProvisioningSettings) HasCloudInitKeyPair() bool`

HasCloudInitKeyPair returns a boolean if a field has been set.

### GetWindowsPassword

`func (o *ProvisioningSettings) GetWindowsPassword() string`

GetWindowsPassword returns the WindowsPassword field if non-nil, zero value otherwise.

### GetWindowsPasswordOk

`func (o *ProvisioningSettings) GetWindowsPasswordOk() (*string, bool)`

GetWindowsPasswordOk returns a tuple with the WindowsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPassword

`func (o *ProvisioningSettings) SetWindowsPassword(v string)`

SetWindowsPassword sets WindowsPassword field to given value.

### HasWindowsPassword

`func (o *ProvisioningSettings) HasWindowsPassword() bool`

HasWindowsPassword returns a boolean if a field has been set.

### SetWindowsPasswordNil

`func (o *ProvisioningSettings) SetWindowsPasswordNil(b bool)`

 SetWindowsPasswordNil sets the value for WindowsPassword to be an explicit nil

### UnsetWindowsPassword
`func (o *ProvisioningSettings) UnsetWindowsPassword()`

UnsetWindowsPassword ensures that no value is present for WindowsPassword, not even an explicit nil
### GetPxeRootPassword

`func (o *ProvisioningSettings) GetPxeRootPassword() string`

GetPxeRootPassword returns the PxeRootPassword field if non-nil, zero value otherwise.

### GetPxeRootPasswordOk

`func (o *ProvisioningSettings) GetPxeRootPasswordOk() (*string, bool)`

GetPxeRootPasswordOk returns a tuple with the PxeRootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxeRootPassword

`func (o *ProvisioningSettings) SetPxeRootPassword(v string)`

SetPxeRootPassword sets PxeRootPassword field to given value.

### HasPxeRootPassword

`func (o *ProvisioningSettings) HasPxeRootPassword() bool`

HasPxeRootPassword returns a boolean if a field has been set.

### SetPxeRootPasswordNil

`func (o *ProvisioningSettings) SetPxeRootPasswordNil(b bool)`

 SetPxeRootPasswordNil sets the value for PxeRootPassword to be an explicit nil

### UnsetPxeRootPassword
`func (o *ProvisioningSettings) UnsetPxeRootPassword()`

UnsetPxeRootPassword ensures that no value is present for PxeRootPassword, not even an explicit nil
### GetDefaultTemplateType

`func (o *ProvisioningSettings) GetDefaultTemplateType() InlineResponse20094Network`

GetDefaultTemplateType returns the DefaultTemplateType field if non-nil, zero value otherwise.

### GetDefaultTemplateTypeOk

`func (o *ProvisioningSettings) GetDefaultTemplateTypeOk() (*InlineResponse20094Network, bool)`

GetDefaultTemplateTypeOk returns a tuple with the DefaultTemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTemplateType

`func (o *ProvisioningSettings) SetDefaultTemplateType(v InlineResponse20094Network)`

SetDefaultTemplateType sets DefaultTemplateType field to given value.

### HasDefaultTemplateType

`func (o *ProvisioningSettings) HasDefaultTemplateType() bool`

HasDefaultTemplateType returns a boolean if a field has been set.

### GetDeployStorageProvider

`func (o *ProvisioningSettings) GetDeployStorageProvider() InlineResponse20040AppDeployInstance`

GetDeployStorageProvider returns the DeployStorageProvider field if non-nil, zero value otherwise.

### GetDeployStorageProviderOk

`func (o *ProvisioningSettings) GetDeployStorageProviderOk() (*InlineResponse20040AppDeployInstance, bool)`

GetDeployStorageProviderOk returns a tuple with the DeployStorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStorageProvider

`func (o *ProvisioningSettings) SetDeployStorageProvider(v InlineResponse20040AppDeployInstance)`

SetDeployStorageProvider sets DeployStorageProvider field to given value.

### HasDeployStorageProvider

`func (o *ProvisioningSettings) HasDeployStorageProvider() bool`

HasDeployStorageProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


