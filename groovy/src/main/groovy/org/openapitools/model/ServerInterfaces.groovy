package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20082LoadBalancerInstanceSslCert;
import org.openapitools.model.PriceSetVolumeType;

@Canonical
class ServerInterfaces {
    
    Long id
    
    String refType
    
    String refId
    
    String name
    
    String internalId
    
    String externalId
    
    String uniqueId
    
    String publicIpAddress
    
    String publicIpv6Address
    
    String ipAddress
    
    String ipv6Address
    
    String ipSubnet
    
    String ipv6Subnet
    
    String description
    
    Boolean dhcp
    
    Boolean active
    
    Boolean poolAssigned
    
    Boolean primaryInterface
    
    InlineResponse20040AppDeployInstance network
    
    String subnet
    
    String networkGroup
    
    String networkPosition
    
    InlineResponse20082LoadBalancerInstanceSslCert networkPool
    
    String networkDomain
    
    PriceSetVolumeType type
    
    String ipMode
    
    String macAddress
}
