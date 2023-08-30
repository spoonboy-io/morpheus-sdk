# Alert
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**AllApps** | **Boolean** |  | [optional] 
**AllChecks** | **Boolean** |  | [optional] 
**AllGroups** | **Boolean** |  | [optional] 
**Active** | **Boolean** |  | [optional] 
**MinSeverity** | **String** |  | [optional] 
**MinDuration** | **Int64** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**Checks** | **Int32[]** |  | [optional] 
**CheckGroups** | **Int32[]** |  | [optional] 
**Apps** | **Int32[]** |  | [optional] 
**Contacts** | [**AddAlertsRequestAlertContactsInner[]**](AddAlertsRequestAlertContactsInner.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$Alert = Initialize-PSOpenAPIToolsAlert  -Id null `
 -Name null `
 -AllApps null `
 -AllChecks null `
 -AllGroups null `
 -Active null `
 -MinSeverity null `
 -MinDuration null `
 -DateCreated null `
 -LastUpdated null `
 -Checks null `
 -CheckGroups null `
 -Apps null `
 -Contacts null
```

- Convert the resource to JSON
```powershell
$Alert | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

