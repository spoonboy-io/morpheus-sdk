# BackupSettings
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupsEnabled** | **Boolean** |  | [optional] 
**CreateBackups** | **Boolean** |  | [optional] 
**BackupAppliance** | **Boolean** |  | [optional] 
**DefaultStorageBucket** | [**ApplianceSettingsEnabledZoneTypesInner**](ApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**DefaultSchedule** | [**BackupSettingsDefaultSchedule**](BackupSettingsDefaultSchedule.md) |  | [optional] 
**RetentionCount** | **Int64** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BackupSettings = Initialize-PSOpenAPIToolsBackupSettings  -BackupsEnabled null `
 -CreateBackups null `
 -BackupAppliance null `
 -DefaultStorageBucket null `
 -DefaultSchedule null `
 -RetentionCount null
```

- Convert the resource to JSON
```powershell
$BackupSettings | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
