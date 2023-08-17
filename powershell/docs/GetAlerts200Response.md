# GetAlerts200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | [**Alert**](Alert.md) |  | [optional] 
**Checks** | [**Check[]**](Check.md) |  | [optional] 
**CheckGroups** | [**CheckGroup[]**](CheckGroup.md) |  | [optional] 
**Apps** | [**CheckApp[]**](CheckApp.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetAlerts200Response = Initialize-PSOpenAPIToolsGetAlerts200Response  -Alert null `
 -Checks null `
 -CheckGroups null `
 -Apps null
```

- Convert the resource to JSON
```powershell
$GetAlerts200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

