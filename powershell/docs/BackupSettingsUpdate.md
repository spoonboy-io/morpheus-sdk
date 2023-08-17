# BackupSettingsUpdate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupsEnabled** | **Boolean** | Use this to enable / disable scheduled backups | [optional] 
**RetentionCount** | **Int64** | Maximum number of successful backups to retain | [optional] 
**CreateBackups** | **Boolean** | Use this to enable / disable create backups | [optional] 
**BackupAppliance** | **Boolean** | When enabled, a Backup will be created to backup the Morpheus appliance database | [optional] 
**UpdateExisting** | **Boolean** | Use this to update existing backups with new settings | [optional] 
**DefaultSchedule** | [**BackupSettingsUpdateDefaultSchedule**](BackupSettingsUpdateDefaultSchedule.md) |  | [optional] 
**ClearDefaultSchedule** | **Boolean** | Use this to clear existing default backup schedule | [optional] 
**DefaultStorageBucket** | [**BackupSettingsUpdateDefaultStorageBucket**](BackupSettingsUpdateDefaultStorageBucket.md) |  | [optional] 
**ClearDefaultStorageBucket** | **Boolean** | Use this to clear default store bucket | [optional] 

## Examples

- Prepare the resource
```powershell
$BackupSettingsUpdate = Initialize-PSOpenAPIToolsBackupSettingsUpdate  -BackupsEnabled null `
 -RetentionCount null `
 -CreateBackups null `
 -BackupAppliance null `
 -UpdateExisting null `
 -DefaultSchedule null `
 -ClearDefaultSchedule null `
 -DefaultStorageBucket null `
 -ClearDefaultStorageBucket null
```

- Convert the resource to JSON
```powershell
$BackupSettingsUpdate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

