package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.BackupSettingsUpdateDefaultSchedule;
import org.openapitools.model.BackupSettingsUpdateDefaultStorageBucket;

@Canonical
class BackupSettingsUpdate {
    /* Use this to enable / disable scheduled backups */
    Boolean backupsEnabled
    /* Maximum number of successful backups to retain */
    Long retentionCount
    /* Use this to enable / disable create backups */
    Boolean createBackups
    /* When enabled, a Backup will be created to backup the Morpheus appliance database */
    Boolean backupAppliance
    /* Use this to update existing backups with new settings */
    Boolean updateExisting
    
    BackupSettingsUpdateDefaultSchedule defaultSchedule
    /* Use this to clear existing default backup schedule */
    Boolean clearDefaultSchedule
    
    BackupSettingsUpdateDefaultStorageBucket defaultStorageBucket
    /* Use this to clear default store bucket */
    Boolean clearDefaultStorageBucket
}
