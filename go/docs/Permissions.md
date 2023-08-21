# Permissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePool** | Pointer to [**ClusterPermissionsResourcePool**](cluster_permissions_resourcePool.md) |  | [optional] 
**ResourcePermissions** | Pointer to [**ResourcePermissions**](resourcePermissions.md) |  | [optional] 
**TenantPermissions** | Pointer to [**ApiZonesZoneIdFoldersIdFolderTenantPermissions**](_api_zones__zoneId__folders__id__folder_tenantPermissions.md) |  | [optional] 

## Methods

### NewPermissions

`func NewPermissions() *Permissions`

NewPermissions instantiates a new Permissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsWithDefaults

`func NewPermissionsWithDefaults() *Permissions`

NewPermissionsWithDefaults instantiates a new Permissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePool

`func (o *Permissions) GetResourcePool() ClusterPermissionsResourcePool`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *Permissions) GetResourcePoolOk() (*ClusterPermissionsResourcePool, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *Permissions) SetResourcePool(v ClusterPermissionsResourcePool)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *Permissions) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *Permissions) GetResourcePermissions() ResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *Permissions) GetResourcePermissionsOk() (*ResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *Permissions) SetResourcePermissions(v ResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *Permissions) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.

### GetTenantPermissions

`func (o *Permissions) GetTenantPermissions() ApiZonesZoneIdFoldersIdFolderTenantPermissions`

GetTenantPermissions returns the TenantPermissions field if non-nil, zero value otherwise.

### GetTenantPermissionsOk

`func (o *Permissions) GetTenantPermissionsOk() (*ApiZonesZoneIdFoldersIdFolderTenantPermissions, bool)`

GetTenantPermissionsOk returns a tuple with the TenantPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantPermissions

`func (o *Permissions) SetTenantPermissions(v ApiZonesZoneIdFoldersIdFolderTenantPermissions)`

SetTenantPermissions sets TenantPermissions field to given value.

### HasTenantPermissions

`func (o *Permissions) HasTenantPermissions() bool`

HasTenantPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


