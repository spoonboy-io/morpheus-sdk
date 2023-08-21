package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class ServicePlanConfigRanges {
    
    String minStorage
    
    String maxStorage
    
    String minPerDiskSize
    
    String maxPerDiskSize
    
    BigDecimal minMemory
    
    BigDecimal maxMemory
    
    String minCores
    
    String maxCores
    
    String minSockets
    
    String maxSockets
    
    String minCoresPerSocket
    
    String maxCoresPerSocket
}
