package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;
import org.openapitools.model.Check;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class Incident {
    
    Long id
    
    ApiBlueprintsIdUpdatePermissionsResourcePermissionSites account
    
    String app
    
    Boolean autoClose
    
    String channelId
    
    List<InlineResponse20040AppDeployInstance> checkGroups = new ArrayList<InlineResponse20040AppDeployInstance>()
    
    List<Check> checks = new ArrayList<Check>()
    
    String comment
    
    String displayName
    
    String duration
    
    Date endDate
    
    Boolean inUptime
    
    Boolean muted
    
    Date lastCheckTime
    
    String lastError
    
    String lastMessage
    
    String name
    
    String resolution
    
    String severity
    
    Long severityId
    
    Date startDate
    
    String status
    
    String visibility
}
