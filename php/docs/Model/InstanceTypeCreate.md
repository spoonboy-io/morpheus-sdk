# # InstanceTypeCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Instance type name |
**labels** | **string[]** | Array of label strings, can be used for filtering. | [optional]
**description** | **string** | Instance type description | [optional]
**code** | **string** | Instance type code | [optional]
**category** | **string** | Category | [optional]
**visibility** | **string** | Visibility | [optional] [default to 'private']
**featured** | **bool** | Featured | [optional]
**has_settings** | **bool** | Enable Settings | [optional]
**has_auto_scale** | **bool** | Enable Scaling (Horizontal) | [optional]
**has_deployment** | **bool** | Supports Deployments | [optional]
**environment_prefix** | **string** | Environment Prefix, can be used to make exported evars unique. | [optional]
**environment_variables** | [**\OpenAPI\Client\Model\ClusterLayoutCreateEnvironmentVariables[]**](ClusterLayoutCreateEnvironmentVariables.md) | Array of instance type env variables. | [optional]
**price_sets** | [**\OpenAPI\Client\Model\InstanceTypeCreatePriceSets[]**](InstanceTypeCreatePriceSets.md) | Array of price set objects | [optional]
**option_types** | **int[]** | Array of instance type option type IDs | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
