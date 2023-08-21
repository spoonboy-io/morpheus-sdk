package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class ClustersWorkerStats {
    
    Long usedStorage
    
    Long maxStorage
    
    Long usedMemory
    
    Long maxMemory
    
    Long usedCpu
    
    BigDecimal cpuUsage
    
    BigDecimal cpuUsagePeak
    
    BigDecimal cpuUsageAvg
}
