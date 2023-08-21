package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionDefaultInstanceType {
    /* `InstanceType` is the code for Default Instance Type Access */
    String permissionCode
    /* The new access level. */
    String access
}
