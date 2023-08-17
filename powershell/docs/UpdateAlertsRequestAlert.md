# UpdateAlertsRequestAlert
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Unique name scoped to your account for the alert | [optional] 
**MinDuration** | **Int32** | Duration in minutes of the delay before sending notification(s) | [optional] [default to 0]
**MinSeverity** | **String** | Severity level threshold for sending notifications. | [optional] [default to "critical"]
**Active** | **Boolean** | Set to false to disable notifications | [optional] [default to $true]
**AllChecks** | **Boolean** | Trigger for all checks | [optional] [default to $false]
**AllGroups** | **Boolean** | Trigger for all check groups | [optional] [default to $false]
**AllApps** | **Boolean** | Trigger for all monitor apps | [optional] [default to $false]
**Checks** | **Int32[]** |  | [optional] 
**Groups** | **Int32[]** |  | [optional] 
**Apps** | **Int32[]** |  | [optional] 
**Contacts** | [**AddAlertsRequestAlertContactsInner[]**](AddAlertsRequestAlertContactsInner.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$UpdateAlertsRequestAlert = Initialize-PSOpenAPIToolsUpdateAlertsRequestAlert  -Name My Alert `
 -MinDuration null `
 -MinSeverity null `
 -Active null `
 -AllChecks null `
 -AllGroups null `
 -AllApps null `
 -Checks null `
 -Groups null `
 -Apps null `
 -Contacts null
```

- Convert the resource to JSON
```powershell
$UpdateAlertsRequestAlert | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

