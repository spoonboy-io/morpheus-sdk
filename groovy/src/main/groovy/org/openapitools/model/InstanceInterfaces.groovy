package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InstanceNetwork;

@Canonical
class InstanceInterfaces {
    
    String id
    
    InstanceNetwork network
    
    String ipAddress
    
    Long networkInterfaceTypeId
    
    String ipMode
}
