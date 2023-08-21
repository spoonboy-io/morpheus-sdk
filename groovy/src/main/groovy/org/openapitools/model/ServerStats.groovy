package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class ServerStats {
    
    Date ts
    
    Long freeSwap
    
    Long usedSwap
    
    Long cpuIdleTime
    
    Long cpuSystemTime
    
    Long cpuUserTime
    
    Long cpuTotalTime
    
    Long maxMemory
    
    Long usedMemory
    
    Long maxStorage
    
    Long usedStorage
    
    Long reservedStorage
    
    BigDecimal cpuUsage
    
    Long freeMemory
    
    Long netTxUsage
    
    Long netRxUsage
    
    Long networkBandwidth
}
