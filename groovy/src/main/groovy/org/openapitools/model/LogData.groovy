package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class LogData {
    
    String typeCode
    
    String message
    
    String level
    
    Date ts
    
    String sourceType
    
    String title
    
    String logSignature
    
    String objectId
    
    Long seq
    
    String id
    
    Boolean signatureVerified
}
