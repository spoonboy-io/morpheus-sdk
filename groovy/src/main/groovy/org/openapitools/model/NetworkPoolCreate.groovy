package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.NetworkPoolCreateIpRanges;
import org.openapitools.model.NetworkPoolCreateType;

@Canonical
class NetworkPoolCreate {
    /* Name */
    String name
    
    NetworkPoolCreateType type
    /* Array of IP range objects. Type 'morpheus' expects startAddress and endAddress. Type 'morpheusipv6' expects a cidrIPv6. */
    List<NetworkPoolCreateIpRanges> ipRanges = new ArrayList<NetworkPoolCreateIpRanges>()
    /* Configuration object with parameters that vary by pool type. */
    Object config
}
