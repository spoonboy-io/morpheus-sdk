# GetInstanceContainers200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Containers** | [**Container[]**](Container.md) |  | [optional] 
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetInstanceContainers200Response = Initialize-PSOpenAPIToolsGetInstanceContainers200Response  -Containers null `
 -Meta null
```

- Convert the resource to JSON
```powershell
$GetInstanceContainers200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

