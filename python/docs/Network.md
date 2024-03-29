# Network


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | Network ID | [optional] 
**name** | **str** | Name | [optional] 
**display_name** | **str** | Network Display Name | [optional] 
**labels** | **[str], none_type** |  | [optional] 
**zone** | [**NetworkZone**](NetworkZone.md) |  | [optional] 
**type** | [**NetworkType**](NetworkType.md) |  | [optional] 
**owner** | [**NetworkOwner**](NetworkOwner.md) |  | [optional] 
**code** | **str, none_type** | Network Code | [optional] 
**ipv4_enabled** | **bool** |  | [optional] 
**ipv6_enabled** | **bool** |  | [optional] 
**category** | **str, none_type** | Network Category | [optional] 
**interface_name** | **str, none_type** |  | [optional] 
**bridge_name** | **str, none_type** |  | [optional] 
**bridge_interface** | **str, none_type** |  | [optional] 
**description** | **str, none_type** | Description | [optional] 
**external_id** | **str, none_type** |  | [optional] 
**internal_id** | **str, none_type** |  | [optional] 
**unique_id** | **str, none_type** |  | [optional] 
**external_type** | **str** |  | [optional] 
**ref_url** | **str, none_type** |  | [optional] 
**ref_type** | **str** |  | [optional] 
**ref_id** | **int** |  | [optional] 
**vlan_id** | **int, none_type** |  | [optional] 
**vswitch_name** | **str, none_type** |  | [optional] 
**dhcp_server** | **bool** |  | [optional] 
**dhcp_ip** | **str, none_type** |  | [optional] 
**dhcp_server_ipv6** | **bool** |  | [optional] 
**gateway** | **str, none_type** | Network Gateway | [optional] 
**netmask** | **str, none_type** |  | [optional] 
**broadcast** | **str, none_type** |  | [optional] 
**subnet_address** | **str, none_type** |  | [optional] 
**dns_primary** | **str, none_type** | Primary DNS Server | [optional] 
**dns_secondary** | **str, none_type** | Secondary DNS Server | [optional] 
**cidr** | **str** | Network CIDR | [optional] 
**gateway_ipv6** | **str, none_type** | IPv6 Network Gateway | [optional] 
**netmask_ipv6** | **str, none_type** |  | [optional] 
**dns_primary_ipv6** | **str, none_type** | Primary IPv6 DNS Server | [optional] 
**dns_secondary_ipv6** | **str, none_type** | Secondary IPv6 DNS Server | [optional] 
**cidr_ipv6** | **str, none_type** | IPv6 Network CIDR | [optional] 
**tftp_server** | **str, none_type** |  | [optional] 
**boot_file** | **str, none_type** |  | [optional] 
**switch_id** | **str, none_type** |  | [optional] 
**fabric_id** | **str, none_type** |  | [optional] 
**network_role** | **str, none_type** |  | [optional] 
**status** | **str** |  | [optional] 
**availability_zone** | **str, none_type** |  | [optional] 
**pool** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}, none_type** |  | [optional] 
**pool_ipv6** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}, none_type** |  | [optional] 
**network_proxy** | [**NetworkNetworkProxy**](NetworkNetworkProxy.md) |  | [optional] 
**network_domain** | [**GetAllNetworkFloatingIps200ResponseAllOfNetworkFloatingIpsInnerNetworkDomain**](GetAllNetworkFloatingIps200ResponseAllOfNetworkFloatingIpsInnerNetworkDomain.md) |  | [optional] 
**search_domains** | **str, none_type** |  | [optional] 
**prefix_length** | **str, none_type** |  | [optional] 
**visibility** | **str** |  | [optional] 
**enable_admin** | **bool** |  | [optional] 
**active** | **bool** |  | [optional] 
**default_network** | **bool** |  | [optional] 
**assign_public_ip** | **bool** |  | [optional] 
**no_proxy** | **str, none_type** |  | [optional] 
**appliance_url_proxy_bypass** | **bool** |  | [optional] 
**zone_pool** | [**ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert**](ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert.md) |  | [optional] 
**allow_static_override** | **bool** |  | [optional] 
**config** | [**NetworkConfig**](NetworkConfig.md) |  | [optional] 
**tenants** | [**[ListDeploys200ResponseAllOfAppDeploysInnerInstance]**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


