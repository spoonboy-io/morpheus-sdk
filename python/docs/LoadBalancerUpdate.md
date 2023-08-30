# LoadBalancerUpdate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Name | [optional] 
**description** | **str** | Description | [optional] 
**enabled** | **bool** | Activate (true) or disable (false) | [optional] 
**config** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | Configuration object with parameters that vary by load balancer type. | [optional] 
**visibility** | **str** | private or public | [optional]  if omitted the server will use the default value of "public"
**tenants** | [**[GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerOwner]**](GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerOwner.md) | Array of tenant account ids that are allowed access | [optional] 
**resource_permission** | [**LoadBalancerCreateResourcePermission**](LoadBalancerCreateResourcePermission.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


