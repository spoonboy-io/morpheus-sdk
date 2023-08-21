package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ResourcePermissionsSites;

@Canonical
class ResourcePermissions {
    
    Boolean all
    
    List<ResourcePermissionsSites> sites = new ArrayList<ResourcePermissionsSites>()
    
    Boolean allPlans
    
    List<ResourcePermissionsSites> plans = new ArrayList<ResourcePermissionsSites>()
}
