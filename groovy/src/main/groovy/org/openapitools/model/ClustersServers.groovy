package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ClustersComputeServerType;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;

@Canonical
class ClustersServers {
    
    Long id
    
    String name
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType typeSet
    
    ClustersComputeServerType computeServerType
}
