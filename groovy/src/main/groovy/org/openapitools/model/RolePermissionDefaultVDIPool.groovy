package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionDefaultVDIPool {
    /* `VdiPools` is the code for Default VDI Pool Access */
    String permissionCode
    /* The new access level. */
    String access
}
