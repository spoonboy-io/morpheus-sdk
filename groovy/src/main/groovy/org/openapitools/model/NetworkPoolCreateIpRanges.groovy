package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class NetworkPoolCreateIpRanges {
    /* Starting IP Address */
    String startAddress
    /* Ending IP Address */
    String endAddress
    /* IPv6 Network CIDR */
    String cidrIPv6
}
