package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class NetworkRouterRouteCreate {
    /* Route name */
    String name
    /* Route description */
    String description
    /* Can be used to enable / disable the route (true, false). Default is off */
    Boolean enabled = false
    /* Can be used to set as default route (true, false). Default is off */
    Boolean defaultRoute = false
    /* Source address or range */
    String source
    /* Destination address or range */
    String destination
    /* MTU */
    String networkMtu
}
