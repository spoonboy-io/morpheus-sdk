# InstanceTypeLayoutUpdate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Layout name | [optional] 
**labels** | **[str], none_type** |  | [optional] 
**instance_version** | **str** | Version of the layout | [optional] 
**description** | **str** | Layout description | [optional] 
**creatable** | **bool** | Can be used to enable / disable the creatability of the layout. | [optional]  if omitted the server will use the default value of True
**provision_type_code** | **str** | Provision type code | [optional] 
**memory_requirement** | **str** | Memory requirement in megabytes | [optional] 
**has_auto_scale** | **bool** | Can be used to enable / disable the horizontal scaling. | [optional]  if omitted the server will use the default value of False
**supports_convert_to_managed** | **bool** | Can be used to enable / disable the supports convert to managed. | [optional]  if omitted the server will use the default value of False
**container_types** | **[int]** | Array of layout node type IDs | [optional] 
**option_types** | **[int]** | Array of layout option type IDs | [optional] 
**spec_templates** | **[int]** | Array of layout spec template IDs | [optional] 
**environment_variables** | [**[ClusterLayoutCreateEnvironmentVariablesInner]**](ClusterLayoutCreateEnvironmentVariablesInner.md) | The environmentVariables parameter is array of env objects | [optional] 
**price_sets** | [**[InstanceTypeCreatePriceSetsInner]**](InstanceTypeCreatePriceSetsInner.md) | Array of price set objects | [optional] 
**permissions** | [**UpdateLayoutPermissionsRequestInstanceTypeLayoutPermissions**](UpdateLayoutPermissionsRequestInstanceTypeLayoutPermissions.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


