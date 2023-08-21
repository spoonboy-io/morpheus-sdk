package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.AppStateInputProviders;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;
import org.openapitools.model.SubnetResourcePermission;

@Canonical
class Subnet {
    
    Long id
    
    String code
    
    String name
    
    List<String> labels = new ArrayList<String>()
    
    Boolean active
    
    String description
    
    String externalId
    
    String uniqueId
    
    String addressPrefix
    
    String cidr
    
    String gateway
    
    String netmask
    
    String subnetAddress
    
    String tftpServer
    
    String bootFile
    
    String pool
    
    Boolean dhcpServer
    
    Boolean hasFloatingIps
    
    String dhcpIp
    
    String dnsPrimary
    
    String dnsSecondary
    
    String dhcpStart
    
    String dhcpEnd
    
    String dhcpRange
    
    String networkProxy
    
    String networkDomain
    
    String searchDomains
    
    Boolean defaultNetwork
    
    Boolean assignPublicIp
    
    String visibility
    
    AppStateInputProviders status
    
    InlineResponse20040AppDeployInstance network
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType type
    
    InlineResponse20040AppDeployInstance account
    
    List<Object> securityGroups = new ArrayList<Object>()
    
    List<InlineResponse20040AppDeployInstance> tenants = new ArrayList<InlineResponse20040AppDeployInstance>()
    
    SubnetResourcePermission resourcePermission
}
