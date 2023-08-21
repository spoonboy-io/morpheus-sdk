package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.HealthThreadsBusyThreads;

@Canonical
class HealthThreads {
    
    Boolean success
    
    List<Object> threadList = new ArrayList<Object>()
    
    List<HealthThreadsBusyThreads> busyThreads = new ArrayList<HealthThreadsBusyThreads>()
    
    List<Object> blockedThreads = new ArrayList<Object>()
    
    List<Object> runningThreads = new ArrayList<Object>()
    
    Long totalCpuTime
    
    Long totalThreads
    
    Long runningWebThreads
    
    String status
}
