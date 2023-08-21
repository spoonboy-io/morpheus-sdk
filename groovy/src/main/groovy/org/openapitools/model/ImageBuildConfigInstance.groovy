package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ImageBuildsConfigPlan;
import org.openapitools.model.ZoneVcenterConfigNetworkServer;

@Canonical
class ImageBuildConfigInstance {
    
    ImageBuildsConfigPlan layout
    
    String type
    
    ZoneVcenterConfigNetworkServer userGroup
}
