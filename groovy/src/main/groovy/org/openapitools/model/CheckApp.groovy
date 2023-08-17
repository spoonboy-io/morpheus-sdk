package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.jackson.nullable.JsonNullable;
import org.openapitools.model.ActivityActivityInnerUser;
import org.openapitools.model.CheckGroupInstance;
import org.openapitools.model.UpdateBlueprintPermissionsRequestResourcePermissionSitesInner;

@Canonical
class CheckApp {
    
    Long id
    
    UpdateBlueprintPermissionsRequestResourcePermissionSitesInner account
    
    Boolean active
    
    CheckGroupInstance app
    
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
    
    ActivityActivityInnerUser createdBy
    
    Date dateCreated
    
    Date lastUpdated
    
    String availability
    
    List<Long> checks
    
    List<Long> checkGroups
}
