package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ContainerPort {
    
    Long id
    
    Long index
    
    Long external
    
    Long internal
    
    String displayName
    
    Boolean primaryPort
    
    Boolean export
    
    Boolean visible
    
    String exportName
    
    String loadBalanceProtocol
    
    Boolean loadBalance
    
    String protocol
    
    Boolean link
    
    String externalIp
    
    String internalIp
}
