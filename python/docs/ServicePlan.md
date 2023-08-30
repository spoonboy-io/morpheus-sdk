# ServicePlan


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | [optional] 
**name** | **str** |  | [optional] 
**code** | **str** |  | [optional] 
**active** | **bool** |  | [optional] 
**sort_order** | **int** |  | [optional] 
**description** | **str** |  | [optional] 
**max_storage** | **float** |  | [optional] 
**max_memory** | **float** |  | [optional] 
**max_cpu** | **float, none_type** |  | [optional] 
**max_cores** | **float, none_type** |  | [optional] 
**max_disks** | **float, none_type** |  | [optional] 
**cores_per_socket** | **float** |  | [optional] 
**custom_cpu** | **bool** |  | [optional] 
**custom_cores** | **bool** |  | [optional] 
**custom_max_storage** | **bool, none_type** |  | [optional] 
**custom_max_data_storage** | **bool, none_type** |  | [optional] 
**custom_max_memory** | **bool, none_type** |  | [optional] 
**add_volumes** | **bool, none_type** |  | [optional] 
**memory_option_source** | **str, none_type** |  | [optional] 
**cpu_option_source** | **str, none_type** |  | [optional] 
**date_created** | **datetime** |  | [optional] 
**last_updated** | **datetime** |  | [optional] 
**region_code** | **str, none_type** |  | [optional] 
**visibility** | **str** |  | [optional] 
**editable** | **bool** |  | [optional] 
**provision_type** | [**GuidanceVmwareSizingPlanBeforeActionProvisionType**](GuidanceVmwareSizingPlanBeforeActionProvisionType.md) |  | [optional] 
**tenants** | **str** |  | [optional] 
**price_sets** | [**[GuidanceVmwareSizingPlanBeforeActionPriceSetsInner], none_type**](GuidanceVmwareSizingPlanBeforeActionPriceSetsInner.md) |  | [optional] 
**config** | [**ServicePlanConfig**](ServicePlanConfig.md) |  | [optional] 
**zones** | [**[GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInnerNetwork]**](GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInnerNetwork.md) |  | [optional] 
**permissions** | [**ServicePlanPermissions**](ServicePlanPermissions.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


