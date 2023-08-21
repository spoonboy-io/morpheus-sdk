package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionCatalogItemType {
    /* `id` of the catalog item type */
    Integer catalogItemTypeId
    /* The new access level. */
    String access
}
