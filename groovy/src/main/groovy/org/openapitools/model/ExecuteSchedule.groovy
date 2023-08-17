package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.jackson.nullable.JsonNullable;

@Canonical
class ExecuteSchedule {
    
    Long id
    
    String name
    
    String description
    
    String visibility
    
    Boolean enabled
    
    String scheduleType
    
    String scheduleTimezone
    
    String cron
    
    Date dateCreated
    
    Date lastUpdated
}
