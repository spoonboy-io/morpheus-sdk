package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.HealthCpu;
import org.openapitools.model.HealthDatabase;
import org.openapitools.model.HealthElastic;
import org.openapitools.model.HealthMemory;
import org.openapitools.model.HealthRabbit;
import org.openapitools.model.HealthThreads;

@Canonical
class Health {
    
    Boolean success
    
    String statusMessage
    
    String applianceUrl
    
    String buildVersion
    
    Boolean setupNeeded
    
    Date date
    
    HealthCpu cpu
    
    HealthMemory memory
    
    HealthThreads threads
    
    HealthDatabase database
    
    HealthElastic elastic
    
    HealthRabbit rabbit
}
