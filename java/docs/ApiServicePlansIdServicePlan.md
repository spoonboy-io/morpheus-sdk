

# ApiServicePlansIdServicePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Service plan name |  [optional]
**code** | **String** | Service plan code, must be unique |  [optional]
**description** | **String** | Service plan description |  [optional]
**editable** | **Boolean** | Can be used to enable / disable the editability of the service plan. |  [optional]
**maxStorage** | **BigDecimal** | Max storage size in bytes |  [optional]
**maxMemory** | **BigDecimal** | Max memory size in bytes |  [optional]
**maxCores** | **BigDecimal** | Max cores |  [optional]
**maxDisks** | **BigDecimal** | Max disks allowed |  [optional]
**provisionType** | [**List&lt;ApiServicePlansIdServicePlanProvisionType&gt;**](ApiServicePlansIdServicePlanProvisionType.md) |  |  [optional]
**customCores** | **Boolean** | Can be used to enable / disable customizable cores |  [optional]
**customMaxStorage** | **Boolean** | Can be used to enable / disable customizable storage |  [optional]
**customMaxDataStorage** | **Boolean** | Can be used to enable / disable customizable extra volumes. |  [optional]
**customMaxMemory** | **Boolean** | Can be used to enable / disable customizable memory. |  [optional]
**addVolumes** | **Boolean** | Can be used to enable / disable ability to add volumes |  [optional]
**sortOrder** | **BigDecimal** | Sort order |  [optional]
**priceSets** | **List&lt;Long&gt;** | List of price sets to include in service plan |  [optional]
**config** | [**ApiServicePlansIdServicePlanConfig**](ApiServicePlansIdServicePlanConfig.md) |  |  [optional]



