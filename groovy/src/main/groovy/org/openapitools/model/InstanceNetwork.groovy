package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20082LoadBalancerInstanceSslCert;

@Canonical
class InstanceNetwork {
    
    Long id
    
    Integer group
    
    String subnet
    
    Boolean dhcpServer
    
    String name
    
    InlineResponse20082LoadBalancerInstanceSslCert pool
}
