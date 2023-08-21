package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionDefaultBlueprint {
    /* `AppTemplate` is the code for Default Blueprint Access */
    String permissionCode
    /* The new access level. */
    String access
}
