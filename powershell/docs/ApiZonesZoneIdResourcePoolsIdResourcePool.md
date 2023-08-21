# ApiZonesZoneIdResourcePoolsIdResourcePool
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **Boolean** | Activate &#x60;true&#x60; or disable &#x60;false&#x60; the datastore | [optional] 
**Visibility** | **String** | Setting &#x60;private&#x60; or &#x60;public&#x60; | [optional] [default to "private"]
**DisplayName** | **String** | Optional Display Name (VMware only) | [optional] 
**Inventory** | **Boolean** | Enable &#x60;True&#x60; or disable &#x60;False&#x60; inventory sync for resource pool during cloud refresh | [optional] 
**TenantPermissions** | [**ApiZonesZoneIdFoldersIdFolderTenantPermissions[]**](ApiZonesZoneIdFoldersIdFolderTenantPermissions.md) |  | [optional] 
**ResourcePermissions** | [**ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions**](ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiZonesZoneIdResourcePoolsIdResourcePool = Initialize-PSOpenAPIToolsApiZonesZoneIdResourcePoolsIdResourcePool  -Active null `
 -Visibility null `
 -DisplayName null `
 -Inventory null `
 -TenantPermissions null `
 -ResourcePermissions null
```

- Convert the resource to JSON
```powershell
$ApiZonesZoneIdResourcePoolsIdResourcePool | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

