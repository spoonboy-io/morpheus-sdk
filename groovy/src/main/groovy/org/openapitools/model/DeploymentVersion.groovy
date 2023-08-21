package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class DeploymentVersion {
    
    Long id
    
    String deployType
    
    Long deploymentId
    
    String fetchUrl
    
    String gitUrl
    
    String gitRef
    
    String userVersion
    
    String version
    
    String status
    
    Date dateCreated
    
    Date lastUpdated
}
