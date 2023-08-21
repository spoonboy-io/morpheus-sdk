package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.Creds;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;
import org.openapitools.model.IntegrationSaltMasterConfig;

@Canonical
class IntegrationSaltMaster {
    
    Long id
    
    String name
    
    Boolean enabled
    
    String type
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType integrationType
    
    String url
    
    String port
    
    String username
    
    String password
    
    String passwordHash
    
    String path
    
    String version
    
    String windowsVersion
    
    String repoUrl
    
    String serviceMode
    
    Boolean isPlugin
    
    IntegrationSaltMasterConfig config
    
    String status
    
    Date statusDate
    
    String statusMessage
    
    String lastSync
    
    String lastSyncDuration
    
    Creds credential
}
