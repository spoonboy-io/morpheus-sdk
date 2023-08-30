# AddCluster200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | [**Cluster**](Cluster.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddCluster200Response = Initialize-PSOpenAPIToolsAddCluster200Response  -Cluster null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddCluster200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

