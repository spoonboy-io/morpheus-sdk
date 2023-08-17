# AddAlerts200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**Alert** | [**Alert**](Alert.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddAlerts200Response = Initialize-PSOpenAPIToolsAddAlerts200Response  -Success null `
 -Alert null
```

- Convert the resource to JSON
```powershell
$AddAlerts200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

