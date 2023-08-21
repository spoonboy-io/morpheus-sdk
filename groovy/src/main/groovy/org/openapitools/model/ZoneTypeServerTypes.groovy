package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ZoneTypeOptionTypes;
import org.openapitools.model.ZoneTypeProvisionType;

@Canonical
class ZoneTypeServerTypes {
    
    Long id
    
    String code
    
    String name
    
    String description
    
    String nodeType
    
    String platform
    
    Boolean enabled
    
    Boolean selectable
    
    Boolean externalDelete
    
    Boolean managed
    
    Boolean controlPower
    
    Boolean controlSuspend
    
    Boolean creatable
    
    Boolean hasAgent
    
    Boolean vmHypervisor
    
    Boolean containerHypervisor
    
    Boolean bareMetalHost
    
    Boolean guestVm
    
    Boolean hasAutomation
    
    ZoneTypeProvisionType provisionType
    
    List<ZoneTypeOptionTypes> optionTypes = new ArrayList<ZoneTypeOptionTypes>()
    
    Long displayOrder
}
