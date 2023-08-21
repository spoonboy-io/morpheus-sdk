package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class DashboardInstanceStats {
    
    BigDecimal maxCpu
    
    BigDecimal maxCores
    
    BigDecimal cpuUsage
    
    BigDecimal cpuUsageAverage
    
    BigDecimal cpuUsagePeak
    
    BigDecimal usedMemory
    
    BigDecimal maxMemory
    
    BigDecimal usedStorage
    
    BigDecimal maxStorage
    
    BigDecimal running
    
    BigDecimal total
    
    BigDecimal totalContainers
}
