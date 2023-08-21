package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class VdiPoolConfigVolumes {
    
    Boolean rootVolume
    
    String name
    
    Long maxStorage
    
    Boolean volumeCustomizable
    
    Boolean readonlyName
    
    Long size
    
    Long storageType
    
    String maxIOPS
    
    String datastoreId
    
    Long vId
}
