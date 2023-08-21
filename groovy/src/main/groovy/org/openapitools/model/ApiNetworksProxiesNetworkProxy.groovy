package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ApiNetworksProxiesNetworkProxyAccount;

@Canonical
class ApiNetworksProxiesNetworkProxy {
    /* Name */
    String name
    /* Proxy Host */
    String proxyHost
    /* Proxy Port */
    String proxyPort
    /* Proxy Username */
    String proxyUser
    /* Proxy Password */
    String proxyPassword
    /* Proxy Domain */
    String proxyDomain
    /* Proxy Workstation */
    String proxyWorkstation
    /* Visibility */
    String visibility = "private"
    
    ApiNetworksProxiesNetworkProxyAccount account
}
