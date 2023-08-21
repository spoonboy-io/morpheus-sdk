# MorpheusApi.ApiJobsIdJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A name for the Job | [optional] 
**labels** | **[String]** | Array of label strings, can be used for filtering. | [optional] 
**enabled** | **Boolean** | Use this to set enabled state | [optional] [default to true]
**task** | [**ApiJobsIdJobTask**](ApiJobsIdJobTask.md) |  | [optional] 
**workflow** | [**ApiJobsIdJobTask**](ApiJobsIdJobTask.md) |  | [optional] 
**scanPath** | **String** | Scan Checklist. Only applies to type scap-package. | [optional] 
**securityProfile** | **String** | Security Profile. Only applies to type scap-package. | [optional] 
**targetType** | **String** | Target type where job will execute | [optional] 
**targets** | [**[ApiJobsIdJobTargets]**](ApiJobsIdJobTargets.md) |  | [optional] 
**instanceLabel** | **String** | Instance Label. Only applicable if &#x60;targetType&#x60; is &#x60;instance-label&#x60;. | [optional] 
**serverLabel** | **String** | Server Label. Only applicable if &#x60;targetType&#x60; is &#x60;server-label&#x60;. | [optional] 
**scheduleMode** | [**OneOfstringlong**](OneOfstringlong.md) |  | [optional] 
**customOptions** | **Object** | Map of options to be used as values in the workflow tasks. These correspond to option types. | [optional] 
**customConfig** | **String** | Job custom configuration (String in JSON format) | [optional] 
**dateTime** | **Date** | Date and Time to execute the job. Use UTC time in the format 2020-02-15T05:00:00Z. Required when scheduleMode is &#39;dateTime&#39;. | [optional] 
**run** | **Boolean** | If true, executes job | [optional] 



## Enum: TargetTypeEnum


* `appliance` (value: `"appliance"`)

* `instance` (value: `"instance"`)

* `instance-label` (value: `"instance-label"`)

* `server` (value: `"server"`)

* `server-label` (value: `"server-label"`)




