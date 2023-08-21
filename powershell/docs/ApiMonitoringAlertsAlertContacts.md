# ApiMonitoringAlertsAlertContacts
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Method** | **String** |  | [optional] 
**Notify** | **Boolean** |  | [optional] 
**Close** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiMonitoringAlertsAlertContacts = Initialize-PSOpenAPIToolsApiMonitoringAlertsAlertContacts  -Id null `
 -Name null `
 -Method null `
 -Notify null `
 -Close null
```

- Convert the resource to JSON
```powershell
$ApiMonitoringAlertsAlertContacts | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

