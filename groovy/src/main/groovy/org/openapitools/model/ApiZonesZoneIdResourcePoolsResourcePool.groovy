package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.AnyOfAwsResourcePoolConfigCloudFoundryResourcePoolConfig;
import org.openapitools.model.ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions;
import org.openapitools.model.ApiZonesZoneIdFoldersIdFolderTenantPermissions;

@Canonical
class ApiZonesZoneIdResourcePoolsResourcePool {
    /* Name of Resource Pool */
    String name
    /* Set as the Default Pool */
    Boolean defaultPool = false
    /* Set as the Default Image Target */
    Boolean defaultImage = false
    /* Activate `true` or disable `false` the datastore */
    Boolean active = true
    /* Setting `private` or `public` */
    String visibility = VisibilityEnum.PRIVATE
    /* Optional Display Name (VMware only) */
    String displayName
    /* Enable `True` or disable `False` inventory sync for resource pool during cloud refresh */
    Boolean inventory = true
    
    AnyOfAwsResourcePoolConfigCloudFoundryResourcePoolConfig config = null
    
    List<ApiZonesZoneIdFoldersIdFolderTenantPermissions> tenantPermissions = new ArrayList<ApiZonesZoneIdFoldersIdFolderTenantPermissions>()
    
    ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions resourcePermissions
}
