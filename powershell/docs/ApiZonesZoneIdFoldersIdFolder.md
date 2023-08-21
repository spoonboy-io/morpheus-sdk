# ApiZonesZoneIdFoldersIdFolder
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultFolder** | **Boolean** |  | [optional] [default to $false]
**DefaultImage** | **Boolean** |  | [optional] [default to $false]
**Active** | **Boolean** | Activate &#x60;true&#x60; or disable &#x60;false&#x60; the folder | [optional] 
**Visibility** | **String** | Setting &#x60;private&#x60; or &#x60;public&#x60; | [optional] [default to "private"]
**TenantPermissions** | [**ApiZonesZoneIdFoldersIdFolderTenantPermissions[]**](ApiZonesZoneIdFoldersIdFolderTenantPermissions.md) |  | [optional] 
**ResourcePermissions** | [**ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions**](ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiZonesZoneIdFoldersIdFolder = Initialize-PSOpenAPIToolsApiZonesZoneIdFoldersIdFolder  -DefaultFolder null `
 -DefaultImage null `
 -Active null `
 -Visibility null `
 -TenantPermissions null `
 -ResourcePermissions null
```

- Convert the resource to JSON
```powershell
$ApiZonesZoneIdFoldersIdFolder | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

