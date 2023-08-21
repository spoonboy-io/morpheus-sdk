package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ApiLoadBalancerPoolsLoadBalancerPoolIdNodesLoadBalancerNode {
    /* Name */
    String name
    /* Description */
    String description
    /* IP Address */
    String ipAddress
    /* Port number */
    Integer port
    /* Weight applied balance algoritm */
    Integer weight
    /* Configuration object with parameters that vary by type. */
    Object config
}
