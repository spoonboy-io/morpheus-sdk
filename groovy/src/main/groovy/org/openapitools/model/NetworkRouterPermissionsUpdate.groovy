package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.NetworkRouterPermissionsUpdateTenantPermissions;

@Canonical
class NetworkRouterPermissionsUpdate {
    /* Sets visibility - public, private */
    String visibility
    
    NetworkRouterPermissionsUpdateTenantPermissions tenantPermissions
}
