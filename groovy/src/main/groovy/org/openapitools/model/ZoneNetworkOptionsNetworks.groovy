package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ZoneNetworkOptionsNetworks {
    
    String id
    
    String name
    
    Boolean dhcpServer
    
    Boolean allowStaticOverride
    
    String pool
}
