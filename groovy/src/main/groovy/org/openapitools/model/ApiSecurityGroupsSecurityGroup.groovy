package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.AnyOfsecurityGroupLocationAzureCustomOptionssecurityGroupLocationAwsCustomOptionssecurityGroupLocationOpenstackCustomOptions;
import org.openapitools.model.ApiSecurityGroupsSecurityGroupTenantPermissions;
import org.openapitools.model.ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions;

@Canonical
class ApiSecurityGroupsSecurityGroup {
    /* Name for your security group */
    String name
    /* Optional description field */
    String description
    /* Scoped Cloud ID */
    Long zoneId
    /* Set to `false` to disable a security group. */
    Boolean active
    
    AnyOfsecurityGroupLocationAzureCustomOptionssecurityGroupLocationAwsCustomOptionssecurityGroupLocationOpenstackCustomOptions customOptions = null
    
    List<ApiSecurityGroupsSecurityGroupTenantPermissions> tenantPermissions = new ArrayList<ApiSecurityGroupsSecurityGroupTenantPermissions>()
    
    ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions resourcePermissions
}
