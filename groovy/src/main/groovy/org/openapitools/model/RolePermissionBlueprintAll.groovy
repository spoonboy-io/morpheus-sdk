package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionBlueprintAll {
    /* Apply to all blueprints (appTemplate) */
    Boolean allAppTemplates
    /* The new access level. */
    String access
}
