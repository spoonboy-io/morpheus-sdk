package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionGroupAll {
    /* Apply to all groups (site) */
    Boolean allGroups
    /* The new access level. */
    String access
}
