package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ClusterDatastorePermissionsResourcePermissions;
import org.openapitools.model.InlineResponse20095NetworkRouterPermissionsTenantPermissions;

@Canonical
class ClusterDatastorePermissions {
    
    ClusterDatastorePermissionsResourcePermissions resourcePermissions
    
    InlineResponse20095NetworkRouterPermissionsTenantPermissions tenantPermissions
}
