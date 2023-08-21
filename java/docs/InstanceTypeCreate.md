

# InstanceTypeCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Instance type name | 
**labels** | **List&lt;String&gt;** | Array of label strings, can be used for filtering. |  [optional]
**description** | **String** | Instance type description |  [optional]
**code** | **String** | Instance type code |  [optional]
**category** | **String** | Category |  [optional]
**visibility** | [**VisibilityEnum**](#VisibilityEnum) | Visibility |  [optional]
**featured** | **Boolean** | Featured |  [optional]
**hasSettings** | **Boolean** | Enable Settings |  [optional]
**hasAutoScale** | **Boolean** | Enable Scaling (Horizontal) |  [optional]
**hasDeployment** | **Boolean** | Supports Deployments |  [optional]
**environmentPrefix** | **String** | Environment Prefix, can be used to make exported evars unique. |  [optional]
**environmentVariables** | [**List&lt;ClusterLayoutCreateEnvironmentVariables&gt;**](ClusterLayoutCreateEnvironmentVariables.md) | Array of instance type env variables. |  [optional]
**priceSets** | [**List&lt;InstanceTypeCreatePriceSets&gt;**](InstanceTypeCreatePriceSets.md) | Array of price set objects |  [optional]
**optionTypes** | **List&lt;Long&gt;** | Array of instance type option type IDs |  [optional]



## Enum: VisibilityEnum

Name | Value
---- | -----
PRIVATE | &quot;private&quot;
PUBLIC | &quot;public&quot;



