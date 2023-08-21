# ClusterNamespaceUpdatePermissions
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePermissions** | [**ClusterNamespaceCreateResourcePermissions**](ClusterNamespaceCreateResourcePermissions.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterNamespaceUpdatePermissions = Initialize-PSOpenAPIToolsClusterNamespaceUpdatePermissions  -ResourcePermissions null
```

- Convert the resource to JSON
```powershell
$ClusterNamespaceUpdatePermissions | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

