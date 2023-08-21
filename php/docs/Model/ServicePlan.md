# # ServicePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | [optional]
**name** | **string** |  | [optional]
**code** | **string** |  | [optional]
**active** | **bool** |  | [optional]
**sort_order** | **int** |  | [optional]
**description** | **string** |  | [optional]
**max_storage** | **float** |  | [optional]
**max_memory** | **float** |  | [optional]
**max_cpu** | **float** |  | [optional]
**max_cores** | **float** |  | [optional]
**max_disks** | **float** |  | [optional]
**cores_per_socket** | **float** |  | [optional]
**custom_cpu** | **bool** |  | [optional]
**custom_cores** | **bool** |  | [optional]
**custom_max_storage** | **bool** |  | [optional]
**custom_max_data_storage** | **bool** |  | [optional]
**custom_max_memory** | **bool** |  | [optional]
**add_volumes** | **bool** |  | [optional]
**memory_option_source** | **string** |  | [optional]
**cpu_option_source** | **string** |  | [optional]
**date_created** | [**\DateTime**](\DateTime.md) |  | [optional]
**last_updated** | [**\DateTime**](\DateTime.md) |  | [optional]
**region_code** | **string** |  | [optional]
**visibility** | **string** |  | [optional]
**editable** | **bool** |  | [optional]
**provision_type** | [**\OpenAPI\Client\Model\GuidanceVmwareSizingPlanBeforeActionProvisionType**](GuidanceVmwareSizingPlanBeforeActionProvisionType.md) |  | [optional]
**tenants** | **string** |  | [optional]
**price_sets** | [**\OpenAPI\Client\Model\GuidanceVmwareSizingPlanBeforeActionPriceSets[]**](GuidanceVmwareSizingPlanBeforeActionPriceSets.md) |  | [optional]
**config** | [**\OpenAPI\Client\Model\ServicePlanConfig**](ServicePlanConfig.md) |  | [optional]
**zones** | [**\OpenAPI\Client\Model\InlineResponse20094Network[]**](InlineResponse20094Network.md) |  | [optional]
**permissions** | [**\OpenAPI\Client\Model\ServicePlanPermissions**](ServicePlanPermissions.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
