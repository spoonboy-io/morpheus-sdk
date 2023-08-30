# GetCheckApps200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorApp** | [**CheckApp**](CheckApp.md) |  | [optional] 
**CheckGroups** | [**CheckGroup[]**](CheckGroup.md) |  | [optional] 
**Checks** | [**Check[]**](Check.md) |  | [optional] 
**OpenIncidents** | [**Incident[]**](Incident.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetCheckApps200Response = Initialize-PSOpenAPIToolsGetCheckApps200Response  -MonitorApp null `
 -CheckGroups null `
 -Checks null `
 -OpenIncidents null
```

- Convert the resource to JSON
```powershell
$GetCheckApps200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

