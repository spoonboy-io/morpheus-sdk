package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;

@Canonical
class InlineResponse20095NetworkRouterNetworkServerIntegration {
    
    Long id
    
    String name
    
    Boolean enabled
    
    String type
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType integrationType
    
    String url
    
    String port
    
    String username
    
    String password
    
    String refType
    
    String refId
    
    Boolean isPlugin
    
    Object config
    
    String status
    
    Date statusDate
    
    String statusMessage
    
    Date lastSync
    
    Long lastSyncDuration
}
