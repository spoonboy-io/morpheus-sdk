package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ImageBuildsConfigNetwork;

@Canonical
class ImageBuildsConfigNetworkInterfaces {
    
    String ipMode
    
    Boolean primaryInterface
    
    Boolean showNetworkPoolLabel
    
    Boolean showNetworkDhcpLabel
    
    ImageBuildsConfigNetwork network
    
    Long networkInterfaceTypeId
    
    String networkInterfaceTypeIdName
}
