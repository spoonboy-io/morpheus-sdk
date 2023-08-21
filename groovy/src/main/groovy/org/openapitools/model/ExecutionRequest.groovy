package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ExecutionRequest {
    
    Long id
    
    String uniqueId
    
    String containerId
    
    String serverId
    
    Long instanceId
    
    String resourceId
    
    String appId
    
    String stdOut
    
    String stdErr
    
    Long exitCode
    
    String status
    
    Date expiresAt
    
    Long createdById
    
    String statusMessage
    
    String errorMessage
    
    Object config
    
    String rawData
}
