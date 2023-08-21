package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ClusterMastersVolumes {
    
    Long id
    
    String name
    
    String controllerId
    
    String controllerMountPoint
    
    Boolean resizeable
    
    Boolean planResizable
    
    Boolean rootVolume
    
    String unitNumber
    
    Long typeId
    
    Boolean configurableIOPS
    
    String datastoreId
    
    Long maxStorage
    
    Long displayOrder
    
    String maxIOPS
    
    String uuid
}
