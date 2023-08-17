package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import org.openapitools.jackson.nullable.JsonNullable;
import org.openapitools.model.CheckCheckType;
import org.openapitools.model.CheckConfig;
import org.openapitools.model.CheckContainer;
import org.openapitools.model.CheckCreatedBy;
import org.openapitools.model.UpdateBlueprintPermissionsRequestResourcePermissionSitesInner;

@Canonical
class Check {
    
    Long id
    
    UpdateBlueprintPermissionsRequestResourcePermissionSitesInner account
    
    Boolean active
    
    String apiKey
    
    BigDecimal availability
    
    String checkAgent
    
    Long checkInterval
    
    String checkSpec
    
    CheckCheckType checkType
    
    CheckConfig config
    
    CheckContainer container
    
    Boolean createIncident
    
    Boolean muted
    
    CheckCreatedBy createdBy
    
    Date dateCreated
    
    String description
    
    Date endDate
    
    Long health
    
    Boolean inUptime
    
    String lastBoxStats
    
    String lastCheckStatus
    
    String lastError
    
    Date lastErrorDate
    
    String lastMessage
    
    String lastMetric
    
    Date lastRunDate
    
    String lastStats
    
    Date lastSuccessDate
    
    Long lastTimer
    
    Date lastUpdated
    
    Date lastWarningDate
    
    String name
    
    Date nextRunDate
    
    Long outageTime
    
    String severity
    
    Date startDate
    
    Boolean deleted
}
