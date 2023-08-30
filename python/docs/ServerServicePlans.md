# ServerServicePlans


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
**max_data_storage** | **int** |  | [optional] 
**custom_cpu** | **bool** |  | [optional] 
**custom_max_memory** | **bool** |  | [optional] 
**custom_max_storage** | **bool** |  | [optional] 
**custom_max_data_storage** | **bool** |  | [optional] 
**custom_cores_per_socket** | **bool** |  | [optional] 
**cores_per_socket** | **int, none_type** |  | [optional] 
**storage_types** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}], none_type** |  | [optional] 
**root_storage_types** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}], none_type** |  | [optional] 
**add_volumes** | **bool** |  | [optional] 
**customize_volume** | **bool** |  | [optional] 
**root_disk_customizable** | **bool** |  | [optional] 
**host_disk_mode** | **str, none_type** |  | [optional] 
**has_datastore** | **str, none_type** |  | [optional] 
**lvm_supported** | **str, none_type** |  | [optional] 
**min_disk** | **str, none_type** |  | [optional] 
**max_disk** | **str, none_type** |  | [optional] 
**datastores** | [**ServerServicePlansDatastores**](ServerServicePlansDatastores.md) |  | [optional] 
**supports_auto_datastore** | **bool** |  | [optional] 
**auto_options** | [**[GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner]**](GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner.md) |  | [optional] 
**cpu_options** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}], none_type** |  | [optional] 
**memory_options** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}], none_type** |  | [optional] 
**root_custom_size_options** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** |  | [optional] 
**custom_size_options** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** |  | [optional] 
**custom_cores** | **bool** |  | [optional] 
**max_disks** | **str, none_type** |  | [optional] 
**memory_size_type** | **str** |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


