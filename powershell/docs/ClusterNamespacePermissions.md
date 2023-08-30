# ClusterNamespacePermissions
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePermissions** | [**ClusterNamespacePermissionsResourcePermissions**](ClusterNamespacePermissionsResourcePermissions.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterNamespacePermissions = Initialize-PSOpenAPIToolsClusterNamespacePermissions  -ResourcePermissions null
```

- Convert the resource to JSON
```powershell
$ClusterNamespacePermissions | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

