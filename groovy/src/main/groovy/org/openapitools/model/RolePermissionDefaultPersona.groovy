package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionDefaultPersona {
    /* `Persona` is the code for Default Persona Access */
    String permissionCode
    /* The new access level. */
    String access
}
