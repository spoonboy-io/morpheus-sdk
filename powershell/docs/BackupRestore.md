# BackupRestore
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** | Backup Result ID | [optional] 
**BackupResultId** | **Int64** |  | [optional] 
**BackupId** | **Int64** |  | [optional] 
**Backup** | [**BackupRestoreBackup**](BackupRestoreBackup.md) |  | [optional] 
**ContainerId** | **Int64** |  | [optional] 
**Container** | [**BackupRestoreContainer**](BackupRestoreContainer.md) |  | [optional] 
**Instance** | [**BackupInstance**](BackupInstance.md) |  | [optional] 
**RestoreToNew** | **Boolean** |  | [optional] 
**Status** | **String** |  | [optional] 
**ErrorMessage** | **String** |  | [optional] 
**StartDate** | **System.DateTime** |  | [optional] 
**EndDate** | **System.DateTime** |  | [optional] 
**DurationMillis** | **Int64** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**ExternalStatusRef** | **String** |  | [optional] 
**DateCreated** | **System.DateTime** | Date Created | [optional] 
**LastUpdated** | **System.DateTime** | Last Updated | [optional] 

## Examples

- Prepare the resource
```powershell
$BackupRestore = Initialize-PSOpenAPIToolsBackupRestore  -Id null `
 -BackupResultId null `
 -BackupId null `
 -Backup null `
 -ContainerId null `
 -Container null `
 -Instance null `
 -RestoreToNew null `
 -Status null `
 -ErrorMessage null `
 -StartDate null `
 -EndDate null `
 -DurationMillis null `
 -ExternalId null `
 -ExternalStatusRef null `
 -DateCreated null `
 -LastUpdated null
```

- Convert the resource to JSON
```powershell
$BackupRestore | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

