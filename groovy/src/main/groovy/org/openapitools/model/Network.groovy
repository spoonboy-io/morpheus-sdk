package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse200108NetworkFloatingIpNetworkDomain;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20082LoadBalancerInstanceSslCert;
import org.openapitools.model.NetworkConfig;
import org.openapitools.model.NetworkNetworkProxy;
import org.openapitools.model.NetworkOwner;
import org.openapitools.model.NetworkType;
import org.openapitools.model.NetworkZone;

@Canonical
class Network {
    /* Network ID */
    Long id
    /* Name */
    String name
    /* Network Display Name */
    String displayName
    
    List<String> labels = new ArrayList<String>()
    
    NetworkZone zone
    
    NetworkType type
    
    NetworkOwner owner
    /* Network Code */
    String code
    
    Boolean ipv4Enabled
    
    Boolean ipv6Enabled
    /* Network Category */
    String category
    
    String interfaceName
    
    String bridgeName
    
    String bridgeInterface
    /* Description */
    String description
    
    String externalId
    
    String internalId
    
    String uniqueId
    
    String externalType
    
    String refUrl
    
    String refType
    
    Long refId
    
    Long vlanId
    
    String vswitchName
    
    Boolean dhcpServer
    
    String dhcpIp
    
    Boolean dhcpServerIPv6
    /* Network Gateway */
    String gateway
    
    String netmask
    
    String broadcast
    
    String subnetAddress
    /* Primary DNS Server */
    String dnsPrimary
    /* Secondary DNS Server */
    String dnsSecondary
    /* Network CIDR */
    String cidr
    /* IPv6 Network Gateway */
    String gatewayIPv6
    
    String netmaskIPv6
    /* Primary IPv6 DNS Server */
    String dnsPrimaryIPv6
    /* Secondary IPv6 DNS Server */
    String dnsSecondaryIPv6
    /* IPv6 Network CIDR */
    String cidrIPv6
    
    String tftpServer
    
    String bootFile
    
    String switchId
    
    String fabricId
    
    String networkRole
    
    String status
    
    String availabilityZone
    
    Object pool
    
    Object poolIPv6
    
    NetworkNetworkProxy networkProxy
    
    InlineResponse200108NetworkFloatingIpNetworkDomain networkDomain
    
    String searchDomains
    
    String prefixLength
    
    String visibility
    
    Boolean enableAdmin
    
    Boolean active
    
    Boolean defaultNetwork
    
    Boolean assignPublicIp
    
    String noProxy
    
    Boolean applianceUrlProxyBypass
    
    InlineResponse20082LoadBalancerInstanceSslCert zonePool
    
    Boolean allowStaticOverride
    
    NetworkConfig config
    
    List<InlineResponse20040AppDeployInstance> tenants = new ArrayList<InlineResponse20040AppDeployInstance>()
}
