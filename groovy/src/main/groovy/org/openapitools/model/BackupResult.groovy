package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.BackupJobBackups;
import org.openapitools.model.InlineResponse200108NetworkFloatingIpCreatedBy;

@Canonical
class BackupResult {
    /* Backup Result ID */
    Long id
    
    BackupJobBackups backup
    
    String backupSetId
    
    Long instanceId
    
    Long containerId
    
    Long serverId
    
    String status
    
    String errorMessage
    
    Date startDate
    
    Date endDate
    
    Long durationMillis
    
    Long sizeInBytes
    
    Long sizeInMb
    
    String volumePath
    
    String resultArchive
    
    String resultPath
    
    String externalId
    
    String snapshotId
    
    String snapshotExternalId
    
    InlineResponse200108NetworkFloatingIpCreatedBy createdBy
    /* Date Created */
    Date dateCreated
    /* Last Updated */
    Date lastUpdated
}
