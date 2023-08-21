package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;

@Canonical
class ApiSubnetsResourcePermission {
    /* Pass true to allow access all groups */
    Boolean all
    /* Array of groups ID objects that are allowed access */
    List<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites> sites = new ArrayList<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites>()
    
    Boolean allPlans
    
    List<Object> plans = new ArrayList<Object>()
}
