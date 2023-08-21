package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class NetworkConfig {
    
    String vlanIDs
    
    String connectedGateway
    
    String subnetIpManagementType
    
    String subnetIpServerId
    
    String dhcpRange
    
    String subnetDhcpServerAddress
    
    String subnetDhcpLeaseTime
}
