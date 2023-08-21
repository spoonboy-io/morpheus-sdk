package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions {
    /* Array of tenant account ids that are allowed access */
    List<Long> accounts = new ArrayList<Long>()
    /* Array of tenant account ids which should use the data store as the Default */
    List<Long> defaultTarget = new ArrayList<Long>()
    /* Array of tenant account ids which should use the data store as the Image Target */
    List<Long> defaultStore = new ArrayList<Long>()
}
