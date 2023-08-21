# MorpheusApi.InstanceTypeLayoutCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Layout name | 
**labels** | **[String]** |  | [optional] 
**instanceVersion** | **String** | Version of the layout | 
**description** | **String** | Layout description | [optional] 
**creatable** | **Boolean** | Can be used to enable / disable the creatability of the layout. | [optional] [default to true]
**provisionTypeCode** | **String** | Provision type code | 
**memoryRequirement** | **String** | Memory requirement in megabytes | [optional] 
**hasAutoScale** | **Boolean** | Can be used to enable / disable the horizontal scaling. | [optional] [default to false]
**supportsConvertToManaged** | **Boolean** | Can be used to enable / disable the supports convert to managed. | [optional] [default to false]
**containerTypes** | **[Number]** | Array of layout node type IDs | [optional] 
**optionTypes** | **[Number]** | Array of layout option type IDs | [optional] 
**specTemplates** | **[Number]** | Array of layout spec template IDs | [optional] 
**environmentVariables** | [**[ClusterLayoutCreateEnvironmentVariables]**](ClusterLayoutCreateEnvironmentVariables.md) | The environmentVariables parameter is array of env objects | [optional] 
**priceSets** | [**[InstanceTypeCreatePriceSets]**](InstanceTypeCreatePriceSets.md) | Array of price set objects | [optional] 
**permissions** | [**ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions**](ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions.md) |  | [optional] 


