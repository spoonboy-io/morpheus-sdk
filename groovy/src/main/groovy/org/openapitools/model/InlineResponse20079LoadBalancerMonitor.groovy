package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancer;

@Canonical
class InlineResponse20079LoadBalancerMonitor {
    
    Long id
    
    InlineResponse20079LoadBalancerMonitorLoadBalancer loadBalancer
    
    String name
    
    String code
    
    String category
    
    String visibility
    
    String description
    
    String monitorType
    
    Long monitorInterval
    
    Long monitorTimeout
    
    String sendData
    
    String sendVersion
    
    String sendType
    
    String receiveData
    
    String receiveCode
    
    String disabledData
    
    String monitorUsername
    
    String monitorPassword
    
    String monitorDestination
    
    Boolean monitorReverse
    
    Boolean monitorTransparent
    
    Boolean monitorAdaptive
    
    String aliasAddress
    
    Long aliasPort
    
    String internalId
    
    String externalId
    
    String monitorSource
    
    String status
    
    String statusMessage
    
    Date statusDate
    
    Boolean enabled
    
    Long maxRetry
    
    Long fallCount
    
    Long riseCount
    
    String dataLength
    
    Object config
    
    String createdBy
    
    Date dateCreated
    
    Date lastUpdated
}
