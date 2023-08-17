# ListApps200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stats** | [**AppStats**](AppStats.md) |  | [optional] 
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Apps** | [**App[]**](App.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ListApps200Response = Initialize-PSOpenAPIToolsListApps200Response  -Stats null `
 -Meta null `
 -Apps null
```

- Convert the resource to JSON
```powershell
$ListApps200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

