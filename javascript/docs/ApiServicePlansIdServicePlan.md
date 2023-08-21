# MorpheusApi.ApiServicePlansIdServicePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Service plan name | [optional] 
**code** | **String** | Service plan code, must be unique | [optional] 
**description** | **String** | Service plan description | [optional] 
**editable** | **Boolean** | Can be used to enable / disable the editability of the service plan. | [optional] [default to true]
**maxStorage** | **Number** | Max storage size in bytes | [optional] 
**maxMemory** | **Number** | Max memory size in bytes | [optional] 
**maxCores** | **Number** | Max cores | [optional] 
**maxDisks** | **Number** | Max disks allowed | [optional] 
**provisionType** | [**[ApiServicePlansIdServicePlanProvisionType]**](ApiServicePlansIdServicePlanProvisionType.md) |  | [optional] 
**customCores** | **Boolean** | Can be used to enable / disable customizable cores | [optional] [default to false]
**customMaxStorage** | **Boolean** | Can be used to enable / disable customizable storage | [optional] [default to false]
**customMaxDataStorage** | **Boolean** | Can be used to enable / disable customizable extra volumes. | [optional] [default to false]
**customMaxMemory** | **Boolean** | Can be used to enable / disable customizable memory. | [optional] [default to false]
**addVolumes** | **Boolean** | Can be used to enable / disable ability to add volumes | [optional] [default to false]
**sortOrder** | **Number** | Sort order | [optional] 
**priceSets** | **[Number]** | List of price sets to include in service plan | [optional] 
**config** | [**ApiServicePlansIdServicePlanConfig**](ApiServicePlansIdServicePlanConfig.md) |  | [optional] 


