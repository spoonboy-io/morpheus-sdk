package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionDefaultTaskSet {
    /* `TaskSet` is the code for Default Workflow Access */
    String permissionCode
    /* The new access level. */
    String access
}
