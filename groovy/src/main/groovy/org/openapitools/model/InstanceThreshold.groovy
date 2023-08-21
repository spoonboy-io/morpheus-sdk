package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class InstanceThreshold {
    
    Long id
    
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
    
    String minNetwork
    
    Boolean networkEnabled
    
    Boolean iopsEnabled
    
    String minIops
    
    String maxIops
    
    String comment
    
    Long zoneId
    
    Date dateCreated
    
    Date lastUpdated
}
