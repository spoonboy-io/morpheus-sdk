# # InstanceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**instance** | [**\OpenAPI\Client\Model\InstanceCreateInstance**](InstanceCreateInstance.md) |  |
**zone_id** | **int** | The Cloud ID to provision the instance onto. | [optional]
**evars** | [**\OpenAPI\Client\Model\ApiServersIdMakeManagedServerTags[]**](ApiServersIdMakeManagedServerTags.md) | Environment Variables, an array of objects that have name and value. | [optional]
**copies** | **int** | Number of copies to provision. | [optional] [default to 1]
**layout_size** | **int** | Apply a multiply factor of containers/vms within the instance. | [optional] [default to 1]
**service_plan_options** | **object** | Map of custom options depending on selected service plan. | [optional]
**security_groups** | **object[]** | Key for security group configuration. It should be passed as an array of objects containing the id of the security group to assign the instance to. | [optional]
**volumes** | [**\OpenAPI\Client\Model\InstanceCreateVolume[]**](InstanceCreateVolume.md) | The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of | [optional]
**network_interfaces** | [**\OpenAPI\Client\Model\InstanceCreateNetwork[]**](InstanceCreateNetwork.md) | The networkInterfaces parameter is for network configuration.  The Options API &#x60;/api/options/zoneNetworkOptions?zoneId&#x3D;5&amp;provisionTypeId&#x3D;10&#x60; can be used to see which options are available. | [optional]
**config** | [**AnyOfInstancesConfigAzureInstancesConfigVMWareInstancesConfigGCPInstancesConfigAWSObject**](AnyOfInstancesConfigAzureInstancesConfigVMWareInstancesConfigGCPInstancesConfigAWSObject.md) |  |
**labels** | **string[]** | Array of strings (keywords). | [optional]
**tags** | [**\OpenAPI\Client\Model\ApiServersIdMakeManagedServerTags[]**](ApiServersIdMakeManagedServerTags.md) | Metadata tags, Array of objects having a name and value. | [optional]
**metadata** | [**\OpenAPI\Client\Model\ApiServersIdMakeManagedServerTags[]**](ApiServersIdMakeManagedServerTags.md) | Alias for &#x60;tags&#x60;. | [optional]
**ports** | [**\OpenAPI\Client\Model\InstanceCreatePorts[]**](InstanceCreatePorts.md) | The ports parameter is for port configuration.  The layout may have default ports, which are defined in node types, that are always configured. This parameter will be for additional custom ports to be opened. | [optional]
**task_set_id** | **int** | The Workflow ID to execute. | [optional]
**task_set_name** | **string** | The Workflow Name to execute. | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
