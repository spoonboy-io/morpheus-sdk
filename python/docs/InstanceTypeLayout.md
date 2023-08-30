# InstanceTypeLayout


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | [optional] 
**instance_type** | [**GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInnerNetwork**](GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInnerNetwork.md) |  | [optional] 
**account** | [**ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert**](ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert.md) |  | [optional] 
**code** | **str** |  | [optional] 
**name** | **str** |  | [optional] 
**labels** | **[str], none_type** | Array of label strings, can be used for filtering. | [optional] 
**instance_version** | **str** |  | [optional] 
**description** | **str, none_type** |  | [optional] 
**creatable** | **bool** |  | [optional] 
**memory_requirement** | **int, none_type** |  | [optional] 
**sort_order** | **int** |  | [optional] 
**supports_convert_to_managed** | **bool, none_type** |  | [optional] 
**provision_type** | [**ProvisionType**](ProvisionType.md) |  | [optional] 
**task_sets** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}], none_type** |  | [optional] 
**container_types** | [**[ContainerType]**](ContainerType.md) |  | [optional] 
**mounts** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}], none_type** |  | [optional] 
**ports** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}], none_type** |  | [optional] 
**option_types** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}], none_type** |  | [optional] 
**environment_variables** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}], none_type** |  | [optional] 
**price_sets** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}], none_type** |  | [optional] 
**spec_templates** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}], none_type** |  | [optional] 
**tfvar_secret** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}], none_type** |  | [optional] 
**permissions** | [**InstanceTypeLayoutPermissions**](InstanceTypeLayoutPermissions.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


