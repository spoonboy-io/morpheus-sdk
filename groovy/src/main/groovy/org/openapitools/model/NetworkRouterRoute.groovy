package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class NetworkRouterRoute {
    
    Long id
    
    String name
    
    String code
    
    String description
    
    String priority
    
    String routeType
    
    String source
    
    String sourceType
    
    String destination
    
    String destinationType
    
    Boolean defaultRoute
    
    String networkMtu
    
    String externalInterface
    
    String internalId
    
    String externalId
    
    String uniqueId
    
    String providerId
    
    String externalType
    
    Boolean enabled
    
    Boolean visible
}
