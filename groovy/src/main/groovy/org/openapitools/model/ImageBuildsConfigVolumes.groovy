package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ImageBuildsConfigVolumes {
    
    Boolean volumeCustomizable
    
    Long vId
    
    Boolean readonlyName
    
    Long size
    
    String maxIOPS
    
    String name
    
    Boolean rootVolume
    
    Long storageType
    
    String datastoreId
    
    Long maxStorage
}
