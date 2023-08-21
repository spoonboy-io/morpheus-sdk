package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ClusterContainersAvailableActions;
import org.openapitools.model.ClusterHistoryCreatedBy;
import org.openapitools.model.ProcessEvents;

@Canonical
class Process {
    
    Long id
    
    Long accountId
    
    String uniqueId
    
    ClusterContainersAvailableActions processType
    
    String displayName
    
    String description
    
    String subType
    
    String subId
    
    Long zoneId
    
    String integrationId
    
    String appId
    
    Long instanceId
    
    Long containerId
    
    Long serverId
    
    String containerName
    
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
    
    List<ProcessEvents> events = new ArrayList<ProcessEvents>()
}
