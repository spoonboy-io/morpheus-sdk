package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;

@Canonical
class ClusterPermissionsResourcePermissions {
    
    Boolean defaultStore
    
    Boolean allPlans
    
    Boolean defaultTarget
    
    Boolean canManage
    
    Boolean all
    
    ApiBlueprintsIdUpdatePermissionsResourcePermissionSites account
    
    List<Object> sites = new ArrayList<Object>()
    
    List<Object> plans = new ArrayList<Object>()
}
