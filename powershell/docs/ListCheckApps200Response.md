# ListCheckApps200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 
**MonitorApps** | [**CheckApp[]**](CheckApp.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ListCheckApps200Response = Initialize-PSOpenAPIToolsListCheckApps200Response  -Meta null `
 -MonitorApps null
```

- Convert the resource to JSON
```powershell
$ListCheckApps200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

