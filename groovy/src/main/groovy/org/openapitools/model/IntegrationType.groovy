package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class IntegrationType {
    
    Long id
    
    String code
    
    String name
    
    String description
    
    String category
    
    Boolean enabled
    
    Boolean creatable
    
    Boolean hasCMDB
    
    Boolean hasCMDBDiscovery
    
    Boolean hasCM
    
    Boolean hasDNS
    
    Boolean hasApprovals
    
    Boolean hasDeleteApprovals
    
    Boolean hasReconfigureApprovals
    
    Boolean isPlugin
}
