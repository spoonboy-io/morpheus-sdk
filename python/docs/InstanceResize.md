# InstanceResize


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**instance** | [**InstanceResizeInstance**](InstanceResizeInstance.md) |  | [optional] 
**service_plan_options** | [**ServicePlanOptions**](ServicePlanOptions.md) |  | [optional] 
**volumes** | [**[InstanceCreateVolume]**](InstanceCreateVolume.md) | Can be used to grow just the logical volume of the instance instead of choosing a plan | [optional] 
**delete_original_volumes** | **bool** | Delete the original volumes after resizing. (Amazon only) | [optional]  if omitted the server will use the default value of False
**network_interfaces** | [**[InstanceCreateNetwork]**](InstanceCreateNetwork.md) | Key for network configuration. Include id to update an existing interface. The existing interfaces and their id can be retrieved with the hosts API. | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


