package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ClusterContainersAvailableActions;
import org.openapitools.model.ClusterHistoryCreatedBy;
import org.openapitools.model.ClusterHistoryEvents;

@Canonical
class ClusterHistoryItem {
    
    Long id
    
    Long accountId
    
    String uniqueId
    
    ClusterContainersAvailableActions processType
    
    String description
    
    String subType
    
    String subId
    
    String zoneId
    
    String integrationId
    
    Long appId
    
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
    
    List<ClusterHistoryEvents> events = new ArrayList<ClusterHistoryEvents>()
}
