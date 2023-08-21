package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.VdiAllocation;

@Canonical
class Vdi {
    
    Long id
    
    String logo
    
    String name
    
    String description
    
    String status
    
    String allocationStatus
    
    VdiAllocation allocation
}
