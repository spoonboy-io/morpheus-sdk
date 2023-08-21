# MorpheusApi.Backup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Number** | Backup ID | [optional] 
**name** | **String** | Name | [optional] 
**locationType** | **String** | Source Type (instance, server, storage) | [optional] 
**instance** | [**BackupInstance**](BackupInstance.md) |  | [optional] 
**containerId** | **Number** |  | [optional] 
**job** | [**BackupJob**](BackupJob.md) |  | [optional] 
**schedule** | [**BackupSchedule**](BackupSchedule.md) |  | [optional] 
**retentionCount** | **Number** |  | [optional] 
**backupType** | [**BackupBackupType**](BackupBackupType.md) |  | [optional] 
**storageProvider** | [**BackupStorageProvider**](BackupStorageProvider.md) |  | [optional] 
**backupProvider** | [**BackupBackupProvider**](BackupBackupProvider.md) |  | [optional] 
**backupRespository** | [**BackupBackupRespository**](BackupBackupRespository.md) |  | [optional] 
**cronExpression** | **String** | Cron Expression | [optional] 
**nextFire** | **Date** | Next Fire | [optional] 
**lastStatus** | **String** | Last Status | [optional] 
**lastResult** | [**BackupLastResult**](BackupLastResult.md) |  | [optional] 
**stats** | [**BackupStats**](BackupStats.md) |  | [optional] 
**enabled** | **Boolean** | Enabled | [optional] 
**dateCreated** | **Date** | Date Created | [optional] 
**lastUpdated** | **Date** | Last Updated | [optional] 


