# ZoneType


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | [optional] 
**name** | **str** |  | [optional] 
**code** | **str** |  | [optional] 
**enabled** | **bool** |  | [optional] 
**provision** | **bool** |  | [optional] 
**auto_capacity** | **bool** |  | [optional] 
**migration_target** | **bool** |  | [optional] 
**has_datastores** | **bool** |  | [optional] 
**has_networks** | **bool** |  | [optional] 
**has_resource_pools** | **bool** |  | [optional] 
**has_security_groups** | **bool** |  | [optional] 
**has_containers** | **bool** |  | [optional] 
**has_bare_metal** | **bool** |  | [optional] 
**has_services** | **bool** |  | [optional] 
**has_functions** | **bool** |  | [optional] 
**has_jobs** | **bool** |  | [optional] 
**has_discovery** | **bool** |  | [optional] 
**has_cloud_init** | **bool** |  | [optional] 
**has_folders** | **bool** |  | [optional] 
**has_floating_ips** | **bool** |  | [optional] 
**has_marketplace** | **bool** |  | [optional] 
**can_create_resource_pools** | **bool** |  | [optional] 
**can_delete_resource_pools** | **bool** |  | [optional] 
**can_create_datastores** | **bool** |  | [optional] 
**can_create_networks** | **bool** |  | [optional] 
**can_choose_container_mode** | **bool** |  | [optional] 
**provision_requires_resource_pool** | **bool** |  | [optional] 
**supports_distributed_worker** | **bool** |  | [optional] 
**cloud** | **str** |  | [optional] 
**provision_types** | **[int]** |  | [optional] 
**zone_instance_type_layout_id** | **int** |  | [optional] 
**server_types** | [**[ZoneTypeServerTypesInner]**](ZoneTypeServerTypesInner.md) |  | [optional] 
**option_types** | [**[ZoneTypeOptionTypesInner]**](ZoneTypeOptionTypesInner.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


