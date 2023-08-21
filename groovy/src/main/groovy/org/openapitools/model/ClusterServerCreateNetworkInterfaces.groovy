package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ClusterServerCreateNetwork;

@Canonical
class ClusterServerCreateNetworkInterfaces {
    
    ClusterServerCreateNetwork network
    /* The id of type of the network interface. */
    Long networkInterfaceTypeId
    /* The ip address. Not applicable when using DHCP or IP Pools. */
    String ipAddress
}
