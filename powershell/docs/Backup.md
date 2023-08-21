# Backup
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** | Backup ID | [optional] 
**Name** | **String** | Name | [optional] 
**LocationType** | **String** | Source Type (instance, server, storage) | [optional] 
**Instance** | [**BackupInstance**](BackupInstance.md) |  | [optional] 
**ContainerId** | **Int64** |  | [optional] 
**Job** | [**BackupJob**](BackupJob.md) |  | [optional] 
**Schedule** | [**BackupSchedule**](BackupSchedule.md) |  | [optional] 
**RetentionCount** | **Int64** |  | [optional] 
**BackupType** | [**BackupBackupType**](BackupBackupType.md) |  | [optional] 
**StorageProvider** | [**BackupStorageProvider**](BackupStorageProvider.md) |  | [optional] 
**BackupProvider** | [**BackupBackupProvider**](BackupBackupProvider.md) |  | [optional] 
**BackupRespository** | [**BackupBackupRespository**](BackupBackupRespository.md) |  | [optional] 
**CronExpression** | **String** | Cron Expression | [optional] 
**NextFire** | **System.DateTime** | Next Fire | [optional] 
**LastStatus** | **String** | Last Status | [optional] 
**LastResult** | [**BackupLastResult**](BackupLastResult.md) |  | [optional] 
**Stats** | [**BackupStats**](BackupStats.md) |  | [optional] 
**Enabled** | **Boolean** | Enabled | [optional] 
**DateCreated** | **System.DateTime** | Date Created | [optional] 
**LastUpdated** | **System.DateTime** | Last Updated | [optional] 

## Examples

- Prepare the resource
```powershell
$Backup = Initialize-PSOpenAPIToolsBackup  -Id null `
 -Name null `
 -LocationType null `
 -Instance null `
 -ContainerId null `
 -Job null `
 -Schedule null `
 -RetentionCount null `
 -BackupType null `
 -StorageProvider null `
 -BackupProvider null `
 -BackupRespository null `
 -CronExpression null `
 -NextFire null `
 -LastStatus null `
 -LastResult null `
 -Stats null `
 -Enabled null `
 -DateCreated null `
 -LastUpdated null
```

- Convert the resource to JSON
```powershell
$Backup | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

