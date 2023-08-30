# Subnet


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | [optional] 
**code** | **str, none_type** |  | [optional] 
**name** | **str** |  | [optional] 
**labels** | **[str], none_type** |  | [optional] 
**active** | **bool** |  | [optional] 
**description** | **str, none_type** |  | [optional] 
**external_id** | **str** |  | [optional] 
**unique_id** | **str, none_type** |  | [optional] 
**address_prefix** | **str, none_type** |  | [optional] 
**cidr** | **str** |  | [optional] 
**gateway** | **str, none_type** |  | [optional] 
**netmask** | **str** |  | [optional] 
**subnet_address** | **str** |  | [optional] 
**tftp_server** | **str, none_type** |  | [optional] 
**boot_file** | **str, none_type** |  | [optional] 
**pool** | **str, none_type** |  | [optional] 
**dhcp_server** | **bool** |  | [optional] 
**has_floating_ips** | **bool** |  | [optional] 
**dhcp_ip** | **str, none_type** |  | [optional] 
**dns_primary** | **str, none_type** |  | [optional] 
**dns_secondary** | **str, none_type** |  | [optional] 
**dhcp_start** | **str** |  | [optional] 
**dhcp_end** | **str** |  | [optional] 
**dhcp_range** | **str, none_type** |  | [optional] 
**network_proxy** | **str, none_type** |  | [optional] 
**network_domain** | **str, none_type** |  | [optional] 
**search_domains** | **str, none_type** |  | [optional] 
**default_network** | **bool** |  | [optional] 
**assign_public_ip** | **bool** |  | [optional] 
**visibility** | **str** |  | [optional] 
**status** | [**AppStateInputProvidersInner**](AppStateInputProvidersInner.md) |  | [optional] 
**network** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**type** | [**ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType.md) |  | [optional] 
**account** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**security_groups** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}]** |  | [optional] 
**tenants** | [**[ListDeploys200ResponseAllOfAppDeploysInnerInstance]**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**resource_permission** | [**SubnetResourcePermission**](SubnetResourcePermission.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


