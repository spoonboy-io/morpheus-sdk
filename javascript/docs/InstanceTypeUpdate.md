# MorpheusApi.InstanceTypeUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Instance type name | [optional] 
**description** | **String** | Instance type description | [optional] 
**labels** | **[String]** | Array of label strings, can be used for filtering. | [optional] 
**code** | **String** | Instance type code | [optional] 
**category** | **String** | Category | [optional] 
**visibility** | **String** | Visibility | [optional] [default to &#39;private&#39;]
**featured** | **Boolean** | Featured | [optional] 
**hasSettings** | **Boolean** | Enable Settings | [optional] 
**hasAutoScale** | **Boolean** | Enable Scaling (Horizontal) | [optional] 
**hasDeployment** | **Boolean** | Supports Deployments | [optional] 
**environmentPrefix** | **String** | Environment Prefix, can be used to make exported evars unique. | [optional] 
**environmentVariables** | [**[ClusterLayoutCreateEnvironmentVariables]**](ClusterLayoutCreateEnvironmentVariables.md) | Array of instance type env variables. | [optional] 
**priceSets** | [**[InstanceTypeCreatePriceSets]**](InstanceTypeCreatePriceSets.md) | Array of price set objects | [optional] 
**optionTypes** | **[Number]** | Array of instance type option type IDs | [optional] 



## Enum: VisibilityEnum


* `private` (value: `"private"`)

* `public` (value: `"public"`)




