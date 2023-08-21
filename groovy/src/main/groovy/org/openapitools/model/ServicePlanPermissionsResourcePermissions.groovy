package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;
import org.openapitools.model.ResourcePermissionsSites;

@Canonical
class ServicePlanPermissionsResourcePermissions {
    
    Boolean defaultStore
    
    Boolean allPlans
    
    Boolean defaultTarget
    
    Boolean canManage
    
    Boolean all
    
    ApiBlueprintsIdUpdatePermissionsResourcePermissionSites account
    
    List<ResourcePermissionsSites> sites = new ArrayList<ResourcePermissionsSites>()
    
    List<Object> plans = new ArrayList<Object>()
}
