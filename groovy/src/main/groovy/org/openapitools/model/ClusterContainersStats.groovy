package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class ClusterContainersStats {
    
    Date ts
    
    Boolean running
    
    BigDecimal userCpuUsage
    
    BigDecimal systemCpuUsage
    
    Long usedMemory
    
    Long maxMemory
    
    Long cacheMemory
    
    Long maxStorage
    
    Long usedStorage
    
    Long readIOPS
    
    Long writeIOPS
    
    Long totalIOPS
    
    Long netTxUsage
    
    Long netRxUsage
}
