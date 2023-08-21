package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.Creds;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;

@Canonical
class IntegrationDockerRepo {
    
    Long id
    
    String name
    
    Boolean enabled
    
    String type
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType integrationType
    
    String url
    
    String username
    
    String password
    
    String passwordHash
    
    Boolean isPlugin
    
    Object config
    
    String status
    
    Date statusDate
    
    String statusMessage
    
    String lastSync
    
    String lastSyncDuration
    
    Creds credential
}
