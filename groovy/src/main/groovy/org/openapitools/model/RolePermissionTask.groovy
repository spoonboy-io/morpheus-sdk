package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionTask {
    /* `id` of the task */
    Integer taskId
    /* The new access level. */
    String access
}
