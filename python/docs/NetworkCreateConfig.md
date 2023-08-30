# NetworkCreateConfig

Configuration object. Settings vary by type.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**resource_group_id** | **str** | Resource Group Name | [optional] 
**subnet_name** | **str** | Subnet Name | [optional] 
**subnet_cidr** | **str** | The subnet&#39;s address range in CIDR notation (e.g. 192.168.1.0/24). It must be contained by the address space of the virtual network. | [optional] 
**availability_zone** | **str** | Availability Zone Name | [optional] 
**cidr** | **str** | Network CIDR | [optional] 
**assign_public_ip** | **bool** | Assign public IPs by default. | [optional] 
**zone_pool** | [**NetworkTypeGcpConfigZonePool**](NetworkTypeGcpConfigZonePool.md) |  | [optional] 
**mtu** | **str** | GCP MTU | [optional]  if omitted the server will use the default value of "1460"
**auto_create** | **bool** | Auto create subnets | [optional]  if omitted the server will use the default value of True
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


