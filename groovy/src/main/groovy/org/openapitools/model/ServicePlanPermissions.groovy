package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20095NetworkRouterPermissionsTenantPermissions;
import org.openapitools.model.ServicePlanPermissionsResourcePermissions;

@Canonical
class ServicePlanPermissions {
    
    ServicePlanPermissionsResourcePermissions resourcePermissions
    
    InlineResponse20095NetworkRouterPermissionsTenantPermissions tenantPermissions
}
