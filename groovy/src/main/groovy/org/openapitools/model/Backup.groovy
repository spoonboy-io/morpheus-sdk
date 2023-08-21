package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.BackupBackupProvider;
import org.openapitools.model.BackupBackupRespository;
import org.openapitools.model.BackupBackupType;
import org.openapitools.model.BackupInstance;
import org.openapitools.model.BackupJob;
import org.openapitools.model.BackupLastResult;
import org.openapitools.model.BackupSchedule;
import org.openapitools.model.BackupStats;
import org.openapitools.model.BackupStorageProvider;

@Canonical
class Backup {
    /* Backup ID */
    Long id
    /* Name */
    String name
    /* Source Type (instance, server, storage) */
    String locationType
    
    BackupInstance instance
    
    Long containerId
    
    BackupJob job
    
    BackupSchedule schedule
    
    Long retentionCount
    
    BackupBackupType backupType
    
    BackupStorageProvider storageProvider
    
    BackupBackupProvider backupProvider
    
    BackupBackupRespository backupRespository
    /* Cron Expression */
    String cronExpression
    /* Next Fire */
    Date nextFire
    /* Last Status */
    String lastStatus
    
    BackupLastResult lastResult
    
    BackupStats stats
    /* Enabled */
    Boolean enabled
    /* Date Created */
    Date dateCreated
    /* Last Updated */
    Date lastUpdated
}
