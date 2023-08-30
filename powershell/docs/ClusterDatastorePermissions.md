# ClusterDatastorePermissions
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePermissions** | [**ClusterDatastorePermissionsResourcePermissions**](ClusterDatastorePermissionsResourcePermissions.md) |  | [optional] 
**TenantPermissions** | [**GetNetworkRouter200ResponseNetworkRouterPermissionsTenantPermissions**](GetNetworkRouter200ResponseNetworkRouterPermissionsTenantPermissions.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterDatastorePermissions = Initialize-PSOpenAPIToolsClusterDatastorePermissions  -ResourcePermissions null `
 -TenantPermissions null
```

- Convert the resource to JSON
```powershell
$ClusterDatastorePermissions | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

