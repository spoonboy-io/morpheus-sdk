package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ClusterNamespaces {
    
    Long id
    
    String name
    
    String description
    
    String regionCode
    
    String externalId
    
    String status
    
    String visibility
    
    Boolean active
}
