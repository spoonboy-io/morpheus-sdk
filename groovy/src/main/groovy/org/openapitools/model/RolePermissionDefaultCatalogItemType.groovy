package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionDefaultCatalogItemType {
    /* `CatalogItemType` is the code for Default Catalog Item Type Access */
    String permissionCode
    /* The new access level. */
    String access
}
