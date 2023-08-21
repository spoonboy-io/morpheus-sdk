package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;

@Canonical
class NetworkInterfaceUpdateSuccessServerCapacityInfo {
    
    Long id
    
    Long maxMemory
    
    ApiBlueprintsIdUpdatePermissionsResourcePermissionSites server
    
    Long usedStorage
    
    String version
    
    String maxCpu
    
    Long usedCores
    
    Long usedMemory
    
    Long maxCores
    
    Long maxStorage
}
