# MorpheusApi.ApiTasksIdExecuteJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A name for the execution job. Can be used to find execution results with &#x60;/api/processes?name&#x3D;&#x60; | [optional] 
**targetType** | **String** | The target context for task execution. This is required for tasks with &#x60;executeTarget&#x60; set to &#x60;resource&#x60;. | [optional] 
**instances** | **[Number]** | Array of Instance IDs. Only applicable if &#x60;targetType&#x60; is instance. | [optional] 
**servers** | **[Number]** | Array of Server IDs. Only applicable if &#x60;targetType&#x60; is &#x60;server&#x60;. | [optional] 
**instanceLabel** | **String** | Instance Label. Only applicable if &#x60;targetType&#x60; is &#x60;instance-label&#x60;. | [optional] 
**serverLabel** | **String** | Server Label. Only applicable if &#x60;targetType&#x60; is &#x60;server-label&#x60;. | [optional] 
**customOptions** | **Object** | Map of options to be used as values in the task. These correspond to option types. | [optional] 
**customConfig** | **String** | String of custom configuration values as JSON. | [optional] 



## Enum: TargetTypeEnum


* `appliance` (value: `"appliance"`)

* `instance` (value: `"instance"`)

* `instance-label` (value: `"instance-label"`)

* `server` (value: `"server"`)

* `server-label` (value: `"server-label"`)




