# ClusterUpdatePermissions
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePermissions** | [**ClusterUpdatePermissionsResourcePermissions**](ClusterUpdatePermissionsResourcePermissions.md) |  | [optional] 
**ResourcePool** | [**ClusterUpdatePermissionsResourcePool**](ClusterUpdatePermissionsResourcePool.md) |  | [optional] 
**TenantPermissions** | [**ClusterUpdatePermissionsTenantPermissions**](ClusterUpdatePermissionsTenantPermissions.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterUpdatePermissions = Initialize-PSOpenAPIToolsClusterUpdatePermissions  -ResourcePermissions null `
 -ResourcePool null `
 -TenantPermissions null
```

- Convert the resource to JSON
```powershell
$ClusterUpdatePermissions | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

