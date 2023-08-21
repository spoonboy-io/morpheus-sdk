# MorpheusApi.ApiTaskSetsTaskSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A unique name for the workflow | 
**description** | **String** | A description of the workflow | [optional] 
**labels** | **[String]** | An array of Category labels for filtering | [optional] 
**type** | **String** | Workflow type | [optional] [default to &#39;provision&#39;]
**visibility** | **String** | private or public | [optional] [default to &#39;private&#39;]
**optionTypes** | **[Number]** | List of option type IDs for use with operational workflow configuration. | [optional] 
**tasks** | [**ApiTaskSetsTaskSetTasks**](ApiTaskSetsTaskSetTasks.md) |  | [optional] 



## Enum: TypeEnum


* `provision` (value: `"provision"`)

* `operation` (value: `"operation"`)





## Enum: VisibilityEnum


* `private` (value: `"private"`)

* `public` (value: `"public"`)




