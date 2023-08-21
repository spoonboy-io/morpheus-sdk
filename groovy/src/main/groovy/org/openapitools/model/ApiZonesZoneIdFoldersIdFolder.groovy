package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions;
import org.openapitools.model.ApiZonesZoneIdFoldersIdFolderTenantPermissions;

@Canonical
class ApiZonesZoneIdFoldersIdFolder {
    
    Boolean defaultFolder = false
    
    Boolean defaultImage = false
    /* Activate `true` or disable `false` the folder */
    Boolean active
    /* Setting `private` or `public` */
    String visibility = VisibilityEnum.PRIVATE
    
    List<ApiZonesZoneIdFoldersIdFolderTenantPermissions> tenantPermissions = new ArrayList<ApiZonesZoneIdFoldersIdFolderTenantPermissions>()
    
    ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions resourcePermissions
}
