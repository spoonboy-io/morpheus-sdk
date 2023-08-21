package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ClusterServices {
    
    Long id
    
    String name
    
    String type
    
    String code
    
    String externalIp
    
    String internalIp
    
    String externalPort
    
    String internalPort
    
    String status
    
    Date dateCreated
    
    Date lastUpdated
}
