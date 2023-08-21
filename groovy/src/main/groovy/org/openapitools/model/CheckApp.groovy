package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;
import org.openapitools.model.InlineResponse200107NetworkPoolCreatedBy;
import org.openapitools.model.InlineResponse20082LoadBalancerInstanceSslCert;

@Canonical
class CheckApp {
    
    Long id
    
    ApiBlueprintsIdUpdatePermissionsResourcePermissionSites account
    
    Boolean active
    
    InlineResponse20082LoadBalancerInstanceSslCert app
    
    String name
    
    String description
    
    Boolean inUptime
    
    String lastCheckStatus
    
    Date lastWarningDate
    
    Date lastErrorDate
    
    Date lastSuccessDate
    
    Date lastRunDate
    
    String lastError
    
    Long lastTimer
    
    Long health
    
    String history
    
    String severity
    
    Boolean createIncident
    
    Boolean muted
    
    InlineResponse200107NetworkPoolCreatedBy createdBy
    
    Date dateCreated
    
    Date lastUpdated
    
    String availability
    
    List<Long> checks = new ArrayList<Long>()
    
    List<Long> checkGroups = new ArrayList<Long>()
}
