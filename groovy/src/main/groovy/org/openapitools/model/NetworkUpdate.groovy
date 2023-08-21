package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;
import org.openapitools.model.NetworkCreateNetworkDomain;
import org.openapitools.model.NetworkCreateNetworkProxy;
import org.openapitools.model.NetworkCreateResourcePermissions;

@Canonical
class NetworkUpdate {
    /* Display Name */
    String displayName
    /* Array of label strings, can be used for filtering. */
    List<String> labels = new ArrayList<String>()
    /* Description */
    String description
    /* CIDR Network */
    String cidr
    /* Network Gateway */
    String gateway
    /* Primary DNS Server */
    String dnsPrimary
    /* Secondary DNS Server */
    String dnsSecondary
    
    Long vlanId
    /* Network Pool ID */
    Long pool
    /* Allow IP Override */
    Boolean allowStaticOverride
    /* Assign Public IP */
    Boolean assignPublicIp
    /* Activate (true) or disable (false) the network */
    Boolean active
    /* DHCP Server enabled network */
    Boolean dhcpServer
    
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
    Object config
    /* Array of tenant account ids that are allowed access */
    List<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites> tenants = new ArrayList<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites>()
    
    NetworkCreateResourcePermissions resourcePermissions
}
