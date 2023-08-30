# AddClusterNamespace200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | [**ClusterNamespaceCreateSuccess**](ClusterNamespaceCreateSuccess.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddClusterNamespace200Response = Initialize-PSOpenAPIToolsAddClusterNamespace200Response  -Namespace null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddClusterNamespace200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

