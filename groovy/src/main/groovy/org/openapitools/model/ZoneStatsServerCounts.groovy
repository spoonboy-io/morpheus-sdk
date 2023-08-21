package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ZoneStatsServerCounts {
    
    Long all
    
    Long host
    
    Long hypervisor
    
    Long containerHost
    
    Long vm
    
    Long baremetal
    
    Long unmanaged
}
