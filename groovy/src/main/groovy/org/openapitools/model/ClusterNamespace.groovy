package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ClusterNamespacePermissions;

@Canonical
class ClusterNamespace {
    
    Long id
    
    String visibility
    
    String name
    
    String description
    
    String status
    
    String externalId
    
    ClusterNamespacePermissions permissions
}
