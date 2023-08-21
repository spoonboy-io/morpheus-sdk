# # InstanceResize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**instance** | [**\OpenAPI\Client\Model\InstanceResizeInstance**](InstanceResizeInstance.md) |  | [optional]
**service_plan_options** | [**\OpenAPI\Client\Model\ServicePlanOptions**](ServicePlanOptions.md) |  | [optional]
**volumes** | [**\OpenAPI\Client\Model\InstanceCreateVolume[]**](InstanceCreateVolume.md) | Can be used to grow just the logical volume of the instance instead of choosing a plan | [optional]
**delete_original_volumes** | **bool** | Delete the original volumes after resizing. (Amazon only) | [optional] [default to false]
**network_interfaces** | [**\OpenAPI\Client\Model\InstanceCreateNetwork[]**](InstanceCreateNetwork.md) | Key for network configuration. Include id to update an existing interface. The existing interfaces and their id can be retrieved with the hosts API. | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
