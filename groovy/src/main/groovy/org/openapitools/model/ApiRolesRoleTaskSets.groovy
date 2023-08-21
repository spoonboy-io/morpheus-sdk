package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ApiRolesRoleTaskSets {
    /* `id` of the workflow (taskSet) */
    Integer id
    /* The new access level. */
    String access
}
