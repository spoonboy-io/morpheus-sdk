package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;
import org.openapitools.model.ImageBuildsConfigInstanceLayout;

@Canonical
class ImageBuildsConfigInstance {
    
    ImageBuildsConfigInstanceLayout layout
    
    ApiBlueprintsIdUpdatePermissionsResourcePermissionSites site
    
    String name
    
    String type
}
