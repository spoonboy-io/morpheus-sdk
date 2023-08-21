package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class HealthMemory {
    
    Boolean success
    
    Long maxMemory
    
    Long totalMemory
    
    Long freeMemory
    
    Long usedMemory
    
    Long systemMemory
    
    Long committedMemory
    
    Long systemFreeMemory
    
    Long systemSwap
    
    Long systemFreeSwap
    
    Long swapPercent
    
    BigDecimal memoryPercent
    
    BigDecimal systemMemoryPercent
    
    String status
}
