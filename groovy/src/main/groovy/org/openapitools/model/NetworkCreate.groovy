package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.AnyOfnetworkTypeAzureConfignetworkTypeAwsConfignetworkTypeGcpConfigobject;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;
import org.openapitools.model.NetworkCreateNetworkDomain;
import org.openapitools.model.NetworkCreateNetworkProxy;
import org.openapitools.model.NetworkCreateResourcePermissions;
import org.openapitools.model.NetworkCreateSite;
import org.openapitools.model.NetworkCreateType;
import org.openapitools.model.NetworkCreateZone;

@Canonical
class NetworkCreate {
    /* Name */
    String name
    /* Display Name */
    String displayName
    /* Array of label strings, can be used for filtering. */
    List<String> labels = new ArrayList<String>()
    /* Description */
    String description
    
    NetworkCreateSite site
    
    NetworkCreateZone zone
    
    NetworkCreateType type
    
    Boolean ipv4Enabled
    
    Boolean ipv6Enabled
    /* CIDR Network */
    String cidr
    /* Network Gateway */
    String gateway
    /* Primary DNS Server */
    String dnsPrimary
    /* Secondary DNS Server */
    String dnsSecondary
    /* IPv6 Network Gateway */
    String gatewayIPv6
    
    String netmaskIPv6
    /* Primary IPv6 DNS Server */
    String dnsPrimaryIPv6
    /* Secondary IPv6 DNS Server */
    String dnsSecondaryIPv6
    /* IPv6 Network CIDR */
    String cidrIPv6
    
    Long vlanId
    /* Network Pool ID */
    Long pool
    /* IPv6 Network Pool ID */
    Long poolIPv6
    /* Allow IP Override */
    Boolean allowStaticOverride
    /* Assign Public IP */
    Boolean assignPublicIp
    /* Activate (true) or disable (false) the network */
    Boolean active
    /* DHCP Server enabled network */
    Boolean dhcpServer
    /* IPv6 DHCP Server enabled network */
    Boolean dhcpServerIPv6
    
    NetworkCreateNetworkDomain networkDomain
    /* Search Domains */
    String searchDomains
    
    NetworkCreateNetworkProxy networkProxy
    /* Bypass Proxy for Appliance URL */
    Boolean applianceUrlProxyBypass
    /* Comma-separated list of ip addresses or name servers to exclude proxy traversal for. Typically locally routable servers are excluded. */
    String noProxy
    /* Visibility, private or public. */
    String visibility = VisibilityEnum.PRIVATE
    /* Configuration object. Settings vary by type. */
    AnyOfnetworkTypeAzureConfignetworkTypeAwsConfignetworkTypeGcpConfigobject config
    /* Array of tenant account ids that are allowed access */
    List<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites> tenants = new ArrayList<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites>()
    
    NetworkCreateResourcePermissions resourcePermissions
}
