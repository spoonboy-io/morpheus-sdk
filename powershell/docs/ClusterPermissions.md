# ClusterPermissions
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePool** | [**ClusterPermissionsResourcePool**](ClusterPermissionsResourcePool.md) |  | [optional] 
**ResourcePermissions** | [**ClusterPermissionsResourcePermissions**](ClusterPermissionsResourcePermissions.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterPermissions = Initialize-PSOpenAPIToolsClusterPermissions  -ResourcePool null `
 -ResourcePermissions null
```

- Convert the resource to JSON
```powershell
$ClusterPermissions | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

