# ApiIntegrationsIdInventoryInventoryIdInventory
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tenants** | [**ApiBlueprintsIdUpdatePermissionsResourcePermissionSites[]**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of tenant accounts that will use this inventory as Default. Used by jobs set to &#39;Use Tenant Default&#39; | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiIntegrationsIdInventoryInventoryIdInventory = Initialize-PSOpenAPIToolsApiIntegrationsIdInventoryInventoryIdInventory  -Tenants null
```

- Convert the resource to JSON
```powershell
$ApiIntegrationsIdInventoryInventoryIdInventory | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

