

# ApiTaskSetsTaskSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A unique name for the workflow | 
**description** | **String** | A description of the workflow |  [optional]
**labels** | **List&lt;String&gt;** | An array of Category labels for filtering |  [optional]
**type** | [**TypeEnum**](#TypeEnum) | Workflow type |  [optional]
**visibility** | [**VisibilityEnum**](#VisibilityEnum) | private or public |  [optional]
**optionTypes** | **List&lt;Long&gt;** | List of option type IDs for use with operational workflow configuration. |  [optional]
**tasks** | [**ApiTaskSetsTaskSetTasks**](ApiTaskSetsTaskSetTasks.md) |  |  [optional]



## Enum: TypeEnum

Name | Value
---- | -----
PROVISION | &quot;provision&quot;
OPERATION | &quot;operation&quot;



## Enum: VisibilityEnum

Name | Value
---- | -----
PRIVATE | &quot;private&quot;
PUBLIC | &quot;public&quot;



