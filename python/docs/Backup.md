# Backup


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | Backup ID | [optional] 
**name** | **str** | Name | [optional] 
**location_type** | **str** | Source Type (instance, server, storage) | [optional] 
**instance** | [**BackupInstance**](BackupInstance.md) |  | [optional] 
**container_id** | **int, none_type** |  | [optional] 
**job** | [**BackupJob**](BackupJob.md) |  | [optional] 
**schedule** | [**BackupSchedule**](BackupSchedule.md) |  | [optional] 
**retention_count** | **int, none_type** |  | [optional] 
**backup_type** | [**BackupBackupType**](BackupBackupType.md) |  | [optional] 
**storage_provider** | [**BackupStorageProvider**](BackupStorageProvider.md) |  | [optional] 
**backup_provider** | [**BackupBackupProvider**](BackupBackupProvider.md) |  | [optional] 
**backup_respository** | [**BackupBackupRespository**](BackupBackupRespository.md) |  | [optional] 
**cron_expression** | **str, none_type** | Cron Expression | [optional] 
**next_fire** | **datetime, none_type** | Next Fire | [optional] 
**last_status** | **str, none_type** | Last Status | [optional] 
**last_result** | [**BackupLastResult**](BackupLastResult.md) |  | [optional] 
**stats** | [**BackupStats**](BackupStats.md) |  | [optional] 
**enabled** | **bool** | Enabled | [optional] 
**date_created** | **datetime** | Date Created | [optional] 
**last_updated** | **datetime** | Last Updated | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


