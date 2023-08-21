

# InstanceTypeLayoutCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Layout name | 
**labels** | **List&lt;String&gt;** |  |  [optional]
**instanceVersion** | **String** | Version of the layout | 
**description** | **String** | Layout description |  [optional]
**creatable** | **Boolean** | Can be used to enable / disable the creatability of the layout. |  [optional]
**provisionTypeCode** | **String** | Provision type code | 
**memoryRequirement** | **String** | Memory requirement in megabytes |  [optional]
**hasAutoScale** | **Boolean** | Can be used to enable / disable the horizontal scaling. |  [optional]
**supportsConvertToManaged** | **Boolean** | Can be used to enable / disable the supports convert to managed. |  [optional]
**containerTypes** | **List&lt;Long&gt;** | Array of layout node type IDs |  [optional]
**optionTypes** | **List&lt;Long&gt;** | Array of layout option type IDs |  [optional]
**specTemplates** | **List&lt;Long&gt;** | Array of layout spec template IDs |  [optional]
**environmentVariables** | [**List&lt;ClusterLayoutCreateEnvironmentVariables&gt;**](ClusterLayoutCreateEnvironmentVariables.md) | The environmentVariables parameter is array of env objects |  [optional]
**priceSets** | [**List&lt;InstanceTypeCreatePriceSets&gt;**](InstanceTypeCreatePriceSets.md) | Array of price set objects |  [optional]
**permissions** | [**ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions**](ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions.md) |  |  [optional]



