package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class NetworkStaticRouteCreate {
    /* Source CIDR */
    String source
    /* Destination address */
    String destination
}
