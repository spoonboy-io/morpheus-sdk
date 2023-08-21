package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;
import org.openapitools.model.CheckCheckType;
import org.openapitools.model.InlineResponse200107NetworkPoolCreatedBy;
import org.openapitools.model.InlineResponse20082LoadBalancerInstanceSslCert;

@Canonical
class CheckGroup {
    
    Long id
    
    ApiBlueprintsIdUpdatePermissionsResourcePermissionSites account
    
    InlineResponse20082LoadBalancerInstanceSslCert instance
    
    String name
    
    String description
    
    Boolean inUptime
    
    String lastCheckStatus
    
    Date lastWarningDate
    
    Date lastErrorDate
    
    Date lastSuccessDate
    
    Date lastRunDate
    
    String lastError
    
    Long outageTime
    
    Long lastTimer
    
    Long health
    
    String history
    
    Long minHappy
    
    String lastMetric
    
    String severity
    
    Boolean createIncident
    
    Boolean muted
    
    InlineResponse200107NetworkPoolCreatedBy createdBy
    
    Date dateCreated
    
    Date lastUpdated
    
    BigDecimal availability
    
    CheckCheckType checkType
    
    List<Long> checks = new ArrayList<Long>()
}
