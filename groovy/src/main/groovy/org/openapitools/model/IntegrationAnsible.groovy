package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.Creds;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;
import org.openapitools.model.IntegrationAnsibleConfig;

@Canonical
class IntegrationAnsible {
    
    Long id
    
    String name
    
    Boolean enabled
    
    String type
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType integrationType
    
    String url
    
    InlineResponse20040AppDeployInstance serviceKey
    
    Boolean isPlugin
    
    IntegrationAnsibleConfig config
    
    String status
    
    Date statusDate
    
    String statusMessage
    
    String lastSync
    
    String lastSyncDuration
    
    Creds credential
}
