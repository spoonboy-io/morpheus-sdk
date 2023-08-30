# GetClusterNamespaces200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespaces** | [**ClusterNamespaces[]**](ClusterNamespaces.md) |  | [optional] 
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetClusterNamespaces200Response = Initialize-PSOpenAPIToolsGetClusterNamespaces200Response  -Namespaces null `
 -Meta null
```

- Convert the resource to JSON
```powershell
$GetClusterNamespaces200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

