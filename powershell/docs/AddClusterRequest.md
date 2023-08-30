# AddClusterRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | [**ClusterCreate**](ClusterCreate.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddClusterRequest = Initialize-PSOpenAPIToolsAddClusterRequest  -Cluster null
```

- Convert the resource to JSON
```powershell
$AddClusterRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

