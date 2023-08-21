package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionDefaultCloud {
    /* `ComputeZone` is the code for Default Cloud Access */
    String permissionCode
    /* The new access level. */
    String access
}
