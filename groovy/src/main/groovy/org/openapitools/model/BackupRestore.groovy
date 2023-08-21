package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.BackupInstance;
import org.openapitools.model.BackupRestoreBackup;
import org.openapitools.model.BackupRestoreContainer;

@Canonical
class BackupRestore {
    /* Backup Result ID */
    Long id
    
    Long backupResultId
    
    Long backupId
    
    BackupRestoreBackup backup
    
    Long containerId
    
    BackupRestoreContainer container
    
    BackupInstance instance
    
    Boolean restoreToNew
    
    String status
    
    String errorMessage
    
    Date startDate
    
    Date endDate
    
    Long durationMillis
    
    String externalId
    
    String externalStatusRef
    /* Date Created */
    Date dateCreated
    /* Last Updated */
    Date lastUpdated
}
