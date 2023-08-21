package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class NetworkInterfaceUpdateSuccessServerComputeServerType {
    
    Long id
    
    String code
    
    String name
    
    String description
    
    String platform
    
    String nodeType
    
    Boolean managed
    
    Boolean enabled
    
    Boolean vmHypervisor
    
    Boolean containerHypervisor
    
    Long displayOrder
    
    Boolean selectable
    
    Boolean controlPower
    
    Boolean controlSuspend
    
    Boolean hasAgent
    
    Boolean creatable
    
    List<Object> optionTypes = new ArrayList<Object>()
}
