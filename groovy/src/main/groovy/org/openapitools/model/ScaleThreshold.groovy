package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ScaleThreshold {
    
    Long id
    
    String name
    
    String type
    
    Boolean autoUp
    
    Boolean autoDown
    
    Long minCount
    
    Long maxCount
    
    Long scaleIncrement
    
    Boolean cpuEnabled
    
    Long minCpu
    
    Long maxCpu
    
    Boolean memoryEnabled
    
    Long minMemory
    
    Long maxMemory
    
    Boolean diskEnabled
    
    Long minDisk
    
    Long maxDisk
    
    Date dateCreated
    
    Date lastUpdated
}
