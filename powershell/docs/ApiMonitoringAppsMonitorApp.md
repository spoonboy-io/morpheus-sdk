# ApiMonitoringAppsMonitorApp
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Unique name scoped to your account for the check app | 
**Description** | **String** | Optional description field | [optional] 
**InUptime** | **Boolean** | Used to determine if check should affect account wide availability calculations | [optional] [default to $true]
**Severity** | **String** | Severity level of incidents that are created when this check fails | [optional] [default to "critical"]
**Active** | **Boolean** | Used to determine if check app is active | [optional] [default to $true]
**Checks** | **Int32[]** |  | [optional] 
**CheckGroups** | **Int32[]** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiMonitoringAppsMonitorApp = Initialize-PSOpenAPIToolsApiMonitoringAppsMonitorApp  -Name My Check App `
 -Description My cool description `
 -InUptime null `
 -Severity null `
 -Active null `
 -Checks null `
 -CheckGroups null
```

- Convert the resource to JSON
```powershell
$ApiMonitoringAppsMonitorApp | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

