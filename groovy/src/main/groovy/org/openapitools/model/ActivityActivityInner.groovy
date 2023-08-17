package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ActivityActivityInnerUser;

@Canonical
class ActivityActivityInner {
    
    String id
    
    Boolean success
    
    String activityType
    
    String name
    
    String message
    
    String objectType
    
    Long objectId
    
    ActivityActivityInnerUser user
    
    Date ts
}
