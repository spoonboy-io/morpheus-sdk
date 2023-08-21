package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;

@Canonical
class ApiIntegrationsIdInventoryInventoryIdInventory {
    /* Array of tenant accounts that will use this inventory as Default. Used by jobs set to 'Use Tenant Default' */
    List<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites> tenants = new ArrayList<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites>()
}
