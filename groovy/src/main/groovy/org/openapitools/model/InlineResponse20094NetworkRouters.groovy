package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;
import org.openapitools.model.InlineResponse20094Interfaces;
import org.openapitools.model.InlineResponse20094Type;

@Canonical
class InlineResponse20094NetworkRouters {
    
    Long id
    
    String code
    
    String name
    
    String description
    
    String category
    
    Date dateCreated
    
    Date lastUpdated
    
    String routerType
    
    String status
    
    Boolean enabled
    
    String externalIp
    
    String externalId
    
    String providerId
    
    InlineResponse20094Type type
    
    String networkServer
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType zone
    
    String instance
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType externalNetwork
    
    String site
    
    List<InlineResponse20094Interfaces> interfaces = new ArrayList<InlineResponse20094Interfaces>()
}
