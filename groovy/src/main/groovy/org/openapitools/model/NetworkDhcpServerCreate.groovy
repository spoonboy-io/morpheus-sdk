package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.NetworkDhcpServerCreateConfig;

@Canonical
class NetworkDhcpServerCreate {
    /* Server Address for the DHCP Server */
    String serverIpAddress
    /* Optional lease time for the DHCP Server */
    Long leaseTime = 86400l
    /* Name */
    String name
    
    NetworkDhcpServerCreateConfig config
}
