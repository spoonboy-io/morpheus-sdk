# InstanceServicePlan


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | [optional] 
**name** | **str** |  | [optional] 
**value** | **int** |  | [optional] 
**code** | **str** |  | [optional] 
**max_storage** | **int** |  | [optional] 
**max_memory** | **int** |  | [optional] 
**max_cpu** | **int** |  | [optional] 
**max_cores** | **int** |  | [optional] 
**custom_cpu** | **bool** |  | [optional] 
**custom_max_memory** | **bool** |  | [optional] 
**custom_max_storage** | **bool** |  | [optional] 
**custom_max_data_storage** | **bool** |  | [optional] 
**custom_cores_per_socket** | **bool** |  | [optional] 
**cores_per_socket** | **int** |  | [optional] 
**storage_types** | [**[InstanceServicePlanStorageType]**](InstanceServicePlanStorageType.md) |  | [optional] 
**root_storage_types** | [**[InstanceServicePlanStorageType]**](InstanceServicePlanStorageType.md) |  | [optional] 
**add_volumes** | **bool** |  | [optional] 
**customize_volume** | **bool** |  | [optional] 
**root_disk_customizable** | **bool** |  | [optional] 
**no_disks** | **bool** |  | [optional] 
**has_datastore** | **bool** |  | [optional] 
**min_disk** | **int** |  | [optional] 
**max_disk** | **str, none_type** |  | [optional] 
**lvm_supported** | **bool** |  | [optional] 
**datastores** | [**InstanceServicePlanDatastores**](InstanceServicePlanDatastores.md) |  | [optional] 
**supports_auto_datastore** | **bool** |  | [optional] 
**auto_options** | [**[GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner]**](GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner.md) |  | [optional] 
**cpu_options** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}], none_type** |  | [optional] 
**core_options** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}], none_type** |  | [optional] 
**memory_options** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}], none_type** |  | [optional] 
**root_custom_size_options** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}, none_type** |  | [optional] 
**custom_size_options** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}, none_type** |  | [optional] 
**custom_cores** | **bool** |  | [optional] 
**max_disks** | **str, none_type** |  | [optional] 
**memory_size_type** | **str** |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


