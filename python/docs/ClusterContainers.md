# ClusterContainers


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | [optional] 
**uuid** | **str** |  | [optional] 
**account_id** | **int** |  | [optional] 
**instance** | **str, none_type** |  | [optional] 
**container_type** | [**ClusterContainersContainerType**](ClusterContainersContainerType.md) |  | [optional] 
**container_type_set** | [**ClusterContainersContainerTypeSet**](ClusterContainersContainerTypeSet.md) |  | [optional] 
**server** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**cloud** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**name** | **str** |  | [optional] 
**ip** | **str** |  | [optional] 
**internal_ip** | **str** |  | [optional] 
**internal_hostname** | **str** |  | [optional] 
**external_hostname** | **str** |  | [optional] 
**external_domain** | **str** |  | [optional] 
**external_fqdn** | **str** |  | [optional] 
**ports** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}]** |  | [optional] 
**plan** | [**ClusterContainersPlan**](ClusterContainersPlan.md) |  | [optional] 
**date_created** | **datetime, none_type** |  | [optional] 
**last_updated** | **datetime** |  | [optional] 
**stats_enabled** | **bool** |  | [optional] 
**status** | **str** |  | [optional] 
**user_status** | **str** |  | [optional] 
**environment_prefix** | **str, none_type** |  | [optional] 
**config_group** | **str, none_type** |  | [optional] 
**config_id** | **str, none_type** |  | [optional] 
**config_role** | **str, none_type** |  | [optional] 
**stats** | [**ClusterContainersStats**](ClusterContainersStats.md) |  | [optional] 
**runtime_info** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** |  | [optional] 
**container_version** | **str, none_type** |  | [optional] 
**repository_image** | **str, none_type** |  | [optional] 
**plan_category** | **str, none_type** |  | [optional] 
**hostname** | **str, none_type** |  | [optional] 
**domain_name** | **str, none_type** |  | [optional] 
**volume_created** | **bool** |  | [optional] 
**container_created** | **bool** |  | [optional] 
**max_storage** | **str, none_type** |  | [optional] 
**max_memory** | **str, none_type** |  | [optional] 
**max_cores** | **str, none_type** |  | [optional] 
**max_cpu** | **str, none_type** |  | [optional] 
**hourly_price** | **float** |  | [optional] 
**available_actions** | [**[ClusterContainersAvailableActionsInner]**](ClusterContainersAvailableActionsInner.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


