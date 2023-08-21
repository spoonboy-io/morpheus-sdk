package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;
import org.openapitools.model.InlineResponse20095NetworkRouterFirewall;
import org.openapitools.model.InlineResponse20095NetworkRouterNetworkServer;
import org.openapitools.model.InlineResponse20095NetworkRouterPermissions;
import org.openapitools.model.InlineResponse20095NetworkRouterType;

@Canonical
class InlineResponse20095NetworkRouter {
    
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
    
    InlineResponse20095NetworkRouterType type
    
    InlineResponse20095NetworkRouterNetworkServer networkServer
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType zone
    
    String instance
    
    String externalNetwork
    
    String site
    
    List<Object> interfaces = new ArrayList<Object>()
    
    InlineResponse20095NetworkRouterFirewall firewall
    
    List<Object> routes = new ArrayList<Object>()
    
    List<Object> nats = new ArrayList<Object>()
    
    InlineResponse20095NetworkRouterPermissions permissions
}
