package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionTaskSetAll {
    /* Apply to all workflows (taskSets) */
    Boolean allTaskSets
    /* The new access level. */
    String access
}
