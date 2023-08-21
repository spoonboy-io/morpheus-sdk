package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionDefaultTask {
    /* `Task` is the code for Default Task Access */
    String permissionCode
    /* The new access level. */
    String access
}
