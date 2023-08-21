package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20094Network;

@Canonical
class InlineResponse20094Interfaces {
    
    Long id
    
    String name
    
    String code
    
    String interfaceType
    
    String networkPosition
    
    String ipAddress
    
    String cidr
    
    String externalLink
    
    Boolean enabled
    
    InlineResponse20094Network network
}
