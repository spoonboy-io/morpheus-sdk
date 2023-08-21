package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.NetworkInterfaceUpdateSuccessServerZoneNetworkServerType;

@Canonical
class NetworkInterfaceUpdateSuccessServerZoneNetworkServer {
    
    Long id
    
    String name
    
    String description
    
    NetworkInterfaceUpdateSuccessServerZoneNetworkServerType type
    
    String status
    
    Boolean enabled
    
    String visibility
}
