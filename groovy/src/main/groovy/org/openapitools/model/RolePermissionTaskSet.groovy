package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionTaskSet {
    /* `id` of the workflow (taskSet) */
    Integer taskSetId
    /* The new access level. */
    String access
}
