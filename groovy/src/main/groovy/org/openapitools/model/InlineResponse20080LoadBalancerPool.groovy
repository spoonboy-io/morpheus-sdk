package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancer;

@Canonical
class InlineResponse20080LoadBalancerPool {
    
    Long id
    
    InlineResponse20079LoadBalancerMonitorLoadBalancer loadBalancer
    
    String name
    
    String category
    
    String visibility
    
    String description
    
    String internalId
    
    String externalId
    
    Boolean enabled
    
    String vipSticky
    
    String vipBalance
    
    String allowNat
    
    String allowSnat
    
    String vipClientIpMode
    
    String vipServerIpMode
    
    Long minActive
    
    String minInService
    
    String minUpMonitor
    
    String minUpAction
    
    String maxQueueDepth
    
    String maxQueueTime
    
    Long numberActive
    
    Long numberInService
    
    Long healthScore
    
    Long performanceScore
    
    Long healthPenalty
    
    Long securityPenalty
    
    Long errorPenalty
    
    String downAction
    
    String rampTime
    
    String port
    
    String portType
    
    String status
    
    List<InlineResponse20040AppDeployInstance> nodes = new ArrayList<InlineResponse20040AppDeployInstance>()
    
    List<InlineResponse20040AppDeployInstance> monitors = new ArrayList<InlineResponse20040AppDeployInstance>()
    
    List<Object> members = new ArrayList<Object>()
    
    Object config
    
    String createdBy
    
    Date dateCreated
    
    Date lastUpdated
}
