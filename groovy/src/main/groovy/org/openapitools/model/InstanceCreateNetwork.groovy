package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InstanceCreateNetworkNetwork;

@Canonical
class InstanceCreateNetwork {
    
    InstanceCreateNetworkNetwork network
    /* The id of type of the network interface. */
    Long networkInterfaceTypeId
    /* The ip address. Not applicable when using DHCP or IP Pools. */
    String ipAddress
    /* The interface id. Applicable when resizing and you want to identify an interface to update that already exists. */
    Long id
}
