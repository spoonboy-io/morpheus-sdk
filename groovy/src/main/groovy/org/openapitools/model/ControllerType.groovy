package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ControllerType {
    
    Long id
    
    String name
    
    Long displayOrder
    
    String category
    
    Boolean enabled
    
    Boolean creatable
    
    Long maxDevices
}
