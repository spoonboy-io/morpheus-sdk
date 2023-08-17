package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.jackson.nullable.JsonNullable;
import org.openapitools.model.ActivityActivityInnerUser;
import org.openapitools.model.CheckCheckType;
import org.openapitools.model.CheckGroupInstance;
import org.openapitools.model.UpdateBlueprintPermissionsRequestResourcePermissionSitesInner;

@Canonical
class CheckGroup {
    
    Long id
    
    UpdateBlueprintPermissionsRequestResourcePermissionSitesInner account
    
    CheckGroupInstance instance
    
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
    
    ActivityActivityInnerUser createdBy
    
    Date dateCreated
    
    Date lastUpdated
    
    BigDecimal availability
    
    CheckCheckType checkType
    
    List<Long> checks
}
