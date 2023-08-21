package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class NetworkRouterNat {
    
    Integer id
    
    String name
    
    String description
    
    Boolean enabled
    
    String sourceNetwork
    
    String destinationNetwork
    
    String translatedNetwork
    
    String sourcePorts
    
    String destinationPorts
    
    String translatedPorts
    
    Integer priority
    
    String protocol
    
    String matchIpv6DestinationPrefix
    
    String translatedIpv4SourcePrefix
    
    String refType
    
    String refId
    
    String syncSource
    
    String internalId
    
    String externalId
    
    String providerId
    
    Date dateCreated
    
    Date lastUpdated
}
