# BackupSchedule
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** | Schedule ID | [optional] 
**Name** | **String** | Schedule Name | [optional] 
**Cron** | **String** | Schedule Cron Expression | [optional] 

## Examples

- Prepare the resource
```powershell
$BackupSchedule = Initialize-PSOpenAPIToolsBackupSchedule  -Id null `
 -Name null `
 -Cron null
```

- Convert the resource to JSON
```powershell
$BackupSchedule | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

