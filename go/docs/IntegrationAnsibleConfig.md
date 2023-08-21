# IntegrationAnsibleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inventory** | Pointer to **NullableString** |  | [optional] 
**DefaultBranch** | Pointer to **string** |  | [optional] 
**CacheEnabled** | Pointer to **NullableString** |  | [optional] 
**AnsiblePlaybooks** | Pointer to **string** |  | [optional] 
**AnsibleRoles** | Pointer to **string** |  | [optional] 
**AnsibleGroupVars** | Pointer to **string** |  | [optional] 
**AnsibleHostVars** | Pointer to **string** |  | [optional] 
**AnsibleCommandBus** | Pointer to **bool** |  | [optional] 
**AnsibleVerbose** | Pointer to **bool** |  | [optional] 
**AnsibleGalaxyEnabled** | Pointer to **bool** |  | [optional] 
**AnsibleDefaultBranch** | Pointer to **string** |  | [optional] 

## Methods

### NewIntegrationAnsibleConfig

`func NewIntegrationAnsibleConfig() *IntegrationAnsibleConfig`

NewIntegrationAnsibleConfig instantiates a new IntegrationAnsibleConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationAnsibleConfigWithDefaults

`func NewIntegrationAnsibleConfigWithDefaults() *IntegrationAnsibleConfig`

NewIntegrationAnsibleConfigWithDefaults instantiates a new IntegrationAnsibleConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInventory

`func (o *IntegrationAnsibleConfig) GetInventory() string`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *IntegrationAnsibleConfig) GetInventoryOk() (*string, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *IntegrationAnsibleConfig) SetInventory(v string)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *IntegrationAnsibleConfig) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### SetInventoryNil

`func (o *IntegrationAnsibleConfig) SetInventoryNil(b bool)`

 SetInventoryNil sets the value for Inventory to be an explicit nil

### UnsetInventory
`func (o *IntegrationAnsibleConfig) UnsetInventory()`

UnsetInventory ensures that no value is present for Inventory, not even an explicit nil
### GetDefaultBranch

`func (o *IntegrationAnsibleConfig) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *IntegrationAnsibleConfig) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *IntegrationAnsibleConfig) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *IntegrationAnsibleConfig) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetCacheEnabled

`func (o *IntegrationAnsibleConfig) GetCacheEnabled() string`

GetCacheEnabled returns the CacheEnabled field if non-nil, zero value otherwise.

### GetCacheEnabledOk

`func (o *IntegrationAnsibleConfig) GetCacheEnabledOk() (*string, bool)`

GetCacheEnabledOk returns a tuple with the CacheEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheEnabled

`func (o *IntegrationAnsibleConfig) SetCacheEnabled(v string)`

SetCacheEnabled sets CacheEnabled field to given value.

### HasCacheEnabled

`func (o *IntegrationAnsibleConfig) HasCacheEnabled() bool`

HasCacheEnabled returns a boolean if a field has been set.

### SetCacheEnabledNil

`func (o *IntegrationAnsibleConfig) SetCacheEnabledNil(b bool)`

 SetCacheEnabledNil sets the value for CacheEnabled to be an explicit nil

### UnsetCacheEnabled
`func (o *IntegrationAnsibleConfig) UnsetCacheEnabled()`

UnsetCacheEnabled ensures that no value is present for CacheEnabled, not even an explicit nil
### GetAnsiblePlaybooks

`func (o *IntegrationAnsibleConfig) GetAnsiblePlaybooks() string`

GetAnsiblePlaybooks returns the AnsiblePlaybooks field if non-nil, zero value otherwise.

### GetAnsiblePlaybooksOk

`func (o *IntegrationAnsibleConfig) GetAnsiblePlaybooksOk() (*string, bool)`

GetAnsiblePlaybooksOk returns a tuple with the AnsiblePlaybooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsiblePlaybooks

`func (o *IntegrationAnsibleConfig) SetAnsiblePlaybooks(v string)`

SetAnsiblePlaybooks sets AnsiblePlaybooks field to given value.

### HasAnsiblePlaybooks

`func (o *IntegrationAnsibleConfig) HasAnsiblePlaybooks() bool`

HasAnsiblePlaybooks returns a boolean if a field has been set.

### GetAnsibleRoles

`func (o *IntegrationAnsibleConfig) GetAnsibleRoles() string`

GetAnsibleRoles returns the AnsibleRoles field if non-nil, zero value otherwise.

### GetAnsibleRolesOk

`func (o *IntegrationAnsibleConfig) GetAnsibleRolesOk() (*string, bool)`

GetAnsibleRolesOk returns a tuple with the AnsibleRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleRoles

`func (o *IntegrationAnsibleConfig) SetAnsibleRoles(v string)`

SetAnsibleRoles sets AnsibleRoles field to given value.

### HasAnsibleRoles

`func (o *IntegrationAnsibleConfig) HasAnsibleRoles() bool`

HasAnsibleRoles returns a boolean if a field has been set.

### GetAnsibleGroupVars

`func (o *IntegrationAnsibleConfig) GetAnsibleGroupVars() string`

GetAnsibleGroupVars returns the AnsibleGroupVars field if non-nil, zero value otherwise.

### GetAnsibleGroupVarsOk

`func (o *IntegrationAnsibleConfig) GetAnsibleGroupVarsOk() (*string, bool)`

GetAnsibleGroupVarsOk returns a tuple with the AnsibleGroupVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleGroupVars

`func (o *IntegrationAnsibleConfig) SetAnsibleGroupVars(v string)`

SetAnsibleGroupVars sets AnsibleGroupVars field to given value.

### HasAnsibleGroupVars

`func (o *IntegrationAnsibleConfig) HasAnsibleGroupVars() bool`

HasAnsibleGroupVars returns a boolean if a field has been set.

### GetAnsibleHostVars

`func (o *IntegrationAnsibleConfig) GetAnsibleHostVars() string`

GetAnsibleHostVars returns the AnsibleHostVars field if non-nil, zero value otherwise.

### GetAnsibleHostVarsOk

`func (o *IntegrationAnsibleConfig) GetAnsibleHostVarsOk() (*string, bool)`

GetAnsibleHostVarsOk returns a tuple with the AnsibleHostVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleHostVars

`func (o *IntegrationAnsibleConfig) SetAnsibleHostVars(v string)`

SetAnsibleHostVars sets AnsibleHostVars field to given value.

### HasAnsibleHostVars

`func (o *IntegrationAnsibleConfig) HasAnsibleHostVars() bool`

HasAnsibleHostVars returns a boolean if a field has been set.

### GetAnsibleCommandBus

`func (o *IntegrationAnsibleConfig) GetAnsibleCommandBus() bool`

GetAnsibleCommandBus returns the AnsibleCommandBus field if non-nil, zero value otherwise.

### GetAnsibleCommandBusOk

`func (o *IntegrationAnsibleConfig) GetAnsibleCommandBusOk() (*bool, bool)`

GetAnsibleCommandBusOk returns a tuple with the AnsibleCommandBus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleCommandBus

`func (o *IntegrationAnsibleConfig) SetAnsibleCommandBus(v bool)`

SetAnsibleCommandBus sets AnsibleCommandBus field to given value.

### HasAnsibleCommandBus

`func (o *IntegrationAnsibleConfig) HasAnsibleCommandBus() bool`

HasAnsibleCommandBus returns a boolean if a field has been set.

### GetAnsibleVerbose

`func (o *IntegrationAnsibleConfig) GetAnsibleVerbose() bool`

GetAnsibleVerbose returns the AnsibleVerbose field if non-nil, zero value otherwise.

### GetAnsibleVerboseOk

`func (o *IntegrationAnsibleConfig) GetAnsibleVerboseOk() (*bool, bool)`

GetAnsibleVerboseOk returns a tuple with the AnsibleVerbose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleVerbose

`func (o *IntegrationAnsibleConfig) SetAnsibleVerbose(v bool)`

SetAnsibleVerbose sets AnsibleVerbose field to given value.

### HasAnsibleVerbose

`func (o *IntegrationAnsibleConfig) HasAnsibleVerbose() bool`

HasAnsibleVerbose returns a boolean if a field has been set.

### GetAnsibleGalaxyEnabled

`func (o *IntegrationAnsibleConfig) GetAnsibleGalaxyEnabled() bool`

GetAnsibleGalaxyEnabled returns the AnsibleGalaxyEnabled field if non-nil, zero value otherwise.

### GetAnsibleGalaxyEnabledOk

`func (o *IntegrationAnsibleConfig) GetAnsibleGalaxyEnabledOk() (*bool, bool)`

GetAnsibleGalaxyEnabledOk returns a tuple with the AnsibleGalaxyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleGalaxyEnabled

`func (o *IntegrationAnsibleConfig) SetAnsibleGalaxyEnabled(v bool)`

SetAnsibleGalaxyEnabled sets AnsibleGalaxyEnabled field to given value.

### HasAnsibleGalaxyEnabled

`func (o *IntegrationAnsibleConfig) HasAnsibleGalaxyEnabled() bool`

HasAnsibleGalaxyEnabled returns a boolean if a field has been set.

### GetAnsibleDefaultBranch

`func (o *IntegrationAnsibleConfig) GetAnsibleDefaultBranch() string`

GetAnsibleDefaultBranch returns the AnsibleDefaultBranch field if non-nil, zero value otherwise.

### GetAnsibleDefaultBranchOk

`func (o *IntegrationAnsibleConfig) GetAnsibleDefaultBranchOk() (*string, bool)`

GetAnsibleDefaultBranchOk returns a tuple with the AnsibleDefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleDefaultBranch

`func (o *IntegrationAnsibleConfig) SetAnsibleDefaultBranch(v string)`

SetAnsibleDefaultBranch sets AnsibleDefaultBranch field to given value.

### HasAnsibleDefaultBranch

`func (o *IntegrationAnsibleConfig) HasAnsibleDefaultBranch() bool`

HasAnsibleDefaultBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


