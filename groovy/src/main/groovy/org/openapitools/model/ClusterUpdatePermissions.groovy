package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ClusterUpdatePermissionsResourcePermissions;
import org.openapitools.model.ClusterUpdatePermissionsResourcePool;
import org.openapitools.model.ClusterUpdatePermissionsTenantPermissions;

@Canonical
class ClusterUpdatePermissions {
    
    ClusterUpdatePermissionsResourcePermissions resourcePermissions
    
    ClusterUpdatePermissionsResourcePool resourcePool
    
    ClusterUpdatePermissionsTenantPermissions tenantPermissions
}
