# # NetworkCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Name |
**display_name** | **string** | Display Name | [optional]
**labels** | **string[]** | Array of label strings, can be used for filtering. | [optional]
**description** | **string** | Description | [optional]
**site** | [**\OpenAPI\Client\Model\NetworkCreateSite**](NetworkCreateSite.md) |  |
**zone** | [**\OpenAPI\Client\Model\NetworkCreateZone**](NetworkCreateZone.md) |  |
**type** | [**\OpenAPI\Client\Model\NetworkCreateType**](NetworkCreateType.md) |  | [optional]
**ipv4_enabled** | **bool** |  | [optional]
**ipv6_enabled** | **bool** |  | [optional]
**cidr** | **string** | CIDR Network | [optional]
**gateway** | **string** | Network Gateway | [optional]
**dns_primary** | **string** | Primary DNS Server | [optional]
**dns_secondary** | **string** | Secondary DNS Server | [optional]
**gateway_i_pv6** | **string** | IPv6 Network Gateway | [optional]
**netmask_i_pv6** | **string** |  | [optional]
**dns_primary_i_pv6** | **string** | Primary IPv6 DNS Server | [optional]
**dns_secondary_i_pv6** | **string** | Secondary IPv6 DNS Server | [optional]
**cidr_i_pv6** | **string** | IPv6 Network CIDR | [optional]
**vlan_id** | **int** |  | [optional]
**pool** | **int** | Network Pool ID | [optional]
**pool_i_pv6** | **int** | IPv6 Network Pool ID | [optional]
**allow_static_override** | **bool** | Allow IP Override | [optional]
**assign_public_ip** | **bool** | Assign Public IP | [optional]
**active** | **bool** | Activate (true) or disable (false) the network | [optional]
**dhcp_server** | **bool** | DHCP Server enabled network | [optional]
**dhcp_server_i_pv6** | **bool** | IPv6 DHCP Server enabled network | [optional]
**network_domain** | [**\OpenAPI\Client\Model\NetworkCreateNetworkDomain**](NetworkCreateNetworkDomain.md) |  | [optional]
**search_domains** | **string** | Search Domains | [optional]
**network_proxy** | [**\OpenAPI\Client\Model\NetworkCreateNetworkProxy**](NetworkCreateNetworkProxy.md) |  | [optional]
**appliance_url_proxy_bypass** | **bool** | Bypass Proxy for Appliance URL | [optional]
**no_proxy** | **string** | Comma-separated list of ip addresses or name servers to exclude proxy traversal for. Typically locally routable servers are excluded. | [optional]
**visibility** | **string** | Visibility, private or public. | [optional] [default to 'private']
**config** | [**AnyOfNetworkTypeAzureConfigNetworkTypeAwsConfigNetworkTypeGcpConfigObject**](AnyOfNetworkTypeAzureConfigNetworkTypeAwsConfigNetworkTypeGcpConfigObject.md) | Configuration object. Settings vary by type. | [optional]
**tenants** | [**\OpenAPI\Client\Model\ApiBlueprintsIdUpdatePermissionsResourcePermissionSites[]**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of tenant account ids that are allowed access | [optional]
**resource_permissions** | [**\OpenAPI\Client\Model\NetworkCreateResourcePermissions**](NetworkCreateResourcePermissions.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
