package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ApiLoadBalancersLoadBalancerIdProfilesLoadBalancerProfile {
    /* Name */
    String name
    /* Description */
    String description
    /* Service Type */
    String serviceType
    /* Configuration object with parameters that vary by type. */
    Object config
}
