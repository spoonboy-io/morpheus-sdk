package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class ContainerStats {
    
    String ts
    
    Boolean running
    
    BigDecimal userCpuUsage
    
    BigDecimal systemCpuUsage
    
    Integer usedMemory
    
    Integer maxMemory
    
    Integer cacheMemory
    
    Integer maxStorage
    
    Integer usedStorage
    
    Integer readIOPS
    
    BigDecimal writeIOPS
    
    BigDecimal totalIOPS
    
    Integer netTxUsage
    
    Integer netRxUsage
}
