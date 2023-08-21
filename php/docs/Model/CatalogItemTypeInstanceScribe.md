# # CatalogItemTypeInstanceScribe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**group** | [**\OpenAPI\Client\Model\CatalogItemTypeInstanceScribeGroup**](CatalogItemTypeInstanceScribeGroup.md) |  |
**cloud** | [**\OpenAPI\Client\Model\CatalogItemTypeInstanceScribeCloud**](CatalogItemTypeInstanceScribeCloud.md) |  |
**type** | **string** | The type of instance by code we want to fetch. |
**name** | **string** | Name of the instance to be created. |
**config** | [**AnyOfInstancesConfigAzureInstancesConfigVMWareInstancesConfigGCPInstancesConfigAWSObject**](AnyOfInstancesConfigAzureInstancesConfigVMWareInstancesConfigGCPInstancesConfigAWSObject.md) |  |
**volumes** | [**\OpenAPI\Client\Model\InstanceCreateVolume[]**](InstanceCreateVolume.md) | The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of |
**host_name** | **string** | Hostname of the instance to be created.  Can be the same as the instance name. | [optional]
**environment** | **string** | Environment code | [optional]
**layout** | [**\OpenAPI\Client\Model\CatalogItemTypeInstanceScribeLayout**](CatalogItemTypeInstanceScribeLayout.md) |  |
**plan** | [**\OpenAPI\Client\Model\CatalogItemTypeInstanceScribePlan**](CatalogItemTypeInstanceScribePlan.md) |  |
**version** | **string** | Version of the layout to create. | [optional]
**evars** | [**\OpenAPI\Client\Model\ApiServersIdMakeManagedServerTags[]**](ApiServersIdMakeManagedServerTags.md) | Environment Variables, an array of objects that have name and value. | [optional]
**service_plan_options** | [**\OpenAPI\Client\Model\ServicePlanOptions**](ServicePlanOptions.md) |  | [optional]
**security_groups** | [**\OpenAPI\Client\Model\CatalogItemTypeInstanceScribeSecurityGroups[]**](CatalogItemTypeInstanceScribeSecurityGroups.md) | Key for security group configuration. It should be passed as an array of objects containing the id of the security group to assign the instance to. | [optional]
**network_interfaces** | [**\OpenAPI\Client\Model\InstanceCreateNetwork[]**](InstanceCreateNetwork.md) | The networkInterfaces parameter is for network configuration.  The Options API &#x60;/api/options/zoneNetworkOptions?zoneId&#x3D;5&amp;provisionTypeId&#x3D;10&#x60; can be used to see which options are available. | [optional]
**labels** | **string[]** | Array of strings (keywords). | [optional]
**tags** | [**\OpenAPI\Client\Model\ApiServersIdMakeManagedServerTags[]**](ApiServersIdMakeManagedServerTags.md) | Metadata tags, Array of objects having a name and value. | [optional]
**metadata** | [**\OpenAPI\Client\Model\ApiServersIdMakeManagedServerTags[]**](ApiServersIdMakeManagedServerTags.md) | Alias for &#x60;tags&#x60;. | [optional]
**ports** | [**\OpenAPI\Client\Model\CatalogItemTypeInstanceScribePorts[]**](CatalogItemTypeInstanceScribePorts.md) | The ports parameter is for port configuration.  The layout may have default ports, which are defined in node types, that are always configured. This parameter will be for additional custom ports to be opened. | [optional]
**task_set_id** | **int** | The Workflow ID to execute. | [optional]
**task_set_name** | **string** | The Workflow Name to execute. | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
