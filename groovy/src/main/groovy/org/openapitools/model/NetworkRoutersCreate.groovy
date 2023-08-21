package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.NetworkRoutersCreateNetworkServer;
import org.openapitools.model.NetworkRoutersCreateSite;
import org.openapitools.model.NetworkRoutersCreateType;
import org.openapitools.model.NetworkRoutersCreateZone;

@Canonical
class NetworkRoutersCreate {
    /* Name */
    String name
    
    NetworkRoutersCreateType type
    
    NetworkRoutersCreateSite site
    /* Can be used to enable / disable the network router (true, false). Default is on */
    Boolean enabled
    
    NetworkRoutersCreateZone zone
    
    NetworkRoutersCreateNetworkServer networkServer
}
