# ListBlueprints200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Blueprints** | [**Blueprint[]**](Blueprint.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ListBlueprints200Response = Initialize-PSOpenAPIToolsListBlueprints200Response  -Meta null `
 -Blueprints null
```

- Convert the resource to JSON
```powershell
$ListBlueprints200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

