# MorpheusApi.ApiTaskSetsIdTaskSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A unique name for the workflow | [optional] 
**description** | **String** | A description of the workflow | [optional] 
**labels** | **[String]** | An array of Category labels for filtering | [optional] 
**type** | **String** | Workflow type | [optional] [default to &#39;provision&#39;]
**optionTypes** | **[Number]** | List of option type IDs for use with operational workflow configuration. | [optional] 
**tasks** | [**ApiTaskSetsTaskSetTasks**](ApiTaskSetsTaskSetTasks.md) |  | [optional] 



## Enum: TypeEnum


* `provision` (value: `"provision"`)

* `operation` (value: `"operation"`)




