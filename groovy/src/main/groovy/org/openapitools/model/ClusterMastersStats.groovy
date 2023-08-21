package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class ClusterMastersStats {
    
    Long usedStorage
    
    Long reservedStorage
    
    Long maxStorage
    
    Long usedMemory
    
    Long reservedMemory
    
    Long maxMemory
    
    BigDecimal cpuUsage
}
