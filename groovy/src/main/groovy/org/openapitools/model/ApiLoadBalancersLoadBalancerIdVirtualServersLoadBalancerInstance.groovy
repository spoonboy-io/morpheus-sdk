package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance {
    /* VIP Name */
    String vipName
    /* Description */
    String description
    /* VIP Address */
    String vipAddress
    /* VIP Port */
    String vipPort
    /* VIP Protocol */
    String vipProtocol
    /* VIP Hostname */
    String vipHostname
    /* SSL Client Certificate ID */
    Long sslCert
    /* SSL Server Certificate ID */
    Long sslServerCert
    /* Configuration object with parameters that vary by type. */
    Object config
}
