package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class AppStateWorkloads {
    
    String refType
    
    Long refId
    
    String refName
    
    String subRefName
    
    Date stateDate
    
    String status
    
    Boolean iacDrift
}
