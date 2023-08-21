package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionGroup {
    /* `id` of the group (site) */
    Integer groupId
    /* The new access level. */
    String access
}
