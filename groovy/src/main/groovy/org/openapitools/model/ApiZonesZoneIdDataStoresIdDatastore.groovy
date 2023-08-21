package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions;
import org.openapitools.model.ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions;

@Canonical
class ApiZonesZoneIdDataStoresIdDatastore {
    /* Activate `true` or disable `false` the datastore */
    Boolean active
    /* Setting `private` or `public` */
    String visibility = VisibilityEnum.PRIVATE
    
    List<ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions> tenantPermissions = new ArrayList<ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions>()
    
    ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions resourcePermissions
}
