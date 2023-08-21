package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class InstanceCreateNetworkNetwork {
    /* id of the network to be used. A network group can be specified instead by prefixing its ID with `networkGroup-`. */
    String id
}
