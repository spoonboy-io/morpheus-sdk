package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ImageBuildConfigVolumes {
    
    Long id
    
    Long size
    
    String maxIOPS
    
    String name
    
    Boolean rootVolume
    
    Long storageType
    
    String datastoreId
}
