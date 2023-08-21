package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class Deployments {
    
    Long id
    
    String name
    
    String description
    
    Long accountId
    
    String externalId
    
    Date dateCreated
    
    Date lastUpdated
    
    Long versionCount
}
