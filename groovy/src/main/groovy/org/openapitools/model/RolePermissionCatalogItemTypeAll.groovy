package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionCatalogItemTypeAll {
    /* Apply to all catalog item types */
    Boolean allCatalogItemTypes
    /* The new access level. */
    String access
}
