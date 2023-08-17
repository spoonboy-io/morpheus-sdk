

# BackupSettingsUpdate


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**backupsEnabled** | **Boolean** | Use this to enable / disable scheduled backups |  [optional] |
|**retentionCount** | **Long** | Maximum number of successful backups to retain |  [optional] |
|**createBackups** | **Boolean** | Use this to enable / disable create backups |  [optional] |
|**backupAppliance** | **Boolean** | When enabled, a Backup will be created to backup the Morpheus appliance database |  [optional] |
|**updateExisting** | **Boolean** | Use this to update existing backups with new settings |  [optional] |
|**defaultSchedule** | [**BackupSettingsUpdateDefaultSchedule**](BackupSettingsUpdateDefaultSchedule.md) |  |  [optional] |
|**clearDefaultSchedule** | **Boolean** | Use this to clear existing default backup schedule |  [optional] |
|**defaultStorageBucket** | [**BackupSettingsUpdateDefaultStorageBucket**](BackupSettingsUpdateDefaultStorageBucket.md) |  |  [optional] |
|**clearDefaultStorageBucket** | **Boolean** | Use this to clear default store bucket |  [optional] |



