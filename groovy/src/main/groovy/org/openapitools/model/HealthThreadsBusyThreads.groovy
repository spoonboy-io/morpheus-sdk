package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;

@Canonical
class HealthThreadsBusyThreads {
    
    Long threadId
    
    String name
    
    Long cpuTime
    
    Long blockedTime
    
    String lockName
    
    Long lockOwnerId
    
    String lockOwnerName
    
    String state
    
    Long waitedCount
    
    Long waitedTime
    
    Boolean isInNative
    
    Boolean isSuspended
    
    List<Object> lockedMonitors = new ArrayList<Object>()
    
    List<Object> lockedSynchronizers = new ArrayList<Object>()
    
    String lockInfo
    
    String currentLines
    
    BigDecimal cpuPercent
}
