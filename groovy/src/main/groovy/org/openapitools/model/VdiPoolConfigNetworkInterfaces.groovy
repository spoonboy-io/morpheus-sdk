package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.VdiPoolConfigNetwork;

@Canonical
class VdiPoolConfigNetworkInterfaces {
    
    Boolean primaryInterface
    
    VdiPoolConfigNetwork network
    
    String ipMode
    
    Boolean showNetworkPoolLabel
    
    Boolean showNetworkDhcpLabel
}
