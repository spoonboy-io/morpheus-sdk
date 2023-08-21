package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.CredentialConfig;
import org.openapitools.model.CredentialUser;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;
import org.openapitools.model.InlineResponse20082LoadBalancerInstanceSslCert;

@Canonical
class Credential {
    
    Long id
    
    String name
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType type
    
    String integration
    
    String description
    
    String username
    
    String password
    
    String passwordHash
    
    InlineResponse20082LoadBalancerInstanceSslCert authKey
    
    String authPath
    
    String externalId
    
    String refType
    
    String refId
    
    String category
    
    String scope
    
    String status
    
    String statusMessage
    
    Date statusDate
    
    Boolean enabled
    
    InlineResponse20040AppDeployInstance account
    
    CredentialUser user
    
    Date dateCreated
    
    Date lastUpdated
    
    CredentialConfig config
}
