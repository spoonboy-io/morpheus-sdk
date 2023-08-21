package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;

@Canonical
class BackupSettings {
    
    Boolean backupsEnabled
    
    Boolean createBackups
    
    Boolean backupAppliance
    
    InlineResponse20040AppDeployInstance defaultStorageBucket
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType defaultSchedule
    
    Long retentionCount
}
