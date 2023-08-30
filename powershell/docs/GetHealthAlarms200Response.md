# GetHealthAlarms200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alarm** | [**Alarm**](Alarm.md) |  | [optional] 
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetHealthAlarms200Response = Initialize-PSOpenAPIToolsGetHealthAlarms200Response  -Alarm null `
 -Meta null
```

- Convert the resource to JSON
```powershell
$GetHealthAlarms200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

