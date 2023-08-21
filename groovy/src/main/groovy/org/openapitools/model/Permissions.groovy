package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ApiZonesZoneIdFoldersIdFolderTenantPermissions;
import org.openapitools.model.ClusterPermissionsResourcePool;
import org.openapitools.model.ResourcePermissions;

@Canonical
class Permissions {
    
    ClusterPermissionsResourcePool resourcePool
    
    ResourcePermissions resourcePermissions
    
    ApiZonesZoneIdFoldersIdFolderTenantPermissions tenantPermissions
}
