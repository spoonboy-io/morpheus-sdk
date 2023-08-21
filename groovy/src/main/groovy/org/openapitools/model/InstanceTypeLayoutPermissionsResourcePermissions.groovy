package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.CheckContainer;

@Canonical
class InstanceTypeLayoutPermissionsResourcePermissions {
    
    Boolean defaultStore
    
    Boolean allPlans
    
    Boolean defaultTarget
    
    Boolean canManage
    
    Boolean all
    
    CheckContainer account
    
    List<Object> sites = new ArrayList<Object>()
    
    List<Object> plans = new ArrayList<Object>()
}
