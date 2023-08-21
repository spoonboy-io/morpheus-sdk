package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionTaskAll {
    /* Apply to all tasks */
    Boolean allTasks
    /* The new access level. */
    String access
}
