package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.NetworkInterfaceUpdateSuccessNetworkInterface;
import org.openapitools.model.NetworkInterfaceUpdateSuccessServer;

@Canonical
class NetworkInterfaceUpdateSuccess {
    
    NetworkInterfaceUpdateSuccessNetworkInterface networkInterface
    
    String interfaceType
    
    Long netId
    
    NetworkInterfaceUpdateSuccessServer server
}
