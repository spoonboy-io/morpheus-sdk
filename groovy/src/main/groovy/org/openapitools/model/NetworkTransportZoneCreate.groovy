package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;

@Canonical
class NetworkTransportZoneCreate {
    /* Network transport zone name */
    String name
    /* Network transport zone description */
    String description
    /* private or public */
    String visibility
    /* Array of tenant account ids that are allowed access */
    List<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites> tenants = new ArrayList<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites>()
}
