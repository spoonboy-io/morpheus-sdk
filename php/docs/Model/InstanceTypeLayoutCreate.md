# # InstanceTypeLayoutCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Layout name |
**labels** | **string[]** |  | [optional]
**instance_version** | **string** | Version of the layout |
**description** | **string** | Layout description | [optional]
**creatable** | **bool** | Can be used to enable / disable the creatability of the layout. | [optional] [default to true]
**provision_type_code** | **string** | Provision type code |
**memory_requirement** | **string** | Memory requirement in megabytes | [optional]
**has_auto_scale** | **bool** | Can be used to enable / disable the horizontal scaling. | [optional] [default to false]
**supports_convert_to_managed** | **bool** | Can be used to enable / disable the supports convert to managed. | [optional] [default to false]
**container_types** | **int[]** | Array of layout node type IDs | [optional]
**option_types** | **int[]** | Array of layout option type IDs | [optional]
**spec_templates** | **int[]** | Array of layout spec template IDs | [optional]
**environment_variables** | [**\OpenAPI\Client\Model\ClusterLayoutCreateEnvironmentVariables[]**](ClusterLayoutCreateEnvironmentVariables.md) | The environmentVariables parameter is array of env objects | [optional]
**price_sets** | [**\OpenAPI\Client\Model\InstanceTypeCreatePriceSets[]**](InstanceTypeCreatePriceSets.md) | Array of price set objects | [optional]
**permissions** | [**\OpenAPI\Client\Model\ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions**](ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
