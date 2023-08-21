package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class InstanceVolumes {
    
    Long controllerId
    
    String datastoreId
    
    Long displayOrder
    
    Long id
    
    String uuid
    
    String maxIOPS
    
    Long maxStorage
    
    String name
    
    String shortName
    
    Boolean resizeable
    
    Boolean planResizable
    
    Boolean rootVolume
    
    Long size
    
    Long storageType
    
    String unitNumber
    
    String controllerMountPoint
}
