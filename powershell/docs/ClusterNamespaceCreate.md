# ClusterNamespaceCreate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Namespace name | 
**Description** | **String** | Namespace description | [optional] 
**Active** | **Boolean** | Namespace active | [optional] [default to $false]
**ResourcePermissions** | [**ClusterNamespaceCreateResourcePermissions**](ClusterNamespaceCreateResourcePermissions.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterNamespaceCreate = Initialize-PSOpenAPIToolsClusterNamespaceCreate  -Name null `
 -Description null `
 -Active null `
 -ResourcePermissions null
```

- Convert the resource to JSON
```powershell
$ClusterNamespaceCreate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

