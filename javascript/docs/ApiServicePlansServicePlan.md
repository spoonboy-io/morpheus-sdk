# MorpheusApi.ApiServicePlansServicePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Service plan name | 
**code** | **String** | Service plan code, must be unique | 
**description** | **String** | Service plan description | [optional] 
**editable** | **Boolean** | Can be used to enable / disable the editability of the service plan. | [optional] [default to true]
**maxStorage** | **Number** | Max storage size in bytes | 
**maxMemory** | **Number** | Max memory size in bytes | 
**maxCores** | **Number** | Max cores | [optional] 
**maxDisks** | **Number** | Max disks allowed | [optional] 
**provisionType** | [**[ApiServicePlansServicePlanProvisionType]**](ApiServicePlansServicePlanProvisionType.md) |  | 
**customCores** | **Boolean** | Can be used to enable / disable customizable cores | [optional] [default to false]
**customMaxStorage** | **Boolean** | Can be used to enable / disable customizable storage | [optional] [default to false]
**customMaxDataStorage** | **Boolean** | Can be used to enable / disable customizable extra volumes. | [optional] [default to false]
**customMaxMemory** | **Boolean** | Can be used to enable / disable customizable memory. | [optional] [default to false]
**addVolumes** | **Boolean** | Can be used to enable / disable ability to add volumes | [optional] [default to false]
**sortOrder** | **Number** | Sort order | [optional] 
**priceSets** | **[Number]** | List of price sets to include in service plan | [optional] 
**config** | [**ApiServicePlansServicePlanConfig**](ApiServicePlansServicePlanConfig.md) |  | [optional] 


