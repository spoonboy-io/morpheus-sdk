package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;

@Canonical
class ClusterDatastorePermissionsResourcePermissions {
    
    Boolean allGroups
    
    Boolean defaultStore
    
    Boolean allPlans
    
    Boolean defaultTarget
    
    String morpheusResourceType
    
    Long morpheusResourceId
    
    Boolean canManage
    
    Boolean all
    
    ApiBlueprintsIdUpdatePermissionsResourcePermissionSites account
    
    List<Object> sites = new ArrayList<Object>()
    
    List<Object> plans = new ArrayList<Object>()
}
