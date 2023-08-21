package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;

@Canonical
class ClusterPods {
    
    Long id
    
    String name
    
    String code
    
    String description
    
    String category
    
    String resourceLevel
    
    String resourceType
    
    Boolean managed
    
    String status
    
    Date lastUpdated
    
    ApiBlueprintsIdUpdatePermissionsResourcePermissionSites owner
    
    Long totalCpuUsage
    
    Object stats
}
