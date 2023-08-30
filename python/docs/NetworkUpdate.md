# NetworkUpdate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**display_name** | **str** | Display Name | [optional] 
**labels** | **[str], none_type** | Array of label strings, can be used for filtering. | [optional] 
**description** | **str, none_type** | Description | [optional] 
**cidr** | **str** | CIDR Network | [optional] 
**gateway** | **str** | Network Gateway | [optional] 
**dns_primary** | **str** | Primary DNS Server | [optional] 
**dns_secondary** | **str** | Secondary DNS Server | [optional] 
**vlan_id** | **int** |  | [optional] 
**pool** | **int, none_type** | Network Pool ID | [optional] 
**allow_static_override** | **bool** | Allow IP Override | [optional] 
**assign_public_ip** | **bool** | Assign Public IP | [optional] 
**active** | **bool** | Activate (true) or disable (false) the network | [optional] 
**dhcp_server** | **bool** | DHCP Server enabled network | [optional] 
**network_domain** | [**NetworkCreateNetworkDomain**](NetworkCreateNetworkDomain.md) |  | [optional] 
**search_domains** | **str** | Search Domains | [optional] 
**network_proxy** | [**NetworkCreateNetworkProxy**](NetworkCreateNetworkProxy.md) |  | [optional] 
**appliance_url_proxy_bypass** | **bool** | Bypass Proxy for Appliance URL | [optional] 
**no_proxy** | **str, none_type** | Comma-separated list of ip addresses or name servers to exclude proxy traversal for. Typically locally routable servers are excluded. | [optional] 
**visibility** | **str** | Visibility, private or public. | [optional]  if omitted the server will use the default value of "private"
**config** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | Configuration object. Settings vary by type. | [optional] 
**tenants** | [**[UpdateBlueprintPermissionsRequestResourcePermissionSitesInner]**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) | Array of tenant account ids that are allowed access | [optional] 
**resource_permissions** | [**NetworkCreateResourcePermissions**](NetworkCreateResourcePermissions.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


