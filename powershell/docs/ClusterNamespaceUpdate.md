# ClusterNamespaceUpdate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Namespace name | [optional] 
**Description** | **String** | Namespace description | [optional] 
**Active** | **Boolean** | Namespace active | [optional] [default to $false]
**Permissions** | [**ClusterNamespaceUpdatePermissions**](ClusterNamespaceUpdatePermissions.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterNamespaceUpdate = Initialize-PSOpenAPIToolsClusterNamespaceUpdate  -Name null `
 -Description null `
 -Active null `
 -Permissions null
```

- Convert the resource to JSON
```powershell
$ClusterNamespaceUpdate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

