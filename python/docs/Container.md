# Container


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | [optional] 
**uuid** | **str** |  | [optional] 
**account_id** | **int** |  | [optional] 
**instance** | [**ContainerInstance**](ContainerInstance.md) |  | [optional] 
**container_type** | [**ContainerContainerType**](ContainerContainerType.md) |  | [optional] 
**container_type_set** | [**ContainerContainerTypeSet**](ContainerContainerTypeSet.md) |  | [optional] 
**server** | [**ContainerInstance**](ContainerInstance.md) |  | [optional] 
**cloud** | [**ContainerInstance**](ContainerInstance.md) |  | [optional] 
**name** | **str** |  | [optional] 
**ip** | **str** |  | [optional] 
**internal_ip** | **str** |  | [optional] 
**internal_hostname** | **str** |  | [optional] 
**external_hostname** | **str** |  | [optional] 
**external_domain** | **str** |  | [optional] 
**external_fqdn** | **str** |  | [optional] 
**ports** | [**[ContainerPort], none_type**](ContainerPort.md) |  | [optional] 
**plan** | [**ContainerPlan**](ContainerPlan.md) |  | [optional] 
**date_created** | **datetime** |  | [optional] 
**last_updated** | **datetime** |  | [optional] 
**stats_enabled** | **bool** |  | [optional] 
**status** | **str** |  | [optional] 
**user_status** | **str** |  | [optional] 
**environment_prefix** | **str, none_type** |  | [optional] 
**stats** | [**ContainerStats**](ContainerStats.md) |  | [optional] 
**runtime_info** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** |  | [optional] 
**container_version** | **str, none_type** |  | [optional] 
**repository_image** | **str, none_type** |  | [optional] 
**plan_category** | **str, none_type** |  | [optional] 
**hostname** | **str** |  | [optional] 
**domain_name** | **str, none_type** |  | [optional] 
**volume_created** | **bool** |  | [optional] 
**container_created** | **bool** |  | [optional] 
**max_storage** | **int** |  | [optional] 
**max_memory** | **int** |  | [optional] 
**max_cores** | **int** |  | [optional] 
**max_cpu** | **int, none_type** |  | [optional] 
**available_actions** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}], none_type** |  | [optional] 
**config_group** | **str, none_type** |  | [optional] 
**config_id** | **str, none_type** |  | [optional] 
**config_role** | **str, none_type** |  | [optional] 
**hourly_cost** | **float** |  | [optional] 
**hourly_price** | **float** |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


