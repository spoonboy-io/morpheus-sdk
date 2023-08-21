package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class VdiGateway {
    
    Long id
    
    String name
    
    String description
    
    String gatewayUrl
    
    String apiKey
    
    Date dateCreated
    
    Date lastUpdated
}
