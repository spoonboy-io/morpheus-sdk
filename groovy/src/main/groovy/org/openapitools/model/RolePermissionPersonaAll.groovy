package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionPersonaAll {
    /* Apply to all personas */
    Boolean allPersonas
    /* The new access level. */
    String access
}
