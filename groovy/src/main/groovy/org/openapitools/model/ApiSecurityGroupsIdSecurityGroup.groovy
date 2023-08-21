package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiSecurityGroupsSecurityGroupTenantPermissions;
import org.openapitools.model.ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions;

@Canonical
class ApiSecurityGroupsIdSecurityGroup {
    /* Name for your security group */
    String name
    /* Optional description field */
    String description
    /* Set to `false` to disable a security group. */
    Boolean active
    
    List<ApiSecurityGroupsSecurityGroupTenantPermissions> tenantPermissions = new ArrayList<ApiSecurityGroupsSecurityGroupTenantPermissions>()
    
    ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions resourcePermissions
}
