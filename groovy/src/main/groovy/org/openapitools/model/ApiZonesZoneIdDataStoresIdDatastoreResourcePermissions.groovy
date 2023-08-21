package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiZonesZoneIdDataStoresIdDatastoreResourcePermissionsSites;

@Canonical
class ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions {
    /* Pass `true` to allow access all groups */
    Boolean all = true
    /* Array of groups that are allowed access */
    List<ApiZonesZoneIdDataStoresIdDatastoreResourcePermissionsSites> sites = new ArrayList<ApiZonesZoneIdDataStoresIdDatastoreResourcePermissionsSites>()
    /* Pass true to allow access all plans */
    Boolean allPlans = true
    /* Array of plans that are allowed access */
    List<ApiZonesZoneIdDataStoresIdDatastoreResourcePermissionsSites> plans = new ArrayList<ApiZonesZoneIdDataStoresIdDatastoreResourcePermissionsSites>()
}
