package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.jackson.nullable.JsonNullable;

@Canonical
class AppStateWorkloadsInner {
    
    String refType
    
    Long refId
    
    String refName
    
    String subRefName
    
    Date stateDate
    
    String status
    
    Boolean iacDrift
}
