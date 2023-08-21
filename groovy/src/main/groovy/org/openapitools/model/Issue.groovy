package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class Issue {
    
    Long id
    
    String attachmentType
    
    String app
    
    Boolean available
    
    String check
    
    InlineResponse20040AppDeployInstance checkGroup
    
    String checkStatus
    
    Date endDate
    
    Long health
    
    Boolean inUptime
    
    ApiBlueprintsIdUpdatePermissionsResourcePermissionSites incident
    
    Date lastCheckTime
    
    String lastError
    
    String lastMessage
    
    String name
    
    String severity
    
    Long severityId
    
    Date startDate
    
    String status
}
