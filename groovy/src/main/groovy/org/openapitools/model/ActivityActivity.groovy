package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse200107NetworkPoolCreatedBy;

@Canonical
class ActivityActivity {
    
    String id
    
    Boolean success
    
    String activityType
    
    String name
    
    String message
    
    String objectType
    
    Long objectId
    
    InlineResponse200107NetworkPoolCreatedBy user
    
    Date ts
}
