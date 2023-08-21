package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.JobExecutionJob;

@Canonical
class JobExecution {
    
    Long id
    
    String name
    
    String process
    
    JobExecutionJob job
    
    String description
    
    Date dateCreated
    
    Date startDate
    
    Date endDate
    
    Long duration
    
    String resultData
    
    String status
    
    String statusMessage
    
    String createdBy
}
