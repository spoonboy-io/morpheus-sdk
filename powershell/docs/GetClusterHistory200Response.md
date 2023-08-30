# GetClusterHistory200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Processes** | [**ClusterHistory[]**](ClusterHistory.md) |  | [optional] 
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetClusterHistory200Response = Initialize-PSOpenAPIToolsGetClusterHistory200Response  -Processes null `
 -Meta null
```

- Convert the resource to JSON
```powershell
$GetClusterHistory200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

