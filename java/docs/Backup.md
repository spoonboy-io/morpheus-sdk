

# Backup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Long** | Backup ID |  [optional]
**name** | **String** | Name |  [optional]
**locationType** | **String** | Source Type (instance, server, storage) |  [optional]
**instance** | [**BackupInstance**](BackupInstance.md) |  |  [optional]
**containerId** | **Long** |  |  [optional]
**job** | [**BackupJob**](BackupJob.md) |  |  [optional]
**schedule** | [**BackupSchedule**](BackupSchedule.md) |  |  [optional]
**retentionCount** | **Long** |  |  [optional]
**backupType** | [**BackupBackupType**](BackupBackupType.md) |  |  [optional]
**storageProvider** | [**BackupStorageProvider**](BackupStorageProvider.md) |  |  [optional]
**backupProvider** | [**BackupBackupProvider**](BackupBackupProvider.md) |  |  [optional]
**backupRespository** | [**BackupBackupRespository**](BackupBackupRespository.md) |  |  [optional]
**cronExpression** | **String** | Cron Expression |  [optional]
**nextFire** | **OffsetDateTime** | Next Fire |  [optional]
**lastStatus** | **String** | Last Status |  [optional]
**lastResult** | [**BackupLastResult**](BackupLastResult.md) |  |  [optional]
**stats** | [**BackupStats**](BackupStats.md) |  |  [optional]
**enabled** | **Boolean** | Enabled |  [optional]
**dateCreated** | **OffsetDateTime** | Date Created |  [optional]
**lastUpdated** | **OffsetDateTime** | Last Updated |  [optional]



