

# ApiServicePlansServicePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Service plan name | 
**code** | **String** | Service plan code, must be unique | 
**description** | **String** | Service plan description |  [optional]
**editable** | **Boolean** | Can be used to enable / disable the editability of the service plan. |  [optional]
**maxStorage** | **BigDecimal** | Max storage size in bytes | 
**maxMemory** | **BigDecimal** | Max memory size in bytes | 
**maxCores** | **BigDecimal** | Max cores |  [optional]
**maxDisks** | **BigDecimal** | Max disks allowed |  [optional]
**provisionType** | [**List&lt;ApiServicePlansServicePlanProvisionType&gt;**](ApiServicePlansServicePlanProvisionType.md) |  | 
**customCores** | **Boolean** | Can be used to enable / disable customizable cores |  [optional]
**customMaxStorage** | **Boolean** | Can be used to enable / disable customizable storage |  [optional]
**customMaxDataStorage** | **Boolean** | Can be used to enable / disable customizable extra volumes. |  [optional]
**customMaxMemory** | **Boolean** | Can be used to enable / disable customizable memory. |  [optional]
**addVolumes** | **Boolean** | Can be used to enable / disable ability to add volumes |  [optional]
**sortOrder** | **BigDecimal** | Sort order |  [optional]
**priceSets** | **List&lt;Long&gt;** | List of price sets to include in service plan |  [optional]
**config** | [**ApiServicePlansServicePlanConfig**](ApiServicePlansServicePlanConfig.md) |  |  [optional]



