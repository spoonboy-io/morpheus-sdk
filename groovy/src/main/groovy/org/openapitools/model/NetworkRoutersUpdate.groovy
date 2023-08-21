package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.NetworkRoutersUpdateNetworkServer;
import org.openapitools.model.NetworkRoutersUpdateSite;
import org.openapitools.model.NetworkRoutersUpdateType;
import org.openapitools.model.NetworkRoutersUpdateZone;

@Canonical
class NetworkRoutersUpdate {
    /* Name */
    String name
    
    NetworkRoutersUpdateType type
    
    NetworkRoutersUpdateSite site
    /* Can be used to enable / disable the network router (true, false). Default is on */
    Boolean enabled
    
    NetworkRoutersUpdateZone zone
    
    NetworkRoutersUpdateNetworkServer networkServer
}
