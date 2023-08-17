package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.Arrays;

@Canonical
class AppStats {
    
    Long usedMemory
    
    Long maxMemory
    
    Long usedStorage
    
    Long maxStorage
    
    Long running
    
    Long total
    
    BigDecimal cpuUsage
    
    Long instanceCount
    
    List<Long> instanceDayCount
    
    Long instanceDayCountTotal
}
