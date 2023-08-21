package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancer;

@Canonical
class InlineResponse20081LoadBalancerProfile {
    
    Long id
    
    InlineResponse20079LoadBalancerMonitorLoadBalancer loadBalancer
    
    String name
    
    String category
    
    String serviceType
    
    String serviceTypeDisplay
    
    String visibility
    
    String description
    
    String internalId
    
    String externalId
    
    String proxyType
    
    String redirectRewrite
    
    String persistenceType
    
    String sslEnabled
    
    String sslCert
    
    String accountCertificate
    
    Boolean enabled
    
    String redirectUrl
    
    Boolean insertXforwardedFor
    
    String persistenceCookieName
    
    String persistenceExpiresIn
    
    Boolean editable
    
    Object config
    
    String createdBy
    
    Date dateCreated
    
    Date lastUpdated
}
