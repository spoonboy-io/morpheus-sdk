package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import org.openapitools.model.DashboardUser;

@Canonical
class DashboardActivity {
    
    String id
    
    Boolean success
    
    String activityType
    
    String name
    
    String message
    
    String objectType
    
    BigDecimal objectId
    
    DashboardUser user
    
    String ts
}
