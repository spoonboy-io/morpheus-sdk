package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ImageBuildConfigConfig;
import org.openapitools.model.ImageBuildConfigInstance;
import org.openapitools.model.ImageBuildConfigNetworkInterfaces;
import org.openapitools.model.ImageBuildConfigVolumes;
import org.openapitools.model.ImageBuildsConfigPlan;

@Canonical
class ImageBuildConfig {
    
    ImageBuildConfigInstance instance
    
    List<ImageBuildConfigNetworkInterfaces> networkInterfaces = new ArrayList<ImageBuildConfigNetworkInterfaces>()
    
    List<ImageBuildConfigVolumes> volumes = new ArrayList<ImageBuildConfigVolumes>()
    
    List<Object> storageControllers = new ArrayList<Object>()
    
    Long zoneId
    
    ImageBuildConfigConfig config
    
    ImageBuildsConfigPlan plan
}
