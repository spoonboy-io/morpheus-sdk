# UpdateHostResizeRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**server** | [**UpdateHostResizeRequestServer**](UpdateHostResizeRequestServer.md) |  | [optional] 
**service_plan_options** | [**UpdateHostResizeRequestServicePlanOptions**](UpdateHostResizeRequestServicePlanOptions.md) |  | [optional] 
**volumes** | [**[InstanceCreateVolume]**](InstanceCreateVolume.md) | List of volumes with their new sizes. | [optional] 
**delete_original_volumes** | **bool** | Delete the original volumes after resizing. (Amazon only) | [optional]  if omitted the server will use the default value of False
**network_interfaces** | [**[InstanceCreateNetwork]**](InstanceCreateNetwork.md) | Key for network configurations. Include id to update an existing interface. | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


