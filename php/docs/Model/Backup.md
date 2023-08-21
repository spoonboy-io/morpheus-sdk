# # Backup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | Backup ID | [optional]
**name** | **string** | Name | [optional]
**location_type** | **string** | Source Type (instance, server, storage) | [optional]
**instance** | [**\OpenAPI\Client\Model\BackupInstance**](BackupInstance.md) |  | [optional]
**container_id** | **int** |  | [optional]
**job** | [**\OpenAPI\Client\Model\BackupJob**](BackupJob.md) |  | [optional]
**schedule** | [**\OpenAPI\Client\Model\BackupSchedule**](BackupSchedule.md) |  | [optional]
**retention_count** | **int** |  | [optional]
**backup_type** | [**\OpenAPI\Client\Model\BackupBackupType**](BackupBackupType.md) |  | [optional]
**storage_provider** | [**\OpenAPI\Client\Model\BackupStorageProvider**](BackupStorageProvider.md) |  | [optional]
**backup_provider** | [**\OpenAPI\Client\Model\BackupBackupProvider**](BackupBackupProvider.md) |  | [optional]
**backup_respository** | [**\OpenAPI\Client\Model\BackupBackupRespository**](BackupBackupRespository.md) |  | [optional]
**cron_expression** | **string** | Cron Expression | [optional]
**next_fire** | [**\DateTime**](\DateTime.md) | Next Fire | [optional]
**last_status** | **string** | Last Status | [optional]
**last_result** | [**\OpenAPI\Client\Model\BackupLastResult**](BackupLastResult.md) |  | [optional]
**stats** | [**\OpenAPI\Client\Model\BackupStats**](BackupStats.md) |  | [optional]
**enabled** | **bool** | Enabled | [optional]
**date_created** | [**\DateTime**](\DateTime.md) | Date Created | [optional]
**last_updated** | [**\DateTime**](\DateTime.md) | Last Updated | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
