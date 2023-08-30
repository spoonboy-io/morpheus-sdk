# InstanceTypeCreate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Instance type name | 
**labels** | **[str], none_type** | Array of label strings, can be used for filtering. | [optional] 
**description** | **str** | Instance type description | [optional] 
**code** | **str** | Instance type code | [optional] 
**category** | **str** | Category | [optional] 
**visibility** | **str** | Visibility | [optional]  if omitted the server will use the default value of "private"
**featured** | **bool** | Featured | [optional] 
**has_settings** | **bool** | Enable Settings | [optional] 
**has_auto_scale** | **bool** | Enable Scaling (Horizontal) | [optional] 
**has_deployment** | **bool** | Supports Deployments | [optional] 
**environment_prefix** | **str** | Environment Prefix, can be used to make exported evars unique. | [optional] 
**environment_variables** | [**[ClusterLayoutCreateEnvironmentVariablesInner]**](ClusterLayoutCreateEnvironmentVariablesInner.md) | Array of instance type env variables. | [optional] 
**price_sets** | [**[InstanceTypeCreatePriceSetsInner]**](InstanceTypeCreatePriceSetsInner.md) | Array of price set objects | [optional] 
**option_types** | **[int]** | Array of instance type option type IDs | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


