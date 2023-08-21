# ApiZonesZoneIdResourcePoolsResourcePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of Resource Pool | 
**DefaultPool** | Pointer to **bool** | Set as the Default Pool | [optional] [default to false]
**DefaultImage** | Pointer to **bool** | Set as the Default Image Target | [optional] [default to false]
**Active** | Pointer to **bool** | Activate &#x60;true&#x60; or disable &#x60;false&#x60; the datastore | [optional] [default to true]
**Visibility** | Pointer to **string** | Setting &#x60;private&#x60; or &#x60;public&#x60; | [optional] [default to "private"]
**DisplayName** | Pointer to **string** | Optional Display Name (VMware only) | [optional] 
**Inventory** | Pointer to **bool** | Enable &#x60;True&#x60; or disable &#x60;False&#x60; inventory sync for resource pool during cloud refresh | [optional] [default to true]
**Config** | [**AnyOfAwsResourcePoolConfigCloudFoundryResourcePoolConfig**](anyOf&lt;AwsResourcePoolConfig,CloudFoundryResourcePoolConfig&gt;.md) |  | 
**TenantPermissions** | Pointer to [**[]ApiZonesZoneIdFoldersIdFolderTenantPermissions**](ApiZonesZoneIdFoldersIdFolderTenantPermissions.md) |  | [optional] 
**ResourcePermissions** | Pointer to [**ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions**](_api_zones__zoneId__data_stores__id__datastore_resourcePermissions.md) |  | [optional] 

## Methods

### NewApiZonesZoneIdResourcePoolsResourcePool

`func NewApiZonesZoneIdResourcePoolsResourcePool(name string, config AnyOfAwsResourcePoolConfigCloudFoundryResourcePoolConfig, ) *ApiZonesZoneIdResourcePoolsResourcePool`

NewApiZonesZoneIdResourcePoolsResourcePool instantiates a new ApiZonesZoneIdResourcePoolsResourcePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiZonesZoneIdResourcePoolsResourcePoolWithDefaults

`func NewApiZonesZoneIdResourcePoolsResourcePoolWithDefaults() *ApiZonesZoneIdResourcePoolsResourcePool`

NewApiZonesZoneIdResourcePoolsResourcePoolWithDefaults instantiates a new ApiZonesZoneIdResourcePoolsResourcePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) SetName(v string)`

SetName sets Name field to given value.


### GetDefaultPool

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) GetDefaultPool() bool`

GetDefaultPool returns the DefaultPool field if non-nil, zero value otherwise.

### GetDefaultPoolOk

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) GetDefaultPoolOk() (*bool, bool)`

GetDefaultPoolOk returns a tuple with the DefaultPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPool

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) SetDefaultPool(v bool)`

SetDefaultPool sets DefaultPool field to given value.

### HasDefaultPool

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) HasDefaultPool() bool`

HasDefaultPool returns a boolean if a field has been set.

### GetDefaultImage

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) GetDefaultImage() bool`

GetDefaultImage returns the DefaultImage field if non-nil, zero value otherwise.

### GetDefaultImageOk

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) GetDefaultImageOk() (*bool, bool)`

GetDefaultImageOk returns a tuple with the DefaultImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImage

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) SetDefaultImage(v bool)`

SetDefaultImage sets DefaultImage field to given value.

### HasDefaultImage

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) HasDefaultImage() bool`

HasDefaultImage returns a boolean if a field has been set.

### GetActive

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetVisibility

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDisplayName

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetInventory

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) GetInventory() bool`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) GetInventoryOk() (*bool, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) SetInventory(v bool)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetConfig

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) GetConfig() AnyOfAwsResourcePoolConfigCloudFoundryResourcePoolConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) GetConfigOk() (*AnyOfAwsResourcePoolConfigCloudFoundryResourcePoolConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) SetConfig(v AnyOfAwsResourcePoolConfigCloudFoundryResourcePoolConfig)`

SetConfig sets Config field to given value.


### GetTenantPermissions

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) GetTenantPermissions() []ApiZonesZoneIdFoldersIdFolderTenantPermissions`

GetTenantPermissions returns the TenantPermissions field if non-nil, zero value otherwise.

### GetTenantPermissionsOk

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) GetTenantPermissionsOk() (*[]ApiZonesZoneIdFoldersIdFolderTenantPermissions, bool)`

GetTenantPermissionsOk returns a tuple with the TenantPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantPermissions

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) SetTenantPermissions(v []ApiZonesZoneIdFoldersIdFolderTenantPermissions)`

SetTenantPermissions sets TenantPermissions field to given value.

### HasTenantPermissions

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) HasTenantPermissions() bool`

HasTenantPermissions returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) GetResourcePermissions() ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) GetResourcePermissionsOk() (*ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) SetResourcePermissions(v ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *ApiZonesZoneIdResourcePoolsResourcePool) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


