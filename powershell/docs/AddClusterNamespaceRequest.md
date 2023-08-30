# AddClusterNamespaceRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | [**ClusterNamespaceCreate**](ClusterNamespaceCreate.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddClusterNamespaceRequest = Initialize-PSOpenAPIToolsAddClusterNamespaceRequest  -Namespace null
```

- Convert the resource to JSON
```powershell
$AddClusterNamespaceRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

