package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionInstanceType {
    /* `id` of the instance type */
    Integer instanceTypeId
    /* The new access level. */
    String access
}
