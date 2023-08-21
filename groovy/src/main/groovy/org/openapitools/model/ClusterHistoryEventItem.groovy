package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ClusterContainersAvailableActions;
import org.openapitools.model.ClusterHistoryCreatedBy;

@Canonical
class ClusterHistoryEventItem {
    
    Long id
    
    Long processId
    
    Long accountId
    
    String uniqueId
    
    ClusterContainersAvailableActions processType
    
    String description
    
    String refType
    
    Long refId
    
    String subType
    
    String subId
    
    String zoneId
    
    String integrationId
    
    String instanceId
    
    String containerId
    
    Long serverId
    
    String containerName
    
    String displayName
    
    String status
    
    String reason
    
    Long percent
    
    Long statusEta
    
    String message
    
    String output
    
    String error
    
    Date startDate
    
    Date endDate
    
    Long duration
    
    Date dateCreated
    
    Date lastUpdated
    
    ClusterHistoryCreatedBy createdBy
    
    ClusterHistoryCreatedBy updatedBy
}
