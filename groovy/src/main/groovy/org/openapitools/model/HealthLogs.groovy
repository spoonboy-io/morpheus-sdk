package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class HealthLogs {
    
    String typeCode
    
    Date ts
    
    String level
    
    String sourceType
    
    String message
    
    String hostname
    
    String title
    
    String logSignature
    
    String objectId
    
    Long seq
    
    String id
    
    Boolean signatureVerified
}
