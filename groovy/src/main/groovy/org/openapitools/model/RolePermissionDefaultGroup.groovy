package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionDefaultGroup {
    /* `ComputeSite` is the code for Default Group Access */
    String permissionCode = PermissionCodeEnum.COMPUTESITE
    /* The new access level. */
    String access
}
