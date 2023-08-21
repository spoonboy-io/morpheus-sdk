package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancer;
import org.openapitools.model.InlineResponse20082LoadBalancerInstanceSslCert;

@Canonical
class InlineResponse20082LoadBalancerInstance {
    
    Long id
    
    InlineResponse20079LoadBalancerMonitorLoadBalancer loadBalancer
    
    String instance
    
    String description
    
    String internalId
    
    String externalId
    
    Date dateCreated
    
    Date lastUpdated
    
    Boolean active
    
    Boolean sticky
    
    String sslEnabled
    
    Boolean externalAddress
    
    String backendPort
    
    String vipType
    
    String vipAddress
    
    String vipHostname
    
    String vipProtocol
    
    String vipScheme
    
    String vipMode
    
    String vipName
    
    Long vipPort
    
    String vipSticky
    
    String vipBalance
    
    String servicePort
    
    String sourceAddress
    
    InlineResponse20082LoadBalancerInstanceSslCert sslCert
    
    String sslMode
    
    String sslRedirectMode
    
    Boolean vipShared
    
    String vipDirectAddress
    
    String serverName
    
    String poolName
    
    Boolean removing
    
    String vipSource
    
    String extraConfig
    
    String serviceAccess
    
    String networkId
    
    String subnetId
    
    String externalPortId
    
    String status
    
    String vipStatus
}
