package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

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
