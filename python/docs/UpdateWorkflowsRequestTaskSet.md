# UpdateWorkflowsRequestTaskSet


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | A unique name for the workflow | [optional] 
**description** | **str** | A description of the workflow | [optional] 
**labels** | **[str]** | An array of Category labels for filtering | [optional] 
**type** | **str** | Workflow type | [optional]  if omitted the server will use the default value of "provision"
**option_types** | **[int]** | List of option type IDs for use with operational workflow configuration. | [optional] 
**tasks** | [**AddWorkflowsRequestTaskSetTasks**](AddWorkflowsRequestTaskSetTasks.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


