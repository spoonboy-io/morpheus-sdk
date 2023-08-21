package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionInstanceTypeAll {
    /* Apply to all instance types */
    Boolean allInstanceTypes
    /* The new access level. */
    String access
}
