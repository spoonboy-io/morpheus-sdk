package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class NetworkDhcpRelayCreate {
    
    String name
    
    List<String> serverIpAddresses = new ArrayList<String>()
}
