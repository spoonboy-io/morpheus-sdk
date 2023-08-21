package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse200106NetworkPoolIpRanges;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20094Network;

@Canonical
class InlineResponse200106NetworkPool {
    
    Long id
    
    InlineResponse20094Network type
    
    InlineResponse20040AppDeployInstance account
    
    String category
    
    String code
    
    String name
    
    String displayName
    
    String internalId
    
    String externalId
    
    String dnsDomain
    
    String dnsSearchPath
    
    String hostPrefix
    
    String httpProxy
    
    List<String> dnsServers = new ArrayList<String>()
    
    List<String> dnsSuffixList = new ArrayList<String>()
    
    Boolean dhcpServer
    
    String dhcpIp
    
    String gateway
    
    String netmask
    
    String subnetAddress
    
    Long ipCount
    
    Long freeCount
    
    Boolean poolEnabled
    
    String tftpServer
    
    String bootFile
    
    String refType
    
    String refId
    
    String parentType
    
    String parentId
    
    String poolGroup
    
    List<InlineResponse200106NetworkPoolIpRanges> ipRanges = new ArrayList<InlineResponse200106NetworkPoolIpRanges>()
}
