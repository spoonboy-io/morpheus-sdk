package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ApiLoadBalancersLoadBalancerIdMonitorsLoadBalancerMonitor {
    /* Name */
    String name
    /* Description */
    String description
    
    String monitorType
    
    Long monitorTimeout
    /* Configuration object with parameters that vary by type. */
    Object config
}
