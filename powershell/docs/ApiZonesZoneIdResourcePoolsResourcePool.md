# ApiZonesZoneIdResourcePoolsResourcePool
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Name of Resource Pool | 
**DefaultPool** | **Boolean** | Set as the Default Pool | [optional] [default to $false]
**DefaultImage** | **Boolean** | Set as the Default Image Target | [optional] [default to $false]
**Active** | **Boolean** | Activate &#x60;true&#x60; or disable &#x60;false&#x60; the datastore | [optional] [default to $true]
**Visibility** | **String** | Setting &#x60;private&#x60; or &#x60;public&#x60; | [optional] [default to "private"]
**DisplayName** | **String** | Optional Display Name (VMware only) | [optional] 
**Inventory** | **Boolean** | Enable &#x60;True&#x60; or disable &#x60;False&#x60; inventory sync for resource pool during cloud refresh | [optional] [default to $true]
**Config** | [**AnyOfAwsResourcePoolConfigCloudFoundryResourcePoolConfig**](AnyOfAwsResourcePoolConfigCloudFoundryResourcePoolConfig.md) |  | 
**TenantPermissions** | [**ApiZonesZoneIdFoldersIdFolderTenantPermissions[]**](ApiZonesZoneIdFoldersIdFolderTenantPermissions.md) |  | [optional] 
**ResourcePermissions** | [**ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions**](ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiZonesZoneIdResourcePoolsResourcePool = Initialize-PSOpenAPIToolsApiZonesZoneIdResourcePoolsResourcePool  -Name null `
 -DefaultPool null `
 -DefaultImage null `
 -Active null `
 -Visibility null `
 -DisplayName null `
 -Inventory null `
 -Config null `
 -TenantPermissions null `
 -ResourcePermissions null
```

- Convert the resource to JSON
```powershell
$ApiZonesZoneIdResourcePoolsResourcePool | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

