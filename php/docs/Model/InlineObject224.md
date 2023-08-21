# # InlineObject224

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**server** | [**\OpenAPI\Client\Model\ApiServersIdResizeServer**](ApiServersIdResizeServer.md) |  | [optional]
**service_plan_options** | [**\OpenAPI\Client\Model\ApiServersIdResizeServicePlanOptions**](ApiServersIdResizeServicePlanOptions.md) |  | [optional]
**volumes** | [**\OpenAPI\Client\Model\InstanceCreateVolume[]**](InstanceCreateVolume.md) | List of volumes with their new sizes. | [optional]
**delete_original_volumes** | **bool** | Delete the original volumes after resizing. (Amazon only) | [optional] [default to false]
**network_interfaces** | [**\OpenAPI\Client\Model\InstanceCreateNetwork[]**](InstanceCreateNetwork.md) | Key for network configurations. Include id to update an existing interface. | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
