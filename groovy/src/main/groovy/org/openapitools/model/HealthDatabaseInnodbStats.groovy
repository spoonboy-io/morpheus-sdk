package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class HealthDatabaseInnodbStats {
    
    Long largeMemory
    
    Long dictionaryMemory
    
    Long bufferPoolSize
    
    Long freeBuffers
    
    Long databasePages
    
    Long oldPages
    
    Long pendingReads
    
    BigDecimal insertsPerSecond
    
    BigDecimal updatesPerSecond
    
    BigDecimal deletesPerSecond
    
    BigDecimal readsPerSecond
    
    Long bufferHitRate
}
