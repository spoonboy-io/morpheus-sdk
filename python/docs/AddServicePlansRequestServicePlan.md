# AddServicePlansRequestServicePlan


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Service plan name | 
**code** | **str** | Service plan code, must be unique | 
**max_storage** | **float** | Max storage size in bytes | 
**max_memory** | **float** | Max memory size in bytes | 
**provision_type** | [**[AddServicePlansRequestServicePlanProvisionTypeInner]**](AddServicePlansRequestServicePlanProvisionTypeInner.md) |  | 
**description** | **str** | Service plan description | [optional] 
**editable** | **bool** | Can be used to enable / disable the editability of the service plan. | [optional]  if omitted the server will use the default value of True
**max_cores** | **float** | Max cores | [optional] 
**max_disks** | **float** | Max disks allowed | [optional] 
**custom_cores** | **bool** | Can be used to enable / disable customizable cores | [optional]  if omitted the server will use the default value of False
**custom_max_storage** | **bool** | Can be used to enable / disable customizable storage | [optional]  if omitted the server will use the default value of False
**custom_max_data_storage** | **bool** | Can be used to enable / disable customizable extra volumes. | [optional]  if omitted the server will use the default value of False
**custom_max_memory** | **bool** | Can be used to enable / disable customizable memory. | [optional]  if omitted the server will use the default value of False
**add_volumes** | **bool** | Can be used to enable / disable ability to add volumes | [optional]  if omitted the server will use the default value of False
**sort_order** | **float** | Sort order | [optional] 
**price_sets** | **[int]** | List of price sets to include in service plan | [optional] 
**config** | [**AddServicePlansRequestServicePlanConfig**](AddServicePlansRequestServicePlanConfig.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


