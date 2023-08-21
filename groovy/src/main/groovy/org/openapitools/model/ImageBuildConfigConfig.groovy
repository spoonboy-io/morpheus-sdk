package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ImageBuildConfigConfig {
    
    Long template
    
    String vmwareFolderId
    
    Long resourcePoolId
    
    String nestedVirtualization
}
