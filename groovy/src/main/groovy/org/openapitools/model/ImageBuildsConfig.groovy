package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ImageBuildsConfigConfig;
import org.openapitools.model.ImageBuildsConfigInstance;
import org.openapitools.model.ImageBuildsConfigNetworkInterfaces;
import org.openapitools.model.ImageBuildsConfigPlan;
import org.openapitools.model.ImageBuildsConfigVolumes;

@Canonical
class ImageBuildsConfig {
    
    ImageBuildsConfigInstance instance
    
    List<ImageBuildsConfigNetworkInterfaces> networkInterfaces = new ArrayList<ImageBuildsConfigNetworkInterfaces>()
    
    List<ImageBuildsConfigVolumes> volumes = new ArrayList<ImageBuildsConfigVolumes>()
    
    List<Object> storageControllers = new ArrayList<Object>()
    
    Long zoneId
    
    ImageBuildsConfigConfig config
    
    ImageBuildsConfigPlan plan
}
