package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class GuidanceVmwareSizingResourceStats {
    
    String ts
    
    Long freeMemory
    
    Long usedMemory
    
    Long freeSwap
    
    Long usedSwap
    
    Long cpuIdleTime
    
    Long cpuSystemTime
    
    Long cpuUserTime
    
    Long cpuTotalTime
    
    BigDecimal cpuUsage
    
    Long usedStorage
    
    Long reservedStorage
    
    Long maxStorage
    
    Long netTxUsage
    
    Long netRxUsage
    
    Long networkBandwidth
}
