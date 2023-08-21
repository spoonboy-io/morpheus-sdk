package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionPersona {
    /* code of the Persona, eg. `standard` or `serviceCatalog` */
    String personaCode
    /* The new access level. */
    String access
}
