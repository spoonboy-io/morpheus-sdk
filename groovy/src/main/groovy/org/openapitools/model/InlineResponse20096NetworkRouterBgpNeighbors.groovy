package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20096Config;

@Canonical
class InlineResponse20096NetworkRouterBgpNeighbors {
    
    Long id
    
    String ipAddress
    
    String forwardingAddress
    
    String protocolAddress
    
    String remoteAs
    
    Long weight
    
    Long keepAlive
    
    Long holdDown
    
    String password
    
    String routeFilteringType
    
    String routeFilteringIn
    
    String routeFilteringOut
    
    Boolean bfdEnabled
    
    Long bfdInterval
    
    Long bfdMultiple
    
    Boolean allowAsIn
    
    Long hopLimit
    
    String restartMode
    
    String providerId
    
    String syncSource
    
    String internalId
    
    String externalId
    
    String refType
    
    String refId
    
    InlineResponse20096Config config
    
    Date dateCreated
    
    Date lastUpdated
}
