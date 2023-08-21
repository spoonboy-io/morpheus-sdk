package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool {
    /* Name */
    String name
    /* Description */
    String description
    /* Balance Algorithm */
    String vipBalance
    /* Min Active Members */
    Long minActive
    /* Configuration object with parameters that vary by type. */
    Object config
}
