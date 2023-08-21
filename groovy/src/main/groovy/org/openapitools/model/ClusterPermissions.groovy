package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ClusterPermissionsResourcePermissions;
import org.openapitools.model.ClusterPermissionsResourcePool;

@Canonical
class ClusterPermissions {
    
    ClusterPermissionsResourcePool resourcePool
    
    ClusterPermissionsResourcePermissions resourcePermissions
}
