# BackupSettingsUpdate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**backups_enabled** | **bool** | Use this to enable / disable scheduled backups | [optional] 
**retention_count** | **int** | Maximum number of successful backups to retain | [optional] 
**create_backups** | **bool** | Use this to enable / disable create backups | [optional] 
**backup_appliance** | **bool** | When enabled, a Backup will be created to backup the Morpheus appliance database | [optional] 
**update_existing** | **bool** | Use this to update existing backups with new settings | [optional] 
**default_schedule** | [**BackupSettingsUpdateDefaultSchedule**](BackupSettingsUpdateDefaultSchedule.md) |  | [optional] 
**clear_default_schedule** | **bool** | Use this to clear existing default backup schedule | [optional] 
**default_storage_bucket** | [**BackupSettingsUpdateDefaultStorageBucket**](BackupSettingsUpdateDefaultStorageBucket.md) |  | [optional] 
**clear_default_storage_bucket** | **bool** | Use this to clear default store bucket | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


