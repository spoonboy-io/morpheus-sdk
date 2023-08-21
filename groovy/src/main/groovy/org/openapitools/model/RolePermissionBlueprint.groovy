package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionBlueprint {
    /* `id` of the blueprint (appTemplate) */
    Integer appTemplateId
    /* The new access level. */
    String access
}
