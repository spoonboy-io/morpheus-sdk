package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20083LoadBalancerNodeCreatedBy;

@Canonical
class UsagesActivity {
    
    String id
    
    Boolean success
    
    String activityType
    
    String name
    
    String message
    
    String objectType
    
    Long objectId
    
    InlineResponse20083LoadBalancerNodeCreatedBy user
    
    Date ts
}
