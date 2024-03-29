package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class GuidanceVmwareSizingPlanBeforeActionProvisionType {
    
    Long id
    
    String name
    
    String code
    
    Boolean rootDiskCustomizable
    
    Boolean addVolumes
    
    Boolean customizeVolume
    
    Boolean hasConfigurableCpuSockets
}
